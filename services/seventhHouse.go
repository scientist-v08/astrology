package services

import (
	"slices"

	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
	"github.com/scientist-v08/bphs/utils"
)

func SeventhHouseEffects(
	reqBody *models.HouseEffectsReqBody,
	housePlacements *map[string]int16,
	guruAspects *[]models.AllRaashis,
	shaniAspects *[]models.AllRaashis,
	kujaAspects *[]models.AllRaashis,
) []string {
	// Create a variable that collects all the effects of the second house
	seventhHouseEffects := make([]string, 0)

	// Get info
	seventhHouse := constants.ReverseNumericalMappingsStore[(*reqBody).Ascendant][7]
	seventhLord := constants.RaashyadhipatiMapStore[string(seventhHouse)]
	seventhLordPlacementNumber := (*housePlacements)[string(seventhLord)]
	seventhLordHousePlacementRaashi := utils.GetHouseName((*reqBody).Ascendant, seventhLordPlacementNumber)
	shukraPlacementNumber := (*housePlacements)["Shukra"]
	shkraHousePlacementRaashi := utils.GetHouseName((*reqBody).Ascendant, shukraPlacementNumber)
	suryaPlacementNumber := (*housePlacements)["Surya"]
	budhaPlacementNumber := (*housePlacements)["Budha"]
	chandraPlacementRaashi := utils.GetHouseName((*reqBody).Ascendant, budhaPlacementNumber)
	chandraPlacementNumber := (*housePlacements)["Chandra"]
	budhaPlacementRaashi := utils.GetHouseName((*reqBody).Ascendant, chandraPlacementNumber)
	shaniPlacementNumber := (*housePlacements)["Shani"]
	kujaPlacementNumber := (*housePlacements)["Kuja"]
	isShukraAfflicted := shukraPlacementNumber == suryaPlacementNumber ||
		shukraPlacementNumber == shaniPlacementNumber ||
		shukraPlacementNumber == kujaPlacementNumber ||
		(shukraPlacementNumber == budhaPlacementNumber && !(*reqBody).BudhaUnafflicted) ||
		(shukraPlacementNumber == chandraPlacementNumber && !(*reqBody).ChandraWaxing)
	isSeventhLordAspectedByBenefic := (slices.Contains(*guruAspects, seventhLordHousePlacementRaashi) && seventhLord != constants.Guru) ||
		(seventhLord != constants.Shukra && constants.OppositeAspectsStore[shkraHousePlacementRaashi] == seventhLordHousePlacementRaashi) ||
		(seventhLord != constants.Chandra && constants.OppositeAspectsStore[chandraPlacementRaashi] == seventhLordHousePlacementRaashi && (*reqBody).ChandraWaxing) ||
		(seventhLord != constants.Budha && constants.OppositeAspectsStore[budhaPlacementRaashi] == seventhLordHousePlacementRaashi && (*reqBody).BudhaUnafflicted)
	isSeventhLordConjunctBenefic := (seventhLordPlacementNumber == (*housePlacements)["Guru"] && seventhLord != constants.Guru) ||
		(seventhLord != constants.Shukra && shukraPlacementNumber == seventhLordPlacementNumber) ||
		(seventhLord != constants.Chandra && chandraPlacementNumber == seventhLordPlacementNumber && (*reqBody).ChandraWaxing) ||
		(seventhLord != constants.Budha && budhaPlacementNumber == seventhLordPlacementNumber && (*reqBody).BudhaUnafflicted)
	isMaleficInTwelfth := shaniPlacementNumber == 12 ||
		kujaPlacementNumber == 12 ||
		(budhaPlacementNumber == 12 && !(*reqBody).BudhaUnafflicted) ||
		suryaPlacementNumber == 12 ||
		(chandraPlacementNumber == 12 && !(*reqBody).ChandraWaxing)
	isMaleficInSeventh := shaniPlacementNumber == 7 ||
		kujaPlacementNumber == 7 ||
		(budhaPlacementNumber == 7 && !(*reqBody).BudhaUnafflicted) ||
		suryaPlacementNumber == 7 ||
		(chandraPlacementNumber == 7 && !(*reqBody).ChandraWaxing)
	isAfflictedChandraInFive := chandraPlacementNumber == 5 && !(*reqBody).ChandraWaxing
	secondHouse := constants.ReverseNumericalMappingsStore[(*reqBody).Ascendant][2]
	secondLord := constants.RaashyadhipatiMapStore[string(secondHouse)]
	secondLordPlacement := (*housePlacements)[string(secondLord)]

	// Start appending
	if seventhLordPlacementNumber == 7 || constants.ExaltationMapStore[seventhLord] == seventhLordHousePlacementRaashi {
		seventhHouseEffects = append(seventhHouseEffects, "The native will derive full happiness thorugh his wife and marriage")
	}

	if shukraPlacementNumber == 2 && (secondLord != constants.Kuja && secondLordPlacement == kujaPlacementNumber) {
		seventhHouseEffects = append(seventhHouseEffects, "The mairrage of the native is likely to be delayed")
	}
	if shukraPlacementNumber == 5 && slices.Contains([]int16{5,9}, (*housePlacements)["RahuPlacement"]) {
		seventhHouseEffects = append(seventhHouseEffects, "The mairrage of the native is likely to be delayed")
	}

	if constants.ExaltationMapStore[seventhLord] != seventhLordHousePlacementRaashi && !slices.Contains(constants.GrahaOwnedRaashiMapStore[seventhLord], seventhLordHousePlacementRaashi) {
		if utils.IsDusthana(seventhLordPlacementNumber) {
			seventhHouseEffects = append(seventhHouseEffects, "The native's spouse maybe sickly depending on the strength of the 7th lord")
			seventhHouseEffects = append(seventhHouseEffects, "The native's spouse is likely to insult the native")
		}
	}
	if constants.ExaltationMapStore[constants.Shukra] != shkraHousePlacementRaashi && !slices.Contains(constants.GrahaOwnedRaashiMapStore[constants.Shukra], shkraHousePlacementRaashi) {
		if utils.IsDusthana(shukraPlacementNumber) {
			seventhHouseEffects = append(seventhHouseEffects, "The native's spouse maybe sickly depending on the strength of Venus")
		}
	}

	if isShukraAfflicted {
		seventhHouseEffects = append(seventhHouseEffects, "The native will face problems in marriage/relationships or may face problems in marital life")
	}
	if utils.IsLordCombust(seventhLord, reqBody) || slices.Contains(constants.EnemySignsMappingStore[seventhLord], seventhLordHousePlacementRaashi) || constants.DebilitationMapStore[seventhLord] == seventhLordHousePlacementRaashi {
		seventhHouseEffects = append(seventhHouseEffects, "The native will face problems in marriage/relationships or may face problems in marital life")
	}

	if constants.ExaltationMapStore[seventhLord] == seventhLordHousePlacementRaashi {
		seventhHouseEffects = append(seventhHouseEffects, "The native is likely to get multiple spouses")
	}

	if slices.Contains([]models.AllRaashis{constants.Tula, constants.Vrushabha, constants.Kumbha, constants.Makara}, seventhLordHousePlacementRaashi) &&
		(isSeventhLordAspectedByBenefic || isSeventhLordConjunctBenefic) {
		seventhHouseEffects = append(seventhHouseEffects, "The native is likely to get multiple spouses")
	}

	if suryaPlacementNumber == 7 {
		seventhHouseEffects = append(seventhHouseEffects, "The native will hate the opposite gender and is likely to befriend the opposite gender with the intention to have sex with them")
	}

	if kujaPlacementNumber == 7 || budhaPlacementNumber == 7 || (*housePlacements)["Guru"] == 7 || shaniPlacementNumber == 7 || shukraPlacementNumber == 7 || chandraPlacementNumber == 7 || (*housePlacements)["RahuPlacement"] == 7 || (*housePlacements)["KetuPlacement"] == 7 {
		seventhHouseEffects = append(seventhHouseEffects, "The native is likely to befriend the opposite gender with the intention to have sex with them")
	}

	if isMaleficInSeventh && isMaleficInTwelfth && isAfflictedChandraInFive {
		seventhHouseEffects = append(seventhHouseEffects, "The spouse of the native will be the one to dominate the relationship and the spouse is likely to hate the family of the native")
	}

	if kujaPlacementNumber == 7 {
		seventhHouseEffects = append(seventhHouseEffects, "The spouse is likely to be of questionable character due to Kuja in the 7th")
	}
	if shaniPlacementNumber == 7 {
		seventhHouseEffects = append(seventhHouseEffects, "The spouse is likely to be of questionable character due to Shani in the 7th")
	}

	return seventhHouseEffects
}