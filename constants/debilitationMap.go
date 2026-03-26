package constants

import "github.com/scientist-v08/bphs/models"

var DebilitationMapStore = models.SpecialPosition{
	Shani: Mesha,
	Kuja: Karkataka,    
	Shukra: Kanya,  
	Budha: Meena,   
	Chandra: Vruschika, 
	Surya: Tula,   
	Guru: Makara,    
}

var ReverseDebilitationMapStore = models.ReverseSpecialPosition{
	Mesha: Shani,
	Karkataka: Kuja,    
	Kanya: Shukra,  
	Meena: Budha,   
	Vruschika: Chandra, 
	Tula: Surya,   
	Makara: Guru,
}