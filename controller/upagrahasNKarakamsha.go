package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/scientist-v08/bphs/models"
	"github.com/scientist-v08/bphs/services"
	"github.com/scientist-v08/bphs/utils"
)

type upagrahaConfig struct {
	name string
	key  string
	fn   func(int16) []string
}

func UpagraHandler(c *gin.Context) {
	// 1. Bind the request body to the variable and check if the expected values have arrived
	var req models.UpagrahasNKarakamshasReqBody

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

	// 2. Get the house palcement of each of the upagraha
	housePlacements := map[string]int16{
		"dhuma": utils.GetHouse(req.Ascendant, req.Dhuma),
		"vyatipata":  utils.GetHouse(req.Ascendant, req.Vyatipata),
		"parivesha":  utils.GetHouse(req.Ascendant, req.Parivesha),
		"indrachapa": utils.GetHouse(req.Ascendant, req.Indrachapa),
		"upaketu":    utils.GetHouse(req.Ascendant, req.Upaketu),
		"gulika":     utils.GetHouse(req.Ascendant, req.Gulika),
		"pranapada":  utils.GetHouse(req.Ascendant, req.Pranapada),
	}

	// 3. Obtain the results from the service
	configs := []upagrahaConfig{
		{"Dhuma", "dhuma", services.DhumaEffectsService},
		{"Vyatipata", "vyatipata", services.VyatipatiEffectsService},
		{"Parivesha", "parivesha", services.ParidhiEffectsService},
		{"Indrachapa", "indrachapa", services.IndraChapaEffectsService},
		{"Upaketu", "upaketu", services.UpaketuEffectsService},
		{"Gulika", "gulika", services.GulikaEffectsService},
		{"Pranapada", "pranapada", services.PranapadaEffectsService},
	}

	var resultsOfUpagrahas []models.UpagrahasNKarakamshasResBody
	for _, cfg := range configs {
		resultsOfUpagrahas = append(resultsOfUpagrahas,
			models.UpagrahasNKarakamshasResBody{
				Upagraha: cfg.name,
				Effects:  cfg.fn(housePlacements[cfg.key]),
			},
		)
	}

	// 4. Proceed with the response
	c.JSON(http.StatusOK, gin.H{
		"upgrahaEffects": resultsOfUpagrahas,
	})
}

func KarakamshaHandler(c *gin.Context) {
	// 1. Obtain the values from the request body and validate the values
	var req models.KarakamshaReqBody

	if err := c.ShouldBindJSON(&req); err != nil {
		var errs []string
		
		// err will be validator.ValidationErrors if validation failed
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			for _, e := range validationErrs {
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

	// 2. Get the house palcement of each of the graha
	housePlacements := map[string]int16{
		"Surya": utils.GetHouse(req.Karakamsha, req.Surya),
		"Budha": utils.GetHouse(req.Karakamsha, req.Budha),
		"Shukra": utils.GetHouse(req.Karakamsha, req.Shukra),
		"Chandra": utils.GetHouse(req.Karakamsha, req.Chandra),
		"Rahu": utils.GetHouse(req.Karakamsha, req.Rahu),
		"Ketu": utils.GetHouse(req.Karakamsha, req.Ketu),
		"Kuja": utils.GetHouse(req.Karakamsha, req.Kuja),
		"Guru": utils.GetHouse(req.Karakamsha, req.Guru),
		"Shani": utils.GetHouse(req.Karakamsha, req.Shani),
		"Lagna": utils.GetHouse(req.Karakamsha, req.Ascendant),
	}

	// 3. Return the results
	c.JSON(http.StatusOK, gin.H{
		"results": services.KarakamshaEffectsService(&housePlacements, &req),
	})
}