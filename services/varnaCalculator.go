package services

import (
	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
)

func VarnaCalculatorFunc(groomRaashi string, brideRaashi string) models.KutaTypeRes {
	varnaScore := constants.RaashiVarnaMapStore[groomRaashi] - constants.RaashiVarnaMapStore[brideRaashi]
	if varnaScore < 0 {
		varnaScore = varnaScore * -1
	}
	var additionalComments string
	if constants.RaashiVarnaMapStore[groomRaashi] < constants.RaashiVarnaMapStore[brideRaashi] {
		additionalComments = "But the man is likely to dominate the woman"
	} else {
		additionalComments = "But the woman is likely to dominate the man"
	}

	switch varnaScore{
	case 0,2:
		return models.KutaTypeRes{
			Index: 0,
			Score: 1,
			Comments: "Both the man and the woman will respect each others role in the relationship." + additionalComments,
		}
	default:
		return models.KutaTypeRes{
			Index: 0,
			Score: 0,
			Comments: "Both the man and the woman may not respect each others role in the relationship." + additionalComments,
		}
	}
}