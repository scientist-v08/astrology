package services

import (
	"strings"

	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
)

type pair struct {
    groom string
    bride string
}

func RajjuAndOtherDoshasCalculator(groomNakshatra string, brideNakshatra string, groomChandra string, brideChandra string) string {
	groomRajju := constants.NakshatraRajjuMap[models.Nakshatra(groomNakshatra)]
	brideRajju := constants.NakshatraRajjuMap[models.Nakshatra(brideNakshatra)]
	groomFirstHouse := constants.RaashyadhipatiMapStore[groomChandra]
	brideFirstHouse := constants.RaashyadhipatiMapStore[brideChandra]
	var rajjuAndOtherComments strings.Builder
	isSameRuler := groomFirstHouse == brideFirstHouse
	forbiddenPairs := []pair{
		{"Krittika", "Ashlesha"},
		{"Ashlesha", "Swati"},
		{"Chitra", "Purva Ashadha"},
		{"Anuradha", "Dhanishta"},
		{"Dhanishta", "Bharani"},
		{"Shatabhisha", "Krittika"},
		{"Revati", "Ardra"},
		{"Jyeshtha", "Uttara Phalguni"},
		{"Vishakha", "Shravana"},
		{"Ashwini", "Shravana"},
	}

	if groomRajju == brideRajju {
		switch groomRajju {
		case "Paada":
			rajjuAndOtherComments.WriteString("Rajju dosha: Paada (feet) - May cause long distance relationship / separation\n")
		case "Ooru":
			rajjuAndOtherComments.WriteString("Rajju dosha: Ooru (Thigh) - May cause decline of wealth / financial problems\n")
		case "Nabhi":
			rajjuAndOtherComments.WriteString("Rajju dosha: Nabhi (navel) - May cause loss of children / progeny issues\n")
		case "Kanta":
			rajjuAndOtherComments.WriteString("Rajju dosha: Kanta (neck) - May cause loss of spouse / serious danger to partner\n")
		case "Sira":
			rajjuAndOtherComments.WriteString("Rajju dosha: Sira (head) - May cause loss of spouse / very severe\n")
		}
		if isSameRuler {
			rajjuAndOtherComments.WriteString("But this pair is exempt from this dosha\n")
		}
	} else {
		rajjuAndOtherComments.WriteString("No Rajju dosha.")
	}

	for _, p := range forbiddenPairs {
		if (groomNakshatra == p.groom && brideNakshatra == p.bride) {
			rajjuAndOtherComments.WriteString("This pair is not recommended irrespective of the AstaKuta score (forbidden / prohibited nakshatra combination)")
		}
	}

	if (groomNakshatra == "Anuradha" && brideNakshatra == "Rohini") || (brideNakshatra == "Anuradha" && groomNakshatra == "Rohini") {
		rajjuAndOtherComments.WriteString("According to Astakuta this pair is the second / third best. However, throughout history astrologers have observed that this is the #1 pair in terms of personality compatibility")
	}

	return rajjuAndOtherComments.String()
}