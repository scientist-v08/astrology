package shadbala

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/mshafiee/swephgo"
	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/dto"
)

type KaalaBalaResult struct {
	Nathonnata map[dto.Raashyadhipati]float64 `json:"nathonnata"`
	Paksha     map[dto.Raashyadhipati]float64 `json:"paksha"`
	Tribhaga   map[dto.Raashyadhipati]float64 `json:"tribhaga"`
	Ayana      map[dto.Raashyadhipati]float64 `json:"ayana"`
	Abda       map[dto.Raashyadhipati]float64 `json:"abda"`
	Masa       map[dto.Raashyadhipati]float64 `json:"masa"`
	Vara       map[dto.Raashyadhipati]float64 `json:"vara"`
	Hora       map[dto.Raashyadhipati]float64 `json:"hora"`
	// Yuddha     map[models.Raashyadhipati]float64 `json:"yuddha"`
	Total      map[dto.Raashyadhipati]float64 `json:"total"`
}

// weekdayFromJD returns the planetary lord of the civil day containing the JD
// (Sunday = Surya … Saturday = Shani). Uses the standard astronomical convention.
func weekdayFromJD(jd, lon float64) dto.Raashyadhipati {
	// Floor(JD + 0.5) gives the civil day number; JD 0.0 was a Monday.
	jdn := int64(math.Floor(jd + lon/360.0 + 0.5))
	rem := (jdn + 1) % 7
	// We want: 0 = Sunday, 1 = Monday … 6 = Saturday
	switch rem {
	case 0:
		return constants.Surya
	case 1:
		return constants.Chandra
	case 2:
		return constants.Kuja
	case 3:
		return constants.Budha
	case 4:
		return constants.Guru
	case 5:
		return constants.Shukra
	default:
		return constants.Shani
	}
}

// weekdayLord is the same as weekdayFromJD but used for Hora sequence start.
func weekdayLord(jd, lon float64) dto.Raashyadhipati { return weekdayFromJD(jd, lon) }

func calculateAbdaBala(meshaSankrantiJD, lon float64) map[dto.Raashyadhipati]float64 {
	res := make(map[dto.Raashyadhipati]float64, 7)
	for _, g := range shadbalaGrahas {
		res[g] = 0
	}
	if meshaSankrantiJD > 0 {
		lord := weekdayFromJD(meshaSankrantiJD, lon)
		res[lord] = 15.0
	}
	return res
}

func calculateMasaBala(solarMonthEntryJD, lon float64) map[dto.Raashyadhipati]float64 {
	res := make(map[dto.Raashyadhipati]float64, 7)
	for _, g := range shadbalaGrahas {
		res[g] = 0
	}
	if solarMonthEntryJD > 0 {
		lord := weekdayFromJD(solarMonthEntryJD, lon)
		res[lord] = 30.0
	}
	return res
}

func calculateVaraBala(sunriseJD, lon float64) map[dto.Raashyadhipati]float64 {
	res := make(map[dto.Raashyadhipati]float64, 7)
	for _, g := range shadbalaGrahas {
		res[g] = 0
	}
	lord := weekdayFromJD(sunriseJD, lon)
	res[lord] = 45.0
	return res
}

func calculateSolarEvent(startJD float64, event int, lat, lon float64) (float64, error) {
	geoPos := []float64{lon, lat, 0}
	tret := make([]float64, 1)
	serr := make([]byte, 256)

	ret := swephgo.RiseTrans(
		startJD,
		swephgo.SeSun,
		[]byte{},
		swephgo.SeflgSwieph,
		event|swephgo.SeBitHinduRising,
		geoPos,
		0,
		15,
		tret,
		serr,
	)
	if ret != 0 {
		return 0, fmt.Errorf("solar event failed: %s (code %d)",
			strings.TrimRight(string(serr), "\x00"), ret)
	}
	return tret[0], nil
}

func getBirthSunriseInterval(ctx *chartContext, birthJD, lat, lon float64) (float64, float64, error) {
	var startJD float64
	if birthJD < ctx.sunriseJD {
		prev, err := calculateSolarEvent(ctx.sunriseJD-1.0, swephgo.SeCalcRise, lat, lon)
		if err != nil {
			return 0, 0, err
		}
		startJD = prev
	} else {
		startJD = ctx.sunriseJD
	}
	next, err := calculateSolarEvent(startJD+1e-6, swephgo.SeCalcRise, lat, lon)
	if err != nil {
		return 0, 0, err
	}
	return startJD, next, nil
}

func calculateHoraBala(birthJD, sunriseJD, nextSunriseJD, lon float64) (map[dto.Raashyadhipati]float64, error) {
	res := make(map[dto.Raashyadhipati]float64, 7)
	for _, g := range shadbalaGrahas {
		res[g] = 0
	}

	interval := nextSunriseJD - sunriseJD
	if interval <= 0 {
		return nil, fmt.Errorf("invalid sunrise interval")
	}
	if birthJD < sunriseJD || birthJD >= nextSunriseJD {
		return nil, fmt.Errorf("birth JD outside Hora interval")
	}

	fraction := (birthJD - sunriseJD) / interval
	horaIndex := int(math.Floor(fraction * 24.0))
	if horaIndex > 23 {
		horaIndex = 23
	}

	// Classical Chaldean order
	horaSequence := []dto.Raashyadhipati{
		constants.Shani, constants.Guru, constants.Kuja,
		constants.Surya, constants.Shukra, constants.Budha, constants.Chandra,
	}

	dayLord := weekdayLord(sunriseJD, lon)
	startIndex := -1
	for i, g := range horaSequence {
		if g == dayLord {
			startIndex = i
			break
		}
	}
	if startIndex == -1 {
		return nil, fmt.Errorf("invalid weekday lord")
	}

	lord := horaSequence[(startIndex+horaIndex)%7]
	res[lord] = 60.0
	return res, nil
}

// calculateNathonnataBala implements the continuous (not binary) form used by
// most modern software following Jaya Sekhar / BV Raman.
func calculateNathonnataBala(ctx *chartContext, birthJD, lat, lon float64) (map[dto.Raashyadhipati]float64, error) {
	res := make(map[dto.Raashyadhipati]float64, 7)

	// Local midnight of the civil day of birth
	birthUTC := julianDayToTime(birthJD)
	// Use a fixed offset for the location (caller can improve with real TZ)
	loc := time.FixedZone("LOC", int((lon/15.0)*3600))
	birthLocal := birthUTC.In(loc)
	localMidnight := time.Date(birthLocal.Year(), birthLocal.Month(), birthLocal.Day(),
		0, 0, 0, 0, loc)
	midnightJD := julianDayUT(localMidnight)

	sunriseJD, err := calculateSolarEvent(midnightJD, swephgo.SeCalcRise, lat, lon)
	if err != nil {
		return nil, fmt.Errorf("sunrise: %w", err)
	}
	sunsetJD, err := calculateSolarEvent(midnightJD, swephgo.SeCalcSet, lat, lon)
	if err != nil {
		return nil, fmt.Errorf("sunset: %w", err)
	}
	prevSunsetJD, err := calculateSolarEvent(midnightJD-1.0, swephgo.SeCalcSet, lat, lon)
	if err != nil {
		return nil, fmt.Errorf("previous sunset: %w", err)
	}
	nextSunriseJD, err := calculateSolarEvent(sunsetJD+1e-6, swephgo.SeCalcRise, lat, lon)
	if err != nil {
		return nil, fmt.Errorf("next sunrise: %w", err)
	}

	// Update context for downstream use
	ctx.sunriseJD = sunriseJD
	ctx.sunsetJD = sunsetJD
	ctx.isDayBirth = birthJD >= sunriseJD && birthJD < sunsetJD

	// Distance from local midnight (mid-point of the night)
	var nightStart, nightEnd float64
	if birthJD < sunriseJD {
		nightStart, nightEnd = prevSunsetJD, sunriseJD
	} else {
		nightStart, nightEnd = sunsetJD, nextSunriseJD
	}
	nightLen := nightEnd - nightStart
	if nightLen <= 0 {
		return nil, fmt.Errorf("invalid nighttime duration")
	}
	localMidnightJD := nightStart + nightLen/2.0

	distanceHours := math.Abs(birthJD-localMidnightJD) * 24.0
	if distanceHours > 12.0 {
		distanceHours = 24.0 - distanceHours
	}

	unnata := math.Min(60.0, distanceHours*5.0) // 0 at midnight → 60 at noon
	nata := 60.0 - unnata

	res[constants.Budha] = 60.0 // always full
	for _, g := range shadbalaGrahas {
		switch {
		case g == constants.Budha:
			// already set
		case isDayPlanet(g):
			res[g] = unnata
		case isNightPlanet(g):
			res[g] = nata
		}
	}
	return res, nil
}

func calculatePakshaBala(ctx *chartContext) map[dto.Raashyadhipati]float64 {
	res := make(map[dto.Raashyadhipati]float64, 7)
	elong := angularSeparation(ctx.longitude[constants.Chandra], ctx.longitude[constants.Surya]) // 0..180
	base := elong / 3.0

	for _, g := range shadbalaGrahas {
		switch g {
		case constants.Chandra:
			res[g] = 2 * base
		case constants.Budha, constants.Guru, constants.Shukra:
			res[g] = base
		default: // Surya, Kuja, Shani
			res[g] = 60 - base
		}
	}
	return res
}

// Mercury is a malefic when it shares a sign with a malefic.
/*func mercuryIsBenefic(ctx *chartContext) bool {
	sign := func(g models.Raashyadhipati) int { return int(ctx.longitude[g] / 30) }
	m := sign(constants.Budha)
	for _, g := range []models.Raashyadhipati{constants.Surya, constants.Kuja, constants.Shani} {
		if sign(g) == m {
			return false
		}
	}
	return true
}*/

func calculateTribhagaBala(ctx *chartContext, birthJD, lat, lon float64) (map[dto.Raashyadhipati]float64, error) {
	res := make(map[dto.Raashyadhipati]float64, 7)
	for _, g := range shadbalaGrahas {
		res[g] = 0
	}
	res[constants.Guru] = 60.0 // always

	sunriseJD := ctx.sunriseJD
	sunsetJD := ctx.sunsetJD
	if sunsetJD <= sunriseJD {
		return nil, fmt.Errorf("invalid sunrise/sunset")
	}

	var periodStart, periodEnd float64
	if birthJD >= sunriseJD && birthJD < sunsetJD {
		ctx.isDayBirth = true
		periodStart, periodEnd = sunriseJD, sunsetJD
	} else {
		ctx.isDayBirth = false
		if birthJD < sunriseJD {
			prevSunset, err := calculateSolarEvent(sunriseJD-1.0, swephgo.SeCalcSet, lat, lon)
			if err != nil {
				return nil, err
			}
			periodStart, periodEnd = prevSunset, sunriseJD
		} else {
			nextSunrise, err := calculateSolarEvent(sunsetJD+1e-6, swephgo.SeCalcRise, lat, lon)
			if err != nil {
				return nil, err
			}
			periodStart, periodEnd = sunsetJD, nextSunrise
		}
	}

	periodLen := periodEnd - periodStart
	if periodLen <= 0 {
		return nil, fmt.Errorf("invalid day/night duration")
	}

	fraction := (birthJD - periodStart) / periodLen
	part := int(math.Floor(fraction * 3.0))
	if part < 0 {
		part = 0
	}
	if part > 2 {
		part = 2
	}

	if ctx.isDayBirth {
		switch part {
		case 0:
			res[constants.Budha] = 60
		case 1:
			res[constants.Surya] = 60
		case 2:
			res[constants.Shani] = 60
		}
	} else {
		switch part {
		case 0:
			res[constants.Chandra] = 60
		case 1:
			res[constants.Shukra] = 60
		case 2:
			res[constants.Kuja] = 60
		}
	}
	return res, nil
}

func calculateAyanaBala(ctx *chartContext) map[dto.Raashyadhipati]float64 {
	res := make(map[dto.Raashyadhipati]float64, 7)
	const maxDecl = 23.45 // 23°27'

	for _, g := range shadbalaGrahas {
		decl := ctx.declinations[g]

		// "favourable" is positive on the side where the planet gains strength
		var favourable float64
		switch g {
		case constants.Surya, constants.Kuja, constants.Guru, constants.Shukra:
			favourable = decl // north is strong
		case constants.Chandra, constants.Shani:
			favourable = -decl // south is strong
		case constants.Budha:
			favourable = math.Abs(decl) // always positive
		}

		ratio := math.Max(-1, math.Min(1, favourable/maxDecl))
		bala := 30 + 30*ratio

		if g == constants.Surya {
			bala *= 2 // Parashara doubles the Sun's Ayana Bala
		}
		res[g] = bala
	}
	return res
}

// Yuddha Bala – simplified classical form used by most software.
// Winner = higher ecliptic latitude (or Venus always wins). Difference of
// preliminary strengths is transferred.
func calculateYuddhaBala(
	data dto.PlanetaryData,
	prelim map[dto.Raashyadhipati]float64,
) map[dto.Raashyadhipati]float64 {

	res := make(map[dto.Raashyadhipati]float64, 5)
	planets := []dto.Raashyadhipati{
		constants.Kuja, constants.Budha, constants.Guru,
		constants.Shukra, constants.Shani,
	}
	for _, g := range planets {
		res[g] = 0
	}

	for i := 0; i < len(planets); i++ {
		for j := i + 1; j < len(planets); j++ {
			p1, p2 := planets[i], planets[j]
			lon1, ok1 := data.Longitudes[p1]
			lon2, ok2 := data.Longitudes[p2]
			lat1, okL1 := data.EclipticLatitudes[p1]
			lat2, okL2 := data.EclipticLatitudes[p2]
			if !ok1 || !ok2 || !okL1 || !okL2 {
				continue
			}

			if angularSeparation(lon1, lon2) > 1.0 {
				continue
			}

			var winner, loser dto.Raashyadhipati
			if p1 == constants.Shukra {
				winner, loser = p1, p2
			} else if p2 == constants.Shukra {
				winner, loser = p2, p1
			} else if lat1 > lat2 {
				winner, loser = p1, p2
			} else if lat2 > lat1 {
				winner, loser = p2, p1
			} else {
				continue // no clear winner
			}

			diff := math.Abs(prelim[winner] - prelim[loser])
			res[winner] += diff
			res[loser] -= diff
		}
	}
	return res
}

func calculateKaalaBala(ctx *chartContext, planets dto.PlanetaryData) KaalaBalaResult {
	nath, _ := calculateNathonnataBala(ctx, ctx.birthDateTime, ctx.reqLatitude, ctx.reqLongitude)
	pak := calculatePakshaBala(ctx)
	tri, _ := calculateTribhagaBala(ctx, ctx.birthDateTime, ctx.reqLatitude, ctx.reqLongitude)
	aya := calculateAyanaBala(ctx)
	abda := calculateAbdaBala(ctx.meshaJD, ctx.reqLongitude)
	masa := calculateMasaBala(ctx.solarMonthStartJD, ctx.reqLongitude)
	
	// Hora needs the sunrise interval that actually contains the birth
	sunriseStart, nextSunrise, err := getBirthSunriseInterval(ctx, ctx.birthDateTime, ctx.reqLatitude, ctx.reqLongitude)
	vara := calculateVaraBala(sunriseStart, ctx.reqLongitude)
	var hora map[dto.Raashyadhipati]float64
	if err == nil {
		hora, _ = calculateHoraBala(ctx.birthDateTime, sunriseStart, nextSunrise, ctx.reqLongitude)
	} else {
		hora = make(map[dto.Raashyadhipati]float64)
		for _, g := range shadbalaGrahas {
			hora[g] = 0
		}
	}

	// Preliminary total for Yuddha (exclude Ayana & Yuddha itself – classical)
	prelim := make(map[dto.Raashyadhipati]float64, 7)
	for _, g := range shadbalaGrahas {
		prelim[g] = nath[g] + pak[g] + tri[g] + abda[g] + masa[g] + vara[g] + hora[g]
	}
	// yuddha := calculateYuddhaBala(planets, prelim)

	total := make(map[dto.Raashyadhipati]float64, 7)
	for _, g := range shadbalaGrahas {
		total[g] = nath[g] + pak[g] + tri[g] + aya[g] +
			abda[g] + masa[g] + vara[g] + hora[g]
	}

	return KaalaBalaResult{
		Nathonnata: nath,
		Paksha:     pak,
		Tribhaga:   tri,
		Ayana:      aya,
		Abda:       abda,
		Masa:       masa,
		Vara:       vara,
		Hora:       hora,
		Total:      total,
	}
}