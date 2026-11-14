package shadbala

import (
	"github.com/scientist-v08/bphs/constants"
	"github.com/scientist-v08/bphs/models"
)

// Sign indices below follow constants.RaashiOrder (0 = Aries ... 11 = Pisces).

// ---------------------------------------------------------------------------
// Sign lords
// ---------------------------------------------------------------------------

var signLord = [12]models.Raashyadhipati{
	constants.Kuja, constants.Shukra, constants.Budha, constants.Chandra,
	constants.Surya, constants.Budha, constants.Shukra, constants.Kuja,
	constants.Guru, constants.Shani, constants.Shani, constants.Guru,
}

// ---------------------------------------------------------------------------
// Dignity
// ---------------------------------------------------------------------------

type Dignity string

const (
	Moolatrikona Dignity = "moolatrikona"
	Swakshetra   Dignity = "swakshetra"
	AdhiMitra    Dignity = "adhimitra"
	Mitra        Dignity = "mitra"
	Sama         Dignity = "sama"
	Shatru       Dignity = "shatru"
	AdhiShatru   Dignity = "adhishatru"
)

// Virupas per dignity (BPHS). The book quotes 45/30/20/15/10/4; the 7th
// grade, adhishatru = 2, is from BPHS and is included here.
var dignityVirupas = map[Dignity]float64{
	Moolatrikona: 45,
	Swakshetra:   30,
	AdhiMitra:    22.5,
	Mitra:        15,
	Sama:         7.5,
	Shatru:       3.75,
	AdhiShatru:   1.875,
}

// Moolatrikona ranges (sign index, start deg, end deg) - applies in Rasi only.
type moolaRange struct {
	sign       int
	start, end float64
}

var moolatrikona = map[models.Raashyadhipati]moolaRange{
	constants.Surya:   {4, 0, 20},  // Leo 0-20
	constants.Chandra: {1, 3, 30},  // Taurus 3-30
	constants.Kuja:    {0, 0, 12},  // Aries 0-12
	constants.Budha:   {5, 15, 20}, // Virgo 15-20
	constants.Guru:    {8, 0, 10},  // Sagittarius 0-10
	constants.Shukra:  {6, 0, 15},  // Libra 0-15
	constants.Shani:   {10, 0, 20}, // Aquarius 0-20
}

// ---------------------------------------------------------------------------
// Friendship
// ---------------------------------------------------------------------------

// Natural (naisargika) friends and enemies per BPHS. Anyone not listed is neutral.
var naturalFriends = map[models.Raashyadhipati][]models.Raashyadhipati{
	constants.Surya:   {constants.Chandra, constants.Kuja, constants.Guru},
	constants.Chandra: {constants.Surya, constants.Budha},
	constants.Kuja:    {constants.Surya, constants.Chandra, constants.Guru},
	constants.Budha:   {constants.Surya, constants.Shukra},
	constants.Guru:    {constants.Surya, constants.Chandra, constants.Kuja},
	constants.Shukra:  {constants.Budha, constants.Shani},
	constants.Shani:   {constants.Budha, constants.Shukra},
}

var naturalEnemies = map[models.Raashyadhipati][]models.Raashyadhipati{
	constants.Surya:   {constants.Shukra, constants.Shani},
	constants.Chandra: {},
	constants.Kuja:    {constants.Budha},
	constants.Budha:   {constants.Chandra},
	constants.Guru:    {constants.Budha, constants.Shukra},
	constants.Shukra:  {constants.Surya, constants.Chandra},
	constants.Shani:   {constants.Surya, constants.Chandra, constants.Kuja},
}

func contains(list []models.Raashyadhipati, g models.Raashyadhipati) bool {
	for _, x := range list {
		if x == g {
			return true
		}
	}
	return false
}

// naturalRelation: +1 friend, 0 neutral, -1 enemy.
func naturalRelation(from, to models.Raashyadhipati) int {
	switch {
	case contains(naturalFriends[from], to):
		return 1
	case contains(naturalEnemies[from], to):
		return -1
	}
	return 0
}

// temporaryRelation (tatkalika): +1 if `to` sits in the 2nd,3rd,4th,10th,11th
// or 12th sign from `from` in the Rasi chart, otherwise -1.
func temporaryRelation(fromSign, toSign int) int {
	house := ((toSign-fromSign)%12+12)%12 + 1
	switch house {
	case 2, 3, 4, 10, 11, 12:
		return 1
	}
	return -1
}

// compoundRelation combines natural + temporary friendship (panchadha maitri):
//
//	friend+friend  -> adhimitra   neutral+friend -> mitra
//	friend+enemy   -> sama        neutral+enemy  -> shatru
//	enemy+friend   -> sama        enemy+enemy    -> adhishatru
func compoundRelation(g, lord models.Raashyadhipati, rasiSign map[models.Raashyadhipati]int) Dignity {
	score := naturalRelation(g, lord) + temporaryRelation(rasiSign[g], rasiSign[lord])
	switch score {
	case 2:
		return AdhiMitra
	case 1:
		return Mitra
	case 0:
		return Sama
	case -1:
		return Shatru
	default:
		return AdhiShatru
	}
}

// dignityIn returns the dignity of graha g placed in `sign` of some varga.
// Moolatrikona is only considered in the Rasi chart (needs the exact degree).
func dignityIn(g models.Raashyadhipati, sign int, rasiLon float64, isRasi bool, rasiSign map[models.Raashyadhipati]int) Dignity {
	if isRasi {
		if m, ok := moolatrikona[g]; ok && sign == m.sign {
			d := degreesInSign(rasiLon)
			if d >= m.start && d < m.end {
				return Moolatrikona
			}
		}
	}
	lord := signLord[sign]
	if lord == g {
		return Swakshetra
	}
	return compoundRelation(g, lord, rasiSign)
}

// ---------------------------------------------------------------------------
// Divisional charts computed from longitude.
// (D1 and D9 are NOT computed here - they come from the chart context.)
// ---------------------------------------------------------------------------

// horaSign: odd sign -> 1st half Sun (Leo), 2nd half Moon (Cancer);
//
//	even sign -> 1st half Moon (Cancer), 2nd half Sun (Leo).
var umaShambhuHora = [12][2]int{
    {0, 1},   // Aries
    {3, 2},   // Taurus
    {4, 5},   // Gemini
    {7, 6},   // Cancer
    {8, 9},   // Leo
    {11, 10}, // Virgo
    {0, 1},   // Libra
    {3, 2},   // Scorpio
    {4, 5},   // Sagittarius
    {7, 6},   // Capricorn
    {8, 9},   // Aquarius
    {11, 10}, // Pisces
}

func calculateUmaShambhuHora(lon float64) int {
    sign := signOf(lon)
    degree := degreesInSign(lon)

    half := 0
    if degree >= 15 {
        half = 1
    }

    return umaShambhuHora[sign][half]
}

// drekkanaSign: 1st decanate = same sign, 2nd = 5th from it, 3rd = 9th from it.
func umaShambhuDrekkana(lon float64) int {
    sign := signOf(lon)
    degree := degreesInSign(lon)

    // Determine the Drekkana division (0, 1, 2).
    division := int(degree / 10)

    // Reverse division order for even signs.
    divisionOrder := division

    if !isOddSign(sign) {
        divisionOrder = 2 - division
    }

    // Determine the D3 sign.
    divisionalSign := (sign*3 + divisionOrder) % 12

    return divisionalSign
}

// saptamsaSign: 7 parts of 4°17'8.57". Odd sign counts from itself,
// even sign counts from its 7th.
func saptamsaSign(lon float64) int {
	part := int(degreesInSign(lon) / (30.0 / 7))
	if part > 6 {
		part = 6
	}

	start := signOf(lon)

	if isOddSign(start) {
		// Odd sign: original sign, forward counting.
		return (start + part) % 12
	}

	// Even sign: 7th sign, reverse counting.
	return (start + 6 - part + 12) % 12
}

// dwadasamsaSign: 12 parts of 2°30', counted from the sign itself.
func dwadasamsaSign(lon float64) int {
	part := int(degreesInSign(lon) / 2.5)
	if part > 11 {
		part = 11
	}
	return (signOf(lon) + part) % 12
}

// trimsamsaSign returns the sign owned by the trimsamsa lord.
// Odd signs:  Mars 0-5, Saturn 5-10, Jupiter 10-18, Mercury 18-25, Venus 25-30.
// Even signs: Venus 0-5, Mercury 5-12, Jupiter 12-20, Saturn 20-25, Mars 25-30.
// Sign owners used: Mars=Aries(0)/Scorpio(7), Venus=Libra(6)/Taurus(1),
// Saturn=Aquarius(10)/Capricorn(9), Jupiter=Sagittarius(8)/Pisces(11), Mercury=Gemini(2)/Virgo(5).
func trimsamsaSign(lon float64) int {
	d := degreesInSign(lon)
	if isOddSign(signOf(lon)) {
		switch {
		case d < 5:
			return 0 // Mars
		case d < 10:
			return 10 // Saturn
		case d < 18:
			return 8 // Jupiter
		case d < 25:
			return 2 // Mercury
		default:
			return 6 // Venus
		}
	}
	switch {
	case d < 5:
		return 1 // Venus
	case d < 12:
		return 5 // Mercury
	case d < 20:
		return 11 // Jupiter
	case d < 25:
		return 9 // Saturn
	default:
		return 7 // Mars
	}
}

// ---------------------------------------------------------------------------
// Saptavargaja Bala
// ---------------------------------------------------------------------------

// saptavargajaBala sums the dignity points of g across the seven vargas:
// Rasi, Hora, Drekkana, Saptamsa, Navamsa, Dwadasamsa, Trimsamsa.
// Rasi and Navamsa signs come from the chart context (request body);
// the rest are derived from the longitude.
func saptavargajaBala(g models.Raashyadhipati, ctx *chartContext) ([]SaptavargaEntry, float64) {
	lon := ctx.longitude[g]

	vargas := []struct {
		name string
		sign int
	}{
		{"rasi", ctx.rasiSign[g]},
		{"hora", calculateUmaShambhuHora(lon)},
		{"drekkana", umaShambhuDrekkana(lon)},
		{"saptamsa", saptamsaSign(lon)},
		{"navamsa", ctx.navamsaSign[g]},
		{"dwadasamsa", dwadasamsaSign(lon)},
		{"trimsamsa", trimsamsaSign(lon)},
	}

	entries := make([]SaptavargaEntry, 0, len(vargas))
	total := 0.0
	for _, v := range vargas {
		dg := dignityIn(g, v.sign, lon, v.name == "rasi", ctx.rasiSign)
		pts := dignityVirupas[dg]
		entries = append(entries, SaptavargaEntry{Varga: v.name, Sign: v.sign, Dignity: dg, Virupas: pts})
		total += pts
	}
	return entries, total
}
