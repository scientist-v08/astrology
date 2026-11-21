package shadbala

import (
	"math"

	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/dto"
	"github.com/scientist-v08/bphs/utils"
)

func aspectStrength(delta float64) float64 {
	d := math.Mod(delta, 360)
	if d < 0 {
		d += 360
	}
	switch {
	case d < 30:
		return 0
	case d < 60:
		return (d - 30) / 2
	case d < 90:
		return (d - 60) + 15
	case d < 120:
		return 45 - (d-90)/2
	case d < 150:
		return 30 - (d - 120)
	case d < 180:
		return (d - 150) * 2
	case d < 300:
		return (300 - d) / 2
	default:
		return 0
	}
}

func rasiOf(lon float64) int { return int(utils.Normalize360(lon) / 30) }

func isMaleficForDrik(p dto.Raashyadhipati, ctx *chartContext) bool {
	switch p {
	case constants.Surya, constants.Kuja, constants.Shani:
		return true
	case constants.Chandra:
		elong := utils.Normalize360(ctx.longitude[constants.Chandra] - ctx.longitude[constants.Surya])
		return elong < 90 || elong > 270
	case constants.Budha:
		r := rasiOf(ctx.longitude[constants.Budha])
		for _, q := range []dto.Raashyadhipati{constants.Surya, constants.Kuja, constants.Shani} {
			if rasiOf(ctx.longitude[q]) == r {
				return true
			}
		}
		return false
	}
	return false // Guru, Shukra
}

func calculateDrikBala(ctx *chartContext) DrikBalaResult {
	res := DrikBalaResult{Values: make(map[dto.Raashyadhipati]float64)}
	planets := []dto.Raashyadhipati{
		constants.Surya, constants.Chandra, constants.Kuja,
		constants.Budha, constants.Guru, constants.Shukra, constants.Shani,
	}
	for _, rec := range planets {
		var net float64
		for _, asp := range planets {
			if asp == rec {
				continue
			}
			s := aspectStrength(ctx.longitude[rec] - ctx.longitude[asp])
			if isMaleficForDrik(asp, ctx) {
				s = -s
			}
			net += s
		}
		res.Values[rec] = net / 4 // Drik Bala = ¼ of net aspect energy
	}
	return res
}

type DrikBalaResult struct {
	Values map[dto.Raashyadhipati]float64 `json:"values"` // can be negative
}