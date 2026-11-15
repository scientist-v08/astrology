package shadbala

import (
	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
	"github.com/scientist-v08/bphs/utils"
)

// digBalaForPlanet calculates Dig Bala for one planet using the classical
// midpoint method preferred by the book.
func digBalaForPlanet(
	planetLon float64,
	ascLon float64,
	mcLon float64,
	planet models.Raashyadhipati,
) float64 {

	// Cardinal midpoints (equal-house style from Ascendant)
	// 1st = Asc, 4th = Asc+90, 7th = Asc+180, 10th = Asc+270
	mid1 := utils.Normalize360(ascLon)           // East  – Lagna
	mid4 := utils.Normalize360(ascLon + 90)      // South
	mid7 := utils.Normalize360(ascLon + 180)     // West
	mid10 := utils.Normalize360(ascLon + 270)    // North

	var weakPoint float64

	switch planet {
	case constants.Surya, constants.Kuja: // strongest at 10th → weak at 4th
		weakPoint = mid4
	case constants.Chandra, constants.Shukra: // strongest at 4th → weak at 10th
		weakPoint = mid10
	case constants.Guru, constants.Budha: // strongest at 1st → weak at 7th
		weakPoint = mid7
	case constants.Shani: // strongest at 7th → weak at 1st
		weakPoint = mid1
	default:
		return 0 // Rahu/Ketu not calculated in classical Dig Bala
	}

	// Arc from weak point to planet, take the shorter way (≤180)
	diff := utils.Normalize360(planetLon - weakPoint)
	if diff > 180 {
		diff = 360 - diff
	}

	// Dig Bala = arc / 3  → 0 … 60 Virupas
	return diff / 3.0
}

// calculateDigBala returns Dig Bala for all seven planets.
func calculateDigBala(ctx *chartContext) DigBalaResult {
	res := DigBalaResult{
		Values: make(map[models.Raashyadhipati]float64),
	}

	planets := []models.Raashyadhipati{
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
	Values map[models.Raashyadhipati]float64 `json:"values"` // Virupas
}