package utils

import "github.com/scientist-v08/bphs/dto"

func ToHouseEffectsReqBody(chart dto.ChartResult) dto.HouseEffectsReqBody {
	return dto.HouseEffectsReqBody{
		Ascendant: chart.Ascendant,

		SuryaPlacement: string(chart.Rashis["Surya"]),

		BudhaPlacement: dto.GrahaFullDetails{
			Placement:  chart.Rashis["Budha"],
			Combustion: chart.Combustion["Budha"],
		},

		ShukraPlacement: dto.GrahaFullDetails{
			Placement:  chart.Rashis["Shukra"],
			Combustion: chart.Combustion["Shukra"],
		},

		ChandraPlacement: chart.Rashis["Chandra"],
		ChandraWaxing:    chart.ChandraWaxing,

		RahuPlacement: chart.Rashis["Rahu"],
		KetuPlacement: chart.Rashis["Ketu"],

		KujaPlacement: dto.GrahaFullDetails{
			Placement:  chart.Rashis["Kuja"],
			Combustion: chart.Combustion["Kuja"],
		},

		GuruPlacement: dto.GrahaFullDetails{
			Placement:  chart.Rashis["Guru"],
			Combustion: chart.Combustion["Guru"],
		},

		ShaniPlacement: dto.GrahaFullDetails{
			Placement:  chart.Rashis["Shani"],
			Combustion: chart.Combustion["Shani"],
		},

		NavAscendant:        chart.NavAscendant,
		NavSuryaPlacement:   chart.NavamshaRashis["Surya"],
		NavBudhaPlacement:   chart.NavamshaRashis["Budha"],
		NavShukraPlacement:  chart.NavamshaRashis["Shukra"],
		NavChandraPlacement: chart.NavamshaRashis["Chandra"],
		NavRahuPlacement:    chart.NavamshaRashis["Rahu"],
		NavKetuPlacement:    chart.NavamshaRashis["Ketu"],
		NavKujaPlacement:    chart.NavamshaRashis["Kuja"],
		NavGuruPlacement:    chart.NavamshaRashis["Guru"],
		NavShaniPlacement:   chart.NavamshaRashis["Shani"],

		Karakamsha: chart.Atmakaraka.Karakamsha,
		FirstHalf: chart.FirstHalf,
	}
}