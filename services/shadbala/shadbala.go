package shadbala

import (
	"fmt"
	"time"

	"github.com/scientist-v08/bphs/dto"
)

// ---------------------------------------------------------------------------
// Public types
// ---------------------------------------------------------------------------

// Input is everything the Shad Bala calculation needs.
type Input struct {
	// Planets is the value already returned by calculatePlanetaryLongitudes
	// (sidereal longitudes, 0..360). Must also carry Speeds, Declinations
	// and EclipticLatitudes when available.
	Planets dto.PlanetaryData

	// Chart carries the already-computed D1/D9 placements and the ascendant.
	Chart dto.HouseEffectsReqBody

	// For Dig / Drik / Kaala
	AscendantLong     float64
	MsLong            float64 // Midheaven longitude
	SunriseJD         float64
	SunsetJD          float64
	IsDayBirth        bool
	BirthDateTime     time.Time
	Latitude          float64
	Longitude         float64
	MeshaJD           float64 // JD of Mesha Sankranti of the year of birth
	SolarMonthStartJD float64 // JD when Sun entered the current rāśi

	// Optional: pre-computed Chesta elements (mean / true / seeghrochcha).
	// If nil, Chesta Bala for the five tara grahas will be zero.
	ChestaElements map[dto.Raashyadhipati]ChestaElements
}

// Result is what the controller receives. New balas get added as fields.
type Result struct {
	Sthana SthanaBalaResult  `json:"sthana"`
	Dig    DigBalaResult     `json:"dig"`
	Drik   DrikBalaResult    `json:"drik"`
	Kaala  KaalaBalaResult   `json:"kaala"`
	Chesta ChestaBalaResult  `json:"chesta"`
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

	kaala := calculateKaalaBala(ctx, in.Planets)
	chesta := calculateChestaBala(ctx, kaala)

	return &Result{
		Sthana: calculateSthanaBala(ctx),
		Dig:    calculateDigBala(ctx),
		Drik:   calculateDrikBala(ctx),
		Kaala:  kaala,
		Chesta: chesta,
	}, nil
}