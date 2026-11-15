package models

type GrahaPositions map[Raashyadhipati]string

type PlanetaryData struct {
	Rashis       GrahaPositions
	Longitudes   map[Raashyadhipati]float64
	Speeds       map[Raashyadhipati]float64
	Declinations map[Raashyadhipati]float64
}

type NavamshaData struct {
	Ascendant string
	Rashis    GrahaPositions
}

type ChartResult struct {
	Ascendant      string
	NavAscendant   string
	Rashis         GrahaPositions
	NavamshaRashis GrahaPositions
	Combustion     map[Raashyadhipati]bool
	Paksha         string
	ChandraWaxing  bool
	Atmakaraka     AtmakarakaResult
	Longitudes     map[Raashyadhipati]float64
	FirstHalf      bool
	GrahaData      PlanetaryData
	AscendantLong  float64
	MsLong         float64
	SunriseJD      float64
	SunsetJD       float64
	IsDayBirth     bool
}