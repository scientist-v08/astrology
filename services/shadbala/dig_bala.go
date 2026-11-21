package shadbala

import (
	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/dto"
	"github.com/scientist-v08/bphs/utils"
)

// digBalaForPlanet calculates Dig Bala for one planet using the classical
// midpoint method preferred by the book.
func digBalaForPlanet(
	planetLon float64,
	ascLon float64,
	mcLon float64,
	planet dto.Raashyadhipati,
) float64 {

	// Cardinal points using actual Ascendant (1st/7th) and MC (10th/4th)
	mid1 := utils.Normalize360(ascLon)        // 1st cusp (Ascendant / East)
	mid7 := utils.Normalize360(ascLon + 180)  // 7th cusp (Descendant / West)
	mid10 := utils.Normalize360(mcLon)        // 10th cusp (MC / South)
	mid4 := utils.Normalize360(mcLon + 180)   // 4th cusp (IC / North)

	var weakPoint float64

	switch planet {
	case constants.Surya, constants.Kuja: // strongest at 10th (MC) → weak at 4th (IC)
		weakPoint = mid4
	case constants.Chandra, constants.Shukra: // strongest at 4th (IC) → weak at 10th (MC)
		weakPoint = mid10
	case constants.Guru, constants.Budha: // strongest at 1st (Asc) → weak at 7th (Desc)
		weakPoint = mid7
	case constants.Shani: // strongest at 7th (Desc) → weak at 1st (Asc)
		weakPoint = mid1
	default:
		return 0 // Rahu/Ketu not calculated in classical Dig Bala
	}

	// Arc from weak point to planet, take the shorter way (≤180)
	diff := utils.Normalize360(planetLon - weakPoint)
	if diff > 180 {
		diff = 360 - diff
	}

	// Dig Bala = arc / 3  → 0 … 60 Virupas (Shashtiamsas)
	return diff / 3.0
}

// calculateDigBala returns Dig Bala for all seven planets.
func calculateDigBala(ctx *chartContext) DigBalaResult {
	res := DigBalaResult{
		Values: make(map[dto.Raashyadhipati]float64),
	}

	planets := []dto.Raashyadhipati{
		constants.Surya, constants.Chandra, constants.Kuja,
		constants.Budha, constants.Guru, constants.Shukra, constants.Shani,
	}

	for _, p := range planets {
		lon := ctx.longitude[p]
		res.Values[p] = digBalaForPlanet(lon, ctx.ascendantLong, ctx.mclong, p)
	}

	return res
}

type DigBalaResult struct {
	Values map[dto.Raashyadhipati]float64 `json:"values"` // Virupas
}