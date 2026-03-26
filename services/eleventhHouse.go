package services

import (
	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
	"github.com/scientist-v08/bphs/utils"
)

func EleventHouseEffects(
	reqBody *models.HouseEffectsReqBody,
	housePlacements *map[string]int16,
	guruAspects *[]models.AllRaashis,
	shaniAspects *[]models.AllRaashis,
	kujaAspects *[]models.AllRaashis,
) []string {
	// Create a variable that collects all the effects of the eleventh house
	eleventhHouseEffects := make([]string, 0)

	// Get info
	eleventhHouse := constants.ReverseNumericalMappingsStore[(*reqBody).Ascendant][11]
	eleventhLord := constants.RaashyadhipatiMapStore[string(eleventhHouse)]
	eleventhLordPlacementNumber := (*housePlacements)[string(eleventhLord)]
	eleventhLordNatalRaashi := utils.GetHouseName((*reqBody).Ascendant, eleventhLordPlacementNumber)
	isEleventhLordDebilitated := constants.DebilitationMapStore[eleventhLord] == eleventhLordNatalRaashi
	isEleventhLordCombust := utils.IsLordCombust(eleventhLord, reqBody)
	secondHouse := constants.ReverseNumericalMappingsStore[(*reqBody).Ascendant][2]
	secondLord := constants.RaashyadhipatiMapStore[string(secondHouse)]
	secondLordPlacement := (*housePlacements)[string(secondLord)]
	thirdHouse := constants.ReverseNumericalMappingsStore[(*reqBody).Ascendant][3]
	thirdLord := constants.RaashyadhipatiMapStore[string(thirdHouse)]
	thirdLordPlacement := (*housePlacements)[string(thirdLord)]
	isEleventhLordExalted := constants.ExaltationMapStore[eleventhLord] == eleventhLordNatalRaashi
	guruPlacement := (*housePlacements)["Guru"]
	budhPlacement := (*housePlacements)["Budha"]
	chandraPlacement := (*housePlacements)["Chandra"]

	// Start appending
	if eleventhLordPlacementNumber == 11 || utils.IsKendra(eleventhLordPlacementNumber) || utils.IsKona(eleventhLordPlacementNumber) || isEleventhLordExalted {
		eleventhHouseEffects = append(eleventhHouseEffects, "The native will get many gains")
	}
	
	if eleventhLordPlacementNumber == 2 && utils.IsKendra(secondLordPlacement) && secondLordPlacement == guruPlacement {
		eleventhHouseEffects = append(eleventhHouseEffects, "The gains made by the native will be great")
	}
	
	if guruPlacement == 9 && budhPlacement == 9 && chandraPlacement == 9 {
		eleventhHouseEffects = append(eleventhHouseEffects, "The native will be endowed with wealth, grains, fortunes, diamonds and ornaments")
	}
	
	if eleventhLordPlacementNumber == 2 && secondLordPlacement == 11 {
		eleventhHouseEffects = append(eleventhHouseEffects, "The native will amass abundant fortunes after marriage")
	}
	
	if eleventhLordPlacementNumber == 3 && thirdLordPlacement == 11 {
		eleventhHouseEffects = append(eleventhHouseEffects, "The native will amass abundant fortunes through the younger siblings")
	}
	
	if isEleventhLordCombust || isEleventhLordDebilitated || utils.IsDusthana(eleventhLordPlacementNumber) {
		eleventhHouseEffects = append(eleventhHouseEffects, "The native may not make many gains despite numerous efforts")
	}

	// Return all the values
	return eleventhHouseEffects

}