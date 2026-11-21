package constants

import "github.com/scientist-v08/bphs/dto"

var ExaltationMapStore = dto.SpecialPosition{
	Shani: Tula,
	Kuja: Makara,    
	Shukra: Meena,  
	Budha: Kanya,   
	Chandra: Vrushabha, 
	Surya: Mesha,   
	Guru: Karkataka,
}

var ReverseExaltationMapStore = dto.ReverseSpecialPosition{
	Tula: Shani,
	Makara: Kuja,    
	Meena: Shukra,  
	Kanya: Budha,   
	Vrushabha: Chandra, 
	Mesha: Surya,   
	Karkataka: Guru,
}