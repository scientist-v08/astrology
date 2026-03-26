package services

import (
	"slices"

	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
)

func ThirdHouseEffects(
	reqBody *models.HouseEffectsReqBody,
	housePlacements *map[string]int16,
	guruAspects *[]models.AllRaashis,
	shaniAspects *[]models.AllRaashis,
	kujaAspects *[]models.AllRaashis,
) []string {
	// Create a variable that collects all the effects of the second house
	thirdHouseEffects := make([]string, 0)

	// Get info
	thirdHouse := constants.ReverseNumericalMappingsStore[(*reqBody).Ascendant][3]

	// Start appending
	thirdHouseEffects = append(thirdHouseEffects, "The third house is strong and suggests that the long term goal and courage and FIL relations are good.")

	if (*housePlacements)["Budha"] == 3 || 
	(*housePlacements)["Shukra"] == 3 || 
	(*housePlacements)["Guru"] == 3 ||
	((*housePlacements)["Chandra"] == 3) {
		thirdHouseEffects = append(thirdHouseEffects, "The native will have younger siblings and will be courageous")
	}
	if slices.Contains(*guruAspects, thirdHouse) || 
	constants.OppositeAspectsStore[models.AllRaashis((*reqBody).ShukraPlacement.Placement)] == thirdHouse || 
	constants.OppositeAspectsStore[models.AllRaashis((*reqBody).BudhaPlacement.Placement)] == thirdHouse {
		thirdHouseEffects = append(thirdHouseEffects, "The native will have younger siblings and will be courageous")
	}

	if (*housePlacements)["Surya"] == 3 {
		thirdHouseEffects = append(thirdHouseEffects, "The elder sibling will die")
	}

	if (*housePlacements)["Shani"] == 3 {
		thirdHouseEffects = append(thirdHouseEffects, "The younger sibling will die")
	}

	if (*housePlacements)["Kuja"] == 3 {
		thirdHouseEffects = append(thirdHouseEffects, "The elder siblings and the younger siblings will die")
	}

	if len(thirdHouseEffects) > 1 {
		thirdHouseEffects = append(thirdHouseEffects, "The effects of the third house is to be announced after assessing the strength of such yogas")
	}

	// Return all the effects
	return thirdHouseEffects
}