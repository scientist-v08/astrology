package shadbala

import (
	"fmt"
	"math"

	"github.com/mshafiee/swephgo"
	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
	"github.com/scientist-v08/bphs/utils"
)

type ChestaBalaResult struct {
	Values map[models.Raashyadhipati]float64 `json:"values"`
}

type ChestaElements struct {
	MeanLongitude float64
	TrueLongitude float64
	Seeghrochcha  float64
}

func calculateChestaBala(
	ctx *chartContext,
	kaala KaalaBalaResult,
) ChestaBalaResult {
	res := make(map[models.Raashyadhipati]float64, 7)

	// Sun's Ayana Bala and Moon's Paksha Bala are already doubled (0–120)
	// in Kaala Bala, so their Chesta Bala is 0 to avoid double-counting.
	res[constants.Surya] = 0
	res[constants.Chandra] = 0

	elements, err := computeChestaElements(ctx)
	if err != nil {
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
		res[g] = chestaBala(e.Seeghrochcha, e.MeanLongitude, e.TrueLongitude)
	}
	fmt.Println(res)

	return ChestaBalaResult{Values: res}
}

func chestaBala(seeghrochcha, meanLon, trueLon float64) float64 {
	// Circular midpoint to handle 0°/360° wraparound safely.
	avg := (meanLon + trueLon) / 2.0
	if math.Abs(meanLon-trueLon) > 180.0 {
		avg = utils.Normalize360(avg + 180.0)
	}

	kendra := utils.Normalize360(seeghrochcha - avg)
	if kendra > 180.0 {
		kendra = 360.0 - kendra
	}

	return kendra / 3.0
}

func computeChestaElements(
	ctx *chartContext,
) (map[models.Raashyadhipati]ChestaElements, error) {
	jd := ctx.birthDateTime
	ayanamsha := swephgo.GetAyanamsaUt(jd)

	toSidereal := func(lon float64) float64 {
		return utils.Normalize360(lon - ayanamsha)
	}

	meanSun, err := meanSunLongitude(jd)
	if err != nil {
		return nil, fmt.Errorf("mean Sun: %w", err)
	}
	meanMars, err := meanLongitude(jd, swephgo.SeMars)
	if err != nil {
		return nil, fmt.Errorf("mean Mars: %w", err)
	}
	meanMerc, err := meanLongitude(jd, swephgo.SeMercury)
	if err != nil {
		return nil, fmt.Errorf("mean Mercury: %w", err)
	}
	meanJup, err := meanLongitude(jd, swephgo.SeJupiter)
	if err != nil {
		return nil, fmt.Errorf("mean Jupiter: %w", err)
	}
	meanVen, err := meanLongitude(jd, swephgo.SeVenus)
	if err != nil {
		return nil, fmt.Errorf("mean Venus: %w", err)
	}
	meanSat, err := meanLongitude(jd, swephgo.SeSaturn)
	if err != nil {
		return nil, fmt.Errorf("mean Saturn: %w", err)
	}

	meanSun = toSidereal(meanSun)
	meanMars = toSidereal(meanMars)
	meanMerc = toSidereal(meanMerc)
	meanJup = toSidereal(meanJup)
	meanVen = toSidereal(meanVen)
	meanSat = toSidereal(meanSat)

	return map[models.Raashyadhipati]ChestaElements{
		// Outer planets: own mean longitude, Sun as Seeghrochcha
		constants.Kuja: {
			MeanLongitude: meanMars,
			TrueLongitude: ctx.longitude[constants.Kuja],
			Seeghrochcha:  meanSun,
		},
		constants.Guru: {
			MeanLongitude: meanJup,
			TrueLongitude: ctx.longitude[constants.Guru],
			Seeghrochcha:  meanSun,
		},
		constants.Shani: {
			MeanLongitude: meanSat,
			TrueLongitude: ctx.longitude[constants.Shani],
			Seeghrochcha:  meanSun,
		},
		// Inner planets: Sun as mean longitude, own mean heliocentric longitude as Seeghrochcha
		constants.Budha: {
			MeanLongitude: meanSun,
			TrueLongitude: ctx.longitude[constants.Budha],
			Seeghrochcha:  meanMerc,
		},
		constants.Shukra: {
			MeanLongitude: meanSun,
			TrueLongitude: ctx.longitude[constants.Shukra],
			Seeghrochcha:  meanVen,
		},
	}, nil
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