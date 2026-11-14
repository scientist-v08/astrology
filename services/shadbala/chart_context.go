package shadbala

import (
	"fmt"
	"math"

	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
	"github.com/scientist-v08/bphs/utils"
)

// The seven grahas that take part in Shad Bala (Rahu/Ketu are excluded).
var shadbalaGrahas = []models.Raashyadhipati{
	constants.Surya, constants.Chandra, constants.Kuja, constants.Budha,
	constants.Guru, constants.Shukra, constants.Shani,
}

// chartContext is the single, validated view of the chart that every bala
// reads from. Sign indices follow constants.RaashiOrder (0 = first raashi, Aries).
type chartContext struct {
	longitude   map[models.Raashyadhipati]float64 // sidereal, normalised 0..360
	rasiSign    map[models.Raashyadhipati]int     // D1 sign index (from request body)
	navamsaSign map[models.Raashyadhipati]int     // D9 sign index (from request body)
	ascSign     int               // lagna sign index (from request body)
}

// d1Placements / d9Placements pull the placements out of the request body.
func d1Placements(b models.HouseEffectsReqBody) map[models.Raashyadhipati]string {
	return map[models.Raashyadhipati]string{
		constants.Surya:   b.SuryaPlacement,
		constants.Chandra: b.ChandraPlacement,
		constants.Kuja:    b.KujaPlacement.Placement,
		constants.Budha:   b.BudhaPlacement.Placement,
		constants.Guru:    b.GuruPlacement.Placement,
		constants.Shukra:  b.ShukraPlacement.Placement,
		constants.Shani:   b.ShaniPlacement.Placement,
	}
}

func d9Placements(b models.HouseEffectsReqBody) map[models.Raashyadhipati]string {
	return map[models.Raashyadhipati]string{
		constants.Surya:   b.NavSuryaPlacement,
		constants.Chandra: b.NavChandraPlacement,
		constants.Kuja:    b.NavKujaPlacement,
		constants.Budha:   b.NavBudhaPlacement,
		constants.Guru:    b.NavGuruPlacement,
		constants.Shukra:  b.NavShukraPlacement,
		constants.Shani:   b.NavShaniPlacement,
	}
}

// newChartContext resolves placements to sign indices and cross-checks them
// against the longitudes so the two data sources can never silently disagree.
func newChartContext(in Input) (*chartContext, error) {
	c := &chartContext{
		longitude:   make(map[models.Raashyadhipati]float64, len(shadbalaGrahas)),
		rasiSign:    make(map[models.Raashyadhipati]int, len(shadbalaGrahas)),
		navamsaSign: make(map[models.Raashyadhipati]int, len(shadbalaGrahas)),
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

		// D1: body placement must match the longitude.
		rs, err := rashiIndex(d1[g])
		if err != nil {
			return nil, fmt.Errorf("D1 placement of %s: %w", g, err)
		}
		if want := signOf(lon); rs != want {
			return nil, fmt.Errorf("D1 mismatch for %s: placement %q but longitude %.4f is in %q",
				g, d1[g], lon, rashiName(want))
		}
		c.rasiSign[g] = rs

		// D9: body placement must match utils.CalculateNavamsha.
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

	return c, nil
}

// ---- raashi <-> index, backed by constants.RaashiOrder ---------------------
func rashiName(idx int) string { 
	return constants.RaashiOrder[idx]
}

func rashiIndex(name string) (int, error) {
	for i := range constants.RaashiOrder {
		if rashiName(i) == name {
			return i, nil
		}
	}
	return 0, fmt.Errorf("unknown raashi %q", name)
}

// ---- small numeric helpers (same logic as int(longitude/30) elsewhere) ------

// signOf returns the sign index of a longitude: int(lon/30), 0..11.
func signOf(lon float64) int {
	s := int(utils.Normalize360(lon) / 30)
	if s > 11 {
		s = 11
	}
	return s
}

// degreesInSign returns 0..<30 (math.Mod(lon, 30)).
func degreesInSign(lon float64) float64 {
	return math.Mod(utils.Normalize360(lon), 30)
}

// angularDistance returns the shortest arc between two longitudes (0..180).
func angularDistance(a, b float64) float64 {
	d := math.Mod(math.Abs(a-b), 360)
	if d > 180 {
		d = 360 - d
	}
	return d
}

// isOddSign: Aries(0), Gemini(2), Leo(4)... are odd (masculine) signs.
func isOddSign(sign int) bool { 
	return sign%2 == 0 
}
