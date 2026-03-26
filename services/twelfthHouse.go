package services

import (
	"slices"

	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
	"github.com/scientist-v08/bphs/utils"
)

func lordNavamsha(lord models.Raashyadhipati, req *models.HouseEffectsReqBody) string {
	switch lord {
	case constants.Budha:
		return (*req).NavBudhaPlacement
	case constants.Shani:
		return (*req).NavShaniPlacement
	case constants.Shukra:
		return (*req).NavShukraPlacement
	case constants.Surya:
		return (*req).NavSuryaPlacement
	case constants.Chandra:
		return (*req).NavChandraPlacement
	case constants.Kuja:
		return (*req).NavKujaPlacement
	case constants.Guru:
		return (*req).NavGuruPlacement
	default:
		return ""
	}
}

func lordNavamshaPlacement(lord models.Raashyadhipati, req *map[string]int16) int16 {
	switch lord {
	case constants.Budha:
		return (*req)["NavBudha"]
	case constants.Shani:
		return (*req)["NavShani"]
	case constants.Shukra:
		return (*req)["NavShukra"]
	case constants.Surya:
		return (*req)["NavSurya"]
	case constants.Chandra:
		return (*req)["NavChandra"]
	case constants.Kuja:
		return (*req)["NavKuja"]
	case constants.Guru:
		return (*req)["NavGuru"]
	default:
		return 0
	}
}

func TwelfthHouseEffects(
	reqBody *models.HouseEffectsReqBody,
	housePlacements *map[string]int16,
	guruAspects *[]models.AllRaashis,
	shaniAspects *[]models.AllRaashis,
	kujaAspects *[]models.AllRaashis,
) []string {
	// Create a variable to collect the effects of the 12th house
	twelfthHouseEffects := make([]string, 0)

	// Get info
	twelfthHouse := constants.ReverseNumericalMappingsStore[(*reqBody).Ascendant][12]
	twelfthLord := constants.RaashyadhipatiMapStore[string(twelfthHouse)]
	twelfthLordPlacementNumber := (*housePlacements)[string(twelfthLord)]
	twelfthLordNatalRaashi := utils.GetHouseName((*reqBody).Ascendant, twelfthLordPlacementNumber)
	guruPlacement := (*housePlacements)["Guru"]
	budhPlacement := (*housePlacements)["Budha"]
	chandraPlacement := (*housePlacements)["Chandra"]
	chandraNatalRaashi := utils.GetHouseName((*reqBody).Ascendant, chandraPlacement)
	shukraPlacement := (*housePlacements)["Shukra"]
	isTwelfthLordWithBenefic := twelfthLord != constants.Guru && guruPlacement == twelfthLordPlacementNumber ||
		twelfthLord != constants.Shukra && shukraPlacement == twelfthLordPlacementNumber ||
		twelfthLord != constants.Budha && budhPlacement == twelfthLordPlacementNumber && (*reqBody).BudhaUnafflicted ||
		twelfthLord != constants.Chandra && chandraPlacement == twelfthLordPlacementNumber && (*reqBody).ChandraWaxing
	isABeneficInTwelve := guruPlacement == 12 ||
		shukraPlacement == 12 ||
		budhPlacement == 12 && (*reqBody).BudhaUnafflicted ||
		chandraPlacement == 12 && (*reqBody).ChandraWaxing
	isTwelfthLordInOwnHouse := slices.Contains(constants.GrahaOwnedRaashiMapStore[twelfthLord], twelfthHouse)
	ascendantLord := constants.RaashyadhipatiMapStore[(*reqBody).Ascendant]
	ascendantLordPlacement := (*housePlacements)[string(ascendantLord)]
	isTwelfthLordExalted := constants.ExaltationMapStore[twelfthLord] == twelfthLordNatalRaashi
	isChandraExaltedOrInOwnRaashiOrOwnNavamsha := slices.Contains([]models.AllRaashis{constants.ChandraUccha, constants.Karkataka}, chandraNatalRaashi) || (*reqBody).NavChandraPlacement == string(constants.Karkataka)

	// Start appending
	twelfthHouseEffects = append(twelfthHouseEffects, "The 12th house themes will be prominant in the life of the native")

	if isTwelfthLordWithBenefic || isTwelfthLordExalted || isTwelfthLordInOwnHouse || isABeneficInTwelve {
		twelfthHouseEffects = append(twelfthHouseEffects, "There will be expenses on good accounts")
	}
	if twelfthLord == constants.Chandra && 
	(isChandraExaltedOrInOwnRaashiOrOwnNavamsha || slices.Contains([]int16{11,9,5}, chandraPlacement)) {
		twelfthHouseEffects = append(twelfthHouseEffects, "The native will own sandalwood or sandalwood like substances, will own a house and a nice bed")
	}
	
	if slices.Contains([]int16{6,8} ,twelfthLordPlacementNumber) || 
	slices.Contains(constants.EnemySignsMappingStore[twelfthLord], models.AllRaashis(lordNavamsha(twelfthLord, reqBody))) ||
	constants.DebilitationMapStore[twelfthLord] == models.AllRaashis(lordNavamsha(twelfthLord, reqBody)) ||
	utils.IsDusthana(lordNavamshaPlacement(twelfthLord, housePlacements)) {
		twelfthHouseEffects = append(twelfthHouseEffects, "The native will be devoid of happiness from wife, be troubled by expenses")
	}
	
	if utils.IsKendra(twelfthLordPlacementNumber) || utils.IsKona(twelfthLordPlacementNumber) {
		twelfthHouseEffects = append(twelfthHouseEffects, "The native will beget a spouse")
	}
	
	if ascendantLordPlacement == 12 && twelfthLordPlacementNumber == 1 && (*housePlacements)["Shukra"] == twelfthLordPlacementNumber {
		twelfthHouseEffects = append(twelfthHouseEffects, "The native will spend money on religious grounds")
	}

	// Return all the results
	return twelfthHouseEffects
}