package constants

import "github.com/scientist-v08/bphs/dto"

const (
    Ashwini     dto.Nakshatra = "Ashwini"
    Bharani     dto.Nakshatra = "Bharani"
    Krittika    dto.Nakshatra = "Krittika"
    Rohini      dto.Nakshatra = "Rohini"
    Mrigashira  dto.Nakshatra = "Mrigashira"
    Ardra       dto.Nakshatra = "Ardra"
    Punarvasu   dto.Nakshatra = "Punarvasu"
    Pushya      dto.Nakshatra = "Pushya"
    Ashlesha    dto.Nakshatra = "Ashlesha"
    Magha       dto.Nakshatra = "Magha"
    PurvaPhalguni  dto.Nakshatra = "Purva Phalguni"
    UttaraPhalguni dto.Nakshatra = "Uttara Phalguni"
    Hasta       dto.Nakshatra = "Hasta"
    Chitra      dto.Nakshatra = "Chitra"
    Swati       dto.Nakshatra = "Swati"
    Vishakha    dto.Nakshatra = "Vishakha"
    Anuradha    dto.Nakshatra = "Anuradha"
    Jyeshtha    dto.Nakshatra = "Jyeshtha"
    Mula        dto.Nakshatra = "Mula"
    PurvaAshadha   dto.Nakshatra = "Purva Ashadha"
    UttaraAshadha  dto.Nakshatra = "Uttara Ashadha"
    Shravana    dto.Nakshatra = "Shravana"
    Dhanishta   dto.Nakshatra = "Dhanishta"
    Shatabhisha dto.Nakshatra = "Shatabhisha"
    PurvaBhadrapada dto.Nakshatra = "Purva Bhadrapada"
    UttaraBhadrapada dto.Nakshatra = "Uttara Bhadrapada"
    Revati      dto.Nakshatra = "Revati"
)

var NakshatraNumberMap = map[dto.Nakshatra]int8{
    Ashwini:          1,
    Bharani:          2,
    Krittika:         3,
    Rohini:           4,
    Mrigashira:       5,
    Ardra:            6,
    Punarvasu:        7,
    Pushya:           8,
    Ashlesha:         9,
    Magha:           10,
    PurvaPhalguni:   11,
    UttaraPhalguni:  12,
    Hasta:           13,
    Chitra:          14,
    Swati:           15,
    Vishakha:        16,
    Anuradha:        17,
    Jyeshtha:        18,
    Mula:            19,
    PurvaAshadha:    20,
    UttaraAshadha:   21,
    Shravana:        22,
    Dhanishta:       23,
    Shatabhisha:     24,
    PurvaBhadrapada: 25,
    UttaraBhadrapada:26,
    Revati:          27,
}

var NakshatraYoniMap = map[dto.Nakshatra]string{
    Ashwini:              "Horse",
    Bharani:              "Elephant",
    Krittika:             "Sheep",
    Rohini:               "Snake",
    Mrigashira:           "Snake",
    Ardra:                "Dog",
    Punarvasu:            "Cat",
    Pushya:               "Sheep",
    Ashlesha:             "Cat",
    Magha:                "Rat",
    PurvaPhalguni:        "Rat",
    UttaraPhalguni:       "Cow",
    Hasta:                "Buffalo",
    Chitra:               "Tiger",
    Swati:                "Buffalo",
    Vishakha:             "Tiger",
    Anuradha:             "Deer",
    Jyeshtha:             "Deer",
    Mula:                 "Dog",
    PurvaAshadha:         "Monkey",
    UttaraAshadha:        "Mongoose",
    Shravana:             "Monkey",
    Dhanishta:            "Lion",
    Shatabhisha:          "Horse",
    PurvaBhadrapada:      "Lion",
    UttaraBhadrapada:     "Cow",
    Revati:               "Elephant",
}

var NakshatraGanaMap = map[dto.Nakshatra]int8{
    Ashwini:              1, // Deva
    Bharani:              2, // Manushya
    Krittika:             3, // Rakshasa
    Rohini:               2,
    Mrigashira:           1,
    Ardra:                2,
    Punarvasu:            1,
    Pushya:               1,
    Ashlesha:             3,
    Magha:                3,
    PurvaPhalguni:        2,
    UttaraPhalguni:       2,
    Hasta:                1,
    Chitra:               3,
    Swati:                1,
    Vishakha:             3,
    Anuradha:             1,
    Jyeshtha:             3,
    Mula:                 3,
    PurvaAshadha:         2,
    UttaraAshadha:        2,
    Shravana:             1,
    Dhanishta:            3,
    Shatabhisha:          3,
    PurvaBhadrapada:      2,
    UttaraBhadrapada:     2,
    Revati:               1,
}

var NakshatraNadiMap = map[dto.Nakshatra]int8{
    Ashwini:              1, // Adi
    Bharani:              2, // Madhya
    Krittika:             3, // Antya
    Rohini:               3,
    Mrigashira:           2,
    Ardra:                1,
    Punarvasu:            1,
    Pushya:               2,
    Ashlesha:             3,
    Magha:                3,
    PurvaPhalguni:        2,
    UttaraPhalguni:       1,
    Hasta:                1,
    Chitra:               2,
    Swati:                3,
    Vishakha:             3,
    Anuradha:             2,
    Jyeshtha:             1,
    Mula:                 1,
    PurvaAshadha:         2,
    UttaraAshadha:        3,
    Shravana:             3,
    Dhanishta:            2,
    Shatabhisha:          1,
    PurvaBhadrapada:      1,
    UttaraBhadrapada:     2,
    Revati:               3,
}

var NakshatraRajjuMap = map[dto.Nakshatra]string{
    Ashwini:              "Paada",
    Bharani:              "Ooru",
    Krittika:             "Nabhi",
    Rohini:               "Kanta",
    Mrigashira:           "Sira",
    Ardra:                "Kanta",
    Punarvasu:            "Nabhi",
    Pushya:               "Ooru",
    Ashlesha:             "Paada",
    Magha:                "Paada",
    PurvaPhalguni:        "Ooru",
    UttaraPhalguni:       "Nabhi",
    Hasta:                "Kanta",
    Chitra:               "Sira",
    Swati:                "Kanta",
    Vishakha:             "Nabhi",
    Anuradha:             "Ooru",
    Jyeshtha:             "Paada",
    Mula:                 "Paada",
    PurvaAshadha:         "Ooru",
    UttaraAshadha:        "Nabhi",
    Shravana:             "Kanta",
    Dhanishta:            "Sira",
    Shatabhisha:          "Kanta",
    PurvaBhadrapada:      "Nabhi",
    UttaraBhadrapada:     "Ooru",
    Revati:               "Paada",
}