package models

type GeoapifyResponse struct {
	Type     string            `json:"type"`
	Features []GeoapifyFeature `json:"features"`
}

type GeoapifyFeature struct {
	Type       string                  `json:"type"`
	Properties GeoapifyPlaceProperties `json:"properties"`
}

type GeoapifyPlaceProperties struct {
	Name        string            `json:"name"`
	Formatted   string            `json:"formatted"`
	Country     string            `json:"country"`
	CountryCode string            `json:"country_code"`
	State       string            `json:"state"`
	City        string            `json:"city"`
	Lat         float64           `json:"lat"`
	Lon         float64           `json:"lon"`
	PlaceID     string            `json:"place_id"`
	Timezone    *GeoapifyTimezone `json:"timezone"`
}

type GeoapifyTimezone struct {
	Name             string `json:"name"`
	NameAlt          string `json:"name_alt"`
	OffsetSTD        string `json:"offset_STD"`
	OffsetDST        string `json:"offset_DST"`
	OffsetSTDSeconds int    `json:"offset_STD_seconds"`
	OffsetDSTSeconds int    `json:"offset_DST_seconds"`
}
