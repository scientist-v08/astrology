package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/dto"
	"github.com/scientist-v08/bphs/services/bhavaLords"
	"github.com/scientist-v08/bphs/services/houses"
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
	return bhavaLords.GetBhavaLordEffect(house, lordPlacement)
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

	c.JSON(http.StatusOK, apiBhavaLordsEffects(req))
}

func apiBhavaLordsEffects(req dto.BhavaLordAndShadBalaReqBody) map[string]any {
	housePlacements := bhavaLords.GetHousePlacements(req)

	sthanaRanks, top3SthanaLords, totalSthana := bhavaLords.CalculateSthanaBalaEffects(req, housePlacements)
	kaalaRanks, top3KaalaLords, totalKaala := bhavaLords.CalculateKaalaBalaEffects(req)
	drikComments := bhavaLords.CalculateDrikBalaEffects(req)
	shadRanks, effects, shadBalaComments := bhavaLords.CalculateShadbalaEffects(req, housePlacements, totalSthana, totalKaala)

	shadBalaComments = shadBalaComments + drikComments

	return map[string]any{
		"effects":          effects,
		"shadRanks":        shadRanks,
		"shadBalaComments": shadBalaComments,
		"top3KaalaLords":   top3KaalaLords,
		"kaalaranks":       kaalaRanks,
		"top3SthanaLords":  top3SthanaLords,
		"sthanaRanks":      sthanaRanks,
	}
}

func sharedBhavaLordsEffects(req dto.BhavaLordAndShadBalaReqBody) map[string]any {
	housePlacements := bhavaLords.GetHousePlacements(req)

	sthanaRanks, top3SthanaLords, totalSthana := bhavaLords.CalculateSthanaBalaEffects(req, housePlacements)
	kaalaRanks, top3KaalaLords, totalKaala := bhavaLords.CalculateKaalaBalaEffects(req)
	digComments := bhavaLords.CalculateDigBalaEffects(req)
	chestaComments := bhavaLords.CalculateChestaBalaEffects(req)
	drikComments := bhavaLords.CalculateDrikBalaEffects(req)
	shadRanks, effects, shadBalaComments := bhavaLords.CalculateShadbalaEffects(req, housePlacements, totalSthana, totalKaala)

	shadBalaComments = shadBalaComments + drikComments + digComments + chestaComments

	return map[string]any{
		"effects":          effects,
		"shadRanks":        shadRanks,
		"shadBalaComments": shadBalaComments,
		"top3KaalaLords":   top3KaalaLords,
		"kaalaranks":       kaalaRanks,
		"top3SthanaLords":  top3SthanaLords,
		"sthanaRanks":      sthanaRanks,
	}
}
