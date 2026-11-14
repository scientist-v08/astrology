package utils

import (
	"fmt"
	"math"

	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
)

func CalculateAtmakaraka(
	longitudes map[models.Raashyadhipati]float64,
) (models.AtmakarakaResult, error) {

	karakaGrahas := []models.Raashyadhipati{
        constants.Surya,
        constants.Chandra,
        constants.Kuja,
        constants.Budha,
        constants.Guru,
        constants.Shukra,
        constants.Shani,
        constants.Rahu,
    }

    var result models.AtmakarakaResult
    highestDegree := -1.0

    for _, graha := range karakaGrahas {

        longitude, exists := longitudes[graha]

        if !exists {
            return models.AtmakarakaResult{}, fmt.Errorf(
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
        return models.AtmakarakaResult{}, err
    }

    result.Karakamsha = karakamsha

    return result, nil
}