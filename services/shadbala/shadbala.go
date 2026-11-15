// Package shadbala calculates the six strengths (Shad Bala) of the seven
// classical grahas (Sun..Saturn) as described in Brihat Parashara Hora Shastra
// and summarised in K. Jaya Sekhar's "System Jaya - Best way to use Shad Bala".
//
// Layout of this package:
//
//	shadbala.go             - orchestrator, the only entry point controllers use
//	chart_context.go        - validated chart data shared by every bala (D1/D9 signs, longitudes, lagna)
//	sthana_bala.go          - Sthana Bala (5 sub-balas)
//	sthana_bala_vargas.go   - divisional charts, dignities, friendship tables
//
// Kaala, Dig, Chesta, Naisargika and Drik bala will be added as sibling files
// and wired into Service.Calculate below.
package shadbala

import (
	"fmt"

	"github.com/scientist-v08/bphs/models"
)

// Input is everything the Shad Bala calculation needs.
type Input struct {
	// Planets is the value already returned by calculatePlanetaryLongitudes
	// (sidereal longitudes, 0..360).
	Planets models.PlanetaryData

	// Chart carries the already-computed D1/D9 placements and the ascendant.
	// TODO: adjust the package if HouseEffectsReqBody does not live in models.
	Chart models.HouseEffectsReqBody

	// For Drik & Dig Bala
	AscendantLong  float64
	MsLong  float64
	SunriseJD      float64
	SunsetJD       float64
	IsDayBirth     bool
}

// Result is what the controller receives. New balas get added as fields.
type Result struct {
	Sthana SthanaBalaResult `json:"sthana"`
	Dig    DigBalaResult    `json:"dig"`
	Drik   DrikBalaResult   `json:"drik"`
}

// Service is the Shad Bala orchestrator.
type Service struct{}

func NewService() *Service { return &Service{} }

// Calculate validates the chart once, then runs every implemented bala.
func (s *Service) Calculate(in Input) (*Result, error) {
	ctx, err := newChartContext(in)
	if err != nil {
		return nil, fmt.Errorf("invalid chart input: %w", err)
	}

	return &Result{
		Sthana: calculateSthanaBala(ctx),
		Dig:    calculateDigBala(ctx),
		Drik:   calculateDrikBala(ctx),
	}, nil
}
