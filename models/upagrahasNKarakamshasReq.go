package models

type UpagrahasNKarakamshasReqBody struct {
	Ascendant  string `json:"ascendant" binding:"required,raashi"`
	Dhuma      string `json:"dhuma" binding:"required,raashi"`
	Vyatipata  string `json:"vyatipata" binding:"required,raashi"`
	Parivesha  string `json:"parivesha" binding:"required,raashi"`
	Indrachapa string `json:"indrachapa" binding:"required,raashi"`
	Upaketu    string `json:"upaketu" binding:"required,raashi"`
	Gulika     string `json:"gulika" binding:"required,raashi"`
	Pranapada  string `json:"pranapada" binding:"required,raashi"`
}

type UpagrahasNKarakamshasResBody struct {
	Upagraha string   `json:"upagraha"`
	Effects  []string `json:"effects"`
}

type KarakamshaReqBody struct {
	Ascendant        string `json:"ascendant" binding:"required,raashi"`
	Karakamsha       string `json:"karakamsha" binding:"required,raashi"`
	Surya            string `json:"surya" binding:"required,raashi"`
	Budha            string `json:"budha" binding:"required,raashi"`
	Shukra           string `json:"shukra" binding:"required,raashi"`
	Chandra          string `json:"chandra" binding:"required,raashi"`
	Rahu             string `json:"rahu" binding:"required,raashi"`
	Ketu             string `json:"ketu" binding:"required,raashi"`
	Kuja             string `json:"kuja" binding:"required,raashi"`
	Guru             string `json:"guru" binding:"required,raashi"`
	Shani            string `json:"shani" binding:"required,raashi"`
	MercuryAfflicted bool   `json:"mercuryAfflicted"`
	MoonWaxing       bool   `json:"moonWaxing"`
}