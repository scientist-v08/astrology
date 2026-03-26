package services

import (
	"fmt"
	"strconv"

	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
)

func getRashiLordDescription(lord models.Raashyadhipati, rashi string) string {
	switch lord {
	case constants.Kuja:
		return fmt.Sprintf("Those born in %s Raashi are driven, often taking decisive actions and showing strong initiative.", rashi)
	case constants.Shukra:
		return fmt.Sprintf("Those born in %s Raashi are drawn to love, beauty, and fostering harmonious relationships.", rashi)
	case constants.Budha:
		return fmt.Sprintf("Those born in %s Raashi excel in communication, intellect, and adaptability.", rashi)
	case constants.Surya:
		return fmt.Sprintf("Those born in %s Raashi radiate authority, vitality, and confident self-expression.", rashi)
	case constants.Guru:
		return fmt.Sprintf("Those born in %s Raashi embody wisdom, growth, and a benevolent spirit.", rashi)
	case constants.Shani:
		return fmt.Sprintf("Those born in %s Raashi are disciplined, often facing delays but embracing responsibilities.", rashi)
	case constants.Chandra:
		return fmt.Sprintf("Those born in %s Raashi are guided by emotions, intuition, and nurturing qualities.", rashi)
	default:
		return fmt.Sprintf("No description available for %s Raashi.", rashi)
	}
}

func MaitriCalculatorFunc(groomChandraHouse string, brideChandrahouse string) models.KutaTypeRes {
	maitriScore := constants.MaitriKuta[constants.RaashyadhipatiMapStore[groomChandraHouse]][constants.RaashyadhipatiMapStore[brideChandrahouse]]
	maitriComments := getRashiLordDescription(constants.RaashyadhipatiMapStore[groomChandraHouse], groomChandraHouse) +
		getRashiLordDescription(constants.RaashyadhipatiMapStore[brideChandrahouse], brideChandrahouse) +
		" The friendship between these two will have a score of " + strconv.Itoa(int(maitriScore))
	return models.KutaTypeRes{
		Index: 4,
		Score: float32(maitriScore),
		Comments: maitriComments,
	}
}