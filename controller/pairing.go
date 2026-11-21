package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/scientist-v08/bphs/dto"
	"github.com/scientist-v08/bphs/services/couple"
)

func evaluateDoshaSamya(groom, bride []dto.DoshaSamyaRes) string {
    var groomTotal, brideTotal float32
    for _, d := range groom { groomTotal += d.DoshaScore }
    for _, d := range bride { brideTotal += d.DoshaScore }

    if groomTotal >= brideTotal {
        diff := groomTotal - brideTotal
        if diff == 0 {
            return "Perfect balance (Dosha Samya achieved)"
        }
        return fmt.Sprintf("Good – Groom has higher/equal dosha (%.2f vs %.2f)", groomTotal, brideTotal)
    }

    return fmt.Sprintf("Imbalance – Bride has higher dosha (%.2f > %.2f). Not recommended without strong remedies.", brideTotal, groomTotal)
}

func PairingHandler(c *gin.Context) {
	var req dto.PairingReqBody
	if err := c.ShouldBindJSON(&req); err != nil {
		// err will be validator.ValidationErrors if validation failed
		var errs []string
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			for _, e := range validationErrs {
				// You can customize messages here
				switch e.Tag() {
				case "raashi":
					errs = append(errs, e.Field()+" must be a valid rāśi")
				case "nakshatra":
					errs = append(errs, e.Field()+" must be a valid nakshatra")
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

	// First get the Kuta scores
	kutaScores := make([]dto.KutaTypeRes, 9)
	kutaScores[0] = couple.VarnaCalculatorFunc(req.GroomChandraPlacement, req.BrideChandraPlacement)
	kutaScores[1] = couple.VashyaCalculatorFunc(req.GroomChandraPlacement, req.BrideChandraPlacement, req.GroomFirstHalf, req.BrideFirstHalf)
	kutaScores[2] = couple.TaraCalculatorFunc(req.GroomNakshatra, req.BrideNakshatra)
	kutaScores[3] = couple.YoniCalculatorFunc(req.GroomNakshatra, req.BrideNakshatra)
	kutaScores[4] = couple.MaitriCalculatorFunc(req.GroomChandraPlacement, req.BrideChandraPlacement)
	kutaScores[5] = couple.GanaCalculatorHandler(req.GroomNakshatra, req.BrideNakshatra, req.GroomChandraPlacement, req.BrideChandraPlacement)
	kutaScores[6] = couple.BhakutaCalculatorFunc(req.GroomChandraPlacement, req.BrideChandraPlacement)
	kutaScores[7] = couple.NadiCalculatorFunc(req.GroomNakshatra, req.BrideNakshatra)

	var totalKutaScore float32
	for i := range 8 {
		totalKutaScore = totalKutaScore + kutaScores[i].Score
	}

	var commentOnTotalKutaScore string
	if totalKutaScore >= 18 {
		commentOnTotalKutaScore = "The total kuta score is good for marriage"
	} else {
		commentOnTotalKutaScore = "The total kuta score is not good for marriage"
	}

	kutaScores[8] = dto.KutaTypeRes{
		Index: 8,
		Score: totalKutaScore,
		Comments: commentOnTotalKutaScore,
	}

	// Second obtain the rajju and other comments
	rajjuAndOtherComments := couple.RajjuAndOtherDoshasCalculator(req.GroomNakshatra, req.BrideNakshatra, req.GroomChandraPlacement, req.BrideChandraPlacement)

	// Third check for Dosha Samya
	groomDosha, brideDosha := couple.DoshaSamyaFunc(&req)

	// Fourth check whose is greater
	whoseIsGreater := evaluateDoshaSamya(groomDosha, brideDosha)

	// Send the response
	c.JSON(http.StatusOK, gin.H{
		"kuta": kutaScores,
		"rajjuAndOther": rajjuAndOtherComments,
		"groomDosha": groomDosha,
		"brideDosha": brideDosha,
		"whoseIsGreater": whoseIsGreater,
	})
}

func TaraGandler(c *gin.Context) {
	var req dto.TaraReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tara": couple.TaraCalculatorFunc(req.GroomNakshatra, req.BrideNakshatra),
	})
}