package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/scientist-v08/bphs/models"
	"github.com/scientist-v08/bphs/services"
)

func ConditionalDasaController(c *gin.Context) {

	// 1. Obtain the request body and validate it.
	var requestBody models.ConditionalDasaReq
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		// err will be validator.ValidationErrors if validation failed
		var errs []string

		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			for _, e := range validationErrs {
				// You can customize messages here
				switch e.Tag() {
				case "raashi":
					errs = append(errs, e.Field()+" must be a valid rāśi")
				default:
					errs = append(errs, e.Error())
				}
			}
		} else {
			errs = append(errs, err.Error())
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "validation failed",
			"details": errs,
		})
		return
	}

	// 2. Pass the request body into the service and obtain the results
	obtainedResults := services.CalculateApplicableConditionalDasa(requestBody)
	if len(obtainedResults) == 0 {
		c.JSON(http.StatusOK, []string{"Vimsottri Dasa is applicable"})
		return
	}

	c.JSON(http.StatusOK, obtainedResults)
}