package validator // or models / pkg / internal — your choice

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var validGrahas = map[string]struct{}{
	"Surya":     {},
	"Chandra": {},
	"Kuja":   {},
	"Budha": {},
	"Guru":     {},
	"Shukra":     {},
	"Shani":      {},
}

// isValidGraha is the actual validation logic
func isValidGraha(fl validator.FieldLevel) bool {
	value, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	_, exists := validGrahas[value]
	return exists
}

func RegisterGrahaValidator() error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		return v.RegisterValidation("graha", isValidGraha)
	}
	return nil
}