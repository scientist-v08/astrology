package constants

import "github.com/scientist-v08/bphs/dto"

const (
	Kuja    dto.Raashyadhipati = "Kuja"
	Shukra  dto.Raashyadhipati = "Shukra"
	Budha   dto.Raashyadhipati = "Budha"
	Chandra dto.Raashyadhipati = "Chandra"
	Surya   dto.Raashyadhipati = "Surya"
	Guru    dto.Raashyadhipati = "Guru"
	Shani   dto.Raashyadhipati = "Shani"
	Rahu 	dto.Raashyadhipati = "Rahu"
	Ketu 	dto.Raashyadhipati = "Ketu"
)

var RaashyadhipatiMapStore = dto.RaashyadhipatiMapping{
	"Mesha":     Kuja,
	"Vrushabha": Shukra,
	"Mithuna":   Budha,
	"Karkataka": Chandra,
	"Simha":     Surya,
	"Kanya":     Budha,
	"Tula":      Shukra,
	"Vruschika": Kuja,
	"Dhanassu":  Guru,
	"Makara":    Shani,
	"Kumbha":    Shani,
	"Meena":     Guru,
}

var MaitriKuta = map[dto.Raashyadhipati]map[dto.Raashyadhipati]int8{
	Kuja: {
		Kuja:    5,
		Shukra:  3,
		Budha:   0,
		Surya:   5,
		Guru:    5,
		Shani:   0,
		Chandra: 3,
	},
	Shukra: {
		Kuja:    3,
		Shukra:  5,
		Budha:   5,
		Surya:   0,
		Guru:    0,
		Shani:   5,
		Chandra: 0,
	},
	Budha: {
		Kuja:    0,
		Shukra:  5,
		Budha:   5,
		Surya:   3,
		Guru:    0,
		Shani:   3,
		Chandra: 0,
	},
	Surya: {
		Kuja:    5,
		Shukra:  0,
		Budha:   3,
		Surya:   5,
		Guru:    5,
		Shani:   0,
		Chandra: 5,
	},
	Guru: {
		Kuja:    5,
		Shukra:  0,
		Budha:   0,
		Surya:   5,
		Guru:    5,
		Shani:   3,
		Chandra: 3,
	},
	Shani: {
		Kuja:    0,
		Shukra:  5,
		Budha:   3,
		Surya:   0,
		Guru:    3,
		Shani:   5,
		Chandra: 0,
	},
	Chandra: {
		Kuja:    3,
		Shukra:  0,
		Budha:   0,
		Surya:   5,
		Guru:    3,
		Shani:   0,
		Chandra: 5,
	},
}