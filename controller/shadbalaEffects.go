package controller

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/scientist-v08/bphs/models"
)

type RankedItem struct {
    idx   int
    value float32
}

func SthanaBalaOrKaalaBalaCalculator(c *gin.Context) {
	var req models.SthanaBalaReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
		})
		return
	}

	if len(req.StanabalaOrKaalabalaOfGrahas) != 7 || len(req.DrigbalaOfGrahas) != 7 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "7 values are to be sent",
		})
		return
	}

	grahas := []string{"Sun", "Moon", "Mars", "Mercury", "Jupiter", "Venus", "Saturn"}
	allGrahaDefined := make([]models.SthanaBalaResField, 7)

	sum := make([]float32, 7)
	for i := range req.DrigbalaOfGrahas {
		sum[i] = req.DrigbalaOfGrahas[i] + req.StanabalaOrKaalabalaOfGrahas[i]
	}

	SthanaBalaMinimums := []float32{165.0,133.0,96.0,165.0,165.0,133.0,96.0}
	KaalaBalaMinimums := []float32{50,30,40,50,50,30,40}
	minimas := make([]float32, 7)
	if req.IsSthana {
		copy(minimas, SthanaBalaMinimums)
	} else {
		copy(minimas, KaalaBalaMinimums)
	}
	effectiveSthanaBala := make([]float32, 7)
	for i := range SthanaBalaMinimums {
		effectiveSthanaBala[i] = sum[i] / minimas[i]
	}

	for i := range grahas {
		allGrahaDefined[i] = models.SthanaBalaResField{
			Graha:        grahas[i],
			EffectiveBala: effectiveSthanaBala[i],
			Rank:         0,
		}
	}

	items := make([]RankedItem, 7)
	for i := range 7 {
		items[i] = RankedItem{i, effectiveSthanaBala[i]}
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].value > items[j].value
	})

	// Assign ranks (1 = strongest)
	for rank, item := range items {
		allGrahaDefined[item.idx].Rank = int8(rank + 1)
	}

	sort.Slice(allGrahaDefined, func(i, j int) bool {
		return allGrahaDefined[i].Rank < allGrahaDefined[j].Rank
	})

	// Proceed with the response
	c.JSON(http.StatusOK, gin.H{
		"rankings": allGrahaDefined,
	})
}

func AyanaBalaRanks(c *gin.Context) {
	var req models.AyanaBalaReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
		})
		return
	}

	if len(req.AyanaBala) != 7 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "7 values are to be sent",
		})
		return
	}

	grahas := []string{"Sun", "Moon", "Mars", "Mercury", "Jupiter", "Venus", "Saturn"}
	allGrahaDefined := make([]models.AyanaBalaResField, 7)
	effectiveAyanaBala := make([]float32, 7)
	AyanaBalaMinimums := []float32{30, 40, 20, 30, 30, 40, 20}

	for i := range 7 {
		effectiveAyanaBala[i] = req.AyanaBala[i] / AyanaBalaMinimums[i]
	}

	for i := range 7 {
		allGrahaDefined[i] = models.AyanaBalaResField{
			Graha:        grahas[i],
			AyanaBala: effectiveAyanaBala[i],
			Rank:         0,
		}
	}

	items := make([]RankedItem, 7)
	for i := range 7 {
		items[i] = RankedItem{i, effectiveAyanaBala[i]}
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].value > items[j].value
	})

	// Assign ranks (1 = strongest)
	for rank, item := range items {
		allGrahaDefined[item.idx].Rank = int8(rank + 1)
	}

	sort.Slice(allGrahaDefined, func(i, j int) bool {
		return allGrahaDefined[i].Rank < allGrahaDefined[j].Rank
	})

	// Proceed with the response
	c.JSON(http.StatusOK, gin.H{
		"rankings": allGrahaDefined,
	})
}

func DigBalaOrChestaBalaCalculator(c *gin.Context) {
	var req models.DigBalaReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
		})
		return
	}

	if len(req.DigbalaOrChestabalaOfGrahas) != 7 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "7 values are to be sent",
		})
		return
	}

	grahas := []string{"Sun", "Moon", "Mars", "Mercury", "Jupiter", "Venus", "Saturn"}
	allGrahaDefined := make([]models.DigOrChestaBalaResField, 7)

	sum := make([]float32, 7)
	for i := range 7 {
		sum[i] = req.DigbalaOrChestabalaOfGrahas[i]
	}

	DigBalaMinimums := []float32{35, 50, 30, 35, 35, 50, 30}
	ChestaBalaMinimums := []float32{112, 100, 67, 112, 112, 100, 67}
	minimas := make([]float32, 7)
	if req.IsDig {
		copy(minimas, DigBalaMinimums)
	} else {
		copy(minimas, ChestaBalaMinimums)
	}
	effectiveSthanaBala := make([]float32, 7)
	for i := range 7 {
		effectiveSthanaBala[i] = sum[i] / minimas[i]
	}

	for i := range grahas {
		allGrahaDefined[i] = models.DigOrChestaBalaResField{
			Graha:        grahas[i],
			EffectiveDigOrChestaBala: effectiveSthanaBala[i],
			Rank:         0,
		}
	}

	items := make([]RankedItem, 7)
	for i := range 7 {
		items[i] = RankedItem{i, effectiveSthanaBala[i]}
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].value > items[j].value
	})

	// Assign ranks (1 = strongest)
	for rank, item := range items {
		allGrahaDefined[item.idx].Rank = int8(rank + 1)
	}

	sort.Slice(allGrahaDefined, func(i, j int) bool {
		return allGrahaDefined[i].Rank < allGrahaDefined[j].Rank
	})

	// Proceed with the response
	c.JSON(http.StatusOK, gin.H{
		"rankings": allGrahaDefined,
	})
}

func IsSthanaGreaterOrKaala(c *gin.Context) {
	var req models.IsSthansOrKaalaGreater

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
		})
		return
	}

	if len(req.SthanaBala) != 7 || len(req.KaalaBala) != 7 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "both sthana_bala and kaala_bala arrays must be having 7 values",
		})
		return
	}

	var totalSthana float32
	var totalKaala float32

	for i := range req.SthanaBala {
		totalSthana += req.SthanaBala[i]
		totalKaala += req.KaalaBala[i]
	}

	result := gin.H{
		"total_sthana_bala":  totalSthana,
		"total_kaala_bala":   totalKaala,
		"difference":         totalSthana - totalKaala,
	}

	var conclusion string
	if totalSthana > totalKaala+0.01 {
		conclusion = "Sthana Bala is greater overall leading to long term success"
	} else if totalKaala > totalSthana+0.01 {
		conclusion = "Kaala Bala is greater overall leading to small term success"
	} else {
		conclusion = "Sthana Bala and Kaala Bala are approximately equal"
	}

	result["conclusion"] = conclusion
	result["sthana_wins"] = totalSthana > totalKaala

	c.JSON(http.StatusOK, result)
}

func PureShadbalaRank(c *gin.Context) {
	var req models.PureShadbalaReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
		})
		return
	}

	if len(req.SthanaBala) != 7 || len(req.KaalaBala) != 7 || len(req.DigBala) != 7 || len(req.ChestaBala) != 7 || len(req.DrigBala) != 7 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "7 values are to be sent for all the types of Balas",
		})
		return
	}

	grahas := []string{"Sun", "Moon", "Mars", "Mercury", "Jupiter", "Venus", "Saturn"}
	naisargikaBala := []float32{60.0,51.4,17.1,25.7,34.3,42.9,8.6}
	allGrahaDefined := make([]models.ShadBalaResField, 7)

	sum := make([]float32, 7)
	effectiveShadbala := make([]float32, 7)
	minimumShadbala := []float32{300, 360, 300, 420, 390, 330, 300}
	for i := range 7 {
		sum[i] = req.SthanaBala[i] + req.KaalaBala[i] + req.DigBala[i] + req.ChestaBala[i] + req.DrigBala[i] + naisargikaBala[i]
	}
	for i := range 7 {
		effectiveShadbala[i] = sum[i] / minimumShadbala[i]
	}

	for i := range 7 {
		allGrahaDefined[i] = models.ShadBalaResField{
			Graha:        grahas[i],
			EffectiveBala: effectiveShadbala[i],
			Rank:         0,
		}
	}

	items := make([]RankedItem, 7)
	for i := range 7 {
		items[i] = RankedItem{i, effectiveShadbala[i]}
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].value > items[j].value
	})

	// Assign ranks (1 = strongest)
	for rank, item := range items {
		allGrahaDefined[item.idx].Rank = int8(rank + 1)
	}

	sort.Slice(allGrahaDefined, func(i, j int) bool {
		return allGrahaDefined[i].Rank < allGrahaDefined[j].Rank
	})

	// Proceed with the response
	c.JSON(http.StatusOK, gin.H{
		"ShadbalaRankings": allGrahaDefined,
	})
}
