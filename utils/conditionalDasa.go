package utils

import "github.com/scientist-v08/bphs/constants"

// raashiIndex returns the 0-based index of a raashi name
func RaashiIndex(name string) (int, bool) {
	for i, r := range constants.RaashiOrder {
		if r == name {
			return i, true
		}
	}
	return -1, false
}

// isMovable (Chara) – starts from itself
func IsMovable(raashi string) bool {
	switch raashi {
	case "Mesha", "Karkataka", "Tula", "Makara":
		return true
	}
	return false
}

// isFixed (Sthira) – starts from 9th
func IsFixed(raashi string) bool {
	switch raashi {
	case "Vrushabha", "Simha", "Vruschika", "Kumbha":
		return true
	}
	return false
}

// isDual (Dwiswabhava) – starts from 5th
func IsDual(raashi string) bool {
	switch raashi {
	case "Mithuna", "Kanya", "Dhanassu", "Meena":
		return true
	}
	return false
}

// isOddSign returns true for odd signs (Mesha, Mithuna, Simha, Tula, Dhanassu, Kumbha)
func IsOddSign(raashi string) bool {
	switch raashi {
	case "Mesha", "Mithuna", "Simha", "Tula", "Dhanassu", "Kumbha":
		return true
	}
	return false
}