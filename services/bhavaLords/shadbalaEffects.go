package bhavaLords

import (
	"fmt"
	"slices"
	"sort"

	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/dto"
	"github.com/scientist-v08/bphs/utils"
)

var grahaOrder = []string{"Surya", "Chandra", "Kuja", "Budha", "Guru", "Shukra", "Shani"}

// rankGrahas sorts the 7 grahas in descending order of strength and returns them ordered by rank.
func rankGrahas(values []float32) []dto.ShadBalaResField {
	ranked := make([]dto.ShadBalaResField, len(grahaOrder))
	items := make([]dto.RankedItem, len(grahaOrder))

	for i := range len(grahaOrder) {
		ranked[i] = dto.ShadBalaResField{
			Graha:         grahaOrder[i],
			EffectiveBala: values[i],
			Rank:          0,
		}
		items[i] = dto.RankedItem{Idx: i, Value: values[i]}
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].Value > items[j].Value
	})

	for rank, item := range items {
		ranked[item.Idx].Rank = int8(rank + 1)
	}

	sort.Slice(ranked, func(i, j int) bool {
		return ranked[i].Rank < ranked[j].Rank
	})

	return ranked
}

// GetBhavaLordEffect returns the effect description for a bhava lord placed in a given house.
func GetBhavaLordEffect(house int8, lordPlacement int8) string {
	switch house {
	case 1:
		return FirstBhavaLordEffect(lordPlacement)
	case 2:
		return SecondBhavaLordEffect(lordPlacement)
	case 3:
		return ThirdBhavaLordEffect(lordPlacement)
	case 4:
		return FourthBhavaLordEffect(lordPlacement)
	case 5:
		return FifthBhavaLordEffect(lordPlacement)
	case 6:
		return SixthBhavaLordEffect(lordPlacement)
	case 7:
		return SeventhBhavaLordEffect(lordPlacement)
	case 8:
		return EighthBhavaLordEffect(lordPlacement)
	case 9:
		return NinthBhavaLordEffect(lordPlacement)
	case 10:
		return TenthBhavaLordEffect(lordPlacement)
	case 11:
		return EleventhBhavaLordEffect(lordPlacement)
	case 12:
		return TwelfthBhavaLordEffect(lordPlacement)
	default:
		return ""
	}
}

// GetHousePlacements maps each graha name to its house number from the chart ascendant.
func GetHousePlacements(req dto.BhavaLordAndShadBalaReqBody) map[string]int16 {
	return map[string]int16{
		"Surya":   utils.GetHouse(req.Ascendant, req.SuryaPlacement),
		"Budha":   utils.GetHouse(req.Ascendant, req.BudhaPlacement),
		"Shukra":  utils.GetHouse(req.Ascendant, req.ShukraPlacement),
		"Chandra": utils.GetHouse(req.Ascendant, req.ChandraPlacement),
		"Rahu":    utils.GetHouse(req.Ascendant, req.RahuPlacement),
		"Ketu":    utils.GetHouse(req.Ascendant, req.KetuPlacement),
		"Kuja":    utils.GetHouse(req.Ascendant, req.KujaPlacement),
		"Guru":    utils.GetHouse(req.Ascendant, req.GuruPlacement),
		"Shani":   utils.GetHouse(req.Ascendant, req.ShaniPlacement),
	}
}

// CalculateSthanaBalaEffects ranks grahas by effective Sthana Bala and returns the top-2 house theme summary and raw total.
func CalculateSthanaBalaEffects(
	req dto.BhavaLordAndShadBalaReqBody,
	housePlacements map[string]int16,
) ([]dto.ShadBalaResField, string, float32) {
	sthanaBalaMinimums := []float32{165.0, 133.0, 96.0, 165.0, 165.0, 133.0, 96.0}
	effectiveSthanaBala := make([]float32, len(grahaOrder))

	for i := range len(grahaOrder) {
		sumOfSthana := req.DrigBala[i] + req.SthanaBala[i]
		effectiveSthanaBala[i] = sumOfSthana / sthanaBalaMinimums[i]
	}

	var totalSthana float32
	for i := range req.SthanaBala {
		totalSthana += req.SthanaBala[i]
	}

	ranks := rankGrahas(effectiveSthanaBala)
	top3SthanaLords := "The native will obtain the highest level of mental satisfaction / success in " +
		constants.Themes[housePlacements[ranks[0].Graha]-1] +
		" and will dedicate his highest thinking towards " +
		constants.Themes[housePlacements[ranks[1].Graha]-1]

	return ranks, top3SthanaLords, totalSthana
}

// CalculateKaalaBalaEffects ranks grahas by effective Kaala Bala and returns the top-3 dasha lords summary and raw total.
func CalculateKaalaBalaEffects(req dto.BhavaLordAndShadBalaReqBody) ([]dto.ShadBalaResField, string, float32) {
	kaalaBalaMinimums := []float32{50, 30, 40, 50, 50, 30, 40}
	effectiveKaalaBala := make([]float32, len(grahaOrder))

	for i := range len(grahaOrder) {
		sumOfKaala := req.DrigBala[i] + req.KaalaBala[i]
		effectiveKaalaBala[i] = sumOfKaala / kaalaBalaMinimums[i]
	}

	var totalKaala float32
	for i := range req.KaalaBala {
		totalKaala += req.KaalaBala[i]
	}

	ranks := rankGrahas(effectiveKaalaBala)
	top3KaalaLords := "The native will obtain the highest success during the dasha/sub-dasha/sub-sub-dasha of the following lords " +
		ranks[0].Graha + "," + ranks[1].Graha + "," + ranks[2].Graha +
		". When the time of the dasha belongs to at least 2 of these lords the highest success will be obtained"

	return ranks, top3KaalaLords, totalKaala
}

// CalculateDigBalaEffects ranks grahas by effective Dig Bala and builds directional commentary.
func CalculateDigBalaEffects(req dto.BhavaLordAndShadBalaReqBody) string {
	digBalaMinimums := []float32{35, 50, 30, 35, 35, 50, 30}
	effectiveDigBala := make([]float32, len(grahaOrder))

	for i := range len(grahaOrder) {
		effectiveDigBala[i] = req.DigBala[i] / digBalaMinimums[i]
	}

	ranks := rankGrahas(effectiveDigBala)

	grahaDirections := map[string]string{
		"Surya":   "East",
		"Chandra": "North-West",
		"Kuja":    "South",
		"Budha":   "North",
		"Guru":    "North-East",
		"Shukra":  "South-East",
		"Shani":   "West",
	}

	var digBalaStr1 string
	dir1 := grahaDirections[ranks[0].Graha]
	if len(ranks) > 0 {
		digBalaStr1 = fmt.Sprintf(
			". Going in the %s direction (ruled by %s with bala %.2f) from the place of birth will give success, material gain, and professional achievements",
			dir1,
			ranks[0].Graha,
			ranks[0].EffectiveBala,
		)
	}

	var digBalaStr2 string
	if len(ranks) > 1 {
		dir2 := grahaDirections[ranks[1].Graha]
		digBalaStr2 = fmt.Sprintf(
			". Going in the %s direction (ruled by %s with bala %.2f) will help the person find increased mental energy, peace, and creative capabilities.",
			dir2,
			ranks[1].Graha,
			ranks[1].EffectiveBala,
		)
	}

	var digBalaStr3 string
	if len(ranks) > 2 {
		dir3 := grahaDirections[ranks[2].Graha]
		digBalaStr3 = fmt.Sprintf(
			" Going in the %s direction (ruled by %s with bala %.2f) is where the person may find a compatible karmic partner or companion in life.",
			dir3,
			ranks[2].Graha,
			ranks[2].EffectiveBala,
		)
	}

	return digBalaStr1 + digBalaStr2 + digBalaStr3
}

// CalculateChestaBalaEffects ranks grahas by effective Chesta Bala and builds dasha timing commentary.
func CalculateChestaBalaEffects(req dto.BhavaLordAndShadBalaReqBody) string {
	chestaBalaMinimums := []float32{112, 100, 67, 112, 112, 100, 67}
	effectiveChestaBala := make([]float32, len(grahaOrder))
	effectiveChestaBalaAct := make([]float32, len(grahaOrder))

	for i := range len(grahaOrder) {
		switch i {
		case 0:
			effectiveChestaBala[i] = float32(req.SuryaAyana)
		case 1:
			effectiveChestaBala[i] = float32(req.ChandraPaksha)
		default:
			effectiveChestaBala[i] = req.ChestaBala[i] * 2
		}
		effectiveChestaBalaAct[i] = effectiveChestaBala[i] / chestaBalaMinimums[i]
	}

	ranks := rankGrahas(effectiveChestaBalaAct)

	var chestaTopStr, chestaBottomStr string
	if len(ranks) >= 7 {
		chestaTopStr = fmt.Sprintf(
			" The top 3 grahas in Chesta Bala are %s(%.2f), %s(%.2f), and %s(%.2f). In their sub-period (Antardasha) or sub-sub-periods (Pratyantardasha), any important travel, launching a project, or making a move to achieve an objective will succeed.",
			ranks[0].Graha,
			ranks[0].EffectiveBala,
			ranks[1].Graha,
			ranks[1].EffectiveBala,
			ranks[2].Graha,
			ranks[2].EffectiveBala,
		)

		chestaBottomStr = fmt.Sprintf(
			" The lowest 3 grahas in Chesta Bala are %s(%.2f), %s(%.2f), and %s(%.2f). In their sub-period (Antardasha) or sub-sub-periods (Pratyantardasha), important travel, launching a project, or making a move to achieve an objective may not be achieved.",
			ranks[4].Graha,
			ranks[4].EffectiveBala,
			ranks[5].Graha,
			ranks[5].EffectiveBala,
			ranks[6].Graha,
			ranks[6].EffectiveBala,
		)
	}

	return chestaTopStr + chestaBottomStr
}

// CalculateDrikBalaEffects evaluates total Drik Bala and returns longevity commentary when negative.
func CalculateDrikBalaEffects(req dto.BhavaLordAndShadBalaReqBody) string {
	var totalDrik float32
	for i := range len(grahaOrder) {
		totalDrik += req.DrigBala[i]
	}

	if totalDrik < float32(0) {
		return ". Due to Drik bala the native is likely to have a long life span. But also consult other longevity methods for confirmation."
	}
	return ""
}

// CalculateShadbalaEffects ranks grahas by Shadbala, computes top-3 bhava lord effects, and builds overall commentary.
func CalculateShadbalaEffects(
	req dto.BhavaLordAndShadBalaReqBody,
	housePlacements map[string]int16,
	totalSthana float32,
	totalKaala float32,
) ([]dto.ShadBalaResField, []string, string) {
	ranks := rankGrahas(req.ShadBala)

	top3Lords := [3]string{ranks[0].Graha, ranks[1].Graha, ranks[2].Graha}

	housesByPlanet := constants.AscendantToRashyadhipatiMapStore[dto.AllRaashis(req.Ascendant)]
	effects := []string{}
	addEffects := func(planetStr string, placement int8) {
		planet := dto.Raashyadhipati(planetStr)
		houses := housesByPlanet[planet]
		for _, houseNum := range houses {
			effect := GetBhavaLordEffect(houseNum, placement)
			if effect != "" {
				effects = append(effects, effect)
			}
		}
	}

	for i := range 3 {
		lord := top3Lords[i]
		placement := housePlacements[lord]
		addEffects(lord, int8(placement))
	}

	isMaleficAmongstTop3 := false
	for i := range 3 {
		if slices.Contains([]string{"Surya", "Kuja", "Shani"}, top3Lords[i]) {
			isMaleficAmongstTop3 = true
			break
		}
	}

	var conclusion string
	if totalSthana > totalKaala+0.01 {
		conclusion = "Sthana Bala is greater overall leading to long term success"
	} else if totalKaala > totalSthana+0.01 {
		conclusion = "Kaala Bala is greater overall leading to small term success"
	} else {
		conclusion = "Sthana Bala and Kaala Bala are approximately equal"
	}

	var shadBalaComments string
	if isMaleficAmongstTop3 {
		shadBalaComments = "The individual will obtain success after a long and hard struggle. " + conclusion
	} else {
		shadBalaComments = "The individual will gain success smoothly. " + conclusion
	}

	return ranks, effects, shadBalaComments
}
