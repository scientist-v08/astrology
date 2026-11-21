package utils

import (
	"fmt"
	"math"

	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/dto"
)

func CalculateAtmakaraka(
	longitudes map[dto.Raashyadhipati]float64,
) (dto.AtmakarakaResult, error) {

	karakaGrahas := []dto.Raashyadhipati{
        constants.Surya,
        constants.Chandra,
        constants.Kuja,
        constants.Budha,
        constants.Guru,
        constants.Shukra,
        constants.Shani,
        constants.Rahu,
    }

    var result dto.AtmakarakaResult
    highestDegree := -1.0

    for _, graha := range karakaGrahas {

        longitude, exists := longitudes[graha]

        if !exists {
            return dto.AtmakarakaResult{}, fmt.Errorf(
                "missing longitude for graha: %v",
                graha,
            )
        }

        longitude = Normalize360(longitude)

        degreeInRashi := math.Mod(longitude, 30)

        // Rahu's degree is counted in reverse.
        if graha == constants.Rahu {
            degreeInRashi = 30 - degreeInRashi
        }

        if degreeInRashi > highestDegree {
            highestDegree = degreeInRashi

            result.Graha = graha
            result.DegreeInRashi = degreeInRashi
        }
    }

    // Karakamsha = Navamsha Rashi of Atmakaraka.
    karakamsha, err := CalculateNavamsha(
        longitudes[result.Graha],
    )

    if err != nil {
        return dto.AtmakarakaResult{}, err
    }

    result.Karakamsha = karakamsha

    return result, nil
}