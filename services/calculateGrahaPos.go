package services

import (
	"fmt"
	"math"
	"time"

	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
	swisseph "github.com/tejzpr/go-swisseph"
)

var rashiNames = []string{
	"Mesha",
	"Vrushabha",
	"Mithuna",
	"Karkataka",
	"Simha",
	"Kanya",
	"Tula",
	"Vruschika",
	"Dhanassu",
	"Makara",
	"Kumbha",
	"Meena",
}

type GrahaPositions map[models.Raashyadhipati]string

type ChartResult struct {
    Rashis     map[models.Raashyadhipati]string
    Combustion map[models.Raashyadhipati]bool
    Paksha     string
}

func angularDistance(a, b float64) float64 {
    d := math.Mod(a-b+540, 360) - 180
    return math.Abs(d)
}

func CalculateGrahaRashis(birthTime time.Time) (ChartResult, error) {

	// Use True Pushya Paksha Ayanamsha.
	swisseph.SetSidMode(swisseph.SidmTruePushya, 0, 0)

	// Convert birth time to UTC.
	t := birthTime.UTC()

	hour := float64(t.Hour()) +
		float64(t.Minute())/60 +
		float64(t.Second())/3600 +
		float64(t.Nanosecond())/3.6e12

	// Calculate Julian Day in Universal Time.
	jd := swisseph.Julday(
		int32(t.Year()),
		int32(t.Month()),
		int32(t.Day()),
		hour,
		swisseph.GregCal,
	)

	// Nine Grahas (Ketu derived from Rahu).
	grahas := map[models.Raashyadhipati]int32{
		constants.Surya:   swisseph.Sun,
		constants.Chandra: swisseph.Moon,
		constants.Kuja:    swisseph.Mars,
		constants.Budha:   swisseph.Mercury,
		constants.Guru:    swisseph.Jupiter,
		constants.Shukra:  swisseph.Venus,
		constants.Shani:   swisseph.Saturn,
		constants.Rahu:    swisseph.MeanNode,
	}

	positions := make(GrahaPositions)

	longitudes := make(map[models.Raashyadhipati]float64)
	speeds := make(map[models.Raashyadhipati]float64)

	flags := swisseph.FlagSwieph | swisseph.FlagSidereal

	for name, planet := range grahas {

		result := swisseph.CalcUT(jd, planet, flags)

		if result.Flag < 0 {
			return nil, fmt.Errorf(
				"failed to calculate %s: %s",
				name, result.Error,
			)
		}

		longitude := swisseph.Degnorm(result.Data[0])

		longitudes[name] = longitude
		speeds[name] = result.Data[3]

		rashiIndex := int(longitude / 30)

		positions[name] = rashiNames[rashiIndex]

		// Ketu is 180 degrees opposite Rahu.
		if name == "Rahu" {
			ketuLongitude := swisseph.Degnorm(longitude + 180)
			ketuIndex := int(ketuLongitude / 30)

			positions["Ketu"] = rashiNames[ketuIndex]
		}
	}
	// Determine combustion.
	thresholds := map[models.Raashyadhipati]float64{
		"Chandra": 12,
		"Kuja": 17,
		"Guru":    11,
		"Shani":   15,
	}

	combustion := make(map[models.Raashyadhipati]bool)

	for graha, threshold := range thresholds {
    		combustion[graha] = angularDistance(
        		longitudes["Surya"],
        		longitudes[graha],
    		) <= threshold
	}

	// Mercury: different limits for direct and retrograde motion.
	mercuryLimit := 14.0
	if speeds["Budha"] < 0 {
    		mercuryLimit = 12.0
	}

	combustion["Budha"] = angularDistance(
    		longitudes["Surya"],
    		longitudes["Budha"],
	) <= mercuryLimit

	// Venus: different limits for direct and retrograde motion.
	venusLimit := 10.0
	if speeds["Shukra"] < 0 {
    		venusLimit = 8.0
	}

	combustion["Shukra"] = angularDistance(
    		longitudes["Surya"],
    		longitudes["Shukra"],
	) <= venusLimit

	// Determine Paksha.
	elongation := swisseph.Degnorm(
    		longitudes["Chandra"] - longitudes["Surya"],
	)

	paksha := "Shukla Paksha"
	if elongation >= 180 {
    		paksha = "Krishna Paksha"
	}

	// Return everything.
	return ChartResult{
    		Rashis:     positions,
    		Combustion: combustion,
    		Paksha:     paksha,
	}, nil
}