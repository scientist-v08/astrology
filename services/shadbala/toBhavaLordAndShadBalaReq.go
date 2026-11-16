package shadbala

import (
	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
)

func ToBhavaLordAndShadBalaReq(houseReq models.HouseEffectsReqBody, res Result) models.BhavaLordAndShadBalaReqBody {
	order := []models.Raashyadhipati{
		constants.Surya,
		constants.Chandra,
		constants.Kuja,
		constants.Budha,
		constants.Guru,
		constants.Shukra,
		constants.Shani,
	}

	// Naisargika Bala in Virupas matching the exact graha order above:
	// Surya (60), Chandra (51.43), Kuja (17.14), Budha (25.71), Guru (34.29), Shukra (42.86), Shani (8.57)
	naisargikaBala := []float64{60.0, 51.43, 17.14, 25.71, 34.29, 42.86, 8.57}

	sthanaBala := make([]float32, len(order))
	kaalaBala := make([]float32, len(order))
	drigBala := make([]float32, len(order))
	shadBala := make([]float32, len(order))

	for i, graha := range order {
		sthana := res.Sthana[graha].Total
		dig := res.Dig.Values[graha]
		kaala := res.Kaala.Total[graha]
		chesta := res.Chesta.Values[graha]
		naisargika := naisargikaBala[i]
		drig := res.Drik.Values[graha]

		// Total Shadbala is the sum of all 6 strengths (Sthana, Dig, Kaala, Chesta, Naisargika, Drik)
		totalShadbala := sthana + dig + kaala + chesta + naisargika + drig

		sthanaBala[i] = float32(sthana)
		kaalaBala[i] = float32(kaala)
		drigBala[i] = float32(drig)
		shadBala[i] = float32(totalShadbala)
	}

	return models.BhavaLordAndShadBalaReqBody{
		Ascendant:        houseReq.Ascendant,
		SuryaPlacement:   houseReq.SuryaPlacement,
		BudhaPlacement:   houseReq.BudhaPlacement.Placement,
		ShukraPlacement:  houseReq.ShukraPlacement.Placement,
		ChandraPlacement: houseReq.ChandraPlacement,
		RahuPlacement:    houseReq.RahuPlacement,
		KetuPlacement:    houseReq.KetuPlacement,
		KujaPlacement:    houseReq.KujaPlacement.Placement,
		GuruPlacement:    houseReq.GuruPlacement.Placement,
		ShaniPlacement:   houseReq.ShaniPlacement.Placement,
		SthanaBala:       sthanaBala,
		KaalaBala:        kaalaBala,
		DrigBala:         drigBala,
		ShadBala:         shadBala,
	}
}