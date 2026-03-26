package services

import (
	"slices"

	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
	"github.com/scientist-v08/bphs/utils"
)

func GeneralEffects(
	reqBody *models.HouseEffectsReqBody,
	housePlacements *map[string]int16,
	guruAspects *[]models.AllRaashis,
	shaniAspects *[]models.AllRaashis,
	kujaAspects *[]models.AllRaashis,
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
	isChandraAspectedByShukra := constants.OppositeAspectsStore[models.AllRaashis((*reqBody).ShukraPlacement.Placement)] == models.AllRaashis((*reqBody).ChandraPlacement)
	isChandraAspectedBySaturn := slices.Contains(*shaniAspects, models.AllRaashis((*reqBody).ChandraPlacement))
	isChandraAspectedByJupiter := slices.Contains(*guruAspects, models.AllRaashis((*reqBody).ChandraPlacement))
	isChandraAspectedByMars := slices.Contains(*kujaAspects, models.AllRaashis((*reqBody).ChandraPlacement))
	isChandraWithSaturn := (*reqBody).ChandraPlacement == (*reqBody).ShaniPlacement.Placement
	isChandraWithMars := (*reqBody).ChandraPlacement == (*reqBody).KujaPlacement.Placement
	isChandraWithRahu := (*reqBody).ChandraPlacement == (*reqBody).RahuPlacement
	isChandraWithKetu := (*reqBody).ChandraPlacement == (*reqBody).KetuPlacement
	isChandraWithGuru := (*reqBody).ChandraPlacement == (*reqBody).GuruPlacement.Placement
	isChandraWithShukra := (*reqBody).ChandraPlacement == (*reqBody).ShukraPlacement.Placement
	isChandraSheltered := isChandraWithGuru || isChandraWithShukra || isChandraAspectedByShukra || isChandraAspectedByJupiter
	isChandraWithBudha := (*reqBody).BudhaPlacement.Placement == (*reqBody).ChandraPlacement
	isChandraOrBudhaInVruschikaOrMeena := slices.Contains([]models.AllRaashis{constants.Vruschika, constants.Meena}, models.AllRaashis((*reqBody).ChandraPlacement)) ||
		slices.Contains([]models.AllRaashis{constants.Vruschika, constants.Meena}, models.AllRaashis((*reqBody).BudhaPlacement.Placement))
	var countChandraAfflictions = 0
	nShukraPlacementNumber := (*housePlacements)["NavShukra"]
	shukraPlacementNumber := (*housePlacements)["Shukra"]
	shkraHousePlacementRaashi := utils.GetHouseName((*reqBody).Ascendant, shukraPlacementNumber)
	shaniPlacementNumber := (*housePlacements)["Shani"]
	kujaPlacementNumber := (*housePlacements)["Kuja"]
	nShkraHousePlacementRaashi := utils.GetHouseName((*reqBody).NavAscendant, nShukraPlacementNumber)
	shukraInRasiOfKuja := slices.Contains([]models.AllRaashis{constants.Mesha, constants.Vruschika}, shkraHousePlacementRaashi)
	shukraInNavamshaOfKuja := slices.Contains([]models.AllRaashis{constants.Mesha, constants.Vruschika}, nShkraHousePlacementRaashi)
	shukraAspectedByKuja := slices.Contains(*kujaAspects, shkraHousePlacementRaashi)
	isShukraConjunctKuja := kujaPlacementNumber == shukraPlacementNumber
	shukraInRasiOfShani := slices.Contains([]models.AllRaashis{constants.Kumbha, constants.Makara}, shkraHousePlacementRaashi)
	shukraInNavamshaOfShani := slices.Contains([]models.AllRaashis{constants.Kumbha, constants.Makara}, nShkraHousePlacementRaashi)
	shukraAspectedByShani := slices.Contains(*shaniAspects, shkraHousePlacementRaashi)
	isShukraConjunctShani := shaniPlacementNumber == shukraPlacementNumber

	// Start appending to string array for general placements
	if (*housePlacements)["Shukra"] == 7 || (*housePlacements)["NavShukra"] == 7 {
		generalEffects = append(generalEffects, "The native will be very horny")
	}

	if shukraInRasiOfKuja || shukraInNavamshaOfKuja || shukraAspectedByKuja || isShukraConjunctKuja {
		generalEffects = append(generalEffects, "The native is likely to develop feelings to lick or kiss the private parts of the opposite gender to satisfy lust")
	}
	if shukraInRasiOfShani || shukraInNavamshaOfShani || shukraAspectedByShani || isShukraConjunctShani {
		generalEffects = append(generalEffects, "The native is likely to develop feelings to lick or kiss the private parts of the same gender to satisfy lust")
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
	(chandraLordPlacement == chandraHousePlacements["Shani"] && chandraLord != models.Raashyadhipati("Shani")) || 
	(chandraLordPlacement == chandraHousePlacements["Kuja"] && chandraLord != models.Raashyadhipati("Kuja")) || 
	(chandraLordPlacement == chandraHousePlacements["Surya"] && chandraLord != models.Raashyadhipati("Surya")) {
		generalEffects = append(generalEffects, "Mental pleasure will diminish due to malefic aspect and dusthana placement of moon sign lord")
	}

	if utils.IsKendra(chandraLordPlacement) || utils.IsKona(chandraLordPlacement) {
		generalEffects = append(generalEffects, "Mental pleasure will be available all the time as moon sign lord is in Kendra/Kona")
	}

	if constants.DebilitationMapStore[chandraLord] == chandraLordHouse {
		if ((chandraHousePlacements["Guru"] == chandraLordPlacement && chandraLord != models.Raashyadhipati("Guru")) || (chandraHousePlacements["Shukra"] == chandraLordPlacement && chandraLord != models.Raashyadhipati("Shukra"))) && (utils.IsKendra(chandraLordPlacement) || utils.IsKona(chandraLordPlacement)) {
			generalEffects = append(generalEffects, "There will be psychological diseases as moon sign lord is neech, but they will disappear")
		} else {
			generalEffects = append(generalEffects, "There will be psychological diseases as moon sign lord is neech")
		}
	}
	if !slices.Contains([]models.Raashyadhipati{constants.Surya, constants.Chandra}, chandraLord) {
		if utils.IsGrahaCombust(
			chandraLord,
			(*reqBody).KujaPlacement.Combustion,
			(*reqBody).ShukraPlacement.Combustion,
			(*reqBody).BudhaPlacement.Combustion,
			(*reqBody).GuruPlacement.Combustion,
			(*reqBody).ShaniPlacement.Combustion,
		) {
			if ((chandraHousePlacements["Guru"] == chandraLordPlacement && chandraLord != models.Raashyadhipati("Guru")) || (chandraHousePlacements["Shukra"] == chandraLordPlacement && chandraLord != models.Raashyadhipati("Shukra"))) && (utils.IsKendra(chandraLordPlacement) || utils.IsKona(chandraLordPlacement)) {
				generalEffects = append(generalEffects, "There will be psychological diseases as moon sign lord is combust, but they will disappear due to Guru/Shukra")
			} else {
				generalEffects = append(generalEffects, "There will be psychological diseases as moon sign lord is combust")
			}
		}
	}
	if slices.Contains(constants.EnemySignsMappingStore[chandraLord], chandraLordHouse) {
		if ((chandraHousePlacements["Guru"] == chandraLordPlacement && chandraLord != models.Raashyadhipati("Guru")) || (chandraHousePlacements["Shukra"] == chandraLordPlacement && chandraLord != models.Raashyadhipati("Shukra"))) && (utils.IsKendra(chandraLordPlacement) || utils.IsKona(chandraLordPlacement)) {
			generalEffects = append(generalEffects, "There will be psychological diseases as moon sign lord is in enemy sign, but they will disappear due to Guru/Shukra")
		} else {
			generalEffects = append(generalEffects, "There will be psychological diseases as moon sign lord is in enemy sign")
		}
	}

	if constants.RaashiQualityStore[(*reqBody).ChandraPlacement] == constants.Chara && slices.Contains(*guruAspects, models.AllRaashis((*reqBody).ChandraPlacement)) {
		generalEffects = append(generalEffects, "Mentally the individual will be satisfied as moon is in Chara and aspected by Guru")
	}

	if !slices.Contains(*guruAspects, models.AllRaashis((*reqBody).ChandraPlacement)) && models.AllRaashis((*reqBody).ChandraPlacement) != constants.OppositeAspectsStore[models.AllRaashis((*reqBody).ShukraPlacement.Placement)] {
		if (*reqBody).GuruPlacement.Placement != (*reqBody).ChandraPlacement && (*reqBody).ShukraPlacement.Placement != (*reqBody).ChandraPlacement {
			if slices.Contains(*kujaAspects, models.AllRaashis((*reqBody).ChandraPlacement)) || slices.Contains(*shaniAspects, models.AllRaashis((*reqBody).ChandraPlacement)) || (*reqBody).ShaniPlacement.Placement == (*reqBody).ChandraPlacement || (*reqBody).KujaPlacement.Placement == (*reqBody).ChandraPlacement {
				generalEffects = append(generalEffects, "The native will not have mental health as moon is aspected/conjunct by a malefic with no protection from a benefic")
			}
		}
	}

	return generalEffects
}