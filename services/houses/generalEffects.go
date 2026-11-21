package houses

import (
	"slices"

	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/dto"
	"github.com/scientist-v08/bphs/utils"
)

func GeneralEffects(
	reqBody *dto.HouseEffectsReqBody,
	housePlacements *map[string]int16,
	guruAspects *[]dto.AllRaashis,
	shaniAspects *[]dto.AllRaashis,
	kujaAspects *[]dto.AllRaashis,
) []string {
	generalEffects := make([]string, 0)

	// Collect required information for Chandra house placements
	chandraHousePlacements := map[string]int16{
		"Surya": utils.GetHouse((*reqBody).ChandraPlacement, (*reqBody).SuryaPlacement),
		"Budha": utils.GetHouse((*reqBody).ChandraPlacement, (*reqBody).BudhaPlacement.Placement),   
		"Shukra": utils.GetHouse((*reqBody).ChandraPlacement, (*reqBody).ShukraPlacement.Placement),  
		"Chandra": utils.GetHouse((*reqBody).ChandraPlacement, (*reqBody).ChandraPlacement), 
		"RahuPlacement": utils.GetHouse((*reqBody).ChandraPlacement, (*reqBody).RahuPlacement),    
		"KetuPlacement": utils.GetHouse((*reqBody).ChandraPlacement, (*reqBody).KetuPlacement),    
		"Kuja": utils.GetHouse((*reqBody).ChandraPlacement, (*reqBody).KujaPlacement.Placement),    
		"Guru": utils.GetHouse((*reqBody).ChandraPlacement, (*reqBody).GuruPlacement.Placement),    
		"Shani": utils.GetHouse((*reqBody).ChandraPlacement, (*reqBody).ShaniPlacement.Placement),   
	}
	chandraLord := constants.RaashyadhipatiMapStore[(*reqBody).ChandraPlacement]
	chandraLordPlacement := chandraHousePlacements[string(chandraLord)]
	chandraLordHouse := utils.GetHouseName((*reqBody).ChandraPlacement, chandraLordPlacement)
	isChandraAspectedByShukra := constants.OppositeAspectsStore[dto.AllRaashis((*reqBody).ShukraPlacement.Placement)] == dto.AllRaashis((*reqBody).ChandraPlacement)
	isChandraAspectedBySaturn := slices.Contains(*shaniAspects, dto.AllRaashis((*reqBody).ChandraPlacement))
	isChandraAspectedByJupiter := slices.Contains(*guruAspects, dto.AllRaashis((*reqBody).ChandraPlacement))
	isChandraAspectedByMars := slices.Contains(*kujaAspects, dto.AllRaashis((*reqBody).ChandraPlacement))
	isChandraWithSaturn := (*reqBody).ChandraPlacement == (*reqBody).ShaniPlacement.Placement
	isChandraWithMars := (*reqBody).ChandraPlacement == (*reqBody).KujaPlacement.Placement
	isChandraWithRahu := (*reqBody).ChandraPlacement == (*reqBody).RahuPlacement
	isChandraWithKetu := (*reqBody).ChandraPlacement == (*reqBody).KetuPlacement
	isChandraWithGuru := (*reqBody).ChandraPlacement == (*reqBody).GuruPlacement.Placement
	isChandraWithShukra := (*reqBody).ChandraPlacement == (*reqBody).ShukraPlacement.Placement
	isChandraSheltered := isChandraWithGuru || isChandraWithShukra || isChandraAspectedByShukra || isChandraAspectedByJupiter
	isChandraWithBudha := (*reqBody).BudhaPlacement.Placement == (*reqBody).ChandraPlacement
	isChandraOrBudhaInVruschikaOrMeena := slices.Contains([]dto.AllRaashis{constants.Vruschika, constants.Meena}, dto.AllRaashis((*reqBody).ChandraPlacement)) ||
		slices.Contains([]dto.AllRaashis{constants.Vruschika, constants.Meena}, dto.AllRaashis((*reqBody).BudhaPlacement.Placement))
	var countChandraAfflictions = 0
	nShukraPlacementNumber := (*housePlacements)["NavShukra"]
	nShaniPlacementNumber := (*housePlacements)["NavShani"]
	nKujaPlacementNumber := (*housePlacements)["NavKuja"]
	shukraPlacementNumber := (*housePlacements)["Shukra"]
	shkraHousePlacementRaashi := utils.GetHouseName((*reqBody).Ascendant, shukraPlacementNumber)
	shaniPlacementNumber := (*housePlacements)["Shani"]
	kujaPlacementNumber := (*housePlacements)["Kuja"]
	nShkraHousePlacementRaashi := utils.GetHouseName((*reqBody).NavAscendant, nShukraPlacementNumber)
	nShaniHousePlacementRaashi := utils.GetHouseName((*reqBody).NavAscendant, nShaniPlacementNumber)
	nKujaHousePlacementRaashi := utils.GetHouseName((*reqBody).NavAscendant, nKujaPlacementNumber)
	shukraInRasiOfKuja := slices.Contains([]dto.AllRaashis{constants.Mesha, constants.Vruschika}, shkraHousePlacementRaashi)
	shukraInNavamshaOfKuja := slices.Contains([]dto.AllRaashis{constants.Mesha, constants.Vruschika}, nShkraHousePlacementRaashi)
	nShaniAspects := utils.GetShaniAspectRaashis(dto.AllRaashis(nShaniHousePlacementRaashi))
	nKujaAspects := utils.GetKujaAspectRaashis(dto.AllRaashis(nKujaHousePlacementRaashi))
	shukraAspectedByKuja := slices.Contains(*kujaAspects, shkraHousePlacementRaashi)
	navShukraAspectedByNavKuja := slices.Contains(nKujaAspects, nShkraHousePlacementRaashi)
	isShukraConjunctKuja := kujaPlacementNumber == shukraPlacementNumber
	isNavShukraConjunctNavKuja := nKujaPlacementNumber == nShukraPlacementNumber
	shukraInRasiOfShani := slices.Contains([]dto.AllRaashis{constants.Kumbha, constants.Makara}, shkraHousePlacementRaashi)
	shukraInNavamshaOfShani := slices.Contains([]dto.AllRaashis{constants.Kumbha, constants.Makara}, nShkraHousePlacementRaashi)
	shukraAspectedByShani := slices.Contains(*shaniAspects, shkraHousePlacementRaashi)
	navShukraAspectedByNavShani := slices.Contains(nShaniAspects, nShkraHousePlacementRaashi)
	isShukraConjunctShani := shaniPlacementNumber == shukraPlacementNumber
	isNavShukraConjunctNavShani := nShaniPlacementNumber == nShukraPlacementNumber
	isAggressiveSex := false
	isGay := false;

	// Start appending to string array for general placements
	if (*housePlacements)["Shukra"] == 7 || (*housePlacements)["NavShukra"] == 7 {
		generalEffects = append(generalEffects, "The native will be very horny")
	}

	if shukraInRasiOfKuja || shukraInNavamshaOfKuja || shukraAspectedByKuja || isShukraConjunctKuja || navShukraAspectedByNavKuja || isNavShukraConjunctNavKuja {
		isAggressiveSex = true
	}
	if shukraInRasiOfShani || shukraInNavamshaOfShani || shukraAspectedByShani || isShukraConjunctShani || navShukraAspectedByNavShani || isNavShukraConjunctNavShani {
		isGay = true
	}
	switch {
	case isAggressiveSex && isGay:
		generalEffects = append(generalEffects, "The native is likely bisexual")
	case isAggressiveSex:
		generalEffects = append(generalEffects, "The native is likely to kiss the private parts of the opposite gender to satisfy lust")
	case isGay:
		generalEffects = append(generalEffects, "The native is gay")
	}

	if isChandraAspectedByJupiter || isChandraWithGuru {
		generalEffects = append(generalEffects, "The native is magnanimous in both heart and mind")
	}

	if isChandraAspectedByShukra || isChandraWithShukra {
		generalEffects = append(generalEffects, "The native is friendly in nature, amicable, fond of beauty and share happiness")
	}

	if utils.IsDusthana((*housePlacements)["Chandra"]) {
		if isChandraSheltered {
			generalEffects = append(generalEffects, "The native's mind may be weak and may not have a healthy attitude, but the native is protected by Guru / Shukra from this")
		} else {
			generalEffects = append(generalEffects, "The native's mind may be weak and may not have a healthy attitude")
		}
	}

	if (*housePlacements)["Surya"] == 7 {
		generalEffects = append(generalEffects, "The native is likely to hate the opposite gender")
	}
	if (*housePlacements)["NavSurya"] == 7 {
		generalEffects = append(generalEffects, "As the native grows up, the native is likely to hate the opposite gender")
	}

	if isChandraWithBudha && isChandraOrBudhaInVruschikaOrMeena {
		generalEffects = append(generalEffects, "The native could harbour horrendous levels of avarice, hate, jeolousy")
	}

	if isChandraAspectedBySaturn || isChandraWithSaturn {
		if isChandraSheltered {
			generalEffects = append(generalEffects, "The native is likely to suffer from persecution complex and are extremely stubborn in life, but the native is protected by Guru / Shukra from this")
		} else {
			generalEffects = append(generalEffects, "The native is likely to suffer from persecution complex and are extremely stubborn in life")
		}
		countChandraAfflictions++
	}

	if isChandraAspectedByMars || isChandraWithMars {
		if isChandraSheltered {
			generalEffects = append(generalEffects, "The native will make for violent moods, behaviour and action, but the native is protected by Guru / Shukra from this")
		} else {
			generalEffects = append(generalEffects, "The native will make for violent moods, behaviour and action")
		}
		countChandraAfflictions++
	}

	if isChandraWithRahu {
		if isChandraSheltered {
			generalEffects = append(generalEffects, "The native suffers from dangerous levels of suspicion, but the native is protected by Guru / Shukra from this")
		} else {
			generalEffects = append(generalEffects, "The native suffers from dangerous levels of suspicion")
		}
		countChandraAfflictions++
	}

	if isChandraWithKetu {
		if isChandraSheltered {
			generalEffects = append(generalEffects, "The native suffers from intense inferiority complex, but the native is protected by Guru / Shukra from this")
		} else {
			generalEffects = append(generalEffects, "The native suffers from intense inferiority complex")
		}
		countChandraAfflictions++
	}

	if countChandraAfflictions > 1 {
		generalEffects = append(generalEffects, "Too many afflictions makes the native dangerous. Intense penance is required")
	}

	if utils.IsDusthana(chandraLordPlacement) || 
	(chandraLordPlacement == chandraHousePlacements["Shani"] && chandraLord != dto.Raashyadhipati("Shani")) || 
	(chandraLordPlacement == chandraHousePlacements["Kuja"] && chandraLord != dto.Raashyadhipati("Kuja")) || 
	(chandraLordPlacement == chandraHousePlacements["Surya"] && chandraLord != dto.Raashyadhipati("Surya")) {
		generalEffects = append(generalEffects, "Mental pleasure will diminish due to malefic conjunction with Chandra lagna lord")
	}

	if utils.IsKendra(chandraLordPlacement) || utils.IsKona(chandraLordPlacement) {
		generalEffects = append(generalEffects, "Mental pleasure will be available all the time as moon sign lord is in Kendra/Kona")
	}

	if constants.DebilitationMapStore[chandraLord] == chandraLordHouse {
		if ((chandraHousePlacements["Guru"] == chandraLordPlacement && chandraLord != dto.Raashyadhipati("Guru")) || (chandraHousePlacements["Shukra"] == chandraLordPlacement && chandraLord != dto.Raashyadhipati("Shukra"))) && (utils.IsKendra(chandraLordPlacement) || utils.IsKona(chandraLordPlacement)) {
			generalEffects = append(generalEffects, "There will be psychological diseases as moon sign lord is neech, but they will disappear")
		} else {
			generalEffects = append(generalEffects, "There will be psychological diseases as moon sign lord is neech")
		}
	}
	if !slices.Contains([]dto.Raashyadhipati{constants.Surya, constants.Chandra}, chandraLord) {
		if utils.IsGrahaCombust(
			chandraLord,
			(*reqBody).KujaPlacement.Combustion,
			(*reqBody).ShukraPlacement.Combustion,
			(*reqBody).BudhaPlacement.Combustion,
			(*reqBody).GuruPlacement.Combustion,
			(*reqBody).ShaniPlacement.Combustion,
		) {
			if ((chandraHousePlacements["Guru"] == chandraLordPlacement && chandraLord != dto.Raashyadhipati("Guru")) || (chandraHousePlacements["Shukra"] == chandraLordPlacement && chandraLord != dto.Raashyadhipati("Shukra"))) && (utils.IsKendra(chandraLordPlacement) || utils.IsKona(chandraLordPlacement)) {
				generalEffects = append(generalEffects, "There will be psychological diseases as moon sign lord is combust, but they will disappear due to Guru/Shukra")
			} else {
				generalEffects = append(generalEffects, "There will be psychological diseases as moon sign lord is combust")
			}
		}
	}
	if slices.Contains(constants.EnemySignsMappingStore[chandraLord], chandraLordHouse) {
		if ((chandraHousePlacements["Guru"] == chandraLordPlacement && chandraLord != dto.Raashyadhipati("Guru")) || (chandraHousePlacements["Shukra"] == chandraLordPlacement && chandraLord != dto.Raashyadhipati("Shukra"))) && (utils.IsKendra(chandraLordPlacement) || utils.IsKona(chandraLordPlacement)) {
			generalEffects = append(generalEffects, "There will be psychological diseases as moon sign lord is in enemy sign, but they will disappear due to Guru/Shukra")
		} else {
			generalEffects = append(generalEffects, "There will be psychological diseases as moon sign lord is in enemy sign")
		}
	}

	if constants.RaashiQualityStore[(*reqBody).ChandraPlacement] == constants.Chara && slices.Contains(*guruAspects, dto.AllRaashis((*reqBody).ChandraPlacement)) {
		generalEffects = append(generalEffects, "Mentally the individual will be satisfied as moon is in Chara and aspected by Guru")
	}

	if !slices.Contains(*guruAspects, dto.AllRaashis((*reqBody).ChandraPlacement)) && dto.AllRaashis((*reqBody).ChandraPlacement) != constants.OppositeAspectsStore[dto.AllRaashis((*reqBody).ShukraPlacement.Placement)] {
		if (*reqBody).GuruPlacement.Placement != (*reqBody).ChandraPlacement && (*reqBody).ShukraPlacement.Placement != (*reqBody).ChandraPlacement {
			if slices.Contains(*kujaAspects, dto.AllRaashis((*reqBody).ChandraPlacement)) || slices.Contains(*shaniAspects, dto.AllRaashis((*reqBody).ChandraPlacement)) || (*reqBody).ShaniPlacement.Placement == (*reqBody).ChandraPlacement || (*reqBody).KujaPlacement.Placement == (*reqBody).ChandraPlacement {
				generalEffects = append(generalEffects, "The native will not have mental health as moon is aspected/conjunct by a malefic with no protection from a benefic")
			}
		}
	}

	return generalEffects
}