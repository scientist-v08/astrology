package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/scientist-v08/bphs/models"
	"github.com/scientist-v08/bphs/services/geoapify"
	"github.com/scientist-v08/bphs/services/shadbala"
	swissEphemeris "github.com/scientist-v08/bphs/services/swiss-ephemeris"
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

	// Validate input.
    if req.BirthDateTime.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "validation failed",
			"details": "birth date and time are required",
		})
		return
    }

    if req.Latitude < -90 ||
        req.Latitude > 90 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "validation failed",
				"details": "latitude must be between -90 and 90",
			})
			return
    }

    if req.Longitude < -180 ||
        req.Longitude > 180 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "validation failed",
				"details": "longitude must be between -180 and 180",
			})
			return
    }

	// 3. Obtain the chart from the DOB
	chart, err := swissEphemeris.CalculateGrahaRashis(req)
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

	// 6. Obtain the balas
	balas, err := shadbala.NewService().Calculate(shadbala.Input{
		Planets: chart.GrahaData, 
		Chart: *result.Horoscope,
		AscendantLong:  chart.AscendantLong,
		MsLong: chart.MsLong,
		SunriseJD:      chart.SunriseJD,
		SunsetJD:       chart.SunsetJD,
		IsDayBirth:     chart.IsDayBirth,
		BirthDateTime:  req.BirthDateTime,
		Latitude:       req.Latitude,
		Longitude:      req.Longitude,
		MeshaJD:           chart.MeshaJD,
		SolarMonthStartJD: chart.SolarMonthStartJD,
	})
	if err != nil {
		fmt.Printf("Shadbala calculation failed: %v\n", err)
		return
	}
	bhavaLords := shadbala.ToBhavaLordAndShadBalaReq(houseEffectsReqBody, *balas)

	bhavaLordsResults := sharedBhavaLordsEffects(bhavaLords)

	// 6. Send back the results
	c.JSON(http.StatusOK, gin.H{
		"result":            result,
		"bhavaLordsResults": bhavaLordsResults,
	})
}

func SearchPlaces(ctx *gin.Context) {

    query := strings.TrimSpace(ctx.Query("text"))

    if len(query) < 3 {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Search query must contain at least 3 characters",
        })
        return
    }

    result, err := geoapify.SearchPlaces(
        ctx.Request.Context(),
        query,
    )

    if err != nil {
        ctx.JSON(http.StatusBadGateway, gin.H{
            "error": "Unable to retrieve place suggestions",
        })
        return
    }

    ctx.JSON(http.StatusOK, result)
}
