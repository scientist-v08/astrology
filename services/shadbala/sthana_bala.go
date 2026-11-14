package shadbala

import (
	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
)

// Sthana Bala = Uchcha + Saptavargaja + Ojayugma-rasyamsa + Kendradi + Drekkana
// All values are in Virupas (60 Virupas = 1 Rupa).

// Minimum Sthana Bala required (Virupas) - Parashara, as tabulated in the book.
var sthanaMinimum = map[models.Raashyadhipati]float64{
	constants.Surya:   165,
	constants.Chandra: 133,
	constants.Kuja:    96,
	constants.Budha:   165,
	constants.Guru:    165,
	constants.Shukra:  133,
	constants.Shani:   96,
}

// Debilitation points (exaltation point + 180), sidereal degrees.
// Exaltation: Sun 10 Ari, Moon 3 Tau, Mars 28 Cap, Mer 15 Vir, Jup 5 Can, Ven 27 Pis, Sat 20 Lib.
var debilitationPoint = map[models.Raashyadhipati]float64{
	constants.Surya:   190,
	constants.Chandra: 213,
	constants.Kuja:    118,
	constants.Budha:   345,
	constants.Guru:    275,
	constants.Shukra:  177,
	constants.Shani:   20,
}

// SaptavargaEntry explains how one varga contributed to Saptavarga Bala.
type SaptavargaEntry struct {
	Varga   string  `json:"varga"`
	Sign    int     `json:"sign"` // index in constants.RaashiOrder
	Dignity Dignity `json:"dignity"`
	Virupas float64 `json:"virupas"`
}

// SthanaBala is the full breakdown for one graha.
type SthanaBala struct {
	Uchcha     float64 `json:"uchcha"`
	Saptavarga float64 `json:"saptavarga"`
	OjaYugma   float64 `json:"ojaYugma"`
	Kendra     float64 `json:"kendra"`
	Drekkana   float64 `json:"drekkana"`

	Total      float64 `json:"total"`      // Virupas
	TotalRupas float64 `json:"totalRupas"` // Total / 60

	Required     float64 `json:"required"`     // minimum per Parashara
	Ratio        float64 `json:"ratio"`        // Total / Required (before Drik Bala)
	MeetsMinimum bool    `json:"meetsMinimum"` // Total >= Required (before Drik Bala)

	SaptavargaDetail []SaptavargaEntry `json:"saptavargaDetail"`
}

// SthanaBalaResult is keyed by graha.
type SthanaBalaResult map[models.Raashyadhipati]SthanaBala

// calculateSthanaBala computes Sthana Bala for the seven grahas.
//
// Note: Jaya Sekhar adds Drik Bala to Sthana Bala before ranking/ratio
// ("effective Sthana Bala"). Ratio/MeetsMinimum here are therefore on the
// raw value; once Drik Bala exists, apply it in the orchestrator.
func calculateSthanaBala(ctx *chartContext) SthanaBalaResult {
	result := make(SthanaBalaResult, len(shadbalaGrahas))

	for _, g := range shadbalaGrahas {
		lon := ctx.longitude[g]

		sb := SthanaBala{
			Uchcha:   uchchaBala(g, lon),
			OjaYugma: ojaYugmaBala(g, ctx.rasiSign[g], ctx.navamsaSign[g]),
			Kendra:   kendradiBala(ctx.rasiSign[g], ctx.ascSign),
			Drekkana: drekkanaBala(g, lon),
		}
		sb.SaptavargaDetail, sb.Saptavarga = saptavargajaBala(g, ctx)

		sb.Total = sb.Uchcha + sb.Saptavarga + sb.OjaYugma + sb.Kendra + sb.Drekkana
		sb.TotalRupas = sb.Total / 60
		sb.Required = sthanaMinimum[g]
		sb.Ratio = sb.Total / sb.Required
		sb.MeetsMinimum = sb.Total >= sb.Required

		result[g] = sb
	}
	return result
}

// uchchaBala: distance from the debilitation point (shortest arc, max 180) / 3.
// 0 at debilitation, 60 at exaltation.
func uchchaBala(g models.Raashyadhipati, lon float64) float64 {
	return angularDistance(lon, debilitationPoint[g]) / 3
}

// ojaYugmaBala: 15 Virupas for the D1 placement and 15 for the D9 placement
// in the "right" polarity. Moon and Venus like even (yugma) signs; the rest
// (Sun, Mars, Jupiter, Mercury, Saturn) like odd (oja) signs.
func ojaYugmaBala(g models.Raashyadhipati, rasiSign, navamsaSign int) float64 {
	female := g == constants.Chandra || g == constants.Shukra
	total := 0.0
	for _, sign := range []int{rasiSign, navamsaSign} {
		if isOddSign(sign) != female { // odd & male, or even & female
			total += 15
		}
	}
	return total
}

// kendradiBala: 60 in kendras (1,4,7,10), 30 in panaparas (2,5,8,11),
// 15 in apoklimas (3,6,9,12), counted from the lagna sign.
func kendradiBala(planetSign, ascSign int) float64 {
	house := ((planetSign-ascSign)%12+12)%12 + 1
	switch house {
	case 1, 4, 7, 10:
		return 60
	case 2, 5, 8, 11:
		return 30
	default:
		return 15
	}
}

// drekkanaBala: 15 Virupas when the planet occupies its
// favoured Drekkana.
// Sun/Mars/Jupiter: 1st Drekkana
// Mercury/Saturn:   2nd Drekkana
// Moon/Venus:       3rd Drekkana
func drekkanaBala(g models.Raashyadhipati, lon float64) float64 {
    decanate := int(degreesInSign(lon) / 10)
	if decanate > 2 {
		decanate = 2
	}
	want := map[models.Raashyadhipati]int{
		constants.Surya: 0, constants.Kuja: 0, constants.Guru: 0,
		constants.Budha: 1, constants.Shani: 1,
		constants.Chandra: 2, constants.Shukra: 2,
	}[g]
	if decanate == want {
		return 15
	}
	return 0
}
