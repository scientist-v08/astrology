package constants

import "github.com/scientist-v08/bphs/models"

var OppositeAspectsStore = map[models.AllRaashis]models.AllRaashis{
    Mesha:     Tula,
    Vrushabha: Vruschika,
    Mithuna:   Dhanassu,
    Karkataka: Makara,
    Simha:     Kumbha,
    Kanya:     Meena,
    Tula:      Mesha,
    Vruschika: Vrushabha,
    Dhanassu:  Mithuna,
    Makara:    Karkataka,
    Kumbha:    Simha,
    Meena:     Kanya,
}