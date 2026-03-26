package controller

import (
	"net/http"
	"slices"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
	"github.com/scientist-v08/bphs/services"
	"github.com/scientist-v08/bphs/utils"
)

func getShadBalaStyleComment(score int8) string {
    switch score {
    case 60:
        return "This house carries the highest strength. Excellent and prominent results can be expected from all matters signified by this house."
    case 50:
        return "Very good strength. Strong and mostly positive results are likely from this house."
    case 40:
        return "Good strength. Favourable results are generally obtained, though not at the highest level."
    case 30:
        return "Moderate strength. Results are mixed — some positive outcomes are possible with effort."
    case 20:
        return "Below average strength. Positive results are limited; challenges are more likely."
    case 10:
        return "Weak strength. Positive manifestations are difficult; problems tend to dominate."
    case 0:
        return "This house is very weak. Positive results are unlikely to manifest; negative tendencies and difficulties are almost certain."
    default:
        return "Unusual strength value — please verify calculation."
    }
}

func getBhavaLordEffect(house int8, lordPlacement int8) string {
	switch house {
	case 1:
		return services.FirstBhavaLordEffect(lordPlacement)
	case 2:
		return services.SecondBhavaLordEffect(lordPlacement)
	case 3:
		return services.ThirdBhavaLordEffect(lordPlacement)
	case 4:
		return services.FourthBhavaLordEffect(lordPlacement)
	case 5:
		return services.FifthBhavaLordEffect(lordPlacement)
	case 6:
		return services.SixthBhavaLordEffect(lordPlacement)
	case 7:
		return services.SeventhBhavaLordEffect(lordPlacement)
	case 8:
		return services.EighthBhavaLordEffect(lordPlacement)
	case 9:
		return services.NinthBhavaLordEffect(lordPlacement)
	case 10:
		return services.TenthBhavaLordEffect(lordPlacement)
	case 11:
		return services.EleventhBhavaLordEffect(lordPlacement)
	case 12:
		return services.TwelfthBhavaLordEffect(lordPlacement)
	default:
		return "" // invalid house
	}
}

func getBhavaEffect(
	house int8,
	req *models.HouseEffectsReqBody,
	housePlacements *map[string]int16,
	guruAspects *[]models.AllRaashis,
	shaniAspects *[]models.AllRaashis,
	kujaAspects *[]models.AllRaashis,
) []string {
   switch house {
	case 1:
		return services.FirstHouseEffects(req, housePlacements, guruAspects, shaniAspects, kujaAspects)
	case 2:
		return services.SecondHouseEffects(req, housePlacements, guruAspects, shaniAspects, kujaAspects)
	case 3:
		return services.ThirdHouseEffects(req, housePlacements, guruAspects, shaniAspects, kujaAspects)
	case 4:
		return services.FourthHouseEffects(req, housePlacements, guruAspects, shaniAspects, kujaAspects)
	case 5:
		return services.FifthHouseEffects(req, housePlacements, guruAspects, shaniAspects, kujaAspects)
	case 6:
		return services.SixthHouseEffects(req, housePlacements, guruAspects, shaniAspects, kujaAspects)
	case 7:
		return services.SeventhHouseEffects(req, housePlacements, guruAspects, shaniAspects, kujaAspects)
	case 8:
		return services.EighthHouseEffects(req, housePlacements, guruAspects, shaniAspects, kujaAspects)
	case 9:
		return services.NinthHouseEffects(req, housePlacements, guruAspects, shaniAspects, kujaAspects)
	case 10:
		return services.TenthHouseEffects(req, housePlacements, guruAspects, shaniAspects, kujaAspects)
	case 11:
		return services.EleventHouseEffects(req, housePlacements, guruAspects, shaniAspects, kujaAspects)
	case 12:
		return services.TwelfthHouseEffects(req, housePlacements, guruAspects, shaniAspects, kujaAspects)
	default:
		return []string{}
	}
}

func buildKarakamshaReqBody(req models.HouseEffectsReqBody) models.KarakamshaReqBody {
	return models.KarakamshaReqBody{
		Ascendant:        req.NavAscendant,
		Karakamsha:       req.Karakamsha,
		Surya:            req.NavSuryaPlacement,
		Budha:            req.NavBudhaPlacement,
		Shukra:           req.NavShukraPlacement,
		Chandra:          req.NavChandraPlacement,
		Rahu:             req.NavRahuPlacement,
		Ketu:             req.NavKetuPlacement,
		Kuja:             req.NavKujaPlacement,
		Guru:             req.NavGuruPlacement,
		Shani:            req.NavShaniPlacement,
		MercuryAfflicted: req.NavBudhaAfflicted,
		MoonWaxing:       req.ChandraWaxing,
	}
}

type IndexedString struct {
	Index int16  `json:"index"`
	Value string `json:"value"`
}

func ToIndexedStrings(input []string) []IndexedString {
	result := make([]IndexedString, len(input))

	for i, v := range input {
		result[i] = IndexedString{
			Index: int16(i),
			Value: v,
		}
	}

	return result
}

func GetAllHouseEffects(c *gin.Context) {
	// 1. Obtain the request body from the post request and validate them.
	var req models.HouseEffectsReqBody

	if err := c.ShouldBindJSON(&req); err != nil {
		// err will be validator.ValidationErrors if validation failed
		var errs []string

		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			for _, e := range validationErrs {
				// You can customize messages here
				switch e.Tag() {
				case "raashi":
					errs = append(errs, e.Field()+" must be a valid rāśi")
				default:
					errs = append(errs, e.Error())
				}
			}
		} else {
			errs = append(errs, err.Error())
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "validation failed",
			"details": errs,
		})
		return
	}

	rahu := models.AllRaashis(req.RahuPlacement)
	ketu := models.AllRaashis(req.KetuPlacement)

	if ketu != constants.OppositeAspectsStore[rahu] {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "validation failed",
			"details": []string{"Rahu and Ketu must always be diagonally opposite to each other"},
		})
		return
	}

	// 2. Compute the different house numbers with respect to the ascendant and the navamsha ascedant and karakamsha.
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
	karakamshaReqBody := buildKarakamshaReqBody(req)
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

	guruAspects := utils.GetGuruAspectRaashis(models.AllRaashis(req.GuruPlacement.Placement))
	shaniAspects := utils.GetShaniAspectRaashis(models.AllRaashis(req.ShaniPlacement.Placement))
	kujaAspects := utils.GetKujaAspectRaashis(models.AllRaashis(req.KujaPlacement.Placement))

	// 1. Determine which proportion array to use
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
		copy(proportions, constants.HouseeffectProprtionStore[models.AllRaashis(asc)])
    }

	// 2. Prepare results
	var result []models.HouseEffectsWithCommentory
	for i := range int8(12) {
        houseNum := i + 1
        proportion := proportions[i]

		if slices.Contains([]int8{60,50,40}, proportion) {

			comment := getShadBalaStyleComment(proportion)
	
			item := models.HouseEffectsWithCommentory{
				HouseNumber: houseNum,
				Theme:       constants.Themes[i],
				Comments:    comment,
				Effects:     getBhavaEffect(houseNum, &req, &housePlacements, &guruAspects, &shaniAspects, &kujaAspects),
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

	otherEffects := services.GeneralEffects(&req, &housePlacements, &guruAspects, &shaniAspects, &kujaAspects)

	var sum int = 0
    for _, v := range proportions {
        sum += int(v)
    }

	percentages := make([]float64, 12)
    for i, val := range proportions {
        percentages[i] = (float64(val) / float64(sum)) * 100
    }

	// Proceed with the response
	c.JSON(http.StatusOK, gin.H{
		"allHouseEffects": result,
		"zeroProportion": zeroProportionHouse,
		"proportions": percentages,
		"generalEffects": otherEffects,
		"karakamshaResults": ToIndexedStrings(services.KarakamshaEffectsService(&karakamshaHousePlacements, &karakamshaReqBody)),
	})
}

func GetBhavaLordEffects(c *gin.Context) {
	var req models.BhavaLordEffectsReqBody

	if err := c.ShouldBindJSON(&req); err != nil {
		// err will be validator.ValidationErrors if validation failed
		var errs []string

		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			for _, e := range validationErrs {
				// You can customize messages here
				switch e.Tag() {
				case "raashi":
					errs = append(errs, e.Field()+" must be a valid graha")
				default:
					errs = append(errs, e.Error())
				}
			}
		} else {
			errs = append(errs, err.Error())
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "validation failed",
			"details": errs,
		})
		return
	}

	ascendant := models.AllRaashis(req.Ascendant)

	// Get the house map for this ascendant
	housesByPlanet := constants.AscendantToRashyadhipatiMapStore[ascendant]
	effects := []string{}

	addEffects := func(planetStr string, placement int8) {
		planet := models.Raashyadhipati(planetStr)
		houses := housesByPlanet[planet]

		for _, houseNum := range houses {
			effect := getBhavaLordEffect(houseNum, placement)
			if effect != "" {
				effects = append(effects, effect)
			}
		}
	}

	addEffects(req.FirstHighestShadBalaLord, req.FirstPlacement)
	addEffects(req.SecondHighestShadBalaLord, req.SecondPlacement)
	addEffects(req.ThirdHighestShadBalaLord, req.ThirdPlacement)

	c.JSON(http.StatusOK, gin.H{
		"ascendant": req.Ascendant,
		"effects":   effects,
	})
}

func BhavaLordEffectsHandler(c *gin.Context) {
	var req models.BhavaLordAndShadBalaReqBody
	if err := c.ShouldBindJSON(&req); err != nil {
		// err will be validator.ValidationErrors if validation failed
		var errs []string
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			for _, e := range validationErrs {
				// You can customize messages here
				switch e.Tag() {
				case "raashi":
					errs = append(errs, e.Field()+" must be a valid rāśi")
				default:
					errs = append(errs, e.Error())
				}
			}
		} else {
			errs = append(errs, err.Error())
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "validation failed",
			"details": errs,
		})
		return
	}

	grahas := []string{"Surya", "Chandra", "Kuja", "Budha", "Guru", "Shukra", "Shani"}
	allGrahaDefined := make([]models.ShadBalaResField, 7)
	allGrahaDefinedSthana := make([]models.ShadBalaResField, 7)
	allGrahaDefinedKaala := make([]models.ShadBalaResField, 7)
	effectiveSthanaBala := make([]float32, 7)
	effectiveKaalaBala := make([]float32, 7)
	SthanaBalaMinimums := []float32{165.0,133.0,96.0,165.0,165.0,133.0,96.0}
	KaalaBalaMinimums := []float32{50,30,40,50,50,30,40}
	sumOfSthana := make([]float32, 7)
	sumOfKaala := make([]float32, 7)
	for i := range req.DrigBala {
		sumOfSthana[i] = req.DrigBala[i] + req.SthanaBala[i]
		sumOfKaala[i] = req.DrigBala[i] + req.KaalaBala[i]
	}
	for i := range 7 {
		effectiveSthanaBala[i] = sumOfSthana[i] / SthanaBalaMinimums[i]
		effectiveKaalaBala[i] = sumOfKaala[i] / KaalaBalaMinimums[i]
	}
	for i := range 7 {
		allGrahaDefined[i] = models.ShadBalaResField{
			Graha:         grahas[i],
			EffectiveBala: req.ShadBala[i],
			Rank:          0,
		}
		allGrahaDefinedSthana[i] = models.ShadBalaResField{
			Graha:         grahas[i],
			EffectiveBala: effectiveSthanaBala[i],
			Rank:          0,
		}
		allGrahaDefinedKaala[i] = models.ShadBalaResField{
			Graha:         grahas[i],
			EffectiveBala: effectiveKaalaBala[i],
			Rank:          0,
		}
	}
	items := make([]RankedItem, 7)
	itemsSthana := make([]RankedItem, 7)
	itemsKaala := make([]RankedItem, 7)
	for i := range 7 {
		items[i] = RankedItem{i, req.ShadBala[i]}
		itemsSthana[i] = RankedItem{i, effectiveSthanaBala[i]}
		itemsKaala[i] = RankedItem{i, effectiveKaalaBala[i]}
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].value > items[j].value
	})
	sort.Slice(itemsSthana, func(i, j int) bool {
		return itemsSthana[i].value > itemsSthana[j].value
	})
	sort.Slice(itemsKaala, func(i, j int) bool {
		return itemsKaala[i].value > itemsKaala[j].value
	})
	for rank, item := range items {
		allGrahaDefined[item.idx].Rank = int8(rank + 1)
	}
	for rank, item := range itemsSthana {
		allGrahaDefinedSthana[item.idx].Rank = int8(rank + 1)
	}
	for rank, item := range itemsKaala {
		allGrahaDefinedKaala[item.idx].Rank = int8(rank + 1)
	}
	sort.Slice(allGrahaDefined, func(i, j int) bool {
		return allGrahaDefined[i].Rank < allGrahaDefined[j].Rank
	})
	sort.Slice(allGrahaDefinedSthana, func(i, j int) bool {
		return allGrahaDefinedSthana[i].Rank < allGrahaDefinedSthana[j].Rank
	})
	sort.Slice(allGrahaDefinedKaala, func(i, j int) bool {
		return allGrahaDefinedKaala[i].Rank < allGrahaDefinedKaala[j].Rank
	})

	top3Lords := [3]string{}
	top3KaalaLords := "The native will obtain the highest success during the dasha/sub-dasha/sub-sub-dasha of the following lords " + allGrahaDefinedKaala[0].Graha + 
		"," + allGrahaDefinedKaala[1].Graha + "," + allGrahaDefinedKaala[2].Graha + 
	    ". When the time of the dasha belongs to at least 2 of these lords the highest success will be obtained"
	for i := range 3 {
		top3Lords[i] = allGrahaDefined[i].Graha
	}
	
	housePlacements := map[string]int16{
		"Surya":   utils.GetHouse(req.Ascendant, req.SuryaPlacement),
		"Budha":   utils.GetHouse(req.Ascendant, req.BudhaPlacement),
		"Shukra":  utils.GetHouse(req.Ascendant, req.ShukraPlacement),
		"Chandra": utils.GetHouse(req.Ascendant, req.ChandraPlacement),
		"Rahu":    utils.GetHouse(req.Ascendant, req.RahuPlacement),
		"Ketu":    utils.GetHouse(req.Ascendant, req.KetuPlacement),
		"Kuja":    utils.GetHouse(req.Ascendant, req.KujaPlacement),
		"Guru":    utils.GetHouse(req.Ascendant, req.GuruPlacement),
		"Shani":   utils.GetHouse(req.Ascendant, req.ShaniPlacement),
	}
	top3SthanaLords := "The native will obtain the highest level of mental satisfaction / success in " + constants.Themes[housePlacements[allGrahaDefinedSthana[0].Graha]-1] +
	 " and will dedicate his highest thinking towards " + constants.Themes[housePlacements[allGrahaDefinedSthana[1].Graha]-1]

	housesByPlanet := constants.AscendantToRashyadhipatiMapStore[models.AllRaashis(req.Ascendant)]
	effects := []string{}
	addEffects := func(planetStr string, placement int8) {
		planet := models.Raashyadhipati(planetStr)
		houses := housesByPlanet[planet]
		for _, houseNum := range houses {
			effect := getBhavaLordEffect(houseNum, placement)
			if effect != "" {
				effects = append(effects, effect)
			}
		}
	}
	for i := range 3 {
		lord := top3Lords[i]
		placement := housePlacements[lord]
		addEffects(lord, int8(placement))
	}
	var isMaleficAmongstTop3 = false
	for i := range 3 {
		lord := top3Lords[i]
		if slices.Contains([]string{"Surya", "Kuja", "Shani"}, lord) {
			isMaleficAmongstTop3 = true
			break
		}
	}
	var totalSthana float32
	var totalKaala float32

	for i := range req.SthanaBala {
		totalSthana += req.SthanaBala[i]
		totalKaala += req.KaalaBala[i]
	}
	var conclusion string
	if totalSthana > totalKaala+0.01 {
		conclusion = "Sthana Bala is greater overall leading to long term success"
	} else if totalKaala > totalSthana+0.01 {
		conclusion = "Kaala Bala is greater overall leading to small term success"
	} else {
		conclusion = "Sthana Bala and Kaala Bala are approximately equal"
	}
	var shadBalaComments = ""
	if isMaleficAmongstTop3 {
		shadBalaComments = "The individual will obtain success after a long and hard struggle. " + conclusion
	} else {
		shadBalaComments = "The individual will gain success smoothly. " + conclusion
	}

	c.JSON(http.StatusOK, gin.H{
		"effects": effects,
		"shadRanks": allGrahaDefined,
		"shadBalaComments": shadBalaComments,
		"top3KaalaLords": top3KaalaLords,
		"kaalaranks": allGrahaDefinedKaala,
		"top3SthanaLords": top3SthanaLords,
		"sthanaRanks": allGrahaDefinedSthana,
	})
}