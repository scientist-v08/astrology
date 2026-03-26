package constants

import "github.com/scientist-v08/bphs/models"

type HouseEffectProportion map[models.AllRaashis][]int8
type SpecialHouseEffectProportion map[string][]int8

var HouseeffectProprtionStore = HouseEffectProportion{
	Mesha: {30,20,10,0,10,20,30,40,50,60,50,40},
	Vrushabha: {30,20,10,0,10,20,30,40,50,60,50,40},
	Mithuna: {60,50,40,30,20,10,0,10,20,30,40,50},
	Karkataka: {30,40,50,60,50,40,30,20,10,0,10,20},
	Simha: {30,20,10,0,10,20,30,40,50,60,50,40},
	Kanya: {60,50,40,30,20,10,0,10,20,30,40,50},
	Tula: {60,50,40,30,20,10,0,10,20,30,40,50},
	Vruschika: {0,10,20,30,40,50,60,50,40,30,20,10},
	Kumbha: {60,50,40,30,20,10,0,10,20,30,40,50},
	Meena: {30,40,50,60,50,40,30,20,10,0,10,20},
}

var SpecialHouseEffectsProportionStore = SpecialHouseEffectProportion{
	"Dhanassu1": {30,20,10,0,10,20,30,40,50,60,50,40},
	"Makara2": {30,20,10,0,10,20,30,40,50,60,50,40},
	"Dhanassu2": {60,50,40,30,20,10,0,10,20,30,40,50},
	"Makara1": {30,40,50,60,50,40,30,20,10,0,10,20},
}