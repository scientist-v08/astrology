package utils

import (
	"fmt"
	"math"

	"github.com/scientist-v08/bphs/constants"
)

func CalculateNavamsha(longitude float64) (string, error) {
	rashiIndex := int(longitude / 30)
	rashi := constants.RaashiOrder[rashiIndex]

	degreeInRashi := math.Mod(longitude, 30)

	const navamshaSize = 30.0 / 9.0
	navNum := int(math.Floor(degreeInRashi / navamshaSize))

	var startIdx int

	switch {
	case IsMovable(rashi):
		startIdx = rashiIndex

	case IsFixed(rashi):
		startIdx = (rashiIndex + 8) % 12

	case IsDual(rashi):
		startIdx = (rashiIndex + 4) % 12

	default:
		return "", fmt.Errorf("unknown nature for raashi: %s", rashi)
	}

	finalIdx := (startIdx + navNum) % 12

	return constants.RaashiOrder[finalIdx], nil
}