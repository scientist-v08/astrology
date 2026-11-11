package services

import (
	"fmt"
	"math"

	"github.com/mshafiee/swephgo"
	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
	"github.com/scientist-v08/bphs/utils"
)

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

func CalculateGrahaRashis(birthDetailsReq models.BirthChartReq) (ChartResult, error) {
	// Close Swiss Ephemeris state when function completes if needed:
	defer swephgo.Close()

	// Use True Pushya Paksha Ayanamsha
	swephgo.SetSidMode(29, 0, 0)

	t := birthDetailsReq.BirthDateTime.UTC()
	hour := float64(t.Hour()) +
		float64(t.Minute())/60.0 +
		float64(t.Second())/3600.0 +
		float64(t.Nanosecond())/3.6e12

	jd := swephgo.Julday(
		t.Year(),
		int(t.Month()),
		t.Day(),
		hour,
		swephgo.SeGregCal,
	)
	cusps := make([]float64, 13)
	ascmc := make([]float64, 10)
	hsys := int('W')

	res := swephgo.HousesEx(
		jd,
		swephgo.SeSidmTruePushya,
		birthDetailsReq.Latitude,
		birthDetailsReq.Longitude,
		hsys,
		cusps,
		ascmc,
	)

	if res < 0 {
		return ChartResult{}, fmt.Errorf("Failed to calculate ascendant")
	}

	ascLongitude := swephgo.Degnorm(ascmc[0])

	ascendant := constants.RaashiOrder[int(ascLongitude/30)]

	navAscendant, err := utils.CalculateNavamsha(ascLongitude)

	if err != nil {
		return ChartResult{}, err
	}

	grahas := map[models.Raashyadhipati]int{
		constants.Surya:   swephgo.SeSun,
		constants.Chandra: swephgo.SeMoon,
		constants.Kuja:    swephgo.SeMars,
		constants.Budha:   swephgo.SeMercury,
		constants.Guru:    swephgo.SeJupiter,
		constants.Shukra:  swephgo.SeVenus,
		constants.Shani:   swephgo.SeSaturn,
		constants.Rahu:    swephgo.SeMeanNode,
	}

	positions := make(GrahaPositions)
	longitudes := make(map[models.Raashyadhipati]float64)
	speeds := make(map[models.Raashyadhipati]float64)

	flags := swephgo.SeflgSwieph | swephgo.SeflgSidereal | swephgo.SeflgSpeed

	for name, planet := range grahas {
		xx := make([]float64, 6)
		serr := make([]byte, 256)

		if iflag := swephgo.CalcUt(jd, planet, flags, xx, serr); iflag < 0 {
			return ChartResult{}, fmt.Errorf("failed to calculate %s: %s", name, string(serr))
		}

		longitude := swephgo.Degnorm(xx[0])
		longitudes[name] = longitude
		speeds[name] = xx[3]

		rashiIndex := int(longitude / 30)
		positions[name] = constants.RaashiOrder[rashiIndex]

		if name == constants.Rahu {
			ketuLongitude := swephgo.Degnorm(longitude + 180)
			ketuIndex := int(ketuLongitude / 30)

			positions[constants.Ketu] = constants.RaashiOrder[ketuIndex]
			longitudes[constants.Ketu] = ketuLongitude
			speeds[constants.Ketu] = xx[3]
		}
	}

	// Combustion
	thresholds := map[models.Raashyadhipati]float64{
		constants.Chandra: 12,
		constants.Kuja:    17,
		constants.Guru:    11,
		constants.Shani:   15,
	}

	combustion := make(map[models.Raashyadhipati]bool)
	for graha, threshold := range thresholds {
		combustion[graha] = angularDistance(
			longitudes[constants.Surya],
			longitudes[graha],
		) <= threshold
	}

	mercuryLimit := 14.0
	if speeds[constants.Budha] < 0 {
		mercuryLimit = 12.0
	}
	combustion[constants.Budha] = angularDistance(
		longitudes[constants.Surya],
		longitudes[constants.Budha],
	) <= mercuryLimit

	venusLimit := 10.0
	if speeds[constants.Shukra] < 0 {
		venusLimit = 8.0
	}
	combustion[constants.Shukra] = angularDistance(
		longitudes[constants.Surya],
		longitudes[constants.Shukra],
	) <= venusLimit

	// Paksha
	elongation := swephgo.Degnorm(
		longitudes[constants.Chandra] - longitudes[constants.Surya],
	)

	paksha := "Shukla Paksha"
	if elongation >= 180 {
		paksha = "Krishna Paksha"
	}

	return ChartResult{
		Rashis:     positions,
		Combustion: combustion,
		Paksha:     paksha,
	}, nil
}