package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
	"github.com/scientist-v08/bphs/services"
	"github.com/scientist-v08/bphs/services/shadbala"
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

	// 6. Obtain the balas
	balas, err := shadbala.NewService().Calculate(shadbala.Input{Planets: chart.GrahaData, Chart: *result.Horoscope})
	if err != nil {
		fmt.Printf("Shadbala calculation failed: %v\n", err)
		return
	}
	order := []models.Raashyadhipati{
		constants.Surya, constants.Chandra, constants.Kuja,
		constants.Budha, constants.Guru, constants.Shukra, constants.Shani,
	}

	headers := []string{"Sun", "Moon", "Mars", "Mercury", "Jupiter", "Venus", "Saturn"}

	// Helper to print a row
	printRow := func(label string, getter func(shadbala.SthanaBala) float64) {
		fmt.Printf("%-12s", label)
		for _, g := range order {
			fmt.Printf("%10.2f", getter(balas.Sthana[g]))
		}
		fmt.Println()
	}

	// Header
	fmt.Printf("%-12s", "")
	for _, h := range headers {
		fmt.Printf("%10s", h)
	}
	fmt.Println()

	// Separator
	fmt.Printf("%-12s", "")
	for range headers {
		fmt.Printf("%10s", "--------")
	}
	fmt.Println()

	// Rows
	printRow("Uchcha", func(b shadbala.SthanaBala) float64 { return b.Uchcha })
	printRow("Saptavarga", func(b shadbala.SthanaBala) float64 { return b.Saptavarga })
	printRow("OjaYugma", func(b shadbala.SthanaBala) float64 { return b.OjaYugma })
	printRow("Kendra", func(b shadbala.SthanaBala) float64 { return b.Kendra })
	printRow("Drekkana", func(b shadbala.SthanaBala) float64 { return b.Drekkana })
	printRow("Total", func(b shadbala.SthanaBala) float64 { return b.Total })

	// 6. Send back the results
	c.JSON(http.StatusOK, result)
}