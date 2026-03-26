package models

type SthanaBalaReq struct {
	StanabalaOrKaalabalaOfGrahas []float32 `json:"sthanaOrKaala" binding:"required"`
	DrigbalaOfGrahas             []float32 `json:"drig" binding:"required"`
	IsSthana                     bool      `json:"isSthana"`
}

type DigBalaReq struct {
	DigbalaOrChestabalaOfGrahas []float32 `json:"digOrChesta" binding:"required"`
	IsDig                       bool      `json:"isDig"`
}

type AyanaBalaReq struct {
	AyanaBala []float32 `json:"ayana" binding:"required"`
}

type SthanaBalaResField struct {
	Graha         string  `json:"graha"`
	EffectiveBala float32 `json:"effectiveBala"`
	Rank          int8    `json:"rank"`
}

type ShadBalaResField struct {
	Graha         string  `json:"graha"`
	EffectiveBala float32 `json:"effectiveBala"`
	Rank          int8    `json:"rank"`
}

type DigOrChestaBalaResField struct {
	Graha                    string  `json:"graha"`
	EffectiveDigOrChestaBala float32 `json:"effectiveBala"`
	Rank                     int8    `json:"rank"`
}

type AyanaBalaResField struct {
	Graha     string  `json:"graha"`
	AyanaBala float32 `json:"ayanaBala"`
	Rank      int8    `json:"rank"`
}

type IsSthansOrKaalaGreater struct {
	SthanaBala []float32 `json:"sthanaBala" binding:"required"`
	KaalaBala  []float32 `json:"kaalaBala" binding:"required"`
}

type PureShadbalaReq struct {
	SthanaBala []float32 `json:"sthanaBala" binding:"required"`
	KaalaBala  []float32 `json:"kaalaBala" binding:"required"`
	DigBala    []float32 `json:"digBala" binding:"required"`
	ChestaBala []float32 `json:"chestaBala" binding:"required"`
	DrigBala   []float32 `json:"drigBala" binding:"required"`
}

type BhavaLordAndShadBalaReqBody struct {
	Ascendant        string    `json:"ascendant" binding:"required,raashi"`
	SuryaPlacement   string    `json:"suryaPlacement" binding:"required,raashi"`
	BudhaPlacement   string    `json:"budhaPlacement" binding:"required,raashi"`
	ShukraPlacement  string    `json:"shukraPlacement" binding:"required,raashi"`
	ChandraPlacement string    `json:"chandraPlacement" binding:"required,raashi"`
	RahuPlacement    string    `json:"rahuPlacement" binding:"required,raashi"`
	KetuPlacement    string    `json:"ketuPlacement" binding:"required,raashi"`
	KujaPlacement    string    `json:"kujaPlacement" binding:"required,raashi"`
	GuruPlacement    string    `json:"guruPlacement" binding:"required,raashi"`
	ShaniPlacement   string    `json:"shaniPlacement" binding:"required,raashi"`
	SthanaBala       []float32 `json:"sthanaBala" binding:"required,len=7"`
	KaalaBala        []float32 `json:"kaalaBala" binding:"required,len=7"`
	DrigBala         []float32 `json:"drigBala" binding:"required,len=7"`
	ShadBala         []float32 `json:"shadBala" binding:"required,len=7"`
}