package controller

import (
	"slices"

	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/dto"
	"github.com/scientist-v08/bphs/services/houses"
	"github.com/scientist-v08/bphs/utils"
)

type HouseEffectsResponse struct {
	AllHouseEffects    []dto.HouseEffectsWithCommentory `json:"allHouseEffects"`
	ZeroProportion     int8                                `json:"zeroProportion"`
	Proportions        []float64                           `json:"proportions"`
	GeneralEffects     []string                            `json:"generalEffects"`
	KarakamshaResults  []IndexedString                     `json:"karakamshaResults"`
	Horoscope          *dto.HouseEffectsReqBody         `json:"horoscope"`
}

func isGrahaAfflicted(
    budha string,
    surya string,
    shani string,
    kuja string,
    chandra string,
    chandraWaxing bool,
    isCombust bool,
    shaniAspects *[]dto.AllRaashis,
    kujaAspects *[]dto.AllRaashis,
) bool {

    // Conjunctions
    isConjunctSurya := budha == surya

    isConjunctShani := budha == shani

    isConjunctKuja := budha == kuja

    isConjunctWaningMoon := budha == chandra && !chandraWaxing

    // Malefic aspects
    isAspectedByShani := slices.Contains(*shaniAspects,dto.AllRaashis(budha))

    isAspectedByKuja := slices.Contains(*kujaAspects,dto.AllRaashis(budha))

    // Opposition by a waning Moon
    isOpposedByWaningMoon := !chandraWaxing && constants.OppositeAspectsStore[dto.AllRaashis(chandra)] == dto.AllRaashis(budha)

    return isConjunctSurya ||
        isConjunctShani ||
        isConjunctKuja ||
        isConjunctWaningMoon ||
        isAspectedByShani ||
        isAspectedByKuja ||
        isOpposedByWaningMoon ||
        isCombust
}

func isBudhaAfflicted(
    req *dto.HouseEffectsReqBody,
    shaniAspects *[]dto.AllRaashis,
    kujaAspects *[]dto.AllRaashis,
) bool {

    return isGrahaAfflicted(
        req.BudhaPlacement.Placement,
        req.SuryaPlacement,
        req.ShaniPlacement.Placement,
        req.KujaPlacement.Placement,
        req.ChandraPlacement,
        req.ChandraWaxing,
        req.BudhaPlacement.Combustion,
        shaniAspects,
        kujaAspects,
    )
}

func isNavBudhaAfflicted(
    req *dto.HouseEffectsReqBody,
    navShaniAspects *[]dto.AllRaashis,
    navKujaAspects *[]dto.AllRaashis,
) bool {

    return isGrahaAfflicted(
        req.NavBudhaPlacement,
        req.NavSuryaPlacement,
        req.NavShaniPlacement,
        req.NavKujaPlacement,
        req.NavChandraPlacement,
        req.ChandraWaxing,
        false, // Navamsha combustion is not separately supplied.
        navShaniAspects,
        navKujaAspects,
    )
}

func calculateHouseEffects(req *dto.HouseEffectsReqBody) HouseEffectsResponse {
	// 1. Compute the different house numbers with respect to the ascendant and the navamsha ascedant and karakamsha.
	housePlacements := map[string]int16{
		"Surya": utils.GetHouse(req.Ascendant, req.SuryaPlacement),
		"Budha": utils.GetHouse(req.Ascendant, req.BudhaPlacement.Placement),   
		"Shukra": utils.GetHouse(req.Ascendant, req.ShukraPlacement.Placement),  
		"Chandra": utils.GetHouse(req.Ascendant, req.ChandraPlacement), 
		"RahuPlacement": utils.GetHouse(req.Ascendant, req.RahuPlacement),    
		"KetuPlacement": utils.GetHouse(req.Ascendant, req.KetuPlacement),    
		"Kuja": utils.GetHouse(req.Ascendant, req.KujaPlacement.Placement),    
		"Guru": utils.GetHouse(req.Ascendant, req.GuruPlacement.Placement),    
		"Shani": utils.GetHouse(req.Ascendant, req.ShaniPlacement.Placement),   
		"NavSurya": utils.GetHouse(req.NavAscendant, req.NavSuryaPlacement),
		"NavBudha": utils.GetHouse(req.NavAscendant, req.NavBudhaPlacement),
		"NavShukra": utils.GetHouse(req.NavAscendant, req.NavShukraPlacement),
		"NavChandra": utils.GetHouse(req.NavAscendant, req.NavChandraPlacement),
		"NavRahuPlacement": utils.GetHouse(req.NavAscendant, req.NavRahuPlacement),
		"NavKetuPlacement": utils.GetHouse(req.NavAscendant, req.NavKetuPlacement),
		"NavKuja": utils.GetHouse(req.NavAscendant, req.NavKujaPlacement),
		"NavGuru": utils.GetHouse(req.NavAscendant, req.NavGuruPlacement),
		"NavShani": utils.GetHouse(req.NavAscendant, req.NavShaniPlacement),
	}
	karakamshaReqBody := buildKarakamshaReqBody(*req)
	karakamshaHousePlacements := map[string]int16{
		"Surya": utils.GetHouse(karakamshaReqBody.Karakamsha, karakamshaReqBody.Surya),
		"Budha": utils.GetHouse(karakamshaReqBody.Karakamsha, karakamshaReqBody.Budha),
		"Shukra": utils.GetHouse(karakamshaReqBody.Karakamsha, karakamshaReqBody.Shukra),
		"Chandra": utils.GetHouse(karakamshaReqBody.Karakamsha, karakamshaReqBody.Chandra),
		"Rahu": utils.GetHouse(karakamshaReqBody.Karakamsha, karakamshaReqBody.Rahu),
		"Ketu": utils.GetHouse(karakamshaReqBody.Karakamsha, karakamshaReqBody.Ketu),
		"Kuja": utils.GetHouse(karakamshaReqBody.Karakamsha, karakamshaReqBody.Kuja),
		"Guru": utils.GetHouse(karakamshaReqBody.Karakamsha, karakamshaReqBody.Guru),
		"Shani": utils.GetHouse(karakamshaReqBody.Karakamsha, karakamshaReqBody.Shani),
		"Lagna": utils.GetHouse(karakamshaReqBody.Karakamsha, karakamshaReqBody.Ascendant),
	}

	guruAspects := utils.GetGuruAspectRaashis(dto.AllRaashis(req.GuruPlacement.Placement))
	shaniAspects := utils.GetShaniAspectRaashis(dto.AllRaashis(req.ShaniPlacement.Placement))
	navShaniAspects := utils.GetShaniAspectRaashis(dto.AllRaashis(req.NavShaniPlacement))
	kujaAspects := utils.GetKujaAspectRaashis(dto.AllRaashis(req.KujaPlacement.Placement))
	navKujaAspects := utils.GetKujaAspectRaashis(dto.AllRaashis(req.NavKujaPlacement))
	// Calculate Mercury affliction.
	req.BudhaUnafflicted = !isBudhaAfflicted(req,&shaniAspects,&kujaAspects)
	req.NavBudhaAfflicted = isNavBudhaAfflicted(req,&navShaniAspects,&navKujaAspects)

	// 2. Determine which proportion array to use
    proportions := make([]int8, 12)
    asc := req.Ascendant
    specialKey := ""
    if asc == "Makara" || asc == "Dhanassu" {
        switch asc {
		case "Makara":
            if req.FirstHalf {
                specialKey = "Makara1"
            } else {
                specialKey = "Makara2"
            }
        case "Dhanassu":
            if req.FirstHalf {
                specialKey = "Dhanassu1"
            } else {
                specialKey = "Dhanassu2"
            }
        }

        if specialKey != "" {
			copy(proportions, constants.SpecialHouseEffectsProportionStore[specialKey])
        }
    } else {
		copy(proportions, constants.HouseeffectProprtionStore[dto.AllRaashis(asc)])
    }

	// 3. Prepare results
	var result []dto.HouseEffectsWithCommentory
	for i := range int8(12) {
        houseNum := i + 1
        proportion := proportions[i]

		if slices.Contains([]int8{60,50,40}, proportion) {

			comment := getShadBalaStyleComment(proportion)
	
			item := dto.HouseEffectsWithCommentory{
				HouseNumber: houseNum,
				Theme:       constants.Themes[i],
				Comments:    comment,
				Effects:     getBhavaEffect(houseNum, req, &housePlacements, &guruAspects, &shaniAspects, &kujaAspects),
			}
	
			result = append(result, item)
		}

    }

	var zeroProportionHouse = int8(0)
	for i, item := range proportions {
		if item == 0 {
			zeroProportionHouse = int8(i + 1)
		}
	}

	otherEffects := houses.GeneralEffects(req, &housePlacements, &guruAspects, &shaniAspects, &kujaAspects)

	var sum int = 0
    for _, v := range proportions {
        sum += int(v)
    }

	percentages := make([]float64, 12)
    for i, val := range proportions {
        percentages[i] = (float64(val) / float64(sum)) * 100
    }

	// 4. Proceed with the response
	return HouseEffectsResponse{
        AllHouseEffects:   result,
        ZeroProportion:    zeroProportionHouse,
        Proportions:       percentages,
        GeneralEffects:    otherEffects,
        KarakamshaResults: ToIndexedStrings(houses.KarakamshaEffectsService(&karakamshaHousePlacements, &karakamshaReqBody)),
        Horoscope:         req,
    }
}