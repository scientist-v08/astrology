package services

import (
	"slices"

	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
	"github.com/scientist-v08/bphs/utils"
)

func NinthHouseEffects(
	reqBody *models.HouseEffectsReqBody,
	housePlacements *map[string]int16,
	guruAspects *[]models.AllRaashis,
	shaniAspects *[]models.AllRaashis,
	kujaAspects *[]models.AllRaashis,
) []string {
	// Create a variable that collects all the effects of the second house
	ninthHouseEffects := make([]string, 0)

	// Get info
	ninthHouse := constants.ReverseNumericalMappingsStore[(*reqBody).Ascendant][9]
	ninthLord := constants.RaashyadhipatiMapStore[string(ninthHouse)]
	ninthLordPlacementNumber := (*housePlacements)[string(ninthLord)]
	sixthHouse := constants.ReverseNumericalMappingsStore[(*reqBody).Ascendant][6]
	sixthLord := constants.RaashyadhipatiMapStore[string(sixthHouse)]
	sixthLordPlacementNumber := (*housePlacements)[string(sixthLord)]
	fifthHouse := constants.ReverseNumericalMappingsStore[(*reqBody).Ascendant][5]
	fifthLord := constants.RaashyadhipatiMapStore[string(fifthHouse)]
	fifthLordPlacement := (*housePlacements)[string(fifthLord)]
	ninthLordHouse := utils.GetHouseName((*reqBody).Ascendant, ninthLordPlacementNumber)
	guruPlacementNumber := (*housePlacements)["Guru"]
	ascendantLord := constants.RaashyadhipatiMapStore[(*reqBody).Ascendant]
	ascendantLordPlacement := (*housePlacements)[string(ascendantLord)]
	ascendantLordHouse := utils.GetHouseName((*reqBody).Ascendant, ascendantLordPlacement)
	ascendantLordIsStrong := utils.IsGrahaStrong(
		!utils.IsLordCombust(ascendantLord, reqBody),
		(constants.DebilitationMapStore[ascendantLord] != ascendantLordHouse),
		!slices.Contains(constants.EnemySignsMappingStore[ascendantLord], ascendantLordHouse),
		!utils.IsDusthana(ascendantLordPlacement),
	)
	ninthLordIsStrong := utils.IsGrahaStrong(
		!utils.IsLordCombust(ninthLord, reqBody), 
		(constants.DebilitationMapStore[ninthLord] != ninthLordHouse),
		!slices.Contains(constants.EnemySignsMappingStore[ninthLord], ninthLordHouse),
		!utils.IsDusthana(ninthLordPlacementNumber),
	)
	shukraPlacementNumber := (*housePlacements)["Shukra"]
	kujaPlacementNumber := (*housePlacements)["Kuja"]
	kujaNatalHouse := utils.GetHouseName((*reqBody).Ascendant, kujaPlacementNumber)
	kujaInTenOrTwelve := slices.Contains([]int16{10, 12}, kujaPlacementNumber)
	kujaNotStrong := constants.ExaltationMapStore[constants.Kuja] != kujaNatalHouse && !slices.Contains([]models.AllRaashis{constants.Mesha, constants.Vruschika}, kujaNatalHouse)
	kujaUnfortunateCondition := kujaInTenOrTwelve && kujaNotStrong
	ninthLordNotGuru := ninthLord != constants.Guru
	ninthLordNotShukra := ninthLord != constants.Shukra
	ninthLordGettingGuruAspect := slices.Contains(*guruAspects, ninthLordHouse)
	ninthLordConjunctGuru := ninthLordPlacementNumber == guruPlacementNumber
	isNinthLordAspectedOrConjunctGuru := ninthLordGettingGuruAspect || ninthLordConjunctGuru
	sunPlacementNumber := (*housePlacements)["Surya"]
	sunNatalRaashi := utils.GetHouseName((*reqBody).Ascendant, sunPlacementNumber)
	shukraNatalRaashi := utils.GetHouseName((*reqBody).Ascendant, shukraPlacementNumber)
	isSunExalted := constants.SuryaUccha == sunNatalRaashi
	isShukraExalted := constants.ShukraUccha == shukraNatalRaashi

	// Start appending
	ninthHouseEffects = append(ninthHouseEffects, "The ninth house is strong and indicates good fortune, strong belief system, and positive relations with father and gurus.")

	if ninthLordPlacementNumber == 9 && !utils.IsLordCombust(ninthLord, reqBody) {
		ninthHouseEffects = append(ninthHouseEffects, "The native will be fortunate")
	}
	if guruPlacementNumber == 9 && utils.IsKendra(ninthLordPlacementNumber) && ascendantLordIsStrong {
		ninthHouseEffects = append(ninthHouseEffects, "The native will be fortunate")
	}

	if shukraPlacementNumber == 9 && ninthLordIsStrong && utils.IsKendra(guruPlacementNumber) {
		ninthHouseEffects = append(ninthHouseEffects, "The native's father will become fortunate")
	}

	if constants.DebilitationMapStore[ninthLord] == ninthLordHouse && kujaUnfortunateCondition {
		ninthHouseEffects = append(ninthHouseEffects, "The native's father is likely poor or the native will disinherit the fathers property or the native can only obtain the father's property through litigation")
	}

	if constants.ExaltationMapStore[ninthLord] == ninthLordHouse && utils.IsKendra(shukraPlacementNumber) && (*housePlacements)["NavGuru"] == 9 {
		ninthHouseEffects = append(ninthHouseEffects, "Navamsha: The father of the native will live for a long time")
	}

	if ninthLordNotGuru && ninthLordGettingGuruAspect && utils.IsKendra(ninthLordPlacementNumber) {
		ninthHouseEffects = append(ninthHouseEffects, "The native's father will be extremely fortunate")
	}

	if isSunExalted && ninthLordPlacementNumber == 11 {
		ninthHouseEffects = append(ninthHouseEffects, "The native will be devoted to his father and be virtuous and be dear to those in leadership positions")
	}

	if isNinthLordAspectedOrConjunctGuru && utils.IsKona(sunPlacementNumber) && ninthLordPlacementNumber == 7 {
		ninthHouseEffects = append(ninthHouseEffects, "The native will be devoted to his father")
	}

	if ascendantLordPlacement == 9 && sixthLordPlacementNumber == ascendantLordPlacement {
		ninthHouseEffects = append(ninthHouseEffects, "The native and the native's father are likely to be enemies")
	}

	if ninthLordNotShukra && isShukraExalted && ninthLordPlacementNumber == shukraPlacementNumber && (*housePlacements)["Shani"] == 3 {
		ninthHouseEffects = append(ninthHouseEffects, "The native will have abundant fortunes")
	}

	if ascendantLordPlacement == 9 && ninthLordPlacementNumber == 1 && guruPlacementNumber == 7 {
		ninthHouseEffects = append(ninthHouseEffects, "The native will gain wealth and vehicles")
	}

	if (*housePlacements)["RahuPlacement"] == 5 && fifthLordPlacement == 8 && constants.DebilitationMapStore[ninthLord] == ninthLordHouse {
		ninthHouseEffects = append(ninthHouseEffects, "The native will be devoid of fortunes")
	}

	if (*housePlacements)["Shani"] == 9 && (*housePlacements)["Chandra"] == 9 && constants.DebilitationMapStore[ascendantLord] == ascendantLordHouse {
		ninthHouseEffects = append(ninthHouseEffects, "The native will be unfortunate")
	}

	return ninthHouseEffects
}