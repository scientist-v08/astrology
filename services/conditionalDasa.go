package services

import (
	"fmt"
	"math"

	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
	"github.com/scientist-v08/bphs/utils"
)

// CalculateNavamshaLagna computes the Navamsha (D9) Lagna raashi
// from the Rasi Ascendant and its degree within the sign.
func CalculateNavamshaLagna(req *models.ConditionalDasaReq) (string, error) {
	asc := req.Ascendant
	deg := float64(req.AscendantDeg.Deg) + (float64(req.AscendantDeg.Min) / 60)

	// Validate degree range (0 ≤ deg < 30)
	if deg < 0 || deg >= 30 {
		return "", fmt.Errorf("ascendantDeg must be in [0, 30): got %v", deg)
	}

	idx, ok := utils.RaashiIndex(asc)
	if !ok {
		return "", fmt.Errorf("invalid ascendant raashi: %s", asc)
	}

	// Determine starting sign for the first Navamsha of this Rasi lagna
	var startIdx int
	switch {
	case utils.IsMovable(asc):
		// Movable → starts from the same sign
		startIdx = idx
	case utils.IsFixed(asc):
		// Fixed → starts from the 9th sign (idx + 8)
		startIdx = (idx + 8) % 12
	case utils.IsDual(asc):
		// Dual → starts from the 5th sign (idx + 4)
		startIdx = (idx + 4) % 12
	default:
		return "", fmt.Errorf("unknown nature for raashi: %s", asc)
	}

	// Each Navamsha = 30° / 9 = 3°20' = 10/3 degrees
	const navamshaSize = 30.0 / 9.0 // ≈ 3.333...

	// Navamsha number (1-based) inside the sign
	navNum := int(math.Floor(deg/navamshaSize)) + 1 // 1 … 9

	// Final Navamsha lagna = start + (navNum-1) signs
	finalIdx := (startIdx + navNum - 1) % 12
	return constants.RaashiOrder[finalIdx], nil
}

// CalculateHoraLagna computes the Hora (D2) Lagna.
// Result is always either "Simha" (Solar Hora) or "Karkataka" (Lunar Hora).
func CalculateHoraLagna(req *models.ConditionalDasaReq) (string, error) {
	asc := req.Ascendant
	deg := float64(req.AscendantDeg.Deg) + (float64(req.AscendantDeg.Min) / 60)

	if deg < 0 || deg >= 30 {
		return "", fmt.Errorf("ascendantDeg must be in [0, 30): got %v", deg)
	}

	if _, ok := utils.RaashiIndex(asc); !ok {
		return "", fmt.Errorf("invalid ascendant raashi: %s", asc)
	}

	// First half (0°–15°) and second half (15°–30°)
	firstHalf := deg < 15.0

	odd := utils.IsOddSign(asc)

	switch {
	case odd && firstHalf:
		return "Simha", nil // Solar Hora
	case odd && !firstHalf:
		return "Karkataka", nil // Lunar Hora
	case !odd && firstHalf:
		return "Karkataka", nil // Lunar Hora
	default: // !odd && !firstHalf
		return "Simha", nil // Solar Hora
	}
}

// CalculateDwadasamsaLagna computes the Dwadasamsa (D12) Lagna raashi.
// Each portion = 2°30'. Counting always starts from the Rasi lagna itself and proceeds forward.
func CalculateDwadasamsaLagna(req *models.ConditionalDasaReq) (string, error) {
	asc := req.Ascendant
	deg := float64(req.AscendantDeg.Deg) + (float64(req.AscendantDeg.Min) / 60)

	if deg < 0 || deg >= 30 {
		return "", fmt.Errorf("ascendantDeg must be in [0, 30): got %v", deg)
	}

	idx, ok := utils.RaashiIndex(asc)
	if !ok {
		return "", fmt.Errorf("invalid ascendant raashi: %s", asc)
	}

	const portionSize = 30.0 / 12.0 // 2.5°

	// 1-based portion number (1 … 12)
	portionNum := int(math.Floor(deg/portionSize)) + 1

	// Start from the same sign and move (portionNum-1) signs forward.
	// Using Mod for safety (handles any future reverse-counting needs).
	finalIdx := int(Mod(int8(idx+(portionNum-1)), 12))

	return constants.RaashiOrder[finalIdx], nil
}

func GetLordHouse(req *models.ConditionalDasaReq, ascLord models.Raashyadhipati) string {
	switch ascLord {
		case    constants.Kuja:
			return req.KujaPlacement
		case 	constants.Shukra:
			return req.ShukraPlacement
		case   	constants.Budha:
			return req.BudhaPlacement
		case 	constants.Chandra:
			return req.ChandraPlacement
		case    constants.Surya:
			return req.SuryaPlacement
		case  	constants.Guru:
			return req.GuruPlacement
		case    constants.Shani:
			return req.ShaniPlacement
	}
	return ""
}

func CalculateApplicableConditionalDasa(req models.ConditionalDasaReq) []string {
	var applicableDasa []string
	// Collect required information for Chandra house placements
	chandraHousePlacements := map[string]int16{
		"Surya": utils.GetHouse(req.ChandraPlacement, req.SuryaPlacement),
	}
	suryaHousePlacements := map[string]int16{
		"Chandra": utils.GetHouse(req.SuryaPlacement, req.ChandraPlacement),
		"Lagna": utils.GetHouse(req.SuryaPlacement, req.Ascendant),
	}
	housePlacements := map[string]int16{
		"Surya": utils.GetHouse(req.Ascendant, req.SuryaPlacement),
		"Budha": utils.GetHouse(req.Ascendant, req.BudhaPlacement),
		"Shukra": utils.GetHouse(req.Ascendant, req.ShukraPlacement),
		"Chandra": utils.GetHouse(req.Ascendant, req.ChandraPlacement),
		"RahuPlacement": utils.GetHouse(req.Ascendant, req.RahuPlacement),
		"KetuPlacement": utils.GetHouse(req.Ascendant, req.KetuPlacement),
		"Kuja": utils.GetHouse(req.Ascendant, req.KujaPlacement),
		"Guru": utils.GetHouse(req.Ascendant, req.GuruPlacement),
		"Shani": utils.GetHouse(req.Ascendant, req.ShaniPlacement),
	}
	hora,_ := CalculateHoraLagna(&req)
	lagnaNavamsha,_ := CalculateNavamshaLagna(&req)
	ascendantLord := constants.RaashyadhipatiMapStore[req.Ascendant]
	houseOfAscendantLord := GetLordHouse(&req, ascendantLord)
	ascLordPlacement := utils.GetHouse(req.Ascendant, houseOfAscendantLord)
	ascLordPlacements := map[string]int16{
		"Rahu": utils.GetHouse(houseOfAscendantLord, req.RahuPlacement),
	}
	dwadasamsaLagna,_ := CalculateDwadasamsaLagna(&req)
	tenthHouse := constants.ReverseNumericalMappingsStore[req.Ascendant][10]
	tenthLord := constants.RaashyadhipatiMapStore[string(tenthHouse)]
	tenthLordPlacement := utils.GetHouse(req.Ascendant, GetLordHouse(&req, tenthLord))
	seventhHouse := constants.ReverseNumericalMappingsStore[req.Ascendant][7]
	seventhLord := constants.RaashyadhipatiMapStore[string(seventhHouse)]
	seventhLordPlacement := utils.GetHouse(req.Ascendant, GetLordHouse(&req, seventhLord))

	// 1. Check for Shodasottari daśā
	if hora == "Simha" && chandraHousePlacements["Surya"] > 6 {
		applicableDasa = append(applicableDasa, "Shodasottari daśā is applicable. Jupiter is the controlling Graha. Surya hora")
	} else if hora == "Karkataka" && suryaHousePlacements["Chandra"] > 6 {
		applicableDasa = append(applicableDasa, "Shodasottari daśā is applicable. Jupiter is the controlling Graha.")
	}

	// 2. Check for Śattriṃśa-sama daśā
	if hora == "Simha" && housePlacements["Surya"] > 6 {
		applicableDasa = append(applicableDasa, "Śattriṃśa-sama is applicable. Mercury is the controlling Graha. Surya Hora")
	} else if hora == "Karkataka" && suryaHousePlacements["Lagna"] > 6 {
		applicableDasa = append(applicableDasa, "Śattriṃśa-sama is applicable. Mercury is the controlling Graha.")
	}

	// 3. Check for Dwadasottari daśā
	if req.Ascendant == "Vrushabha" || lagnaNavamsha == "Tula" {
		applicableDasa = append(applicableDasa, "Dwadasottari daśā is applicable. Ketu is the controlling Graha.")
	}

	// 4. Check for Ashtottari daśā
	if housePlacements["RahuPlacement"] != 1 && (utils.IsKendra(ascLordPlacements["Rahu"]) || utils.IsKona(ascLordPlacements["Rahu"])) {
		applicableDasa = append(applicableDasa, "Ashtottari daśā is applicable. Mars is the controlling Graha.")
	}

	// 5. Check for Panchottari daśā
	if dwadasamsaLagna == "Karkataka" {
		applicableDasa = append(applicableDasa, "Panchottari daśā is applicable. Venus is the controlling Graha.")
	}

	// 6. Check for Satabdika daśā
	if req.Ascendant == lagnaNavamsha {
		applicableDasa = append(applicableDasa, "Satabdika daśā is applicable. Sun is the controlling Graha.")
	}

	// 7. Check for Chaturaaseeti daśā
	if tenthLordPlacement == int16(10) {
		applicableDasa = append(applicableDasa, "Chaturaaseeti daśā is applicable. Saturn is the controlling Graha.")
	}

	// 8. Check for Dwisaptati sama daśā
	if ascLordPlacement == int16(7) || seventhLordPlacement == int16(1) {
		applicableDasa = append(applicableDasa, "Dwisaptati sama daśā is applicable. Rahu is the controlling Graha.")
	}

	// 9. Check for Shashtisama daśā
	if housePlacements["Surya"] == int16(1) {
		applicableDasa = append(applicableDasa, "Shashtisama daśā is applicable. Moon is the controlling Graha.")
	}

	return applicableDasa
}