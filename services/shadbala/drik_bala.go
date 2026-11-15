package shadbala

import (
	"math"

	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
)

// aspectStrength returns the absolute aspect strength (0–60)
// matching the key points given by Jaya Sekhar.
func aspectStrength(delta float64) float64 {
    d := math.Mod(delta, 360)
    if d < 0 {
        d += 360
    }

    switch {
    case d <= 30:
        return 0

    case d <= 75:
        return (d - 30) * (30.0 / 45.0)

    case d <= 90:
        return 30 + (d-75)

    case d <= 120:
    	return 45 - (d-90)*(25.0/30.0)

	case d <= 150:
    	return 20 * (150-d) / 30.0

    case d <= 165:
        return (d - 150) * 2

    case d <= 172.5:
        return 30 + (d-165)*2

    case d <= 180:
        return 45 + (d-172.5)*2

    case d <= 210:
        return 60 - (d-180)*0.5

    case d <= 240:
        return 45 - (d-210)*0.5

    case d <= 300:
        return 30 - (d-240)*0.5

    default:
        return 0
    }
}

// isNaturalMalefic decides the sign of the aspect energy.
func isNaturalMalefic(p models.Raashyadhipati, waxingMoon bool) bool {
	switch p {
	case constants.Surya, constants.Kuja, constants.Shani:
		return true
	case constants.Chandra:
		return !waxingMoon // waning Moon is malefic
	case constants.Budha:
		// Classical: unafflicted Mercury is benefic.
		// For pure Drik we treat it as benefic unless you later pass combustion/affliction.
		return false
	default: // Guru, Shukra
		return false
	}
}

// calculateDrikBala computes net aspect energy received by each planet.
func calculateDrikBala(ctx *chartContext) DrikBalaResult {
	res := DrikBalaResult{
		Values: make(map[models.Raashyadhipati]float64),
	}

	planets := []models.Raashyadhipati{
		constants.Surya, constants.Chandra, constants.Kuja,
		constants.Budha, constants.Guru, constants.Shukra, constants.Shani,
	}

	for _, receiver := range planets {
		var net float64
		recLon := ctx.longitude[receiver]

		for _, aspector := range planets {
			if aspector == receiver {
				continue
			}
			aspLon := ctx.longitude[aspector]
			strength := aspectStrength(recLon - aspLon)

			if isNaturalMalefic(aspector, ctx.chandraWaxing) {
				strength = -strength
			}
			net += strength
		}
		res.Values[receiver] = net
	}

	return res
}

type DrikBalaResult struct {
	Values map[models.Raashyadhipati]float64 `json:"values"` // can be negative
}