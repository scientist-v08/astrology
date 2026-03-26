package constants

type RaashiQuality string
type RaashiQualityiMapping map[string]RaashiQuality

const (
	Chara  RaashiQuality = "Chara"
	Sthira RaashiQuality = "Sthira"
	Dwi    RaashiQuality = "Dwi"
)

var RaashiQualityStore = RaashiQualityiMapping{
	"Mesha":     Chara,
	"Vrushabha": Sthira,
	"Mithuna":   Dwi,
	"Karkataka": Chara,
	"Simha":     Sthira,
	"Kanya":     Dwi,
	"Tula":      Chara,
	"Vruschika": Sthira,
	"Dhanassu":  Dwi,
	"Makara":    Chara,
	"Kumbha":    Sthira,
	"Meena":     Dwi,
}