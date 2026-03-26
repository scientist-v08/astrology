package services

import (
	"slices"

	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
)

func VashyaCalculatorFunc(groomRaashi string, brideRaashi string, isGroomFirstHalf bool, isBrideFirstHalf bool) models.KutaTypeRes {
	var groomVaashya string
	var brideVaashya string
	if slices.Contains([]string{"Dhanassu", "Makara"}, groomRaashi) {
		switch groomRaashi{
			case "Makara":
            if isGroomFirstHalf {
                groomVaashya = constants.RaashiVaashyaMapStore["Makara1"]
            } else {
                groomVaashya = constants.RaashiVaashyaMapStore["Makara2"]
            }
        case "Dhanassu":
            if isGroomFirstHalf {
                groomVaashya = constants.RaashiVaashyaMapStore["Dhanassu1"]
            } else {
                groomVaashya = constants.RaashiVaashyaMapStore["Dhanassu2"]
            }
		}
	} else {
		groomVaashya = constants.RaashiVaashyaMapStore[groomRaashi]
	}
	if slices.Contains([]string{"Dhanassu", "Makara"}, brideRaashi) {
		switch brideRaashi{
			case "Makara":
            if isBrideFirstHalf {
                brideVaashya = constants.RaashiVaashyaMapStore["Makara1"]
            } else {
                brideVaashya = constants.RaashiVaashyaMapStore["Makara2"]
            }
        case "Dhanassu":
            if isBrideFirstHalf {
                brideVaashya = constants.RaashiVaashyaMapStore["Dhanassu1"]
            } else {
                brideVaashya = constants.RaashiVaashyaMapStore["Dhanassu2"]
            }
		}
	} else {
		brideVaashya = constants.RaashiVaashyaMapStore[brideRaashi]
	}
	vashyaScore := constants.VashyaCompatibility[groomVaashya][brideVaashya]
	var comments string
	switch vashyaScore {
      case 2:
        comments =
          "Both the man and the woman will have control over the other.";
      case 1:
        comments =
          "Both the man and the woman will have partial control over the other.";
      default:
        comments =
          "Both the man and the woman may not have control over the other.";
    }
	return models.KutaTypeRes{
		Index: 1,
		Score: float32(vashyaScore),
		Comments: comments,
	}
}