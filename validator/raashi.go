package validator // or models / pkg / internal — your choice

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var validRaashis = map[string]struct{}{
	"Mesha":     {},
	"Vrushabha": {},
	"Mithuna":   {},
	"Karkataka": {},
	"Simha":     {},
	"Kanya":     {},
	"Tula":      {},
	"Vruschika": {},
	"Dhanassu":  {},
	"Makara":    {},
	"Kumbha":    {},
	"Meena":     {},
}

// isValidRaashi is the actual validation logic
func isValidRaashi(fl validator.FieldLevel) bool {
	value, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	_, exists := validRaashis[value]
	return exists
}

func RegisterRaashiValidator() error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		return v.RegisterValidation("raashi", isValidRaashi)
	}
	return nil
}