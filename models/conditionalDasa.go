package models

type GivenDeg struct {
	Deg int32 `json:"deg" binding:"required"`
	Min int32 `json:"min" binding:"required"`
}

type ConditionalDasaReq struct {
	Ascendant        string   `json:"ascendant" binding:"required,raashi"`
	SuryaPlacement   string   `json:"suryaPlacement" binding:"required,raashi"`
	BudhaPlacement   string   `json:"budhaPlacement" binding:"required,raashi"`
	ShukraPlacement  string   `json:"shukraPlacement" binding:"required,raashi"`
	ChandraPlacement string   `json:"chandraPlacement" binding:"required,raashi"`
	RahuPlacement    string   `json:"rahuPlacement" binding:"required,raashi"`
	KetuPlacement    string   `json:"ketuPlacement" binding:"required,raashi"`
	KujaPlacement    string   `json:"kujaPlacement" binding:"required,raashi"`
	GuruPlacement    string   `json:"guruPlacement" binding:"required,raashi"`
	ShaniPlacement   string   `json:"shaniPlacement" binding:"required,raashi"`
	AscendantDeg     GivenDeg `json:"ascendantDeg"`
	SuryaDeg         GivenDeg `json:"suryaDeg"`
	ChandraDeg       GivenDeg `json:"chandraDeg"`
}