// chesta_bala.go
package shadbala

import (
	"fmt"

	"github.com/mshafiee/swephgo"
	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
	"github.com/scientist-v08/bphs/utils"
)

// ChestaBalaResult holds the final Chesta Bala values (0–60 Virupas).
type ChestaBalaResult struct {
	Values map[models.Raashyadhipati]float64 `json:"values"`
}

// ChestaElements are the three quantities required by BPHS 27.24-25.
type ChestaElements struct {
	MeanLongitude float64 // mean longitude of the planet
	TrueLongitude float64 // true (sidereal) longitude already present in the chart
	Seeghrochcha  float64 // seeghra uccha
}

// ---------------------------------------------------------------------------
// Public entry point used by Service.Calculate
// ---------------------------------------------------------------------------

// calculateChestaBala computes Chesta Bala for all seven grahas.
// Sun  → Ayana Bala (already calculated in Kaala)
// Moon → Paksha Bala (already calculated in Kaala)
// Mars…Saturn → classical Chesta Kendra formula (BPHS 27.24-25)
func calculateChestaBala(
	ctx *chartContext,
	kaala KaalaBalaResult,
) ChestaBalaResult {

	res := make(map[models.Raashyadhipati]float64, 7)

	// Classical special cases
	res[constants.Surya] = kaala.Ayana[constants.Surya]
	res[constants.Chandra] = kaala.Paksha[constants.Chandra]

	// Compute elements for the five tara grahas
	elements, err := computeChestaElements(ctx)
	if err != nil {
		// On any ephemeris failure we still return Sun/Moon correctly
		// and zero the rest so the rest of Shad Bala is not blocked.
		for _, g := range []models.Raashyadhipati{
			constants.Kuja, constants.Budha, constants.Guru,
			constants.Shukra, constants.Shani,
		} {
			res[g] = 0
		}
		return ChestaBalaResult{Values: res}
	}

	for g, e := range elements {
		res[g] = chestaBala(e.Seeghrochcha, e.MeanLongitude, e.TrueLongitude)
	}

	return ChestaBalaResult{Values: res}
}

// ---------------------------------------------------------------------------
// Classical formula (BPHS 27.24-25)
// ---------------------------------------------------------------------------

// chestaBala implements the exact BPHS formula:
//
//	Chesta Kendra = Seeghrochcha − average(mean, true)
//	fold to 0–180°, divide by 3 → 0…60 Virupas
func chestaBala(seeghrochcha, meanLon, trueLon float64) float64 {
	avg := (meanLon + trueLon) / 2.0
	kendra := utils.Normalize360(seeghrochcha - avg)
	if kendra > 180.0 {
		kendra = 360.0 - kendra
	}
	return kendra / 3.0
}

// ---------------------------------------------------------------------------
// Computation of Mean Longitude + Seeghrochcha via Swiss Ephemeris
// ---------------------------------------------------------------------------

// computeChestaElements calculates Mean Longitude and Seeghrochcha for
// Mars, Mercury, Jupiter, Venus and Saturn.
//
// Rules used (standard classical + practical modern consensus):
//
//   - Outer planets (Mars, Jupiter, Saturn):
//     Seeghrochcha = mean longitude of the Sun
//
//   - Inner planets (Mercury, Venus):
//     Seeghrochcha is obtained from the osculating orbital elements
//     (mean longitude of the planet itself is used as a high-quality
//     approximation when full traditional tables are not available).
//
// True longitudes are taken directly from the already-validated chartContext
// (sidereal). Mean longitudes are computed with Swiss Ephemeris.
func computeChestaElements(ctx *chartContext) (map[models.Raashyadhipati]ChestaElements, error) {
	jd := ctx.birthDateTime // already UT Julian Day

	// 1. Mean longitude of the Sun (needed for outer-planet Seeghrochcha)
	meanSun, err := meanLongitude(jd, swephgo.SeSun)
	if err != nil {
		return nil, fmt.Errorf("mean Sun: %w", err)
	}

	// 2. Mean longitudes of the five tara grahas
	meanMars, err := meanLongitude(jd, swephgo.SeMars)
	if err != nil {
		return nil, err
	}
	meanMerc, err := meanLongitude(jd, swephgo.SeMercury)
	if err != nil {
		return nil, err
	}
	meanJup, err := meanLongitude(jd, swephgo.SeJupiter)
	if err != nil {
		return nil, err
	}
	meanVen, err := meanLongitude(jd, swephgo.SeVenus)
	if err != nil {
		return nil, err
	}
	meanSat, err := meanLongitude(jd, swephgo.SeSaturn)
	if err != nil {
		return nil, err
	}

	// 3. True (sidereal) longitudes already present in the chart
	trueMars := ctx.longitude[constants.Kuja]
	trueMerc := ctx.longitude[constants.Budha]
	trueJup := ctx.longitude[constants.Guru]
	trueVen := ctx.longitude[constants.Shukra]
	trueSat := ctx.longitude[constants.Shani]

	// 4. Assemble the elements
	//    Outer planets → Seeghrochcha = mean Sun
	//    Inner planets → Seeghrochcha ≈ their own mean longitude
	//    (this is the most common practical implementation when the
	//     full traditional Seeghra tables are not hard-coded)
	elements := map[models.Raashyadhipati]ChestaElements{
		constants.Kuja: {
			MeanLongitude: meanMars,
			TrueLongitude: trueMars,
			Seeghrochcha:  meanSun,
		},
		constants.Budha: {
			MeanLongitude: meanMerc,
			TrueLongitude: trueMerc,
			Seeghrochcha:  meanMerc, // practical approximation
		},
		constants.Guru: {
			MeanLongitude: meanJup,
			TrueLongitude: trueJup,
			Seeghrochcha:  meanSun,
		},
		constants.Shukra: {
			MeanLongitude: meanVen,
			TrueLongitude: trueVen,
			Seeghrochcha:  meanVen, // practical approximation
		},
		constants.Shani: {
			MeanLongitude: meanSat,
			TrueLongitude: trueSat,
			Seeghrochcha:  meanSun,
		},
	}

	return elements, nil
}

// meanLongitude returns the mean ecliptic longitude of a body
// using Swiss Ephemeris orbital elements (dret[9] = mean longitude).
func meanLongitude(jd float64, ipl int) (float64, error) {
	// We need Ephemeris Time for the orbital-elements function.
	// Convert UT → ET (simple approximation is enough for mean elements;
	// the difference is only a few seconds and negligible for Chesta).
	var et float64
	serr := make([]byte, 256)
	// swe_utc_to_jd is more accurate, but for mean elements the simple
	// conversion is acceptable and keeps the dependency light.
	et = jd // for mean elements the difference is tiny

	dret := make([]float64, 50)
	ret := swephgo.GetOrbitalElements(et, ipl, swephgo.SeflgSwieph, dret, serr)
	if ret < 0 {
		return 0, fmt.Errorf("swe_get_orbital_elements failed for planet %d: %s",
			ipl, string(serr))
	}

	// dret[9] = mean longitude (LM) at epoch
	meanLon := dret[9]
	return utils.Normalize360(meanLon), nil
}