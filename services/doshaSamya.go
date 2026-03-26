package services

import (
	"github.com/scientist-v08/bphs/models"
	"github.com/scientist-v08/bphs/utils"
)

type PersonPlacements struct {
	Ascendant       string
	Chandra         string
	Shukra          string
	Surya           string
	Kuja            string
	Shani           string
	Rahu            string
	Ketu            string
	Guru            string
}

func contains(slice []int16, val int16) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func computeDoshaForPerson(p PersonPlacements) []models.DoshaSamyaRes {
	refs := map[string]string{
		"Lagna": p.Ascendant,
		"Moon":  p.Chandra,
		"Venus": p.Shukra,
	}

	malefics := []string{"Surya", "Kuja", "Shani", "Rahu"}

	getPlacement := func(graha string) string {
		switch graha {
		case "Surya":
			return p.Surya
		case "Kuja":
			return p.Kuja
		case "Shani":
			return p.Shani
		case "Rahu":
			return p.Rahu
		default:
			return ""
		}
	}

	var res []models.DoshaSamyaRes

	doshaHouses := []int16{1, 2, 4, 7, 8, 12}

	for refName, refSign := range refs {
		// Compute houses for malefics and benefics relative to this reference
		var ascendantBase float32
		switch refName{
		case "Lagna":
			ascendantBase = 1
		case "Moon":
			ascendantBase = 0.5
		case "Venus":
			ascendantBase = 0.25
		}
		houses := make(map[string]int16)
		for _, graha := range malefics {
			placement := getPlacement(graha)
			if placement != "" {
				houses[graha] = utils.GetHouse(refSign, placement)
			}
		}

		for _, mal := range malefics {
			house, ok := houses[mal]
			if !ok || !contains(doshaHouses, house) {
				continue
			}

			var base float32
			switch mal {
			case "Surya":
				base = 0.5
			case "Shani":
				base = 1.0
			case "Kuja":
				if house == 7 || house == 8 {
					base = 2.0
				} else {
					base = 1.0
				}
			case "Rahu":
				if house == 4 || house == 7 || house == 8 {
					base = 2.0
				} else {
					base = 1.0
				}
			}

			if base > 0 {
				res = append(res, models.DoshaSamyaRes{
					AscendantConsidered: refName,
					GrahaInQuestion:     mal,
					DoshaScore:          base * ascendantBase,
				})
			}
		}
	}

	return res
}

func DoshaSamyaFunc(req *models.PairingReqBody) ([]models.DoshaSamyaRes, []models.DoshaSamyaRes) {
	groomPlacements := PersonPlacements{
		Ascendant: (*req).GroomAscendant,
		Chandra:   (*req).GroomChandraPlacement,
		Shukra:    (*req).GroomShukraPlacement,
		Surya:     (*req).GroomSuryaPlacement,
		Kuja:      (*req).GroomKujaPlacement,
		Shani:     (*req).GroomShaniPlacement,
		Rahu:      (*req).GroomRahuPlacement,
		Ketu:      (*req).GroomKetuPlacement,
		Guru:      (*req).GroomGuruPlacement,
	}

	bridePlacements := PersonPlacements{
		Ascendant: (*req).BrideAscendant,
		Chandra:   (*req).BrideChandraPlacement,
		Shukra:    (*req).BrideShukraPlacement,
		Surya:     (*req).BrideSuryaPlacement,
		Kuja:      (*req).BrideKujaPlacement,
		Shani:     (*req).BrideShaniPlacement,
		Rahu:      (*req).BrideRahuPlacement,
		Ketu:      (*req).BrideKetuPlacement,
		Guru:      (*req).BrideGuruPlacement,
	}

	groomRes := computeDoshaForPerson(groomPlacements)
	brideRes := computeDoshaForPerson(bridePlacements)

	return groomRes, brideRes
}