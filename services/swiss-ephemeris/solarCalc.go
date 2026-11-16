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

func sunSiderealLon(jd float64) (float64, error) {
	xx := make([]float64, 6)
	serr := make([]byte, 256)
	flags := swephgo.SeflgSwieph | swephgo.SeflgSidereal | swephgo.SeflgSpeed
	if swephgo.CalcUt(jd, swephgo.SeSun, flags, xx, serr) < 0 {
		return 0, fmt.Errorf("CalcUt failed: %s", string(bytes.Trim(serr, "\x00")))
	}
	return utils.Normalize360(xx[0]), nil
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

	if (diffLow > 0 && diffHigh > 0) || (diffLow < 0 && diffHigh < 0) {
		return 0, fmt.Errorf(
			"target longitude %.4f° not crossed within ±%.1f days of JD %.4f (lonLow=%.4f°, lonHigh=%.4f°)",
			targetLon, searchWindow, approxJD, lonLow, lonHigh,
		)
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
			low = mid
			diffLow = diffMid
		} else {
			high = mid
			diffHigh = diffMid
		}
	}

	// Best effort after max iterations
	return (low + high) / 2, nil
}

// lastSunIngress returns the JD of the most recent moment on or before
// birthJD at which the Sun's sidereal longitude equalled targetLon.
func lastSunIngress(birthJD, targetLon float64) (float64, error) {
	lon, err := sunSiderealLon(birthJD)
	if err != nil {
		return 0, err
	}
	back := utils.Normalize360(lon - targetLon)       // degrees travelled since ingress
	approx := birthJD - back*(365.25/360.0)            // good to a few days
	jd, err := findSunIngress(approx, targetLon)       // keep searchWindow, or cut to 20
	if err != nil {
		return 0, err
	}
	if jd > birthJD { // only possible at the exact edge
		return findSunIngress(birthJD-1-back*(365.25/360.0), targetLon)
	}
	return jd, nil
}

func getMeshaSankrantiJD(birthJD float64) (float64, error) {
	return lastSunIngress(birthJD, 0)
}

// getSolarMonthStartJD returns the JD when the Sun entered the rāśi
// that it occupies at birthJD.
func getSolarMonthStartJD(birthJD float64, sunLongitudeAtBirth float64) (float64, error) {
	lon, err := sunSiderealLon(birthJD)
	if err != nil {
		return 0, err
	}
	return lastSunIngress(birthJD, math.Floor(lon/30)*30)
}