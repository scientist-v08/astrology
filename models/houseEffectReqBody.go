package models

type GrahaFullDetails struct {
	Placement  string `json:"placement" binding:"required,raashi"`
	Combustion bool   `json:"combustion"`
}

type HouseEffectsReqBody struct {
	Ascendant           string           `json:"ascendant" binding:"required,raashi"`
	SuryaPlacement      string           `json:"suryaPlacement" binding:"required,raashi"`
	BudhaPlacement      GrahaFullDetails `json:"budhaPlacement"`
	ShukraPlacement     GrahaFullDetails `json:"shukraPlacement"`
	ChandraPlacement    string           `json:"chandraPlacement" binding:"required,raashi"`
	ChandraWaxing       bool             `json:"chandraWaxing"`
	RahuPlacement       string           `json:"rahuPlacement" binding:"required,raashi"`
	KetuPlacement       string           `json:"ketuPlacement" binding:"required,raashi"`
	KujaPlacement       GrahaFullDetails `json:"kujaPlacement"`
	GuruPlacement       GrahaFullDetails `json:"guruPlacement"`
	ShaniPlacement      GrahaFullDetails `json:"shaniPlacement"`
	BudhaUnafflicted    bool             `json:"budhaUnafflicted"`
	FirstHalf           bool             `json:"firstHalf"`
	NavAscendant        string           `json:"nascendant" binding:"required,raashi"`
	NavSuryaPlacement   string           `json:"nsuryaPlacement" binding:"required,raashi"`
	NavBudhaPlacement   string           `json:"nbudhaPlacement" binding:"required,raashi"`
	NavShukraPlacement  string           `json:"nshukraPlacement" binding:"required,raashi"`
	NavChandraPlacement string           `json:"nchandraPlacement" binding:"required,raashi"`
	NavRahuPlacement    string           `json:"nrahuPlacement" binding:"required,raashi"`
	NavKetuPlacement    string           `json:"nketuPlacement" binding:"required,raashi"`
	NavKujaPlacement    string           `json:"nkujaPlacement" binding:"required,raashi"`
	NavGuruPlacement    string           `json:"nguruPlacement" binding:"required,raashi"`
	NavShaniPlacement   string           `json:"nshaniPlacement" binding:"required,raashi"`
	Karakamsha          string           `json:"karakamsha" binding:"required,raashi"`
	NavBudhaAfflicted   bool             `json:"navBudhaAfflicted"`
}

type HouseEffectsWithCommentory struct {
	HouseNumber int8     `json:"houseNumber"`
	Theme       string   `json:"theme"`
	Comments    string   `json:"comments"`
	Effects     []string `json:"effects"`
}

type BhavaLordEffectsReqBody struct {
	Ascendant                 string `json:"ascendant" binding:"required,raashi"`
	FirstHighestShadBalaLord  string `json:"firstHighestShadBalaLord" binding:"required,graha"`
	SecondHighestShadBalaLord string `json:"secondHighestShadBalaLord" binding:"required,graha"`
	ThirdHighestShadBalaLord  string `json:"thirdtHighestShadBalaLord" binding:"required,graha"`
	FirstPlacement            int8   `json:"firstPlacement" binding:"required"`
	SecondPlacement           int8   `json:"secondPlacement" binding:"required"`
	ThirdPlacement            int8   `json:"thirdPlacement" binding:"required"`
}