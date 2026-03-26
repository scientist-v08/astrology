package services

import (
	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
	"github.com/scientist-v08/bphs/utils"
)

func EighthHouseEffects(
	reqBody *models.HouseEffectsReqBody,
	housePlacements *map[string]int16,
	guruAspects *[]models.AllRaashis,
	shaniAspects *[]models.AllRaashis,
	kujaAspects *[]models.AllRaashis,
) []string {
	// Create a variable that collects all the effects of the second house
	eighthHouseEffects := make([]string, 0)

	// Get info
	eighthHouse := constants.ReverseNumericalMappingsStore[(*reqBody).Ascendant][8]
	eighthLord := constants.RaashyadhipatiMapStore[string(eighthHouse)]
	eighthLordPlacementNumber := (*housePlacements)[string(eighthLord)]
	tenthHouse := constants.ReverseNumericalMappingsStore[(*reqBody).Ascendant][10]
	tenthLord := constants.RaashyadhipatiMapStore[string(tenthHouse)]
	tenthLordPlacementNumber := (*housePlacements)[string(tenthLord)]
	sixthHouse := constants.ReverseNumericalMappingsStore[(*reqBody).Ascendant][6]
	sixthLord := constants.RaashyadhipatiMapStore[string(sixthHouse)]
	sixthLordPlacementNumber := (*housePlacements)[string(sixthLord)]
	twelveHouse := constants.ReverseNumericalMappingsStore[(*reqBody).Ascendant][12]
	twelveLord := constants.RaashyadhipatiMapStore[string(twelveHouse)]
	twelveLordPlacementNumber := (*housePlacements)[string(twelveLord)]
	ascendantLord := constants.RaashyadhipatiMapStore[(*reqBody).Ascendant]
	ascendantLordPlacement := (*housePlacements)[string(ascendantLord)]
	shaniPlacementNumber := (*housePlacements)["Shani"]
	kujaPlacementNumber := (*housePlacements)["Kuja"]
	suryaPlacementNumber := (*housePlacements)["Surya"]
	budhaPlacementNumber := (*housePlacements)["Budha"]
	chandraPlacementNumber := (*housePlacements)["Chandra"]
	isEighthLordAfflicted := (eighthLordPlacementNumber == suryaPlacementNumber && eighthLord != constants.Surya) ||
		(eighthLord != constants.Shani && eighthLordPlacementNumber == shaniPlacementNumber) ||
		(eighthLord != constants.Kuja && eighthLordPlacementNumber == kujaPlacementNumber) ||
		(eighthLord != constants.Budha && eighthLordPlacementNumber == budhaPlacementNumber && !(*reqBody).BudhaUnafflicted) ||
		(eighthLord != constants.Chandra && eighthLordPlacementNumber == chandraPlacementNumber && !(*reqBody).ChandraWaxing)
	isTenthLordAfflicted := (tenthLord != constants.Surya && tenthLordPlacementNumber == suryaPlacementNumber) ||
		(tenthLord != constants.Shani && tenthLordPlacementNumber == shaniPlacementNumber) ||
		(tenthLord != constants.Kuja && tenthLordPlacementNumber == kujaPlacementNumber) ||
		(tenthLord != constants.Budha && tenthLordPlacementNumber == budhaPlacementNumber && !(*reqBody).BudhaUnafflicted) ||
		(tenthLord != constants.Chandra && tenthLordPlacementNumber == chandraPlacementNumber && !(*reqBody).ChandraWaxing)
	isShaniWithAMalefic := (shaniPlacementNumber == suryaPlacementNumber) ||
		(shaniPlacementNumber == kujaPlacementNumber) ||
		(shaniPlacementNumber == budhaPlacementNumber && !(*reqBody).BudhaUnafflicted) ||
		(shaniPlacementNumber == chandraPlacementNumber && !(*reqBody).ChandraWaxing)
	saturnIsNotTheAscendantLord := ascendantLord != constants.Shani

	// Start appending
	eighthHouseEffects = append(eighthHouseEffects, "Although this house speaks about longevity it is better to calculate the lifespan using Nisargaayu, Ashmaayu and Pindaayu")

	if utils.IsKendra(eighthLordPlacementNumber) {
		eighthHouseEffects = append(eighthHouseEffects, "The native will have a long life")
	}
	if sixthLordPlacementNumber == 12 {
		eighthHouseEffects = append(eighthHouseEffects, "The native will have a long life")
	}
	if twelveLordPlacementNumber == 12 && sixthLordPlacementNumber == 6 {
		eighthHouseEffects = append(eighthHouseEffects, "The native will have a long life")
	}
	if twelveLordPlacementNumber == 8 && sixthLordPlacementNumber == 1 {
		eighthHouseEffects = append(eighthHouseEffects, "The native will have a long life")
	}

	if eighthLordPlacementNumber == 8 && (ascendantLordPlacement == 8 || isEighthLordAfflicted) {
		eighthHouseEffects = append(eighthHouseEffects, "The native will have a short life")
	}
	if shaniPlacementNumber == 8 && (isShaniWithAMalefic || (saturnIsNotTheAscendantLord && ascendantLordPlacement == 8)) {
		eighthHouseEffects = append(eighthHouseEffects, "The native will have a short life")
	}
	if tenthLordPlacementNumber == 8 && (isTenthLordAfflicted || (tenthLord != ascendantLord && ascendantLordPlacement == 8)) {
		eighthHouseEffects = append(eighthHouseEffects, "The native will have a short life")
	}

	return eighthHouseEffects
}