package utils

import "math"

func Normalize360(degree float64) float64 {
	degree = math.Mod(degree, 360)

	if degree < 0 {
		degree += 360
	}

	return degree
}