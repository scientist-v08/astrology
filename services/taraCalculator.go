package services

import (
	"fmt"
	"slices"

	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
)

type TaraResult struct {
    Comments string
    Score    int
}

func Mod(n, m int8) int8 {
    return ((n % m) + m) % m
}

func TaraType(
    taraNumber int8,
    isBride bool, // true if calculating from bride's perspective, false = groom's perspective
    groomNakshatra models.Nakshatra,
    brideNakshatra models.Nakshatra,
) TaraResult {

    // Special handling for Janma Tara (case 1)
    if taraNumber == 1 {
        var comment string
        var score = 0

        if groomNakshatra == brideNakshatra {
            // Same birth star — often considered inauspicious unless mitigated
            if isBride {
                comment = "" // empty when bride's view (some traditions skip comment)
            } else {
                comment = fmt.Sprintf("%s nakshatra aligns with %s's birth star, generally indicating challenges.", 
                    groomNakshatra, brideNakshatra)
            }
        } else {
            if isBride {
                comment = ""
            } else {
                comment = fmt.Sprintf("%s nakshatra may bring challenges to %s and vice versa.", 
                    groomNakshatra, brideNakshatra)
            }
        }

        return TaraResult{
            Comments: comment,
            Score:    score,
        }
    }

    // Standard Tara types (2–9)
    switch taraNumber {
    case 2: // Sampat Tara
        return TaraResult{
            Comments: fmt.Sprintf("%s nakshatra brings wealth, prosperity, and good fortune to %s nakshatra.",
                groomNakshatra, brideNakshatra),
            Score: 3,
        }

    case 3: // Vipat Tara
        return TaraResult{
            Comments: fmt.Sprintf("%s nakshatra may bring obstacles, dangers, or challenges to %s nakshatra.",
                groomNakshatra, brideNakshatra),
            Score: 0,
        }

    case 4: // Kshema Tara
        return TaraResult{
            Comments: fmt.Sprintf("%s nakshatra ensures well-being, growth, and security for %s nakshatra.",
                groomNakshatra, brideNakshatra),
            Score: 3,
        }

    case 5: // Pratyak / Pratyari Tara
        return TaraResult{
            Comments: fmt.Sprintf("%s nakshatra may indicate opposition, enemies, or conflict for %s nakshatra.",
                groomNakshatra, brideNakshatra),
            Score: 0,
        }

    case 6: // Sadhaka Tara
        return TaraResult{
            Comments: fmt.Sprintf("%s nakshatra is supportive, helping %s nakshatra achieve goals and find success.",
                groomNakshatra, brideNakshatra),
            Score: 3,
        }

    case 7: // Vadha / Naidhana Tara
        return TaraResult{
            Comments: fmt.Sprintf("%s nakshatra may bring significant distress, destruction, or serious trouble to %s nakshatra.",
                groomNakshatra, brideNakshatra),
            Score: 0,
        }

    case 8: // Mitra Tara
        return TaraResult{
            Comments: fmt.Sprintf("%s nakshatra fosters friendship, harmony, and mutual understanding with %s nakshatra.",
                groomNakshatra, brideNakshatra),
            Score: 3,
        }

    case 9, 0: // Parama Mitra / Ati Mitra Tara (9 mod 9 = 0)
        return TaraResult{
            Comments: fmt.Sprintf("%s nakshatra brings the highest form of luck and great friendship to %s nakshatra.",
                groomNakshatra, brideNakshatra),
            Score: 3,
        }

    default:
        return TaraResult{
            Comments: "Unable to calculate Tara compatibility",
            Score:    0,
        }
    }
}

func TaraCalculatorFunc(groomNakshatra string, brideNakshatra string) models.KutaTypeRes {
	groomNakshatraNum := constants.NakshatraNumberMap[models.Nakshatra(groomNakshatra)]
	brideNakshatraNum := constants.NakshatraNumberMap[models.Nakshatra(brideNakshatra)]
	groomTara := Mod(brideNakshatraNum - groomNakshatraNum + 1, 9)
	brideTara := Mod(groomNakshatraNum - brideNakshatraNum + 1,9)
	groom := TaraType(groomTara, false, models.Nakshatra(groomNakshatra), models.Nakshatra(brideNakshatra))
	bride := TaraType(brideTara, true, models.Nakshatra(brideNakshatra), models.Nakshatra(groomNakshatra))
	var note string
	if slices.Contains([]string{"Mrigashira", "Magha", "Swati", "Anuradha"}, groomNakshatra) || slices.Contains([]string{"Mrigashira", "Magha", "Swati", "Anuradha"}, brideNakshatra) {
		note = " Note:- Mrigashira, Magha, Swati & Anuradha are exempt from bad luck"
	} else {
		note = ""
	}
    score := (float32(groom.Score) + float32(bride.Score)) / 2
	return models.KutaTypeRes{
		Index: 2,
		Score: score,
		Comments: groom.Comments + bride.Comments + note,
	}
}