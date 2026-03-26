package services

import (
	"slices"

	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
	"github.com/scientist-v08/bphs/utils"
)

func FirstHouseEffects(
	reqBody *models.HouseEffectsReqBody,
	housePlacements *map[string]int16,
	guruAspects *[]models.AllRaashis,
	shaniAspects *[]models.AllRaashis,
	kujaAspects *[]models.AllRaashis,
) []string {
	// Create a variable that collects all the effects of the first house
	firstHouseEffects := make([]string, 0)

	// Get info
	ascendantLord := constants.RaashyadhipatiMapStore[(*reqBody).Ascendant]
	ascendantLordPlacement := (*housePlacements)[string(ascendantLord)]
	ascendantLordHouse := utils.GetHouseName((*reqBody).Ascendant, ascendantLordPlacement)
	doesAscendantLordHaveGuruAspect := slices.Contains(*guruAspects, ascendantLordHouse)
	doesAscendantLordHaveShukraAspect := constants.OppositeAspectsStore[models.AllRaashis((*reqBody).ShukraPlacement.Placement)] == ascendantLordHouse

	// Start appending to string array for lagna placements
	if utils.IsDusthana(ascendantLordPlacement) || 
	(ascendantLordPlacement == (*housePlacements)["Shani"] && ascendantLord != models.Raashyadhipati("Shani")) || 
	(ascendantLordPlacement == (*housePlacements)["Kuja"] && ascendantLord != models.Raashyadhipati("Kuja")) || 
	(ascendantLordPlacement == (*housePlacements)["Surya"] && ascendantLord != models.Raashyadhipati("Surya")) {
		firstHouseEffects = append(firstHouseEffects, "Since the ascendant lord is in a dusthana, physical pleasure will diminish")
	}

	if utils.IsKendra(ascendantLordPlacement) || utils.IsKona(ascendantLordPlacement) {
		firstHouseEffects = append(firstHouseEffects, "Since the ascendant lord is in a kendra/kona, physical pleasure will be available all the time")
	}

	if constants.DebilitationMapStore[ascendantLord] == ascendantLordHouse {
		if (((*housePlacements)["Guru"] == ascendantLordPlacement && ascendantLord != models.Raashyadhipati("Guru")) || ((*housePlacements)["Shukra"] == ascendantLordPlacement && ascendantLord != models.Raashyadhipati("Shukra"))) && (utils.IsKendra(ascendantLordPlacement) || utils.IsKona(ascendantLordPlacement)) {
			firstHouseEffects = append(firstHouseEffects, "There will be diseases as ascendant lord is neech, but they will disappear due to Guru/Shukra")
		} else if (doesAscendantLordHaveGuruAspect || doesAscendantLordHaveShukraAspect) {
			firstHouseEffects = append(firstHouseEffects, "There will be diseases as ascendant lord is neech, but they will disappear due to Guru/Shukra")
		} else {
			firstHouseEffects = append(firstHouseEffects, "There will be diseases as ascendant lord is neech")
		}
	}
	if !slices.Contains([]models.Raashyadhipati{constants.Surya, constants.Chandra}, ascendantLord) {
		if utils.IsGrahaCombust(
			ascendantLord,
			(*reqBody).KujaPlacement.Combustion,
			(*reqBody).ShukraPlacement.Combustion,
			(*reqBody).BudhaPlacement.Combustion,
			(*reqBody).GuruPlacement.Combustion,
			(*reqBody).ShaniPlacement.Combustion,
		) {
			if (((*housePlacements)["Guru"] == ascendantLordPlacement && ascendantLord != models.Raashyadhipati("Guru")) || ((*housePlacements)["Shukra"] == ascendantLordPlacement && ascendantLord != models.Raashyadhipati("Shukra"))) && (utils.IsKendra(ascendantLordPlacement) || utils.IsKona(ascendantLordPlacement)) {
				firstHouseEffects = append(firstHouseEffects, "There will be diseases as ascendant lord is combust, but they will disappear due to Guru/Shukra")
			} else if (doesAscendantLordHaveGuruAspect || doesAscendantLordHaveShukraAspect) {
				firstHouseEffects = append(firstHouseEffects, "There will be diseases as ascendant lord is combust, but they will disappear due to Guru/Shukra")
			} else {
				firstHouseEffects = append(firstHouseEffects, "There will be diseases as ascendant lord is combust")
			}
		}
	}
	if slices.Contains(constants.EnemySignsMappingStore[ascendantLord], ascendantLordHouse) {
		if (((*housePlacements)["Guru"] == ascendantLordPlacement && ascendantLord != models.Raashyadhipati("Guru")) || ((*housePlacements)["Shukra"] == ascendantLordPlacement && ascendantLord != models.Raashyadhipati("Shukra"))) && (utils.IsKendra(ascendantLordPlacement) || utils.IsKona(ascendantLordPlacement)) {
			firstHouseEffects = append(firstHouseEffects, "There will be diseases as ascendant lord is in enemy sign, but they will disappear due to Guru/Shukra")
		} else if (doesAscendantLordHaveGuruAspect || doesAscendantLordHaveShukraAspect) {
			firstHouseEffects = append(firstHouseEffects, "There will be diseases as ascendant lord is in enemy sign, but they will disappear due to Guru/Shukra")
		} else {
			firstHouseEffects = append(firstHouseEffects, "There will be diseases as ascendant lord is in enemy sign")
		}
	}

	if !slices.Contains(*guruAspects, models.AllRaashis((*reqBody).Ascendant)) && models.AllRaashis((*reqBody).Ascendant) != constants.OppositeAspectsStore[models.AllRaashis((*reqBody).ShukraPlacement.Placement)] {
		if models.AllRaashis((*reqBody).GuruPlacement.Placement) != models.AllRaashis((*reqBody).Ascendant) && models.AllRaashis((*reqBody).ShukraPlacement.Placement) != models.AllRaashis((*reqBody).Ascendant) {
			if slices.Contains(*kujaAspects, models.AllRaashis((*reqBody).Ascendant)) || slices.Contains(*shaniAspects, models.AllRaashis((*reqBody).Ascendant)) || (*reqBody).ShaniPlacement.Placement == (*reqBody).Ascendant || (*reqBody).KujaPlacement.Placement == (*reqBody).Ascendant {
				firstHouseEffects = append(firstHouseEffects, "Since the ascendant is aspected/conjunct Kuja or Shani or Surya, the native will not have bodily health")
			}
		}
	}

	if (*housePlacements)["Guru"] == 1 || (*housePlacements)["Shukra"] == 1 || (*housePlacements)["Shukra"] == 7 {
		firstHouseEffects = append(firstHouseEffects, "The native will be embodied with bodily beauty and enjoy pleasures of the body due to Guru/Shukra aspect/conjunct on the ascendant")
	}

	if slices.Contains(*guruAspects, models.AllRaashis((*reqBody).Ascendant)) {
		firstHouseEffects = append(firstHouseEffects, "The native will be embodied with bodily beauty and enjoy pleasures of the body due to Guru/Shukra aspect/conjunct on the ascendant")
	}

	if constants.RaashiQualityStore[(*reqBody).Ascendant] == constants.Chara && slices.Contains(*guruAspects, models.AllRaashis((*reqBody).Ascendant)) {
		firstHouseEffects = append(firstHouseEffects, "Fame, wealth, abundant pleasures and comforts of the body will be acquired as the ascendant is Chara and aspected by Guru")
	}

	if slices.Contains([]int16{4, 7, 10}, (*housePlacements)["Budha"]) {
		firstHouseEffects = append(firstHouseEffects, "Since Budha is in a Kendra, the native will enjoy royal fortunes")
	}
	if slices.Contains([]int16{4, 7, 10}, (*housePlacements)["Shukra"]) {
		firstHouseEffects = append(firstHouseEffects, "Since Shukra is in a Kendra, the native will enjoy royal fortunes")
	}
	if slices.Contains([]int16{4, 7, 10}, (*housePlacements)["Guru"]) {
		firstHouseEffects = append(firstHouseEffects, "Since Guru is in a Kendra, the native will enjoy royal fortunes")
	}

	// Return all the effects of the first house
	return firstHouseEffects
}