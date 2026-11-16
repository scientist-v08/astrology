package swissEphemeris

import (
	"bytes"
	"fmt"
	"math"
	"time"

	"github.com/mshafiee/swephgo"
	"github.com/scientist-v08/bphs/utils"
)

func julianDayUT(t time.Time) float64 {
    utc := t.UTC()

    return 2440587.5 +
        float64(utc.Unix())/86400.0 +
        float64(utc.Nanosecond())/86400e9
}

// findSunIngress finds the Julian Day when the Sun's sidereal longitude
// reaches (or most closely approaches) targetLon (0–360).
// It searches in a window around approxJD.
func findSunIngress(approxJD, targetLon float64) (float64, error) {
	const (
		searchWindow = 45.0   // days (±)
		precision    = 1e-7   // ~0.01 seconds
		maxIter      = 60
	)

	flags := swephgo.SeflgSwieph | swephgo.SeflgSidereal | swephgo.SeflgSpeed

	getLon := func(jd float64) (float64, error) {
		xx := make([]float64, 6)
		serr := make([]byte, 256)
		ret := swephgo.CalcUt(jd, swephgo.SeSun, flags, xx, serr)
		if ret < 0 {
			return 0, fmt.Errorf("CalcUt failed: %s", string(bytes.Trim(serr, "\x00")))
		}
		return utils.Normalize360(xx[0]), nil
	}

	// Circular difference: result is in range [-180, +180]
	circDiff := func(a, b float64) float64 {
		d := utils.Normalize360(a - b)
		if d > 180 {
			d -= 360
		}
		return d
	}

	low := approxJD - searchWindow
	high := approxJD + searchWindow

	lonLow, err := getLon(low)
	if err != nil {
		return 0, err
	}
	lonHigh, err := getLon(high)
	if err != nil {
		return 0, err
	}

	// Check that the target is actually crossed inside the window
	diffLow := circDiff(lonLow, targetLon)
	diffHigh := circDiff(lonHigh, targetLon)

	// If both sides have the same sign and the absolute difference is large,
	// the target is probably not inside this window.
	if (diffLow > 0 && diffHigh > 0) || (diffLow < 0 && diffHigh < 0) {
		// Optional: you can expand the window or return an error
		// For solar motion this almost never happens with ±45 days.
	}

	// Bisection
	for i := 0; i < maxIter; i++ {
		mid := (low + high) / 2
		lonMid, err := getLon(mid)
		if err != nil {
			return 0, err
		}

		diffMid := circDiff(lonMid, targetLon)

		if math.Abs(diffMid) < precision {
			return mid, nil
		}

		// Decide which side to discard using the sign of the difference
		if (diffLow < 0 && diffMid < 0) || (diffLow > 0 && diffMid > 0) {
			// same sign → move the low boundary
			low = mid
			lonLow = lonMid
			diffLow = diffMid
		} else {
			high = mid
			lonHigh = lonMid
			diffHigh = diffMid
		}
	}

	// Best effort after max iterations
	return (low + high) / 2, nil
}

// getMeshaSankrantiJD returns the JD of the most recent Mesha Saṅkrānti
// (Sun entering 0° sidereal) on or before the given birth JD.
func getMeshaSankrantiJD(birthJD float64) (float64, error) {
	// A rough estimate: Sun moves ~1°/day, so go back up to ~370 days
	approx := birthJD - 10 // start a little before birth

	jd, err := findSunIngress(approx, 0.0)
	if err != nil {
		return 0, err
	}

	// If the found ingress is after birth, go one year earlier
	if jd > birthJD {
		jd, err = findSunIngress(jd-365.25, 0.0)
		if err != nil {
			return 0, err
		}
	}

	return jd, nil
}

// getSolarMonthStartJD returns the JD when the Sun entered the rāśi
// that it occupies at birthJD.
func getSolarMonthStartJD(birthJD float64, sunLongitudeAtBirth float64) (float64, error) {
	// Current rāśi start longitude
	rashiStart := math.Floor(sunLongitudeAtBirth/30.0) * 30.0

	// Search near birth
	jd, err := findSunIngress(birthJD, rashiStart)
	if err != nil {
		return 0, err
	}

	// Safety: if we somehow got a future ingress, go one month earlier
	if jd > birthJD {
		jd, err = findSunIngress(jd-35, rashiStart)
		if err != nil {
			return 0, err
		}
	}

	return jd, nil
}