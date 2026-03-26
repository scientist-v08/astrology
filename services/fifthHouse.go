package services

import (
	"slices"

	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
	"github.com/scientist-v08/bphs/utils"
)

func FifthHouseEffects(
	reqBody *models.HouseEffectsReqBody,
	housePlacements *map[string]int16,
	guruAspects *[]models.AllRaashis,
	shaniAspects *[]models.AllRaashis,
	kujaAspects *[]models.AllRaashis,
) []string {
	// Create a variable that collects all the effects of the second house
	fifthHouseEffects := make([]string, 0)

	// Get info
	fifthHouse := constants.ReverseNumericalMappingsStore[(*reqBody).Ascendant][5]
	fifthLord := constants.RaashyadhipatiMapStore[string(fifthHouse)]
	fifthLordPlacement := (*housePlacements)[string(fifthLord)]
	fifthLordHouse := utils.GetHouseName((*reqBody).Ascendant, fifthLordPlacement)
	isFifthLordInInimicalHouse := slices.Contains(constants.EnemySignsMappingStore[fifthLord], fifthLordHouse)
	ascendantLord := constants.RaashyadhipatiMapStore[(*reqBody).Ascendant]
	ascendantLordPlacement := (*housePlacements)[string(ascendantLord)]
	ascendantLordHouse := utils.GetHouseName((*reqBody).Ascendant, ascendantLordPlacement)
	ninthHouse := constants.ReverseNumericalMappingsStore[(*reqBody).Ascendant][9]
	ninthLord := constants.RaashyadhipatiMapStore[string(ninthHouse)]
	ninthLordPlacement := (*housePlacements)[string(ninthLord)]

	// Start appending
	if (utils.IsLordInOwnHouse(fifthLord, fifthLordHouse) && utils.IsLordInOwnHouse(ascendantLord, ascendantLordHouse)) || 
	(utils.IsKendra(fifthLordPlacement) && utils.IsKendra(ascendantLordPlacement)) ||
	(utils.IsKona(fifthLordPlacement) && utils.IsKona(ascendantLordPlacement)) {
		fifthHouseEffects = append(fifthHouseEffects, "The native will enjoy by seeing the happiness of his children")
	}

	if utils.IsDusthana(fifthLordPlacement) {
		fifthHouseEffects = append(fifthHouseEffects, "The native will not have any offsprings")
	}

	if fifthLordPlacement == ascendantLordPlacement && (utils.IsKendra(fifthLordPlacement) || utils.IsKona(fifthLordPlacement)) {
		fifthHouseEffects = append(fifthHouseEffects, "The native will obtain children early in his life")
	}

	if fifthLordPlacement == ascendantLordPlacement && utils.IsDusthana(fifthLordPlacement) {
		fifthHouseEffects = append(fifthHouseEffects, "The native will not obtain children and suffer becuase of it.")
	}

	if fifthLordPlacement == 6 && ascendantLordPlacement == (*housePlacements)["Kuja"] {
		fifthHouseEffects = append(fifthHouseEffects, "The native will lose his first child and subsequently the spouse of the native will become infertile")
	}

	if utils.IsLordCombust(fifthLord, reqBody) && utils.IsDusthana(fifthLordPlacement) && (*housePlacements)["Budha"] == 5 && (*housePlacements)["KetuPlacement"] == 5 {
		fifthHouseEffects = append(fifthHouseEffects, "The native will have only one child")
	}

	if ninthLordPlacement == 1 && utils.IsLordCombust(fifthLord, reqBody) && (*housePlacements)["Budha"] == 5 && (*housePlacements)["KetuPlacement"] == 5 {
		fifthHouseEffects = append(fifthHouseEffects, "A child will be obtained after a lot of difficulty or be childless")
	}

	if utils.IsDusthana(fifthLordPlacement) || isFifthLordInInimicalHouse || utils.IsLordCombust(fifthLord, reqBody) || fifthLordPlacement == 5 {
		fifthHouseEffects = append(fifthHouseEffects, "The native will get children with difficulty")
	}

	if (*housePlacements)["Surya"] == (*housePlacements)["Chandra"] && (*housePlacements)["NavChandra"] == (*housePlacements)["NavSurya"] {
		fifthHouseEffects = append(fifthHouseEffects, "The native will be bought up by more than 1 father and more than 1 mother")
	}

	if (*housePlacements)["Chandra"] == fifthLordPlacement && fifthLord != constants.Chandra {
		fifthHouseEffects = append(fifthHouseEffects, "The native will get more daughters than sons or only daughters")
	}

	if fifthLordHouse == constants.Karkataka || fifthLordHouse == constants.Vruschika || fifthLordHouse == constants.Meena {
		fifthHouseEffects = append(fifthHouseEffects, "If the fifth lord is in the first 10 degrees of Karkataka or second ten degrees of Meena or last ten degrees of Vruschika then the native will get daughters")
	}

	if fifthLordHouse == constants.ExaltationMapStore[fifthLord] || 
	fifthLordPlacement == 2 || 
	fifthLordPlacement == 5 || 
	fifthLordPlacement == 9 ||
	(fifthLordPlacement == (*housePlacements)["Guru"] && fifthLord != constants.Guru) ||
	slices.Contains(*guruAspects, fifthLordHouse) {
		fifthHouseEffects = append(fifthHouseEffects, "The native will definitely get children")
	}

	// Return all the fifth house effects
	return fifthHouseEffects
}