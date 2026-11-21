package constants

import (
	"github.com/scientist-v08/bphs/dto"
)

type EnemySignMappings map[dto.Raashyadhipati][]dto.AllRaashis

var EnemySignsMappingStore = EnemySignMappings{
	Kuja: []dto.AllRaashis{Mithuna, Kanya},
	Surya: []dto.AllRaashis{Tula, Vrushabha, Makara, Kumbha},
	Chandra: []dto.AllRaashis{},
	Budha: []dto.AllRaashis{Karkataka},
	Guru: []dto.AllRaashis{Vrushabha, Mithuna, Tula, Kanya},
	Shukra: []dto.AllRaashis{Karkataka, Simha},
	Shani: []dto.AllRaashis{Mesha, Karkataka, Simha, Vruschika},
}