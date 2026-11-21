package shadbala

import (
	"fmt"
	"math"
	"time"

	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/dto"
	"github.com/scientist-v08/bphs/utils"
)

var shadbalaGrahas = []dto.Raashyadhipati{
	constants.Surya, constants.Chandra, constants.Kuja, constants.Budha,
	constants.Guru, constants.Shukra, constants.Shani,
}

// chartContext is the single, validated view of the chart that every bala
// reads from. Sign indices follow constants.RaashiOrder (0 = Aries).
type chartContext struct {
	longitude         map[dto.Raashyadhipati]float64 // sidereal, 0..360
	rasiSign          map[dto.Raashyadhipati]int     // D1 sign index
	navamsaSign       map[dto.Raashyadhipati]int     // D9 sign index
	ascSign           int
	ascendantLong     float64
	chandraWaxing     bool
	speeds            map[dto.Raashyadhipati]float64
	declinations      map[dto.Raashyadhipati]float64 // tropical
	eclipticLatitudes map[dto.Raashyadhipati]float64
	sunriseJD         float64
	sunsetJD          float64
	isDayBirth        bool
	mclong            float64
	birthDateTime     float64 // JD
	reqLatitude       float64
	reqLongitude      float64
	meshaJD           float64
	solarMonthStartJD float64
}

func d1Placements(b dto.HouseEffectsReqBody) map[dto.Raashyadhipati]string {
	return map[dto.Raashyadhipati]string{
		constants.Surya:   b.SuryaPlacement,
		constants.Chandra: b.ChandraPlacement,
		constants.Kuja:    b.KujaPlacement.Placement,
		constants.Budha:   b.BudhaPlacement.Placement,
		constants.Guru:    b.GuruPlacement.Placement,
		constants.Shukra:  b.ShukraPlacement.Placement,
		constants.Shani:   b.ShaniPlacement.Placement,
	}
}

func d9Placements(b dto.HouseEffectsReqBody) map[dto.Raashyadhipati]string {
	return map[dto.Raashyadhipati]string{
		constants.Surya:   b.NavSuryaPlacement,
		constants.Chandra: b.NavChandraPlacement,
		constants.Kuja:    b.NavKujaPlacement,
		constants.Budha:   b.NavBudhaPlacement,
		constants.Guru:    b.NavGuruPlacement,
		constants.Shukra:  b.NavShukraPlacement,
		constants.Shani:   b.NavShaniPlacement,
	}
}

func newChartContext(in Input) (*chartContext, error) {
	c := &chartContext{
		longitude:         make(map[dto.Raashyadhipati]float64, len(shadbalaGrahas)),
		rasiSign:          make(map[dto.Raashyadhipati]int, len(shadbalaGrahas)),
		navamsaSign:       make(map[dto.Raashyadhipati]int, len(shadbalaGrahas)),
		speeds:            make(map[dto.Raashyadhipati]float64, len(shadbalaGrahas)),
		declinations:      make(map[dto.Raashyadhipati]float64, len(shadbalaGrahas)),
		eclipticLatitudes: make(map[dto.Raashyadhipati]float64, len(shadbalaGrahas)),
	}

	d1 := d1Placements(in.Chart)
	d9 := d9Placements(in.Chart)

	for _, g := range shadbalaGrahas {
		lon, ok := in.Planets.Longitudes[g]
		if !ok {
			return nil, fmt.Errorf("missing longitude for %s", g)
		}
		lon = utils.Normalize360(lon)
		c.longitude[g] = lon

		if in.Planets.Speeds != nil {
			c.speeds[g] = in.Planets.Speeds[g]
		}
		if in.Planets.Declinations != nil {
			c.declinations[g] = in.Planets.Declinations[g]
		}
		// Fixed: previously checked Declinations by mistake
		if in.Planets.EclipticLatitudes != nil {
			c.eclipticLatitudes[g] = in.Planets.EclipticLatitudes[g]
		}

		// D1 validation
		rs, err := rashiIndex(d1[g])
		if err != nil {
			return nil, fmt.Errorf("D1 placement of %s: %w", g, err)
		}
		if want := signOf(lon); rs != want {
			return nil, fmt.Errorf("D1 mismatch for %s: placement %q but longitude %.4f is in %q",
				g, d1[g], lon, rashiName(want))
		}
		c.rasiSign[g] = rs

		// D9 validation
		ns, err := rashiIndex(d9[g])
		if err != nil {
			return nil, fmt.Errorf("D9 placement of %s: %w", g, err)
		}
		navFromLon, err := utils.CalculateNavamsha(lon)
		if err != nil {
			return nil, fmt.Errorf("navamsha of %s: %w", g, err)
		}
		if string(navFromLon) != d9[g] {
			return nil, fmt.Errorf("D9 mismatch for %s: placement %q but longitude gives %q",
				g, d9[g], string(navFromLon))
		}
		c.navamsaSign[g] = ns
	}

	asc, err := rashiIndex(in.Chart.Ascendant)
	if err != nil {
		return nil, fmt.Errorf("ascendant: %w", err)
	}
	c.ascSign = asc
	c.ascendantLong = in.AscendantLong
	c.chandraWaxing = in.Chart.ChandraWaxing
	c.sunriseJD = in.SunriseJD
	c.sunsetJD = in.SunsetJD
	c.isDayBirth = in.IsDayBirth
	c.mclong = in.MsLong
	c.birthDateTime = julianDayUT(in.BirthDateTime)
	c.reqLatitude = in.Latitude
	c.reqLongitude = in.Longitude
	c.meshaJD = in.MeshaJD
	c.solarMonthStartJD = in.SolarMonthStartJD

	return c, nil
}

// ---- helpers ----------------------------------------------------------------

func rashiName(idx int) string { return constants.RaashiOrder[idx] }

func rashiIndex(name string) (int, error) {
	for i := range constants.RaashiOrder {
		if rashiName(i) == name {
			return i, nil
		}
	}
	return 0, fmt.Errorf("unknown raashi %q", name)
}

func signOf(lon float64) int {
	s := int(utils.Normalize360(lon) / 30)
	if s > 11 {
		s = 11
	}
	return s
}

func degreesInSign(lon float64) float64 {
	return math.Mod(utils.Normalize360(lon), 30)
}

func angularDistance(a, b float64) float64 {
	d := math.Mod(math.Abs(a-b), 360)
	if d > 180 {
		d = 360 - d
	}
	return d
}

func angularSeparation(a, b float64) float64 { return angularDistance(a, b) }

func isOddSign(sign int) bool { return sign%2 == 0 }

func isDayPlanet(g dto.Raashyadhipati) bool {
	return g == constants.Surya || g == constants.Guru || g == constants.Shukra
}

func isNightPlanet(g dto.Raashyadhipati) bool {
	return g == constants.Chandra || g == constants.Kuja || g == constants.Shani
}

func julianDayUT(t time.Time) float64 {
	utc := t.UTC()
	return 2440587.5 +
		float64(utc.Unix())/86400.0 +
		float64(utc.Nanosecond())/86400e9
}

func julianDayToTime(jd float64) time.Time {
	unixSeconds := (jd - 2440587.5) * 86400
	sec, nano := math.Modf(unixSeconds)
	return time.Unix(int64(sec), int64(nano*1e9)).UTC()
}
