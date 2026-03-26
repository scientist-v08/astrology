package services

import (
	"strconv"

	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
)

func getYoniSexualFantasy(yoni string) string {
	switch yoni {
	case "Horse":
		return "Seems aloof about sex, yet possesses powerful hidden passion and stamina and craves fast, free, adventurous encounters with real endurance once the session begins"

	case "Elephant":
		return "Slow, deep, long-lasting, very sensual and heavy body contact, enjoys prolonged foreplay and staying inside long"

	case "Sheep":
		return "Gentle, submissive, loves being led and guided, enjoys soft repeated movements and feeling protected"

	case "Snake":
		return "Intense eye contact, hypnotic seduction, twisting and coiling positions, very sensual tongue play, slow penetration with sudden deep thrusts"

	case "Dog":
		return "Loyal but lustful, loves from behind, quick and enthusiastic sessions, high stamina, can go many rounds"

	case "Cat":
		return "Playful, teasing, loves being chased and caught, independent but becomes very demanding once aroused, lots of scratching and biting"

	case "Rat":
		return "Fast, sneaky, opportunistic sex, loves quickies in unusual places, very high frequency, multiple short sessions"

	case "Cow":
		return "Nurturing, slow and steady, sensual massage-like lovemaking, enjoys being mounted and giving complete surrender, long comfortable sessions"

	case "Buffalo":
		return "Powerful, dominant, earthy, heavy pounding, likes raw physical strength and stamina displays, can be very territorial during sex"

	case "Tiger":
		return "Fierce, predatory, dominant, loves the hunt and pouncing, intense eye contact, growling, powerful thrusts, can be rough and passionate"

	case "Deer":
		return "Graceful, shy at first, then very sensitive and responsive, loves gentle but deep lovemaking in beautiful settings, easily startled"

	case "Monkey":
		return "Teases and flirts for a very long time, playful and mischievous foreplay, but once it starts — short, explosive, and restless"

	case "Mongoose":
		return "Fearless, quick attacks, loves fighting for dominance even during sex, intense battles that turn into passionate mating, very courageous"

	case "Lion":
		return "King/Queen energy, very proud, loves being worshipped, BDSM foreplay, powerful and regal positions, roars during climax"

	default:
		return "Unknown Yoni no specific fantasy information available"
	}
}

func YoniCalculatorFunc(groomNakshatra string, brideNakshatra string) models.KutaTypeRes {
	yoniScore := constants.YoniCompatibility[constants.NakshatraYoniMap[models.Nakshatra(groomNakshatra)]][constants.NakshatraYoniMap[models.Nakshatra(brideNakshatra)]]
	return models.KutaTypeRes{
		Index: 3,
		Score: float32(yoniScore),
		Comments: "The grooms nakshatra is " + groomNakshatra + " and its passion animal is " + constants.NakshatraYoniMap[models.Nakshatra(groomNakshatra)]+ 
		" which has the following fantasy:"+ getYoniSexualFantasy(constants.NakshatraYoniMap[models.Nakshatra(groomNakshatra)]) + ". The brides nakshatra is " +
		brideNakshatra + " and its passion animal is " + constants.NakshatraYoniMap[models.Nakshatra(brideNakshatra)] + 
		" which has the following fantasy:"+ getYoniSexualFantasy(constants.NakshatraYoniMap[models.Nakshatra(brideNakshatra)]) +
		". The union gives us a score of " + strconv.Itoa(int(yoniScore)),
	}
}