package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var validNakshatras = map[string]struct{}{
	"Ashwini":           {},
	"Bharani":           {},
	"Krittika":          {},
	"Rohini":            {},
	"Mrigashira":        {},
	"Ardra":             {},
	"Punarvasu":         {},
	"Pushya":            {},
	"Ashlesha":          {},
	"Magha":             {},
	"Purva Phalguni":    {},
	"Uttara Phalguni":   {},
	"Hasta":             {},
	"Chitra":            {},
	"Swati":             {},
	"Vishakha":          {},
	"Anuradha":          {},
	"Jyeshtha":          {},
	"Mula":              {},
	"Purva Ashadha":     {},
	"Uttara Ashadha":    {},
	"Shravana":          {},
	"Dhanishta":         {},
	"Shatabhisha":       {},
	"Purva Bhadrapada":  {},
	"Uttara Bhadrapada": {},
	"Revati":            {},
}

// isValidNakshatra is the actual validation logic
func isValidNakshatra(fl validator.FieldLevel) bool {
	value, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	_, exists := validNakshatras[value]
	return exists
}

func RegisterNakshatraValidator() error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		return v.RegisterValidation("nakshatra", isValidNakshatra)
	}
	return nil
}