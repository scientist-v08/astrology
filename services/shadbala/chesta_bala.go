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

// ChestaElements are the three quantities required by BPHS.
type ChestaElements struct {
	MeanLongitude float64
	TrueLongitude float64
	Seeghrochcha  float64
}

// calculateChestaBala computes Chesta Bala for all seven grahas.
func calculateChestaBala(
	ctx *chartContext,
	kaala KaalaBalaResult,
) ChestaBalaResult {

	res := make(map[models.Raashyadhipati]float64, 7)

	// Sun and Moon special cases.
	res[constants.Surya] = kaala.Ayana[constants.Surya]
	res[constants.Chandra] = kaala.Paksha[constants.Chandra]

	elements, err := computeChestaElements(ctx)
	if err != nil {
		// Preserve the error for debugging rather than silently
		// pretending that the calculated Bala is zero.
		fmt.Printf("Chesta Bala calculation failed: %v\n", err)

		for _, g := range []models.Raashyadhipati{
			constants.Kuja,
			constants.Budha,
			constants.Guru,
			constants.Shukra,
			constants.Shani,
		} {
			res[g] = 0
		}

		return ChestaBalaResult{Values: res}
	}

	for g, e := range elements {
		res[g] = chestaBala(
			e.Seeghrochcha,
			e.MeanLongitude,
			e.TrueLongitude,
		)
	}

	return ChestaBalaResult{Values: res}
}

// chestaBala implements the Chesta Kendra formula.
//
// Chesta Kendra = Seeghrochcha - average(mean, true)
// Fold to 0–180 degrees, divide by 3.
func chestaBala(
	seeghrochcha float64,
	meanLon float64,
	trueLon float64,
) float64 {

	avg := (meanLon + trueLon) / 2.0

	kendra := utils.Normalize360(seeghrochcha - avg)

	if kendra > 180.0 {
		kendra = 360.0 - kendra
	}

	return kendra / 3.0
}

// computeChestaElements calculates Chesta Bala inputs.
//
// Convention:
//   Mars, Jupiter, Saturn:
//       Mean longitude = Mean Sun
//       Seeghrochcha = Mean Sun
//
//   Mercury, Venus:
//       Mean longitude = Mean Sun
//       Seeghrochcha = Own mean heliocentric longitude
//
// NOTE: This is a practical approximation, not yet a verified
// reproduction of Jagannatha Hora's internal calculation.

func computeChestaElements(
    ctx *chartContext,
) (map[models.Raashyadhipati]ChestaElements, error) {

    jd := ctx.birthDateTime

    // 1. Obtain the ayanamsha for the birth moment.
    // Swiss Ephemeris uses the currently configured sidereal mode.
    ayanamsha := swephgo.GetAyanamsaUt(jd)

    // Convert tropical longitude to sidereal longitude.
    toSidereal := func(lon float64) float64 {
        return utils.Normalize360(lon - ayanamsha)
    }

    // 2. Calculate mean longitude of the Sun.
    meanSun, err := meanSunLongitude(jd)
    if err != nil {
        return nil, fmt.Errorf("mean Sun: %w", err)
    }

    // 3. Calculate mean heliocentric longitudes.
    meanMerc, err := meanLongitude(jd, swephgo.SeMercury)
    if err != nil {
        return nil, fmt.Errorf("mean Mercury: %w", err)
    }

    meanVen, err := meanLongitude(jd, swephgo.SeVenus)
    if err != nil {
        return nil, fmt.Errorf("mean Venus: %w", err)
    }

    // 4. Convert tropical mean longitudes to sidereal.
    meanSun = toSidereal(meanSun)
    meanMerc = toSidereal(meanMerc)
    meanVen = toSidereal(meanVen)

    // 5. Obtain true sidereal longitudes from chart context.
    trueMars := ctx.longitude[constants.Kuja]
    trueMerc := ctx.longitude[constants.Budha]
    trueJup := ctx.longitude[constants.Guru]
    trueVen := ctx.longitude[constants.Shukra]
    trueSat := ctx.longitude[constants.Shani]

    // 6. Assemble Chesta Bala elements.
    elements := map[models.Raashyadhipati]ChestaElements{

        constants.Kuja: {
            MeanLongitude: meanSun,
            TrueLongitude: trueMars,
            Seeghrochcha:  meanSun,
        },

        constants.Budha: {
            MeanLongitude: meanSun,
            TrueLongitude: trueMerc,
            Seeghrochcha:  meanMerc,
        },

        constants.Guru: {
            MeanLongitude: meanSun,
            TrueLongitude: trueJup,
            Seeghrochcha:  meanSun,
        },

        constants.Shukra: {
            MeanLongitude: meanSun,
            TrueLongitude: trueVen,
            Seeghrochcha:  meanVen,
        },

        constants.Shani: {
            MeanLongitude: meanSun,
            TrueLongitude: trueSat,
            Seeghrochcha:  meanSun,
        },
    }

    return elements, nil
}


// meanSunLongitude obtains the Sun's mean longitude.
//
// Uses the solar mean longitude from Swiss Ephemeris orbital
// elements of the Earth, then adds 180 degrees.
//
// This is an approximation and should be validated against
// JHora's solar mean-longitude convention.
func meanSunLongitude(jd float64) (float64, error) {

	dret := make([]float64, 50)
	serr := make([]byte, 256)

	// Convert UT to ET.
	deltaT := swephgo.Deltat(jd)
	et := jd + deltaT

	ret := swephgo.GetOrbitalElements(
		et,
		swephgo.SeEarth,
		swephgo.SeflgSwieph,
		dret,
		serr,
	)

	if ret < 0 {
		return 0, fmt.Errorf(
			"swe_get_orbital_elements failed for Earth: %s",
			string(serr),
		)
	}

	// Earth's mean heliocentric longitude + 180 = Sun's
	// mean geocentric longitude.
	return utils.Normalize360(dret[9] + 180.0), nil
}

// meanLongitude returns the mean heliocentric longitude
// from Swiss Ephemeris orbital elements.
//
// dret[9] = mean longitude (LM).
func meanLongitude(jd float64, ipl int) (float64, error) {

	// Convert UT to ET.
	deltaT := swephgo.Deltat(jd)
	et := jd + deltaT

	dret := make([]float64, 50)
	serr := make([]byte, 256)

	ret := swephgo.GetOrbitalElements(
		et,
		ipl,
		swephgo.SeflgSwieph,
		dret,
		serr,
	)

	if ret < 0 {
		return 0, fmt.Errorf(
			"swe_get_orbital_elements failed for planet %d: %s",
			ipl,
			string(serr),
		)
	}

	return utils.Normalize360(dret[9]), nil
}
