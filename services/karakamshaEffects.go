package services

import (
	"slices"

	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
	"github.com/scientist-v08/bphs/utils"
)

func isKendraOrTrikona(house int16) bool {
	return house == 1 || house == 4 || house == 5 ||
		house == 7 || house == 9 || house == 10
}

func toRaashi(r string) models.AllRaashis {
	return models.AllRaashis(r)
}

var validGrahas = [...]string{
    "Surya", "Budha", "Shukra", "Chandra", 
    "Rahu", "Ketu", "Kuja", "Guru", "Shani",
}

func noOfGrahasInBhava(housePlacements *map[string]int16, bhavaNum int16) int8 {
	numberOfGrahas := int8(0)

	for _, val := range validGrahas {
		if (*housePlacements)[val] == bhavaNum {
			numberOfGrahas++
		}
	}

	return numberOfGrahas
}

func KarakamshaEffectsService(housePlacements *map[string]int16, reqBody *models.KarakamshaReqBody) []string {
	var karakamshaEffects []string

	guruAspects := utils.GetGuruAspectRaashis(models.AllRaashis((*reqBody).Guru))
	shaniAspects := utils.GetShaniAspectRaashis(models.AllRaashis((*reqBody).Shani))
	kujaAspects := utils.GetKujaAspectRaashis(models.AllRaashis((*reqBody).Kuja))
	maleficAspectOnKarakamshaExceptKetu := slices.Contains(shaniAspects, toRaashi((*reqBody).Karakamsha)) ||
		slices.Contains(kujaAspects, toRaashi((*reqBody).Karakamsha)) ||
		(constants.OppositeAspectsStore[models.AllRaashis((*reqBody).Surya)] == models.AllRaashis((*reqBody).Karakamsha)) ||
		(constants.OppositeAspectsStore[models.AllRaashis((*reqBody).Budha)] == models.AllRaashis((*reqBody).Karakamsha) && (*reqBody).MercuryAfflicted) ||
		(constants.OppositeAspectsStore[models.AllRaashis((*reqBody).Chandra)] == models.AllRaashis((*reqBody).Karakamsha)  && !(*reqBody).MoonWaxing)
	karakamshaHouseFirstConditionModifiers := "A benefic's aspect will remove evils while that of a melefic will increase the bad effects. In case of both, check the shadbala and see which is stronger and if benefic is stronger no evils and if malefic is stronger than it causes no good."
	beneficsInKarakamsha := (*housePlacements)["Guru"] == 1 || (*housePlacements)["Shukra"] == 1 ||
		((*housePlacements)["Chandra"] == 1 && (*reqBody).MoonWaxing) ||
		((*housePlacements)["Budha"] == 1 && !(*reqBody).MercuryAfflicted)
	noMaleficsInKarakamsha := (*housePlacements)["Surya"] != 1 &&
		(*housePlacements)["Shani"] != 1 &&
		(*housePlacements)["Kuja"] != 1 &&
		(*housePlacements)["Rahu"] != 1 &&
		(*housePlacements)["Ketu"] != 1 &&
		((*housePlacements)["Chandra"] != 1 || (*reqBody).MoonWaxing) &&
		((*housePlacements)["Budha"] != 1 || !(*reqBody).MercuryAfflicted)
	onlyBeneficsInKarakamsha := beneficsInKarakamsha && noMaleficsInKarakamsha
	beneficsAspectLagna := slices.Contains(guruAspects, models.AllRaashis((*reqBody).Ascendant)) ||
		(constants.OppositeAspectsStore[models.AllRaashis((*reqBody).Shukra)] == models.AllRaashis((*reqBody).Ascendant)) ||
		(constants.OppositeAspectsStore[models.AllRaashis((*reqBody).Chandra)] == models.AllRaashis((*reqBody).Ascendant) && (*reqBody).MoonWaxing) ||
		(constants.OppositeAspectsStore[models.AllRaashis((*reqBody).Budha)] == models.AllRaashis((*reqBody).Ascendant) && !(*reqBody).MercuryAfflicted)
	beneficsInLagna := (*housePlacements)["Guru"] == (*housePlacements)["Lagna"] ||
		(*housePlacements)["Shukra"] == (*housePlacements)["Lagna"] ||
		((*housePlacements)["Chandra"] == (*housePlacements)["Lagna"] && (*reqBody).MoonWaxing) ||
		((*housePlacements)["Budha"] == (*housePlacements)["Lagna"] && !(*reqBody).MercuryAfflicted)
	beneficsInAngleOrTrineFromKarakamsha := isKendraOrTrikona((*housePlacements)["Guru"]) ||
		isKendraOrTrikona((*housePlacements)["Shukra"]) ||
		(isKendraOrTrikona((*housePlacements)["Chandra"]) && (*reqBody).MoonWaxing) ||
		(isKendraOrTrikona((*housePlacements)["Budha"]) && !(*reqBody).MercuryAfflicted)
	maleficsInAngleOrTrineFromKarakamsha :=	isKendraOrTrikona((*housePlacements)["Shani"]) ||
		isKendraOrTrikona((*housePlacements)["Mangal"]) ||
		isKendraOrTrikona((*housePlacements)["Rahu"]) ||
		isKendraOrTrikona((*housePlacements)["Ketu"]) ||
		isKendraOrTrikona((*housePlacements)["Surya"]) ||
		(isKendraOrTrikona((*housePlacements)["Chandra"]) && !(*reqBody).MoonWaxing) ||
		(isKendraOrTrikona((*housePlacements)["Budha"]) && (*reqBody).MercuryAfflicted)
	akInMoonMarsVenusSigns := []models.AllRaashis{
		constants.Karkataka,
		constants.Mesha,
		constants.Vruschika,
		constants.Vrushabha,
		constants.Tula,
	}
	secondHouse := constants.ReverseNumericalMappingsStore[(*reqBody).Karakamsha][2]
	isSecondVenusMarsSign := []models.AllRaashis{
		constants.Vrushabha,
		constants.Tula,
		constants.Mesha,
		constants.Vruschika,
	}
	isVenusOrMarsAspectSecondHouse := slices.Contains(kujaAspects, secondHouse) ||
		(constants.OppositeAspectsStore[models.AllRaashis((*reqBody).Shukra)] == secondHouse)
	isMaleficInThirdFromKarakamsha := (*housePlacements)["Surya"] == 3 ||
		(*housePlacements)["Shani"] == 3 ||
		(*housePlacements)["Kuja"] == 3 ||
		(*housePlacements)["Rahu"] == 3 ||
		(*housePlacements)["Ketu"] == 3 ||
		((*housePlacements)["Chandra"] == 3 && !(*reqBody).MoonWaxing) ||
		((*housePlacements)["Budha"] == 3 && (*reqBody).MercuryAfflicted)
	isBeneficsInThirdFromKarakamsha := (*housePlacements)["Guru"] == 3 || (*housePlacements)["Shukra"] == 3 ||
		((*housePlacements)["Chandra"] == 3 && (*reqBody).MoonWaxing) ||
		((*housePlacements)["Budha"] == 3 && !(*reqBody).MercuryAfflicted)
	isMaleficInSixthFromKarakamsha := (*housePlacements)["Surya"] == 6 ||
		(*housePlacements)["Shani"] == 6 ||
		(*housePlacements)["Kuja"] == 6 ||
		(*housePlacements)["Rahu"] == 6 ||
		(*housePlacements)["Ketu"] == 6 ||
		((*housePlacements)["Chandra"] == 6 && !(*reqBody).MoonWaxing) ||
		((*housePlacements)["Budha"] == 6 && (*reqBody).MercuryAfflicted)
	eighthHouse := constants.ReverseNumericalMappingsStore[(*reqBody).Karakamsha][8]
	eightLord := constants.RaashyadhipatiMapStore[string(eighthHouse)]
	eighthLordPlacement := (*housePlacements)[string(eightLord)]
	ninthHouse := constants.ReverseNumericalMappingsStore[(*reqBody).Karakamsha][9]
	isMaleficInEighthFromKarakamsha :=	("Surya" != string(eightLord) && (*housePlacements)["Surya"] == 8) ||
		("Shani" != string(eightLord) && (*housePlacements)["Shani"] == 8) ||
		("Kuja" != string(eightLord) && (*housePlacements)["Kuja"] == 8) ||
		((*housePlacements)["Rahu"] == 8) || ((*housePlacements)["Ketu"] == 8) ||
		("Chandra" != string(eightLord) && (*housePlacements)["Chandra"] == 8 && !(*reqBody).MoonWaxing) ||
		("Budha" != string(eightLord) && (*housePlacements)["Budha"] == 8 && (*reqBody).MercuryAfflicted)
	maleficAspectOnKarakamsha := (string(eightLord) != "Shani" && slices.Contains(shaniAspects, eighthHouse)) ||
		(string(eightLord) != "Kuja" && slices.Contains(kujaAspects, eighthHouse)) ||
		(string(eightLord) != "Surya" && constants.OppositeAspectsStore[models.AllRaashis((*reqBody).Surya)] == eighthHouse) ||
		(string(eightLord) != "Budha" && constants.OppositeAspectsStore[models.AllRaashis((*reqBody).Budha)] == eighthHouse &&	(*reqBody).MercuryAfflicted) ||
		(string(eightLord) != "Chandra" &&	constants.OppositeAspectsStore[models.AllRaashis((*reqBody).Chandra)] == eighthHouse &&	!(*reqBody).MoonWaxing)
	isEighthContainsOrAspectedByMalefic := isMaleficInEighthFromKarakamsha || maleficAspectOnKarakamsha
	isBeneficsInSixthFromKarakamsha := (*housePlacements)["Guru"] == 6 || (*housePlacements)["Shukra"] == 6 ||
		((*housePlacements)["Chandra"] == 6 && (*reqBody).MoonWaxing) ||
		((*housePlacements)["Budha"] == 6 && !(*reqBody).MercuryAfflicted)
	isBeneficsInEighthFromKarakamsha := (*housePlacements)["Guru"] == 8 || (*housePlacements)["Shukra"] == 8 ||
		((*housePlacements)["Chandra"] == 8 && (*reqBody).MoonWaxing) ||
		((*housePlacements)["Budha"] == 8 && !(*reqBody).MercuryAfflicted)
	isBeneficsInNinthFromKarakamsha := (*housePlacements)["Guru"] == 9 || (*housePlacements)["Shukra"] == 9 ||
		((*housePlacements)["Chandra"] == 9 && (*reqBody).MoonWaxing) ||
		((*housePlacements)["Budha"] == 9 && !(*reqBody).MercuryAfflicted)
	beneficsAspectNinth := slices.Contains(guruAspects, ninthHouse) ||
		(constants.OppositeAspectsStore[models.AllRaashis((*reqBody).Shukra)] == ninthHouse) ||
		(constants.OppositeAspectsStore[models.AllRaashis((*reqBody).Chandra)] == ninthHouse && (*reqBody).MoonWaxing) ||
		(constants.OppositeAspectsStore[models.AllRaashis((*reqBody).Budha)] == ninthHouse && !(*reqBody).MercuryAfflicted)
	fourthHouse := constants.ReverseNumericalMappingsStore[(*reqBody).Karakamsha][4]
	fifthHouse := constants.ReverseNumericalMappingsStore[(*reqBody).Karakamsha][5]
	kujaAspectOnFifthFromKarakamsha := slices.Contains(kujaAspects, fifthHouse)
	grahaExaltedInFourth, exaltedGrahaExists := constants.ReverseExaltationMapStore[fourthHouse]
	isAnyGrahaExaltedInFourth := false
	if exaltedGrahaExists {
		if (*housePlacements)[string(grahaExaltedInFourth)] == 4 {
			isAnyGrahaExaltedInFourth = true
		}
	}
	grahaNeechInFourth, exists := constants.ReverseDebilitationMapStore[fourthHouse]
	fourthLord := constants.RaashyadhipatiMapStore[string(fourthHouse)]
	fourthLordPlacement := (*housePlacements)[string(fourthLord)]
	isAnyGrahaNeechBhanghaInFourth := false
	if exists {
		if (*housePlacements)[string(grahaNeechInFourth)] == fourthLordPlacement && fourthLordPlacement == 4 {
			isAnyGrahaNeechBhanghaInFourth = true
		}
	}

	if isBeneficsInNinthFromKarakamsha || beneficsAspectNinth {
		karakamshaEffects = append(karakamshaEffects, "The native will be truthful, devoted to elders and attached to his own religion")
	}

	if isBeneficsInEighthFromKarakamsha || eighthLordPlacement == 8 {
		if isEighthContainsOrAspectedByMalefic {
			karakamshaEffects = append(karakamshaEffects, "The native will have a medium life-span.")
		} else {
			karakamshaEffects = append(karakamshaEffects, "The native will be long lived.")
		}
	} else if isEighthContainsOrAspectedByMalefic {
		karakamshaEffects = append(karakamshaEffects, "The native's life span may not be long.")
	}

	if isMaleficInSixthFromKarakamsha {
		karakamshaEffects = append(karakamshaEffects, "The native will be an agriculturist.")
	}

	if isBeneficsInSixthFromKarakamsha {
		karakamshaEffects = append(karakamshaEffects, "The native will be lazy.")
	}

	if (*housePlacements)["Chandra"] == 7 && (*housePlacements)["Guru"] == 7 {
		karakamshaEffects = append(karakamshaEffects, "The native will beget a very beautiful wife.")
	}

	if (*housePlacements)["Shukra"] == 7 {
		karakamshaEffects = append(karakamshaEffects, "The native will beget a sensuous wife.")
	}
	
	if (*housePlacements)["Budha"] == 7 {
		karakamshaEffects = append(karakamshaEffects, "The native will beget a wife well versed in arts.")
	}

	if (*housePlacements)["Budha"] == 10 && (*housePlacements)["Shukra"] == 10 {
		karakamshaEffects = append(karakamshaEffects, "The native will gain in business and will do many great deeds.")
	}
	
	if (*housePlacements)["Rahu"] == 7 {
		karakamshaEffects = append(karakamshaEffects, "The native is likely to marry an individual as his or her second spouse.")
	}
	
	if (*housePlacements)["Surya"] == 7 {
		karakamshaEffects = append(karakamshaEffects, "The native will beget a wife that is confined to domestic core.")
	}
	
	if (*housePlacements)["Shani"] == 7 {
		karakamshaEffects = append(karakamshaEffects, "The native will beget a wife of higher age bracket or a pious wife or a sick wife.")
	}

	if kujaAspectOnFifthFromKarakamsha {
		karakamshaEffects = append(karakamshaEffects, "The native may get boils or ulcers")
	}

	if (*housePlacements)["Ketu"] == 11 {
		karakamshaEffects = append(karakamshaEffects, "The native may get dysentry or diseases related to impure water")
	}
	
	if (*housePlacements)["Budha"] == 5 {
		karakamshaEffects = append(karakamshaEffects, "The native will be an ascetic of the highest order")
	}
	
	if (*housePlacements)["Surya"] == 5 && (*housePlacements)["Kuja"] == 5 {
		karakamshaEffects = append(karakamshaEffects, "The native may be involved in activities or professions requiring precision tools, sharp instruments, or decisive technical action")
	}
	if (*housePlacements)["Shani"] == 5 {
		karakamshaEffects = append(karakamshaEffects, "The native may engage in activities requiring focus, precision, and strategic targeting")
	}
	if (*housePlacements)["Shukra"] == 5 {
		karakamshaEffects = append(karakamshaEffects, "The native may become a poet and a eloquent speaker.")
	}

	validHouses := []int16{1, 5}
	if slices.Contains(validHouses, (*housePlacements)["Guru"]) &&	slices.Contains(validHouses, (*housePlacements)["Chandra"]) &&	(*housePlacements)["Guru"] == (*housePlacements)["Chandra"] {
		karakamshaEffects = append(karakamshaEffects, "The native will be an author")
	}
	if slices.Contains(validHouses, (*housePlacements)["Shukra"]) &&	slices.Contains(validHouses, (*housePlacements)["Chandra"]) &&	(*housePlacements)["Shukra"] == (*housePlacements)["Chandra"] {
		karakamshaEffects = append(karakamshaEffects, "The native will be an ordinary author")
	}

	if slices.Contains(validHouses, (*housePlacements)["Guru"]) {
		if noOfGrahasInBhava(housePlacements, (*housePlacements)["Guru"]) == 1 {
			karakamshaEffects = append(karakamshaEffects, "The native will be a knower of everything, a writer, be versed in Vedas and Vedanta but not an oratorian or a grammarian.")
		}
	}
	if slices.Contains(validHouses, (*housePlacements)["Kuja"]) {
		if noOfGrahasInBhava(housePlacements, (*housePlacements)["Kuja"]) == 1 {
			karakamshaEffects = append(karakamshaEffects, "The native is a logician")
		}
	}
	if slices.Contains(validHouses, (*housePlacements)["Budha"]) {
		if noOfGrahasInBhava(housePlacements, (*housePlacements)["Budha"]) == 1 {
			karakamshaEffects = append(karakamshaEffects, "The native is a mimamsaka")
		}
	}
	if slices.Contains(validHouses, (*housePlacements)["Shani"]) {
		if noOfGrahasInBhava(housePlacements, (*housePlacements)["Shani"]) == 1 {
			karakamshaEffects = append(karakamshaEffects, "The native is dumb-witted")
		}
	}
	if slices.Contains(validHouses, (*housePlacements)["Surya"]) {
		if noOfGrahasInBhava(housePlacements, (*housePlacements)["Surya"]) == 1 {
			karakamshaEffects = append(karakamshaEffects, "The native is a musician")
		}
	}
	if slices.Contains(validHouses, (*housePlacements)["Chandra"]) {
		if noOfGrahasInBhava(housePlacements, (*housePlacements)["Chandra"]) == 1 {
			karakamshaEffects = append(karakamshaEffects, "The native is a follower of Sankhya philosophy")
		}
	}
	if slices.Contains(validHouses, (*housePlacements)["Rahu"]) {
		if noOfGrahasInBhava(housePlacements, (*housePlacements)["Rahu"]) == 1 {
			karakamshaEffects = append(karakamshaEffects, "The native is an astrologer")
		}
	}
	if slices.Contains(validHouses, (*housePlacements)["Ketu"]) {
		if noOfGrahasInBhava(housePlacements, (*housePlacements)["Ketu"]) == 1 {
			karakamshaEffects = append(karakamshaEffects, "The native is an astrologer")
		}
	}

	if isAnyGrahaExaltedInFourth {
		karakamshaEffects = append(karakamshaEffects, "The native will own large buildings")
	}

	if isAnyGrahaNeechBhanghaInFourth {
		karakamshaEffects = append(karakamshaEffects, "The native will own large buildings after struggles and initial setbacks.")
	}

	if (*housePlacements)["Chandra"] == 4 && (*housePlacements)["Shukra"] == 4 {
		karakamshaEffects = append(karakamshaEffects, "The native will own large buildings")
	}

	if (*housePlacements)["Rahu"] == 4 && (*housePlacements)["Shani"] == 4 {
		karakamshaEffects = append(karakamshaEffects,
			"The native may reside in or acquire property built with durable, heavy materials such as concrete or stone, indicating strong and long-lasting structures")
	}
	
	if (*housePlacements)["Rahu"] == 5 && (*housePlacements)["Kuja"] == 5 {
		karakamshaEffects = append(karakamshaEffects, "The native may suffer from a pulmonary consumption")
	}

	if (*housePlacements)["Kuja"] == 4 && (*housePlacements)["Ketu"] == 4 {
		karakamshaEffects = append(karakamshaEffects,
			"The native may be associated with properties built using engineered or modular materials, such as brick or structured construction, often reflecting practical and functional design")
	}

	if (*housePlacements)["Guru"] == 4 {
		karakamshaEffects = append(karakamshaEffects,
			"The native may own or reside in spacious, well-designed homes with natural or traditional elements, often reflecting comfort and wisdom in living spaces")
	}

	if (*housePlacements)["Surya"] == 4 {
		karakamshaEffects = append(karakamshaEffects,
			"The native may reside in simple or minimalistic dwellings, possibly with temporary or lightweight structures, or environments exposed to natural elements")
	}	
	
	if isBeneficsInThirdFromKarakamsha {
		karakamshaEffects = append(karakamshaEffects, "The native will be timid.")
	}

	if isMaleficInThirdFromKarakamsha {
		karakamshaEffects = append(karakamshaEffects, "The native will be courageous.")
	}

	if slices.Contains(isSecondVenusMarsSign, secondHouse) {
		seconfHouseEffect := "The native will be addicted to others wives. "
		if isVenusOrMarsAspectSecondHouse {
			seconfHouseEffect += "This will continue until the last breath of the native."
		}
		if (*housePlacements)["Ketu"] != 2 {
			karakamshaEffects = append(karakamshaEffects, seconfHouseEffect)
		}
		if (*housePlacements)["Rahu"] == 2 {
			karakamshaEffects = append(karakamshaEffects, "The natives wealth maybe be destroyed")
		}
	}

	if (*housePlacements)["Surya"] == 1 {
		karakamshaEffects = append(karakamshaEffects, "The native will be engaged in royal assignments")
	}

	if (*housePlacements)["Surya"] == 1 && (*housePlacements)["Rahu"] == 1 {
		karakamshaEffects = append(karakamshaEffects, "The native will have fear from snake. A benefic aspect ensures no fear and a malefic aspect worsens this effect")
	}

	if (*housePlacements)["Rahu"] == 1 {
		karakamshaEffects = append(karakamshaEffects, "The native is likely to become an engineer or works involving secrecy or medical technician")
	}

	if (*housePlacements)["Ketu"] == 1 {
		effect := "The native may be associated with works involving large assets, heavy machinery, or specialized technical domains. "
		if maleficAspectOnKarakamshaExceptKetu {
			effect += "The native may have diseases or problems related to the ear."
		}
		karakamshaEffects = append(karakamshaEffects, effect)
	}

	if (*housePlacements)["Shani"] == 1 {
		karakamshaEffects = append(karakamshaEffects, "The native will inherit the profession of his family for livelihood")
	}

	if (*housePlacements)["Shukra"] == 1 {
		karakamshaEffects = append(karakamshaEffects, "The native will be long lived and be sensuous.")
	}

	if (*housePlacements)["Mangal"] == 1 {
		karakamshaEffects = append(karakamshaEffects, "If Kuja has obtained minimum in shadbala then the native may be inclined towards the use of weapons, sharp instruments, or professions involving precision, force, and technical skill")
	}

	if (*housePlacements)["Budha"] == 1 {
		karakamshaEffects = append(karakamshaEffects, "The native will be skillful in arts and trading, be intelligent and educated")
	}

	if (*housePlacements)["Guru"] == 1 {
		karakamshaEffects = append(karakamshaEffects, "The native will be engaged towards good acts and inclined towards spiritualism")
	}

	if onlyBeneficsInKarakamsha && (beneficsAspectLagna || beneficsInLagna) {
		karakamshaEffects = append(karakamshaEffects, "The native will undoubtedly obtain leadership position in a huge organization")
	}

	switch models.AllRaashis((*reqBody).Karakamsha) {
	case constants.Mesha:
		karakamshaEffects = append(karakamshaEffects, "There will be nuisance from rats and cats." + karakamshaHouseFirstConditionModifiers)
	case constants.Vrushabha:
		karakamshaEffects = append(karakamshaEffects, "There will be happiness from pets.")
	case constants.Mithuna:
		karakamshaEffects = append(karakamshaEffects, "The native will be afflicted by itch etc." + karakamshaHouseFirstConditionModifiers)
	case constants.Karkataka:
		karakamshaEffects = append(karakamshaEffects, "The native will have fear from water." + karakamshaHouseFirstConditionModifiers)
	case constants.Simha:
		karakamshaEffects = append(karakamshaEffects, "There will be fear from animals." + karakamshaHouseFirstConditionModifiers)
	case constants.Kanya:
		karakamshaEffects = append(karakamshaEffects, "There will be fear from itch, corpulence, fire etc." + karakamshaHouseFirstConditionModifiers)
	case constants.Tula:
		karakamshaEffects = append(karakamshaEffects, "The native will be a trader having skills in making clothes.")
	case constants.Vruschika:
		karakamshaEffects = append(karakamshaEffects, "The native will have fear from snakes." + karakamshaHouseFirstConditionModifiers)
	case constants.Dhanassu:
		karakamshaEffects = append(karakamshaEffects, "The native is likely to fall from heights or vehicles." + karakamshaHouseFirstConditionModifiers)
	case constants.Makara:
		karakamshaEffects = append(karakamshaEffects, "The native will make gains from conch, pearl, coral etc.")
	case constants.Kumbha:
		karakamshaEffects = append(karakamshaEffects, "The native will create data stores, or setting up charitable trusts, or working in battery tech, or civil engineering.")
	case constants.Meena:
		karakamshaEffects = append(karakamshaEffects, "The native will obtain Moksha")
	}

	if beneficsInAngleOrTrineFromKarakamsha {
		effect := "The native will obtain wealth and learning"
		if maleficsInAngleOrTrineFromKarakamsha {
			effect += ", but results will be mixed due to presence of malefics"
		}
		karakamshaEffects = append(karakamshaEffects, effect)
	}

	if slices.Contains(akInMoonMarsVenusSigns, toRaashi((*reqBody).Karakamsha)) {
		karakamshaEffects = append(karakamshaEffects, "The native may be inclined towards other's spouse")
	} else {
		karakamshaEffects = append(karakamshaEffects, "The native will never be inclined towards other's spouse")
	}

	return karakamshaEffects
}