package constants

import "github.com/scientist-v08/bphs/models"

type EnemySignMappings map[models.Raashyadhipati][]models.AllRaashis

var EnemySignsMappingStore = EnemySignMappings{
	Kuja: []models.AllRaashis{Mithuna, Kanya},
	Surya: []models.AllRaashis{Tula, Vrushabha, Makara, Kumbha},
	Chandra: []models.AllRaashis{},
	Budha: []models.AllRaashis{Karkataka},
	Guru: []models.AllRaashis{Vrushabha, Mithuna, Tula, Kanya},
	Shukra: []models.AllRaashis{Karkataka, Simha},
	Shani: []models.AllRaashis{Mesha, Karkataka, Simha, Vruschika},
}