package models

type PairingReqBody struct {
	// Rashi (Moon sign) - already present
	GroomRaashi string `json:"groomRaashi" binding:"raashi,required"`
	BrideRaashi string `json:"brideRaashi" binding:"raashi,required"`

	// Nakshatra - already present
	GroomNakshatra string `json:"groomNakshatra" binding:"nakshatra,required"`
	BrideNakshatra string `json:"brideNakshatra" binding:"nakshatra,required"`

	// Ascendant (Lagna) for both
	GroomAscendant string `json:"groomAscendant" binding:"required,raashi"`
	BrideAscendant string `json:"brideAscendant" binding:"required,raashi"`

	// Sun (Surya) placement
	GroomSuryaPlacement string `json:"groomSuryaPlacement" binding:"required,raashi"`
	BrideSuryaPlacement string `json:"brideSuryaPlacement" binding:"required,raashi"`

	// Mercury (Budha)
	GroomBudhaPlacement string `json:"groomBudhaPlacement" binding:"required,raashi"`
	BrideBudhaPlacement string `json:"brideBudhaPlacement" binding:"required,raashi"`

	// Venus (Shukra) – very important for marriage/attraction
	GroomShukraPlacement string `json:"groomShukraPlacement" binding:"required,raashi"`
	BrideShukraPlacement string `json:"brideShukraPlacement" binding:"required,raashi"`

	// Moon (Chandra) – already have Raashi, but if you want separate Moon sign placement (sometimes used)
	GroomChandraPlacement string `json:"groomChandraPlacement" binding:"required,raashi"`
	BrideChandraPlacement string `json:"brideChandraPlacement" binding:"required,raashi"`

	// Rahu & Ketu (nodes) – often checked for doshas in matching
	GroomRahuPlacement string `json:"groomRahuPlacement" binding:"required,raashi"`
	BrideRahuPlacement string `json:"brideRahuPlacement" binding:"required,raashi"`
	GroomKetuPlacement string `json:"groomKetuPlacement" binding:"required,raashi"`
	BrideKetuPlacement string `json:"brideKetuPlacement" binding:"required,raashi"`

	// Mars (Kuja/Mangal) – critical for Mangal Dosha check
	GroomKujaPlacement string `json:"groomKujaPlacement" binding:"required,raashi"`
	BrideKujaPlacement string `json:"brideKujaPlacement" binding:"required,raashi"`

	// Jupiter (Guru) – important for husband/wife prospects & wisdom
	GroomGuruPlacement string `json:"groomGuruPlacement" binding:"required,raashi"`
	BrideGuruPlacement string `json:"brideGuruPlacement" binding:"required,raashi"`

	// Saturn (Shani) – longevity, delays, responsibilities
	GroomShaniPlacement string `json:"groomShaniPlacement" binding:"required,raashi"`
	BrideShaniPlacement string `json:"brideShaniPlacement" binding:"required,raashi"`

	// Check if first half or second half for Makara and Dhanu
	GroomFirstHalf bool `json:"groomFirstHalf"`
	BrideFirstHalf bool `json:"brideFirstHalf"`
}

type DoshaSamyaRes struct {
	AscendantConsidered string  `json:"ascendantConsidered"`
	GrahaInQuestion     string  `json:"grahaInQuestion"`
	DoshaScore          float32 `json:"doshaScore"`
}

type KutaTypeRes struct {
	Index    int8    `json:"index"`
	Score    float32 `json:"score"`
	Comments string  `json:"comments"`
}

type TaraReq struct {
	GroomNakshatra string `json:"groomNakshatra" binding:"nakshatra,required"`
	BrideNakshatra string `json:"brideNakshatra" binding:"nakshatra,required"`
}