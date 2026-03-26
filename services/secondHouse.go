package services

import (
	"slices"

	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
	"github.com/scientist-v08/bphs/utils"
)

func SecondHouseEffects(
	reqBody *models.HouseEffectsReqBody,
	housePlacements *map[string]int16,
	guruAspects *[]models.AllRaashis,
	shaniAspects *[]models.AllRaashis,
	kujaAspects *[]models.AllRaashis,
) []string {
	// Create a variable that collects all the effects of the second house
	secondHouseEffects := make([]string, 0)

	// Get info
	secondHouse := constants.ReverseNumericalMappingsStore[(*reqBody).Ascendant][2]
	secondLord := constants.RaashyadhipatiMapStore[string(secondHouse)]
	secondLordPlacement := (*housePlacements)[string(secondLord)]
	secondLordHouse := utils.GetHouseName((*reqBody).Ascendant, secondLordPlacement)
	eleventhHouse := constants.ReverseNumericalMappingsStore[(*reqBody).Ascendant][11]
	eleventhLord := constants.RaashyadhipatiMapStore[string(eleventhHouse)]
	eleventhLordPlacement := (*housePlacements)[string(eleventhLord)]
	eleventhLordHouse := utils.GetHouseName((*reqBody).Ascendant, eleventhLordPlacement)

	// Start appending to the string array based on conditions
	if secondLordPlacement == 2 || utils.IsKendra(secondLordPlacement) || utils.IsKendra(secondLordPlacement) {
		secondHouseEffects = append(secondHouseEffects, "As the second lord is in 2/4/5/7/9/10, wealth will be promoted")
	}

	if utils.IsDusthana(secondLordPlacement) {
		secondHouseEffects = append(secondHouseEffects, "Financial conditions will decline")
	}

	if (*housePlacements)["Shukra"] == 2 || (*housePlacements)["Budha"] == 2 {
		secondHouseEffects = append(secondHouseEffects, "Shukra/Budha will give wealth")
	}

	if (*reqBody).GuruPlacement.Placement == string(secondHouse) || (*reqBody).GuruPlacement.Placement == (*reqBody).KujaPlacement.Placement {
		secondHouseEffects = append(secondHouseEffects, "Guru will make the native wealthy")
	}

	if (*housePlacements)["Guru"] == 2 {
		secondHouseEffects = append(secondHouseEffects, "SGuru(2nd lord) will give wealth but is not wholly auspicious")
	}

	if (*housePlacements)["Surya"] == 2 || (*housePlacements)["Shani"] == 2 || (*housePlacements)["Kuja"] == 2 {
		secondHouseEffects = append(secondHouseEffects, "Surya/Shani/Kuja will destroy wealth")
	}

	if eleventhLordPlacement == 2 && secondLordPlacement == 11 {
		secondHouseEffects = append(secondHouseEffects, "Wealth will be acquired by the native")
	}

	if eleventhLordPlacement == secondLordPlacement && (utils.IsKendra(secondLordPlacement) || utils.IsKona(secondLordPlacement)) {
		secondHouseEffects = append(secondHouseEffects, "Wealth will be acquired by the native")
	}

	if utils.IsKendra(secondLordPlacement) && utils.IsKona(eleventhLordPlacement) {
		secondHouseEffects = append(secondHouseEffects, "The subject will be wealthy")
	}
	if utils.IsKendra(secondLordPlacement) && (slices.Contains(*guruAspects, eleventhLordHouse) || (*reqBody).GuruPlacement.Placement == string(eleventhLordHouse)) {
		secondHouseEffects = append(secondHouseEffects, "The subject will be wealthy")
	}
	if utils.IsKendra(secondLordPlacement) && (constants.OppositeAspectsStore[models.AllRaashis((*reqBody).ShukraPlacement.Placement)] == eleventhLordHouse || (*reqBody).ShukraPlacement.Placement == string(eleventhLordHouse)) {
		secondHouseEffects = append(secondHouseEffects, "The subject will be wealthy")
	}

	if utils.IsDusthana(secondLordPlacement) && utils.IsDusthana(eleventhLordPlacement) {
		if (*housePlacements)["Kuja"] == 2 || (*housePlacements)["Surya"] == 2 || (*housePlacements)["Shani"] == 2 {
			secondHouseEffects = append(secondHouseEffects, "The native will be poor")
		}
		if (*housePlacements)["Kuja"] == 11 && (*housePlacements)["RahuPlacement"] == 2 {
			secondHouseEffects = append(secondHouseEffects, "The native will lose wealth on account of royal punishments")
		}
	}

	if (*housePlacements)["Guru"] == 11 && (*housePlacements)["Shukra"] == 2 && ((*housePlacements)["Budha"] == 12 || ((*housePlacements)["Chandra"] == 12 && (*reqBody).ChandraWaxing)) {
		if secondLordPlacement == (*housePlacements)["Guru"] || secondLordPlacement == (*housePlacements)["Shukra"] || secondLordPlacement == (*housePlacements)["Budha"] || (secondLordPlacement == (*housePlacements)["Chandra"] && (*reqBody).ChandraWaxing) {
			secondHouseEffects = append(secondHouseEffects, "The native will spend money on religious activities")
		}
	}

	if secondLordHouse == constants.ExaltationMapStore[secondLord] {
		secondHouseEffects = append(secondHouseEffects, "The native looks after his people and will help others and will become famous")
	}

	if (*housePlacements)["Kuja"] == 2 || (*housePlacements)["Surya"] == 2 || (*housePlacements)["Shani"] == 2 {
		if secondLordPlacement == (*housePlacements)["Kuja"] || secondLordPlacement == (*housePlacements)["Surya"] || secondLordPlacement == (*housePlacements)["Shani"] {
			secondHouseEffects = append(secondHouseEffects, "The native will be a liar")
		}
	}

	return secondHouseEffects
}