package swissEphemeris

import (
	"bytes"
	"fmt"
	"math"
	"path/filepath"
	"sync"

	"github.com/mshafiee/swephgo"
	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
	"github.com/scientist-v08/bphs/utils"
)

func angularDistance(a, b float64) float64 {
    d := math.Mod(a-b+540, 360) - 180
    return math.Abs(d)
}

func calculatePlanetaryLongitudes(jd float64) (models.PlanetaryData, error) {
    grahas := []struct {
        Name   models.Raashyadhipati
        Planet int
    }{
        {constants.Surya, swephgo.SeSun},
        {constants.Chandra, swephgo.SeMoon},
        {constants.Kuja, swephgo.SeMars},
        {constants.Budha, swephgo.SeMercury},
        {constants.Guru, swephgo.SeJupiter},
        {constants.Shukra, swephgo.SeVenus},
        {constants.Shani, swephgo.SeSaturn},
        {constants.Rahu, swephgo.SeTrueNode},
    }

    data := models.PlanetaryData{
        Rashis:       make(models.GrahaPositions),
        Longitudes:   make(map[models.Raashyadhipati]float64),
        Speeds:       make(map[models.Raashyadhipati]float64),
        Declinations: make(map[models.Raashyadhipati]float64),
    }

    flags := swephgo.SeflgSwieph |
        swephgo.SeflgSidereal |
        swephgo.SeflgSpeed

    eqFlags := swephgo.SeflgSwieph | swephgo.SeflgEquatorial

    for _, graha := range grahas {

        xx := make([]float64, 6)
        serr := make([]byte, 256)

        ret := swephgo.CalcUt(
            jd,
            graha.Planet,
            flags,
            xx,
            serr,
        )

        // Clean the error string (Swiss Ephemeris often leaves null bytes)
		errMsg := string(bytes.Trim(serr, "\x00"))

        if ret < 0 {
            return models.PlanetaryData{}, fmt.Errorf(
                "failed to calculate %s (planet id %d): %s",
				graha.Name, graha.Planet, errMsg,
            )
        }

        longitude := utils.Normalize360(xx[0])

        data.Longitudes[graha.Name] = longitude
        data.Speeds[graha.Name] = xx[3]

        rashiIndex := int(longitude / 30)

        data.Rashis[graha.Name] = constants.RaashiOrder[rashiIndex]

        xxEq := make([]float64, 6)
		serrEq := make([]byte, 256)

		retEq := swephgo.CalcUt(jd, graha.Planet, eqFlags, xxEq, serrEq)
		errMsgEq := string(bytes.Trim(serrEq, "\x00"))
		if retEq < 0 {
			return models.PlanetaryData{}, fmt.Errorf(
				"failed to calculate %s (equatorial): %s", graha.Name, errMsgEq,
			)
		}

		// xxEq[1] is the declination in degrees
		data.Declinations[graha.Name] = xxEq[1]

        // Ketu is exactly opposite Rahu.
        if graha.Name == constants.Rahu {

            ketuLongitude := utils.Normalize360(longitude + 180)

            data.Longitudes[constants.Ketu] = ketuLongitude

            // Ketu's longitudinal speed is opposite Rahu's.
            data.Speeds[constants.Ketu] = -xx[3]

            ketuIndex := int(ketuLongitude / 30)

            data.Rashis[constants.Ketu] =
                constants.RaashiOrder[ketuIndex]
        }
    }

    return data, nil
}

func calculateAscendant(
    jd float64,
    latitude float64,
    longitude float64,
) (float64, float64, error) {

    cusps := make([]float64, 13)
    ascmc := make([]float64, 10)

    hsys := int('E')

    flags := swephgo.SeflgSidereal

    ret := swephgo.HousesEx(
        jd,
        flags,
        latitude,
        longitude,
        hsys,
        cusps,
        ascmc,
    )

    if ret < 0 {
        return 0, 0, fmt.Errorf(
            "failed to calculate ascendant",
        )
    }

    return utils.Normalize360(ascmc[0]), utils.Normalize360(ascmc[1]),nil
}

func calculateNavamshaPlacements(
    ascLongitude float64,
    longitudes map[models.Raashyadhipati]float64,
) (models.NavamshaData, error) {

    data := models.NavamshaData{
        Rashis: make(models.GrahaPositions),
    }

    // Navamsha Ascendant.
    navAscendant, err := utils.CalculateNavamsha(ascLongitude)

    if err != nil {
        return models.NavamshaData{}, fmt.Errorf(
            "failed to calculate Navamsha Ascendant: %w",
            err,
        )
    }

    data.Ascendant = navAscendant

    // Navamsha placements for all nine Grahas.
    for graha, longitude := range longitudes {

        navRashi, err := utils.CalculateNavamsha(longitude)

        if err != nil {
            return models.NavamshaData{}, fmt.Errorf(
                "failed to calculate Navamsha for %v: %w",
                graha,
                err,
            )
        }

        data.Rashis[graha] = navRashi
    }

    return data, nil
}

func calculateCombustion(
    longitudes map[models.Raashyadhipati]float64,
    speeds map[models.Raashyadhipati]float64,
) map[models.Raashyadhipati]bool {

    thresholds := map[models.Raashyadhipati]float64{
        constants.Chandra: 12,
        constants.Kuja:    17,
        constants.Guru:    11,
        constants.Shani:   15,
    }

    combustion := make(
        map[models.Raashyadhipati]bool,
    )

    sunLongitude := longitudes[constants.Surya]

    for graha, threshold := range thresholds {

        combustion[graha] = angularDistance(
            sunLongitude,
            longitudes[graha],
        ) <= threshold
    }

    // Mercury.
    mercuryLimit := 14.0

    if speeds[constants.Budha] < 0 {
        mercuryLimit = 12.0
    }

    combustion[constants.Budha] = angularDistance(
        sunLongitude,
        longitudes[constants.Budha],
    ) <= mercuryLimit

    // Venus.
    venusLimit := 10.0

    if speeds[constants.Shukra] < 0 {
        venusLimit = 8.0
    }

    combustion[constants.Shukra] = angularDistance(
        sunLongitude,
        longitudes[constants.Shukra],
    ) <= venusLimit

    return combustion
}

func calculatePaksha(
    longitudes map[models.Raashyadhipati]float64,
) (string, bool) {

    elongation := utils.Normalize360(
        longitudes[constants.Chandra] -
            longitudes[constants.Surya],
    )

    // Treat the exact conjunction boundary as Krishna.
    // Full Moon (180 degrees) remains Shukla.
    waxing := elongation > 0 && elongation <= 180

    if waxing {
        return "Shukla Paksha", true
    }

    return "Krishna Paksha", false
}

// calculateSunTimes computes the sunrise, sunset, and whether the birth occurred during the day.
func calculateSunTimes(jd, longitude, latitude float64) (sunriseJD, sunsetJD float64, isDayBirth bool, err error) {
	geopos := []float64{
		longitude, // East positive
		latitude,  // North positive
		0,         // altitude (sea level is fine)
	}

	// Sunrise
    rsSunrise := make([]float64, 1)
	serrRise := make([]byte, 256)
	retRise := swephgo.RiseTrans(
		jd, // approximate start
		swephgo.SeSun,
		nil,
		swephgo.SeflgSwieph,
		swephgo.SeCalcRise,
		geopos,
		0, 0, // atmospheric pressure & temperature (0 = default)
		rsSunrise,
		serrRise,
	)
	if retRise < 0 {
		return 0, 0, false, fmt.Errorf("sunrise calculation failed: %s", string(bytes.Trim(serrRise, "\x00")))
	}

	// Sunset
    rsSunset := make([]float64, 1)
	serrSet := make([]byte, 256)
	retSet := swephgo.RiseTrans(
		jd,
		swephgo.SeSun,
		nil,
		swephgo.SeflgSwieph,
		swephgo.SeCalcSet,
		geopos,
		0, 0,
		rsSunset,
		serrSet,
	)
	if retSet < 0 {
		return 0, 0, false, fmt.Errorf("sunset calculation failed: %s", string(bytes.Trim(serrSet, "\x00")))
	}

	isDayBirth = jd >= sunriseJD && jd < sunsetJD

	return sunriseJD, sunsetJD, isDayBirth, nil
}

var swephMutex sync.Mutex

func CalculateGrahaRashis(birthDetailsReq models.BirthChartReq) (models.ChartResult, error) {
	swephMutex.Lock()
    defer swephMutex.Unlock()

    ephePath, err := filepath.Abs("./ephe")
    if err != nil {
        return models.ChartResult{}, fmt.Errorf(
            "failed to resolve ephemeris directory: %w", err,
        )
    }

    // Configure Swiss Ephemeris.
    swephgo.SetEphePath([]byte(ephePath))

    // Set True Pushya Paksha Ayanamsha.
    swephgo.SetSidMode(
        swephgo.SeSidmTruePushya,
        0,
        0,
    )

    // Convert birth time to UTC.
    t := birthDetailsReq.BirthDateTime.UTC()

    hour := float64(t.Hour()) +
        float64(t.Minute())/60.0 +
        float64(t.Second())/3600.0 +
        float64(t.Nanosecond())/3.6e12

    // Calculate Julian Day UT.
    jd := swephgo.Julday(
        t.Year(),
        int(t.Month()),
        t.Day(),
        hour,
        swephgo.SeGregCal,
    )

    // -------------------------------------------------
    // Sunrise & Sunset (required for Nathonnata / Kaala Bala)
    // -------------------------------------------------
    sunriseJD, sunsetJD, isDayBirth, err := calculateSunTimes(
        jd,
        birthDetailsReq.Longitude,
        birthDetailsReq.Latitude,
    )
    if err != nil {
        return models.ChartResult{}, err
    }

    // -------------------------------------------------
    // 1. Planetary longitudes and D1 Rashis.
    // -------------------------------------------------

    planetaryData, err := calculatePlanetaryLongitudes(jd)
    if err != nil {
        return models.ChartResult{}, err
    }
    moonLong, ok := planetaryData.Longitudes[constants.Chandra]
    if !ok {
        return models.ChartResult{}, fmt.Errorf("moon longitude not found")
    }

    moonRashi := constants.RaashiOrder[int(moonLong/30.0)]
    degreeInSign := math.Mod(moonLong, 30.0)

    firstHalf := false
    if moonRashi == "Dhanassu" || moonRashi == "Makara" {
        firstHalf = degreeInSign < 15.0
    }

    // -------------------------------------------------
    // 2. Ascendant.
    // -------------------------------------------------

    ascLongitude, msLongitude, err := calculateAscendant(
        jd,
        birthDetailsReq.Latitude,
        birthDetailsReq.Longitude,
    )

    if err != nil {
        return models.ChartResult{}, err
    }

    ascendant := constants.RaashiOrder[int(ascLongitude/30)]

    // -------------------------------------------------
    // 3. Navamsha placements.
    // -------------------------------------------------

    navamshaData, err := calculateNavamshaPlacements(
        ascLongitude,
        planetaryData.Longitudes,
    )

    if err != nil {
        return models.ChartResult{}, err
    }

    // -------------------------------------------------
    // 4. Combustion.
    // -------------------------------------------------

    combustion := calculateCombustion(
        planetaryData.Longitudes,
        planetaryData.Speeds,
    )

    // -------------------------------------------------
    // 5. Paksha.
    // -------------------------------------------------

    paksha, waxing := calculatePaksha(
        planetaryData.Longitudes,
    )

    // -------------------------------------------------
    // 6. Atmakaraka and Karakamsha.
    // -------------------------------------------------

    atmakaraka, err := utils.CalculateAtmakaraka(
        planetaryData.Longitudes,
    )

    if err != nil {
        return models.ChartResult{}, err
    }

    // -------------------------------------------------
    // Final result.
    // -------------------------------------------------

    return models.ChartResult{
        Ascendant:      ascendant,
        NavAscendant:   navamshaData.Ascendant,
        Rashis:         planetaryData.Rashis,
        NavamshaRashis: navamshaData.Rashis,
        Combustion:     combustion,
        Paksha:         paksha,
        ChandraWaxing:  waxing,
        Atmakaraka:     atmakaraka,
        Longitudes:     planetaryData.Longitudes,
        FirstHalf:      firstHalf,
        GrahaData:      planetaryData,
        AscendantLong:  ascLongitude,
        SunriseJD:  sunriseJD,
        SunsetJD:   sunsetJD,
        IsDayBirth: isDayBirth,
        MsLong: msLongitude,
    }, nil
}