package services

import (
	"slices"

	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
	"github.com/scientist-v08/bphs/utils"
)

func isBudhaAspecting(lagnaLordHouse models.AllRaashis, budhaHouse models.AllRaashis) bool {
	switch lagnaLordHouse {
    case constants.Mesha:
        return budhaHouse == constants.Tula

    case constants.Vruschika:
        return budhaHouse == constants.Vrushabha

    case constants.Mithuna:
        return budhaHouse == constants.Dhanassu

    case constants.Kanya:
        return budhaHouse == constants.Meena

    default:
        return false
    }
}

func relationWithNative(placement int16) string {
	switch placement {
	case 1:
		return "the native"
	case 3:
		return "the native's younger siblings"
	case 4:
		return "the native's mother"
	case 5:
		return "the natives's children"
	case 7:
		return "the native's wife"
	case 9:
		return "the native's father"
	default:
		return ""
	}
}

func SixthHouseEffects(
	reqBody *models.HouseEffectsReqBody,
	housePlacements *map[string]int16,
	guruAspects *[]models.AllRaashis,
	shaniAspects *[]models.AllRaashis,
	kujaAspects *[]models.AllRaashis,
) []string {
	// Create a variable that collects all the effects of the second house
	sixthHouseEffects := make([]string, 0)

	// Get info
	sixthHouse := constants.ReverseNumericalMappingsStore[(*reqBody).Ascendant][6]
	sixthLord := constants.RaashyadhipatiMapStore[string(sixthHouse)]
	sixthLordPlacement := (*housePlacements)[string(sixthLord)]
	sixthLordplacementEffectFor := relationWithNative(sixthLordPlacement)
	eigthHouse := constants.ReverseNumericalMappingsStore[(*reqBody).Ascendant][8]
	eigthLord := constants.RaashyadhipatiMapStore[string(eigthHouse)]
	eigthLordPlacement := (*housePlacements)[string(eigthLord)]
	budhaPlacement := (*housePlacements)["Budha"];
	budhaHouse := utils.GetHouseName((*reqBody).Ascendant, budhaPlacement)
	ascendantLord := constants.RaashyadhipatiMapStore[(*reqBody).Ascendant]
	ascendantLordPlacement := (*housePlacements)[string(ascendantLord)]
	ascendantLordHouse := utils.GetHouseName((*reqBody).Ascendant, ascendantLordPlacement)

	if slices.Contains([]int16{1,6,8}, sixthLordPlacement) {
		sixthHouseEffects = append(sixthHouseEffects, "The native will get ulcers or bruises on the body")
	}

	if slices.Contains([]models.AllRaashis{constants.Mesha, constants.Kanya, constants.Mithuna, constants.Vruschika}, ascendantLordHouse) {
		if isBudhaAspecting(ascendantLordHouse, budhaHouse) {
			sixthHouseEffects = append(sixthHouseEffects, "The native will get facial diseases")
		}
	}

	if eigthLordPlacement == sixthLordPlacement {
		if slices.Contains([]int16{1,3,4,5,7,9}, sixthLordPlacement) {
			if (*housePlacements)["Surya"] == sixthLordPlacement && eigthLord != constants.Surya && sixthLord != constants.Surya {
				sixthHouseEffects = append(sixthHouseEffects, "Tumours, fever will affect " + sixthLordplacementEffectFor)
			}
			if (*housePlacements)["Budha"] == sixthLordPlacement && eigthLord != constants.Budha && sixthLord != constants.Budha {
				sixthHouseEffects = append(sixthHouseEffects, "Bilious diseases will affect " + sixthLordplacementEffectFor)
			}
			if (*housePlacements)["Shukra"] == sixthLordPlacement && eigthLord != constants.Shukra && sixthLord != constants.Shukra {
				sixthHouseEffects = append(sixthHouseEffects, "Diseases caused by sexual union will affect " + sixthLordplacementEffectFor)
			}
			if (*housePlacements)["Chandra"] == sixthLordPlacement && eigthLord != constants.Chandra && sixthLord != constants.Chandra {
				sixthHouseEffects = append(sixthHouseEffects, "Drwoning or respiratory related diseases will affect " + sixthLordplacementEffectFor)
			}
			if (*housePlacements)["Kuja"] == sixthLordPlacement && eigthLord != constants.Kuja && sixthLord != constants.Kuja {
				sixthHouseEffects = append(sixthHouseEffects, "Diseases of blood vessels, hits or wounds will affect " + sixthLordplacementEffectFor)
			}
			if (*housePlacements)["Guru"] == sixthLordPlacement && eigthLord != constants.Guru && sixthLord != constants.Guru {
				sixthHouseEffects = append(sixthHouseEffects, sixthLordplacementEffectFor + " will be free of diseases")
			}
			if (*housePlacements)["Shani"] == sixthLordPlacement && eigthLord != constants.Shani && sixthLord != constants.Shani {
				sixthHouseEffects = append(sixthHouseEffects, "Windy diseases will affect " + sixthLordplacementEffectFor)
			}
			if (*housePlacements)["RahuPlacement"] == sixthLordPlacement {
				sixthHouseEffects = append(sixthHouseEffects, "Danger from criminals for " + sixthLordplacementEffectFor)
			}
			if (*housePlacements)["KetuPlacement"] == sixthLordPlacement {
				sixthHouseEffects = append(sixthHouseEffects, "Diseases of the navel will affect " + sixthLordplacementEffectFor)
			}
		}
	}

	return sixthHouseEffects
}