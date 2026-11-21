package controller

import (
	"fmt"
	"net/http"
	"slices"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/dto"
	"github.com/scientist-v08/bphs/services/bhavaLords"
	"github.com/scientist-v08/bphs/services/houses"
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
		return bhavaLords.FirstBhavaLordEffect(lordPlacement)
	case 2:
		return bhavaLords.SecondBhavaLordEffect(lordPlacement)
	case 3:
		return bhavaLords.ThirdBhavaLordEffect(lordPlacement)
	case 4:
		return bhavaLords.FourthBhavaLordEffect(lordPlacement)
	case 5:
		return bhavaLords.FifthBhavaLordEffect(lordPlacement)
	case 6:
		return bhavaLords.SixthBhavaLordEffect(lordPlacement)
	case 7:
		return bhavaLords.SeventhBhavaLordEffect(lordPlacement)
	case 8:
		return bhavaLords.EighthBhavaLordEffect(lordPlacement)
	case 9:
		return bhavaLords.NinthBhavaLordEffect(lordPlacement)
	case 10:
		return bhavaLords.TenthBhavaLordEffect(lordPlacement)
	case 11:
		return bhavaLords.EleventhBhavaLordEffect(lordPlacement)
	case 12:
		return bhavaLords.TwelfthBhavaLordEffect(lordPlacement)
	default:
		return "" // invalid house
	}
}

func getBhavaEffect(
	house int8,
	req *dto.HouseEffectsReqBody,
	housePlacements *map[string]int16,
	guruAspects *[]dto.AllRaashis,
	shaniAspects *[]dto.AllRaashis,
	kujaAspects *[]dto.AllRaashis,
) []string {
   switch house {
	case 1:
		return houses.FirstHouseEffects(req, housePlacements, guruAspects, shaniAspects, kujaAspects)
	case 2:
		return houses.SecondHouseEffects(req, housePlacements, guruAspects, shaniAspects, kujaAspects)
	case 3:
		return houses.ThirdHouseEffects(req, housePlacements, guruAspects, shaniAspects, kujaAspects)
	case 4:
		return houses.FourthHouseEffects(req, housePlacements, guruAspects, shaniAspects, kujaAspects)
	case 5:
		return houses.FifthHouseEffects(req, housePlacements, guruAspects, shaniAspects, kujaAspects)
	case 6:
		return houses.SixthHouseEffects(req, housePlacements, guruAspects, shaniAspects, kujaAspects)
	case 7:
		return houses.SeventhHouseEffects(req, housePlacements, guruAspects, shaniAspects, kujaAspects)
	case 8:
		return houses.EighthHouseEffects(req, housePlacements, guruAspects, shaniAspects, kujaAspects)
	case 9:
		return houses.NinthHouseEffects(req, housePlacements, guruAspects, shaniAspects, kujaAspects)
	case 10:
		return houses.TenthHouseEffects(req, housePlacements, guruAspects, shaniAspects, kujaAspects)
	case 11:
		return houses.EleventHouseEffects(req, housePlacements, guruAspects, shaniAspects, kujaAspects)
	case 12:
		return houses.TwelfthHouseEffects(req, housePlacements, guruAspects, shaniAspects, kujaAspects)
	default:
		return []string{}
	}
}

func buildKarakamshaReqBody(req dto.HouseEffectsReqBody) dto.KarakamshaReqBody {
	return dto.KarakamshaReqBody{
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
	var req dto.HouseEffectsReqBody

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

	rahu := dto.AllRaashis(req.RahuPlacement)
	ketu := dto.AllRaashis(req.KetuPlacement)

	if ketu != constants.OppositeAspectsStore[rahu] {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "validation failed",
			"details": []string{"Rahu and Ketu must always be diagonally opposite to each other"},
		})
		return
	}

	// 2. Obtain the results
	result := calculateHouseEffects(&req)

	// 3. Proceed with the response
	c.JSON(http.StatusOK, result)
}

func GetBhavaLordEffects(c *gin.Context) {
	var req dto.BhavaLordEffectsReqBody

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

	ascendant := dto.AllRaashis(req.Ascendant)

	// Get the house map for this ascendant
	housesByPlanet := constants.AscendantToRashyadhipatiMapStore[ascendant]
	effects := []string{}

	addEffects := func(planetStr string, placement int8) {
		planet := dto.Raashyadhipati(planetStr)
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
	var req dto.BhavaLordAndShadBalaReqBody
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

	c.JSON(http.StatusOK, sharedBhavaLordsEffects(req))
}

func sharedBhavaLordsEffects(req dto.BhavaLordAndShadBalaReqBody) map[string]any {
	grahas := []string{"Surya", "Chandra", "Kuja", "Budha", "Guru", "Shukra", "Shani"}
	allGrahaDefined := make([]dto.ShadBalaResField, 7)
	allGrahaDefinedSthana := make([]dto.ShadBalaResField, 7)
	allGrahaDefinedKaala := make([]dto.ShadBalaResField, 7)
	allGrahaDefinedDig := make([]dto.ShadBalaResField, 7)
	allGrahaDefinedChesta := make([]dto.ShadBalaResField, 7)
	effectiveSthanaBala := make([]float32, 7)
	effectiveKaalaBala := make([]float32, 7)
	effectiveDigBala := make([]float32, 7)
	effectiveChestaBala := make([]float32, 7)
	effectiveChestaBalaAct := make([]float32, 7)
	SthanaBalaMinimums := []float32{165.0,133.0,96.0,165.0,165.0,133.0,96.0}
	KaalaBalaMinimums := []float32{50,30,40,50,50,30,40}
	digBalaMinimums := []float32{35,50,30,35,35,50,30}
	chestaBalaMinimums := []float32{112,100,67,112,112,100,67}
	sumOfSthana := make([]float32, 7)
	sumOfKaala := make([]float32, 7)
	totalDrik := float32(0)
	for i := range req.DrigBala {
		sumOfSthana[i] = req.DrigBala[i] + req.SthanaBala[i]
		sumOfKaala[i] = req.DrigBala[i] + req.KaalaBala[i]
	}
	for i := range 7 {
		switch i {
		case 0:
			effectiveChestaBala[i] = float32(req.SuryaAyana)
		case 1:
			effectiveChestaBala[i] = float32(req.ChandraPaksha)
		default:
			effectiveChestaBala[i] = req.ChestaBala[i] * 2
		}
	}
	for i := range 7 {
		effectiveSthanaBala[i] = sumOfSthana[i] / SthanaBalaMinimums[i]
		effectiveKaalaBala[i] = sumOfKaala[i] / KaalaBalaMinimums[i]
		effectiveDigBala[i] = req.DigBala[i] / digBalaMinimums[i]
		effectiveChestaBalaAct[i] = effectiveChestaBala[i] / chestaBalaMinimums[i]
		totalDrik += req.DrigBala[i]
	}
	for i := range 7 {
		allGrahaDefined[i] = dto.ShadBalaResField{
			Graha:         grahas[i],
			EffectiveBala: req.ShadBala[i],
			Rank:          0,
		}
		allGrahaDefinedSthana[i] = dto.ShadBalaResField{
			Graha:         grahas[i],
			EffectiveBala: effectiveSthanaBala[i],
			Rank:          0,
		}
		allGrahaDefinedKaala[i] = dto.ShadBalaResField{
			Graha:         grahas[i],
			EffectiveBala: effectiveKaalaBala[i],
			Rank:          0,
		}
		allGrahaDefinedDig[i] = dto.ShadBalaResField{
			Graha:         grahas[i],
			EffectiveBala: effectiveDigBala[i],
			Rank:          0,
		}
		allGrahaDefinedChesta[i] = dto.ShadBalaResField{
			Graha:         grahas[i],
			EffectiveBala: effectiveChestaBalaAct[i],
			Rank:          0,
		}
	}
	items := make([]RankedItem, 7)
	itemsSthana := make([]RankedItem, 7)
	itemsKaala := make([]RankedItem, 7)
	itemsDig := make([]RankedItem, 7)
	itemsChesta := make([]RankedItem, 7)
	for i := range 7 {
		items[i] = RankedItem{i, req.ShadBala[i]}
		itemsSthana[i] = RankedItem{i, effectiveSthanaBala[i]}
		itemsKaala[i] = RankedItem{i, effectiveKaalaBala[i]}
		itemsDig[i] = RankedItem{i, effectiveDigBala[i]}
		itemsChesta[i] = RankedItem{i, effectiveChestaBalaAct[i]}
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
	sort.Slice(itemsDig, func(i, j int) bool {
		return itemsDig[i].value > itemsDig[j].value
	})
	sort.Slice(itemsChesta, func(i, j int) bool {
		return itemsChesta[i].value > itemsChesta[j].value
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
	for rank, item := range itemsDig {
		allGrahaDefinedDig[item.idx].Rank = int8(rank + 1)
	}
	for rank, item := range itemsChesta {
		allGrahaDefinedChesta[item.idx].Rank = int8(rank + 1)
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
	sort.Slice(allGrahaDefinedDig, func(i, j int) bool {
		return allGrahaDefinedDig[i].Rank < allGrahaDefinedDig[j].Rank
	})
	sort.Slice(allGrahaDefinedChesta, func(i, j int) bool {
		return allGrahaDefinedChesta[i].Rank < allGrahaDefinedChesta[j].Rank
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

	housesByPlanet := constants.AscendantToRashyadhipatiMapStore[dto.AllRaashis(req.Ascendant)]
	effects := []string{}
	addEffects := func(planetStr string, placement int8) {
		planet := dto.Raashyadhipati(planetStr)
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
	if totalDrik < float32(0) {
		shadBalaComments += ". Due to Drik bala the native is likely to have a long life span. But also consult other longevity methods for confirmation."
	}

	// Now add comments related to Dig Bala
	grahaDirections := map[string]string{
		"Surya":   "East",
		"Chandra": "North-West",
		"Kuja":    "South",
		"Budha":   "North",
		"Guru":    "North-East",
		"Shukra":  "South-East",
		"Shani":   "West",
	}

	var digBalaStr1 string
	dir1 := grahaDirections[allGrahaDefinedDig[0].Graha]
	if len(allGrahaDefinedDig) > 0 && allGrahaDefinedDig[0].EffectiveBala > 1 {
		digBalaStr1 = fmt.Sprintf(
			". Going in the %s direction (ruled by %s) from the place of birth will give success, material gain, and professional achievements",
			dir1,
			allGrahaDefinedDig[0].Graha,
		)
	} else {
		digBalaStr1 = fmt.Sprintf(
			". Going in the %s direction (ruled by %s) from the place of birth will give success, material gain, and professional achievements however the graha has not obtained the minimum required Dig bala",
			dir1,
			allGrahaDefinedDig[0].Graha,
		)
	}

	var digBalaStr2 string
	if len(allGrahaDefinedDig) > 1 {
		dir2 := grahaDirections[allGrahaDefinedDig[1].Graha]
		digBalaStr2 = fmt.Sprintf(
			". Going in the %s direction (ruled by %s) will help the person find increased mental energy, peace, and creative capabilities.",
			dir2,
			allGrahaDefinedDig[1].Graha,
		)
	}

	var digBalaStr3 string
	if len(allGrahaDefinedDig) > 2 {
		dir3 := grahaDirections[allGrahaDefinedDig[2].Graha]
		digBalaStr3 = fmt.Sprintf(
			" Going in the %s direction (ruled by %s) is where the person may find a compatible karmic partner or companion in life.",
			dir3,
			allGrahaDefinedDig[2].Graha,
		)
	}

	shadBalaComments = shadBalaComments + digBalaStr1 + digBalaStr2 + digBalaStr3

	var chestaTopStr, chestaBottomStr string

	if len(allGrahaDefinedChesta) >= 7 {
		// Top 3 grahas (indices 0, 1, 2)
		chestaTopStr = fmt.Sprintf(
			" The top 3 grahas in Chesta Bala are %s(%.2f), %s(%.2f), and %s(%.2f). In their sub-period (Antardasha) or sub-sub-periods (Pratyantardasha), any important travel, launching a project, or making a move to achieve an objective will succeed.",
			allGrahaDefinedChesta[0].Graha,
			allGrahaDefinedChesta[0].EffectiveBala,
			allGrahaDefinedChesta[1].Graha,
			allGrahaDefinedChesta[1].EffectiveBala,
			allGrahaDefinedChesta[2].Graha,
			allGrahaDefinedChesta[2].EffectiveBala,
		)

		// Bottom 3 grahas (indices 4, 5, 6)
		chestaBottomStr = fmt.Sprintf(
			" The lowest 3 grahas in Chesta Bala are %s(%.2f), %s(%.2f), and %s(%.2f). In their sub-period (Antardasha) or sub-sub-periods (Pratyantardasha), important travel, launching a project, or making a move to achieve an objective may not be achieved.",
			allGrahaDefinedChesta[4].Graha,
			allGrahaDefinedChesta[4].EffectiveBala,
			allGrahaDefinedChesta[5].Graha,
			allGrahaDefinedChesta[5].EffectiveBala,
			allGrahaDefinedChesta[6].Graha,
			allGrahaDefinedChesta[6].EffectiveBala,
		)
	}

	shadBalaComments = shadBalaComments + chestaTopStr + chestaBottomStr

	return map[string]any{
		"effects": effects,
		"shadRanks": allGrahaDefined,
		"shadBalaComments": shadBalaComments,
		"top3KaalaLords": top3KaalaLords,
		"kaalaranks": allGrahaDefinedKaala,
		"top3SthanaLords": top3SthanaLords,
		"sthanaRanks": allGrahaDefinedSthana,
	}
}
