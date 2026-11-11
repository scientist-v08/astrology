package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/scientist-v08/bphs/models"
	"github.com/scientist-v08/bphs/services"
	"github.com/scientist-v08/bphs/utils"
)

func DobForAll(c *gin.Context) {
	// 1. Obtain the request body from the post request and validate them.
	var req models.BirthChartReq

	// 2. Validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
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

	// 3. Obtain the chart from the DOB
	chart, err := services.CalculateGrahaRashis(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 4. Obtain the correct request body format
	houseEffectsReqBody := utils.ToHouseEffectsReqBody(chart)

	// 5. Obtain the results
	result := calculateHouseEffects(&houseEffectsReqBody)

	// 6. Send back the results
	c.JSON(http.StatusOK, result)
}