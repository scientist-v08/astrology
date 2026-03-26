package services

import (
	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
)

func GanaCalculatorHandler(groomNakshatra string, brideNakshatra string, groomChandra string, brideChandra string) models.KutaTypeRes {
	groomGana := constants.NakshatraGanaMap[models.Nakshatra(groomNakshatra)]
	brideGana := constants.NakshatraGanaMap[models.Nakshatra(brideNakshatra)]
	groomChandraRuler := constants.RaashyadhipatiMapStore[groomChandra]
	brideChandraRuler := constants.RaashyadhipatiMapStore[brideChandra]
	var groomComment string
	var brideComment string
	var ganaScore int8
	switch groomGana{
	case 1:
		groomComment = "The groom is born in Deva Gana"
	case 2:
		groomComment = "The groom is born under Manushya Gana"
	default:
		groomComment = "The groom is born under Raakshasa Gana"
	}
	switch brideGana{
	case 1:
		brideComment = "The bride is born in Deva Gana"
	case 2:
		brideComment = "The bride is born under Manushya Gana"
	default:
		brideComment = "The bride is born under Raakshasa Gana"
	}
	if groomGana == brideGana {
		ganaScore = 6
	} else if groomGana == 1 && brideGana == 2 {
		ganaScore = 5
	} else if brideGana == 1 && groomGana == 2 {
		ganaScore = 5
	} else {
		if brideGana == 3 {
			brideComment = brideComment + ". It is often noted that the bride should not be Rakshasa Gana in case of Gana mismatch"
		}
		ganaScore = 0
	}
	if ganaScore == 0 && groomChandraRuler == brideChandraRuler {
		ganaScore = 3
		brideComment = brideComment + ".This particular pair is exempt from the Gana mismatch rule."
	}
	return models.KutaTypeRes{
		Index: 5,
		Score: float32(ganaScore),
		Comments: groomComment + brideComment,
	}
}