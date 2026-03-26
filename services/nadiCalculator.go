package services

import (
	"slices"

	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
)

func NadiCalculatorFunc(groomNakshatra string, brideNakshatra string) models.KutaTypeRes {
	nadiExceptionNakshatras := []string{
		"Rohini",
		"Ardra",
		"Magha",
		"Hasta",
		"Vishakha",
		"Shravana",
		"Uttara Bhadrapada",
		"Revati",
	}

	// Nadi mediocre/partial exception nakshatras
	nadiMediocoreExceptionNakshatras := []string{
		"Ashwini",
		"Kritika",
		"Mrigashira",
		"Punarvasu",
		"Pushya",
		"Purva Phalguni",
		"Uttara Phalguni",
		"Chitra",
		"Anuradha",
		"Purva Ashadha",
		"Uttara Ashadha",
	}

	groomNadi := constants.NakshatraNadiMap[models.Nakshatra(groomNakshatra)]
	brideNadi := constants.NakshatraNadiMap[models.Nakshatra(brideNakshatra)]

	var naadiScore int8
	var naadiComment string

	if groomNadi == brideNadi {
		naadiScore = 0
		naadiComment = `Nadi dosha: This may cause either the husband or the wife or the child to develop some genetic disease 
        although modern medicine can help to mitigate them I am of the opinion prevention is better than cure.`

		if slices.Contains(nadiExceptionNakshatras, groomNakshatra) {
			naadiScore = 7
			naadiComment = `Nadi dosha: Exists. But astrologers have observed that this pair is exempt from the rule. 
            But ensure that the pada of the boy preceeds the pada of the girl if all the padas of a nakshatra is in the same
            Raashi else do viceversa.`
		}
		if slices.Contains(nadiMediocoreExceptionNakshatras, groomNakshatra) {
			naadiScore = 4
			naadiComment = `Nadi dosha: Exists. But astrologers have observed that this pair is partially exempt from the rule. 
            But ensure that the pada of the boy preceeds the pada of the girl if all the padas of a nakshatra is in the same
            Raashi else do viceversa.`
		}
	} else {
		naadiScore = 8
		naadiComment = "No Naadi dosha"
	}

	return models.KutaTypeRes{
		Index: 7,
		Score: float32(naadiScore),
		Comments: naadiComment,
	}
}