package utils

import (
	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/dto"
)

func IsKendra(h int16) bool {
	switch h {
	case 1, 4, 7, 10:
		return true
	default:
		return false
	}
}

func IsKona(h int16) bool {
	switch h {
	case 1, 5, 9:
		return true
	default:
		return false
	}
}

func IsDusthana(h int16) bool {
	switch h {
	case 6, 8, 12:
		return true
	default:
		return false
	}
}

func GetGuruAspectRaashis(lagna dto.AllRaashis) []dto.AllRaashis {
	aspects, exists := constants.BrihaspatiAspectsStore[lagna]
	if !exists {
		return nil // or []AllRaashis{} depending on your preference
	}

	result := make([]dto.AllRaashis, 0, len(aspects))
	for _, item := range aspects {
		result = append(result, item.House)
	}
	return result
}

func GetShaniAspectRaashis(lagna dto.AllRaashis) []dto.AllRaashis {
	aspects, exists := constants.ShaniAspectsStore[lagna]
	if !exists {
		return nil // or []AllRaashis{} depending on your preference
	}

	result := make([]dto.AllRaashis, 0, len(aspects))
	for _, item := range aspects {
		result = append(result, item.House)
	}
	return result
}

func GetKujaAspectRaashis(lagna dto.AllRaashis) []dto.AllRaashis {
	aspects, exists := constants.KujaAspectsStore[lagna]
	if !exists {
		return nil // or []AllRaashis{} depending on your preference
	}

	result := make([]dto.AllRaashis, 0, len(aspects))
	for _, item := range aspects {
		result = append(result, item.House)
	}
	return result
}

func GetHouse(ascendant string, raashiPlacement string) int16 {
	house, _ := constants.NumericalMappingsStore[ascendant][dto.AllRaashis(raashiPlacement)]
	return house
}

func GetHouseName(ascendant string, raashiPlacement int16) dto.AllRaashis {
	house, _ := constants.ReverseNumericalMappingsStore[ascendant][raashiPlacement]
	return house
}

func IsGrahaCombust(
	graha dto.Raashyadhipati,
	kujaCombust bool,
	shukraCombust bool,
	budhCombust bool,
	guruCombust bool,
	shaniCombust bool,
) bool {
	if graha == constants.Kuja && kujaCombust {
		return true
	}
	if graha == constants.Shukra && shukraCombust {
		return true
	}
	if graha == constants.Budha && budhCombust {
		return true
	}
	if graha == constants.Guru && guruCombust {
		return true
	}
	if graha == constants.Shani && shaniCombust {
		return true
	}
	return false
}

func IsLordInOwnHouse(lord dto.Raashyadhipati, houseName dto.AllRaashis) bool {
    switch lord {
    case constants.Kuja:
        return houseName == constants.Mesha || houseName == constants.Vruschika

    case constants.Shukra:
        return houseName == constants.Vrushabha || houseName == constants.Tula

    case constants.Budha:
        return houseName == constants.Mithuna || houseName == constants.Kanya

    case constants.Chandra:
        return houseName == constants.Karkataka

    case constants.Surya:
        return houseName == constants.Simha

    case constants.Guru:
        return houseName == constants.Dhanassu || houseName == constants.Meena

    case constants.Shani:
        return houseName == constants.Makara || houseName == constants.Kumbha

    default:
        return false
    }
}

func IsLordCombust(lord dto.Raashyadhipati, reqBody *dto.HouseEffectsReqBody) bool {
	switch lord {
    case constants.Kuja:
        return (*reqBody).KujaPlacement.Combustion

    case constants.Shukra:
        return (*reqBody).ShukraPlacement.Combustion

    case constants.Budha:
        return (*reqBody).BudhaPlacement.Combustion

    case constants.Chandra:
        return false

    case constants.Surya:
        return false

    case constants.Guru:
        return (*reqBody).GuruPlacement.Combustion

    case constants.Shani:
        return (*reqBody).ShaniPlacement.Combustion

    default:
        return false
    }
}

func IsGrahaStrong(
	isNotCombust bool,
	isNotDebilitated bool,
	isNotInEnemySign bool,
	isNotInDusthana bool,
) bool {

	conditions := []bool{
		isNotCombust,
		isNotDebilitated,
		isNotInEnemySign,
		isNotInDusthana,
	}

	trueCount := 0
	for _, condition := range conditions {
		if condition {
			trueCount++
			if trueCount >= 2 {
				return true
			}
		}
	}

	return false
}
