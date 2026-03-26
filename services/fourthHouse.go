package services

import (
	"slices"

	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
	"github.com/scientist-v08/bphs/utils"
)

func FourthHouseEffects(
	reqBody *models.HouseEffectsReqBody,
	housePlacements *map[string]int16,
	guruAspects *[]models.AllRaashis,
	shaniAspects *[]models.AllRaashis,
	kujaAspects *[]models.AllRaashis,
) []string {
	// Create a variable that collects all the effects of the second house
	fourthHouseEffects := make([]string, 0)

	// Get info
	fourthHouse := constants.ReverseNumericalMappingsStore[(*reqBody).Ascendant][4]
	fourthLord := constants.RaashyadhipatiMapStore[string(fourthHouse)]
	fourthLordPlacement := (*housePlacements)[string(fourthLord)]
	fourthLordHouse := utils.GetHouseName((*reqBody).Ascendant, fourthLordPlacement)
	isFourthLordExalted := constants.ExaltationMapStore[fourthLord] == fourthLordHouse
	tenthHouse := constants.ReverseNumericalMappingsStore[(*reqBody).Ascendant][10]
	tenthLord := constants.RaashyadhipatiMapStore[string(tenthHouse)]
	tenthLordPlacement := (*housePlacements)[string(tenthLord)]
	ascendantLord := constants.RaashyadhipatiMapStore[(*reqBody).Ascendant]
	ascendantLordPlacement := (*housePlacements)[string(ascendantLord)]
	fifthHouse := constants.ReverseNumericalMappingsStore[(*reqBody).Ascendant][5]
	fifthLord := constants.RaashyadhipatiMapStore[string(fifthHouse)]
	fifthLordPlacement := (*housePlacements)[string(fifthLord)]
	fifthLordHouse := utils.GetHouseName((*reqBody).Ascendant, fifthLordPlacement)
	isFifthLordExalted := constants.ExaltationMapStore[fifthLord] == fifthLordHouse
	nfifthHouse := constants.ReverseNumericalMappingsStore[(*reqBody).NavAscendant][5]
	nfifthLord := constants.RaashyadhipatiMapStore[string(nfifthHouse)]
	nfifthLordPlacement := (*housePlacements)[string(nfifthLord)]
	isBeneficAspectOn4 := slices.Contains(*guruAspects, fourthHouse) || 
	constants.OppositeAspectsStore[models.AllRaashis((*reqBody).ShukraPlacement.Placement)] == fourthHouse ||
	(constants.OppositeAspectsStore[models.AllRaashis((*reqBody).BudhaPlacement.Placement)] == fourthHouse && (*reqBody).BudhaUnafflicted) ||
	(constants.OppositeAspectsStore[models.AllRaashis((*reqBody).ChandraPlacement)] == fourthHouse && (*reqBody).ChandraWaxing)
	isFourthLordBenefic := fourthLord == constants.Shukra || (fourthLord == constants.Budha && (*reqBody).BudhaUnafflicted) ||
	fourthLord == constants.Guru || (fourthLord == constants.Chandra && (*reqBody).ChandraWaxing)
	isAscendantLordBenefic := ascendantLord == constants.Shukra || (ascendantLord == constants.Budha && (*reqBody).BudhaUnafflicted) ||
	ascendantLord == constants.Guru || (ascendantLord == constants.Chandra && (*reqBody).ChandraWaxing)
	isFourthLordHavingBeneficAspect := 
    (slices.Contains(*guruAspects, fourthLordHouse) && fourthLord != constants.Guru) || 
    (constants.OppositeAspectsStore[models.AllRaashis((*reqBody).ShukraPlacement.Placement)] == fourthLordHouse && fourthLord != constants.Shukra) ||
    (constants.OppositeAspectsStore[models.AllRaashis((*reqBody).BudhaPlacement.Placement)] == fourthLordHouse && (*reqBody).BudhaUnafflicted && fourthLord != constants.Budha) ||
    (constants.OppositeAspectsStore[models.AllRaashis(reqBody.ChandraPlacement)] == fourthLordHouse && reqBody.ChandraWaxing && fourthLord != constants.Chandra)
	isFourthOccupiedByBenefic := 
		(*reqBody).GuruPlacement.Placement == string(fourthHouse) ||
		(*reqBody).ShukraPlacement.Placement == string(fourthHouse) ||
		((*reqBody).BudhaPlacement.Placement == string(fourthHouse) && (*reqBody).BudhaUnafflicted) ||
		((*reqBody).ChandraPlacement == string(fourthHouse) && (*reqBody).ChandraWaxing)
	isFourthLordConjunctBenefic :=
		(*reqBody).GuruPlacement.Placement == string(fourthLordHouse) && fourthLord != constants.Guru ||
		(*reqBody).ShukraPlacement.Placement == string(fourthLordHouse) && fourthLord != constants.Shukra ||
		((*reqBody).BudhaPlacement.Placement == string(fourthLordHouse) && fourthLord != constants.Budha && (*reqBody).BudhaUnafflicted) ||
		((*reqBody).ChandraPlacement == string(fourthLordHouse) && fourthLord != constants.Chandra && (*reqBody).ChandraWaxing)
	shukraPlacement := (*housePlacements)["Shukra"]
	isShukraAtKendra := utils.IsKendra(shukraPlacement)
	isFourthLordAtKendra := utils.IsKendra(fourthLordPlacement)
	isBudhaExalted := constants.ExaltationMapStore[constants.Budha] == utils.GetHouseName((*reqBody).Ascendant, (*housePlacements)["Budha"])
	eleventhHouse := constants.ReverseNumericalMappingsStore[(*reqBody).Ascendant][11]
	eleventhLord := constants.RaashyadhipatiMapStore[string(eleventhHouse)]
	eleventhLordPlacement := (*housePlacements)[string(eleventhLord)]

	// Start appending
	if (fourthLordPlacement == 4 || ascendantLordPlacement == 4) && isBeneficAspectOn4 {
		fourthHouseEffects = append(fourthHouseEffects, "The native will enjoy full residential comforts")
	}

	if fifthLordPlacement == 5 || nfifthLordPlacement == 5 || isFifthLordExalted {
		fourthHouseEffects = append(fourthHouseEffects, "The native will will be endowed with comforts related to lands and houses")
	}

	if tenthLordPlacement == fourthLordPlacement && (utils.IsKendra(fourthLordPlacement) || utils.IsKona(fourthLordPlacement)) {
		fourthHouseEffects = append(fourthHouseEffects, "The native will acquire beautiful mansions")
	}

	if (*reqBody).BudhaPlacement.Placement == (*reqBody).Ascendant && isFourthLordBenefic && isFourthLordHavingBeneficAspect {
		fourthHouseEffects = append(fourthHouseEffects, "The native will be honoured by relatives")
	}

	if isFourthOccupiedByBenefic && isFourthLordExalted {
		fourthHouseEffects = append(fourthHouseEffects, "The natives mother will be long lived")
	}

	if isShukraAtKendra && isFourthLordAtKendra && isBudhaExalted {
		fourthHouseEffects = append(fourthHouseEffects, "The mother of the native will be happy")
	}

	if (*housePlacements)["Surya"] == 4 && (*housePlacements)["Chandra"] == 9 && (*housePlacements)["Shani"] == 9 && (*housePlacements)["Kuja"] == 11 {
		fourthHouseEffects = append(fourthHouseEffects, "The native will obtain cows and buffaloes")
	}

	if constants.Chara == constants.RaashiQualityStore[string(fourthHouse)] && (*housePlacements)["Kuja"] == fourthLordPlacement && (fourthLordPlacement == 6 || fourthLordPlacement == 8) && fourthLord != constants.Kuja {
		fourthHouseEffects = append(fourthHouseEffects, "The native will be dumb")
	}

	if isAscendantLordBenefic && (*housePlacements)["Shukra"] == fourthLordPlacement && fourthLordPlacement == 11 {
		fourthHouseEffects = append(fourthHouseEffects, "The native will obtain a house early in life")
	}
	if eleventhLordPlacement == 4 && fourthLordPlacement == 11 {
		fourthHouseEffects = append(fourthHouseEffects, "The native will obtain a house early in life")
	}
	
	if isFourthOccupiedByBenefic || isBeneficAspectOn4 || isFourthLordHavingBeneficAspect || isFourthLordConjunctBenefic {
		fourthHouseEffects = append(fourthHouseEffects, "The native will be happy about the vehicles he obtains and will have less accidents")
	}

	return fourthHouseEffects
}