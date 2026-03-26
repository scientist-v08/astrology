package constants

import "github.com/scientist-v08/bphs/models"

var ExaltationMapStore = models.SpecialPosition{
	Shani: Tula,
	Kuja: Makara,    
	Shukra: Meena,  
	Budha: Kanya,   
	Chandra: Vrushabha, 
	Surya: Mesha,   
	Guru: Karkataka,
}

var ReverseExaltationMapStore = models.ReverseSpecialPosition{
	Tula: Shani,
	Makara: Kuja,    
	Meena: Shukra,  
	Kanya: Budha,   
	Vrushabha: Chandra, 
	Mesha: Surya,   
	Karkataka: Guru,
}