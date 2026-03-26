package services

import (
	"slices"

	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
)

func BhakutaCalculatorFunc(groomChandra string, brideChandra string) models.KutaTypeRes {
	groomChandraRuler := constants.RaashyadhipatiMapStore[groomChandra]
	brideChandraRuler := constants.RaashyadhipatiMapStore[brideChandra]
	brideRaashiFromGroomRaashi := constants.NumericalMappingsStore[groomChandra][models.AllRaashis(brideChandra)]
	isGroomAnEvenRaashi := slices.Contains([]string{"Vrushabha", "Karkataka", "Kanya", "Vruschika", "Makara", "Meena"}, groomChandra)
	var loveScore int8
	var comments string

	switch brideRaashiFromGroomRaashi {
	case 1, 7:
		loveScore = 7
		comments = "The couple will fall in love with each other and be happy"
	case 2:
		loveScore = 0
		comments = "The couple are likely to become enemies"
		if groomChandraRuler == brideChandraRuler {
			loveScore = 2
			comments = "The couple are likely to be neutral to each other (eventually)"
		}
	case 3:
		loveScore = 4
		comments = "The longer the couple stay together the greater the happiness the couple obtain"
	case 4:
		loveScore = 4
		comments = "This combination will increase prosperity"
	case 5:
		loveScore = 0
		comments = "Initially the couple will love each other but it is often seen that this combination causes problem in progeny and the couple tend to gradually hate each other"
	case 6, 8:
		loveScore = 0
		comments = "The couple are likely to become enemies"
		if groomChandraRuler == brideChandraRuler {
			loveScore = 2
			comments = "The couple are likely to be neutral to each other (eventually)"
		}
	case 9:
		loveScore = 0
		comments = "Initially the couple will love each other but it is often seen that this combination causes problem in the Dharma followed by the husband and the couple tend to gradually hate each other (Like Bill Gates and Mellinda Gates)"
	case 10:
		loveScore = 4
		comments = "This combination is likely to cause the deterioration of prosperity"
	case 11:
		loveScore = 4
		comments = "This combination is likely to cause the deterioration of happiness"
	case 12:
		loveScore = 0
		comments = "The couple are likely to become enemies"
		if groomChandraRuler == brideChandraRuler {
			loveScore = 2
			comments = "The couple are likely to be neutral to each other (eventually)"
		}
		if isGroomAnEvenRaashi {
			loveScore = 2
			comments = "The couple are likely to be neutral to each other (eventually)"
		}
	}
	return models.KutaTypeRes{
		Index: 6,
		Score: float32(loveScore),
		Comments: comments,
	}
}