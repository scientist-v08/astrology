package services

import (
	"slices"

	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
	"github.com/scientist-v08/bphs/utils"
)

func isTenthOccupiedByNeechGraha(tenthHouse models.AllRaashis, surya int16, shukra int16, chandra int16, guru int16, budha int16, kuja int16) bool {
	switch tenthHouse {

	case constants.Tula: // Surya Neecha
		return surya == 10

	case constants.Vruschika: // Chandra Neecha
		return chandra == 10

	case constants.Karkataka: // Kuja Neecha
		return kuja == 10

	case constants.Meena: // Budha Neecha
		return budha == 10

	case constants.Makara: // Guru Neecha
		return guru == 10

	case constants.Kanya: // Shukra Neecha
		return shukra == 10

	default:
		return false
	}
}

func isTenthLordInFall(houseOccupiedByTenthLord models.AllRaashis, tenthLord models.Raashyadhipati) bool {
	switch tenthLord {

	case constants.Surya:
		return houseOccupiedByTenthLord == constants.SuryaNeecha

	case constants.Chandra:
		return houseOccupiedByTenthLord == constants.ChandraNeecha

	case constants.Kuja:
		return houseOccupiedByTenthLord == constants.KujaNeecha

	case constants.Budha:
		return houseOccupiedByTenthLord == constants.BudhaNeecha

	case constants.Guru:
		return houseOccupiedByTenthLord == constants.GuruNeecha

	case constants.Shukra:
		return houseOccupiedByTenthLord == constants.ShukraNeecha

	case constants.Shani:
		return houseOccupiedByTenthLord == constants.ShaniNeecha

	default:
		return false
	}
}

func TenthHouseEffects(
	reqBody *models.HouseEffectsReqBody,
	housePlacements *map[string]int16,
	guruAspects *[]models.AllRaashis,
	shaniAspects *[]models.AllRaashis,
	kujaAspects *[]models.AllRaashis,
) []string {
	// Create a variable that collects all the effects of the second house
	tenthHouseEffects := make([]string, 0)

	// Get info
	ascendantLord := constants.RaashyadhipatiMapStore[(*reqBody).Ascendant]
	ascendantLordPlacement := (*housePlacements)[string(ascendantLord)]
	tenthHouse := constants.ReverseNumericalMappingsStore[(*reqBody).Ascendant][10]
	tenthLord := constants.RaashyadhipatiMapStore[string(tenthHouse)]
	tenthLordPlacementNumber := (*housePlacements)[string(tenthLord)]
	eleventhHouse := constants.ReverseNumericalMappingsStore[(*reqBody).Ascendant][11]
	eleventhLord := constants.RaashyadhipatiMapStore[string(eleventhHouse)]
	eleventhLordPlacementNumber := (*housePlacements)[string(eleventhLord)]
	tenthLordHouse := utils.GetHouseName((*reqBody).Ascendant, tenthLordPlacementNumber)
	ninthHouse := constants.ReverseNumericalMappingsStore[(*reqBody).Ascendant][9]
	ninthLord := constants.RaashyadhipatiMapStore[string(ninthHouse)]
	ninthLordPlacementNumber := (*housePlacements)[string(ninthLord)]
	seventhHouse := constants.ReverseNumericalMappingsStore[(*reqBody).Ascendant][7]
	seventhLord := constants.RaashyadhipatiMapStore[string(seventhHouse)]
	seventhLordPlacementNumber := (*housePlacements)[string(seventhLord)]
	eighthHouse := constants.ReverseNumericalMappingsStore[(*reqBody).Ascendant][8]
	eighthLord := constants.RaashyadhipatiMapStore[string(eighthHouse)]
	eighthLordPlacementNumber := (*housePlacements)[string(eighthLord)]
	isTenthLordExalted := constants.ExaltationMapStore[tenthLord] == tenthLordHouse
	isTenthLordInOwnRaashi := slices.Contains(constants.GrahaOwnedRaashiMapStore[tenthLord], tenthLordHouse)
	shukraPlacementNumber := (*housePlacements)["Shukra"]
	chandraPlacementNumber := (*housePlacements)["Chandra"]
	budhaPlacementNumber := (*housePlacements)["Budha"]
	shaniPlacement := (*housePlacements)["Shani"]
	kujaPlacement := (*housePlacements)["Kuja"]
	guruPlacement := (*housePlacements)["Guru"]
	suryaPlacement := (*housePlacements)["Surya"]
	isTenthLordConjunctBenefic := (tenthLordPlacementNumber == (*housePlacements)["Guru"] && tenthLord != constants.Guru) ||
		(tenthLord != constants.Shukra && shukraPlacementNumber == tenthLordPlacementNumber) ||
		(tenthLord != constants.Chandra && chandraPlacementNumber == tenthLordPlacementNumber && (*reqBody).ChandraWaxing) ||
		(tenthLord != constants.Budha && budhaPlacementNumber == tenthLordPlacementNumber && (*reqBody).BudhaUnafflicted)
	isTenthLordConjunctMalefic := (tenthLordPlacementNumber == suryaPlacement && tenthLord != constants.Surya) ||
		(tenthLord != constants.Kuja && (*housePlacements)["Kuja"] == tenthLordPlacementNumber) ||
		(tenthLord != constants.Shani && (*housePlacements)["Shani"] == tenthLordPlacementNumber) ||
		(tenthLord != constants.Chandra && chandraPlacementNumber == tenthLordPlacementNumber && !(*reqBody).ChandraWaxing) ||
		(tenthLord != constants.Budha && budhaPlacementNumber == tenthLordPlacementNumber && !(*reqBody).BudhaUnafflicted)
	isSeventhLordConjunctMalefic := (seventhLordPlacementNumber == suryaPlacement && seventhLord != constants.Surya) ||
		(seventhLord != constants.Kuja && (*housePlacements)["Kuja"] == seventhLordPlacementNumber) ||
		(seventhLord != constants.Shani && (*housePlacements)["Shani"] == seventhLordPlacementNumber) ||
		(seventhLord != constants.Chandra && chandraPlacementNumber == seventhLordPlacementNumber && !(*reqBody).ChandraWaxing) ||
		(seventhLord != constants.Budha && budhaPlacementNumber == seventhLordPlacementNumber && !(*reqBody).BudhaUnafflicted)
	isEighthLordConjunctMalefic := (eighthLordPlacementNumber == suryaPlacement && eighthLord != constants.Surya) ||
		(eighthLord != constants.Kuja && (*housePlacements)["Kuja"] == eighthLordPlacementNumber) ||
		(eighthLord != constants.Shani && (*housePlacements)["Shani"] == eighthLordPlacementNumber) ||
		(eighthLord != constants.Chandra && chandraPlacementNumber == eighthLordPlacementNumber && !(*reqBody).ChandraWaxing) ||
		(eighthLord != constants.Budha && budhaPlacementNumber == eighthLordPlacementNumber && !(*reqBody).BudhaUnafflicted)
	isSeventhOccupiedByMalefic := suryaPlacement == 7 || kujaPlacement == 7 || shaniPlacement == 7 || (budhaPlacementNumber == 7 && !(*reqBody).BudhaUnafflicted) || (chandraPlacementNumber == 7 && !(*reqBody).ChandraWaxing)
	isTenthOccupiedByMalefic := suryaPlacement == 10 || kujaPlacement == 10 || shaniPlacement == 10 || (budhaPlacementNumber == 10 && !(*reqBody).BudhaUnafflicted) || (chandraPlacementNumber == 10 && !(*reqBody).ChandraWaxing)
	isEleventhOccupiedByMalefic := suryaPlacement == 11 || kujaPlacement == 11 || shaniPlacement == 11 || (budhaPlacementNumber == 11 && !(*reqBody).BudhaUnafflicted) || (chandraPlacementNumber == 11 && !(*reqBody).ChandraWaxing)
	isTenthLordNotStrong := !utils.IsGrahaStrong(
		!utils.IsLordCombust(tenthLord, reqBody),
		(constants.DebilitationMapStore[tenthLord] != tenthLordHouse),
		!slices.Contains(constants.EnemySignsMappingStore[tenthLord], tenthLordHouse),
		!utils.IsDusthana(tenthLordPlacementNumber),
	)
	guruPlacementHouse := utils.GetHouseName((*reqBody).Ascendant, guruPlacement)
	tenthLordInConjunctionWithGuru := tenthLordPlacementNumber == guruPlacement
	tenthLordInAspectOfGuru := slices.Contains(*guruAspects, tenthLordHouse)
	tenthLordIsEitherConjunctOrAspectGuru := tenthLordInAspectOfGuru || tenthLordInConjunctionWithGuru
	isNavTenthOccupiedByMalefic := (*housePlacements)["NavSurya"] == 10 ||
		(*housePlacements)["NavKuja"] == 10 ||
		(*housePlacements)["NavShani"] == 10 ||
		((*housePlacements)["NavChandra"] == 10 && !(*reqBody).ChandraWaxing) ||
		((*housePlacements)["NavBudha"] == 10 && !(reqBody).BudhaUnafflicted)

	// Start appending
	tenthHouseEffects = append(tenthHouseEffects, "The native will see good results in 10th house themes")

	if isTenthLordExalted || isTenthLordInOwnRaashi {
		tenthHouseEffects = append(tenthHouseEffects, "If the 10th lord is strong (check shadbala) then the native will enjoy complete paternal happiness and will obtain fame and will have a good job in life")
	}

	if isTenthLordNotStrong {
		tenthHouseEffects = append(tenthHouseEffects, "The native will face obstructions in his work")
	}

	if utils.IsKendra((*housePlacements)["RahuPlacement"]) || utils.IsKona((*housePlacements)["RahuPlacement"]) {
		tenthHouseEffects = append(tenthHouseEffects, "The native will perform great religious activities")
	}

	if isTenthLordConjunctBenefic || slices.Contains([]models.AllRaashis{constants.Meena, constants.Dhanassu, constants.Vrushabha, constants.Tula}, tenthLordHouse) {
		tenthHouseEffects = append(tenthHouseEffects, "The native will gain through royal patronage and business")
	}
	if tenthLord != constants.Chandra && (*reqBody).ChandraWaxing && tenthLordHouse == constants.Karkataka {
		tenthHouseEffects = append(tenthHouseEffects, "The native will gain through royal patronage and business")
	}
	if tenthLord != constants.Budha && (*reqBody).BudhaUnafflicted && (tenthLordHouse == constants.Mithuna || tenthLordHouse == constants.Kanya) {
		tenthHouseEffects = append(tenthHouseEffects, "The native will gain through royal patronage and business")
	}

	if isTenthLordConjunctMalefic || slices.Contains([]models.AllRaashis{constants.Kumbha, constants.Makara, constants.Vruschika, constants.Mesha, constants.Simha}, tenthLordHouse) {
		tenthHouseEffects = append(tenthHouseEffects, "The native will lose through those in leadership positions and in business")
	}
	if tenthLord != constants.Chandra && !(*reqBody).ChandraWaxing && tenthLordHouse == constants.Karkataka {
		tenthHouseEffects = append(tenthHouseEffects, "The native will lose through those in leadership positions and in business")
	}
	if tenthLord != constants.Budha && !(*reqBody).BudhaUnafflicted && (tenthLordHouse == constants.Mithuna || tenthLordHouse == constants.Kanya) {
		tenthHouseEffects = append(tenthHouseEffects, "The native will lose through those in leadership positions and in business")
	}

	if isTenthOccupiedByMalefic && isEleventhOccupiedByMalefic {
		tenthHouseEffects = append(tenthHouseEffects, "The native will obtain only bad jobs and will hate his co-workers")
	}

	if tenthLordPlacementNumber == 8 && (*housePlacements)["RahuPlacement"] == 8 {
		tenthHouseEffects = append(tenthHouseEffects, "The native will hate his co workers and be a fool and do bad jobs")
	}

	if shaniPlacement == 7 && kujaPlacement == 7 && isSeventhLordConjunctMalefic && tenthLordPlacementNumber == 7 {
		tenthHouseEffects = append(tenthHouseEffects, "The native will be fond of sex and eating a lot of food")
	}

	if isTenthLordExalted && tenthLordPlacementNumber == guruPlacement && ninthLordPlacementNumber == 10 {
		tenthHouseEffects = append(tenthHouseEffects, "The native will be endowed with honour, wealth and valour")
	}

	if eleventhLordPlacementNumber == 10 && tenthLordPlacementNumber == 1 {
		tenthHouseEffects = append(tenthHouseEffects, "The native will lead a happy life")
	}

	if eleventhLordPlacementNumber == tenthLordPlacementNumber && utils.IsKendra(tenthLordPlacementNumber) {
		tenthHouseEffects = append(tenthHouseEffects, "The native will lead a happy life")
	}

	if tenthLordHouse == constants.Meena && (tenthLord != constants.Guru && guruPlacementHouse == constants.Meena) {
		tenthHouseEffects = append(tenthHouseEffects, "If the tenth lord has strength (check shadbala) then the native will obtain good clothes, ornaments and happiness")
	}

	if (*housePlacements)["RahuPlacement"] == 11 && (*housePlacements)["Surya"] == 11 && (*housePlacements)["Shani"] == 11 && kujaPlacement == 11 {
		tenthHouseEffects = append(tenthHouseEffects, "The native will cease his duties")
	}

	if guruPlacementHouse == constants.Meena && utils.GetHouseName((*reqBody).Ascendant ,(*housePlacements)["Shukra"]) == constants.Meena && utils.GetHouseName((*reqBody).Ascendant ,(*housePlacements)["Chandra"]) == constants.ChandraUccha {
		tenthHouseEffects = append(tenthHouseEffects, "If the ascendant lord is strong (check shadbala) then the native will be learned and wealthy and is likely to obtain liberation through Gnana yoga")
	}

	if tenthLordPlacementNumber == 11 && eleventhLordPlacementNumber == 1 && (*housePlacements)["Shukra"] == 10 {
		tenthHouseEffects = append(tenthHouseEffects, "The native will be endowed with precious stones")
	}

	if tenthLordIsEitherConjunctOrAspectGuru && (utils.IsKendra(tenthLordPlacementNumber) || utils.IsKona(tenthLordPlacementNumber)) && isTenthLordExalted {
		tenthHouseEffects = append(tenthHouseEffects, "The native will be endowed with worthy jobs")
	}

	if isTenthOccupiedByNeechGraha(tenthHouse, suryaPlacement, shukraPlacementNumber, chandraPlacementNumber, guruPlacement, budhaPlacementNumber, kujaPlacement) && 
		shaniPlacement == 10 && isNavTenthOccupiedByMalefic {
		tenthHouseEffects = append(tenthHouseEffects, "The native will be bereft of good acts")
	}

	if isEighthLordConjunctMalefic && eighthLordPlacementNumber == 10 && tenthLordPlacementNumber == 8 {
		tenthHouseEffects = append(tenthHouseEffects, "The native will indulge in bad acts")
	}

	if isTenthLordInFall(tenthLordHouse, tenthLord) && isTenthOccupiedByMalefic && isSeventhOccupiedByMalefic {
		tenthHouseEffects = append(tenthHouseEffects, "Obstructions to the natives acts will crop up")
	}

	if tenthLordPlacementNumber == 1 && tenthLordPlacementNumber == ascendantLordPlacement && (utils.IsKendra(chandraPlacementNumber) || utils.IsKona(chandraPlacementNumber)) {
		tenthHouseEffects = append(tenthHouseEffects, "The native will be interested in good jobs")
	}

	if (*housePlacements)["Chandra"] == 10 && utils.IsKona(tenthLordPlacementNumber) && utils.IsKendra(ascendantLordPlacement) {
		tenthHouseEffects = append(tenthHouseEffects, "The native will obtain fame")
	}

	if eleventhLordPlacementNumber == 10 && !isTenthLordNotStrong && tenthLordInAspectOfGuru {
		tenthHouseEffects = append(tenthHouseEffects, "The native will obtain fame")
	}
	
	if tenthLordPlacementNumber == 9 && ascendantLordPlacement == 10 && (*housePlacements)["Chandra"] == 5 {
		tenthHouseEffects = append(tenthHouseEffects, "The native will obtain fame")
	}

	return tenthHouseEffects
}