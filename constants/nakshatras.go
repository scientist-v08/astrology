package constants

import "github.com/scientist-v08/bphs/models"

const (
    Ashwini     models.Nakshatra = "Ashwini"
    Bharani     models.Nakshatra = "Bharani"
    Krittika    models.Nakshatra = "Krittika"
    Rohini      models.Nakshatra = "Rohini"
    Mrigashira  models.Nakshatra = "Mrigashira"
    Ardra       models.Nakshatra = "Ardra"
    Punarvasu   models.Nakshatra = "Punarvasu"
    Pushya      models.Nakshatra = "Pushya"
    Ashlesha    models.Nakshatra = "Ashlesha"
    Magha       models.Nakshatra = "Magha"
    PurvaPhalguni  models.Nakshatra = "Purva Phalguni"
    UttaraPhalguni models.Nakshatra = "Uttara Phalguni"
    Hasta       models.Nakshatra = "Hasta"
    Chitra      models.Nakshatra = "Chitra"
    Swati       models.Nakshatra = "Swati"
    Vishakha    models.Nakshatra = "Vishakha"
    Anuradha    models.Nakshatra = "Anuradha"
    Jyeshtha    models.Nakshatra = "Jyeshtha"
    Mula        models.Nakshatra = "Mula"
    PurvaAshadha   models.Nakshatra = "Purva Ashadha"
    UttaraAshadha  models.Nakshatra = "Uttara Ashadha"
    Shravana    models.Nakshatra = "Shravana"
    Dhanishta   models.Nakshatra = "Dhanishta"
    Shatabhisha models.Nakshatra = "Shatabhisha"
    PurvaBhadrapada models.Nakshatra = "Purva Bhadrapada"
    UttaraBhadrapada models.Nakshatra = "Uttara Bhadrapada"
    Revati      models.Nakshatra = "Revati"
)

var NakshatraNumberMap = map[models.Nakshatra]int8{
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

var NakshatraYoniMap = map[models.Nakshatra]string{
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

var NakshatraGanaMap = map[models.Nakshatra]int8{
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

var NakshatraNadiMap = map[models.Nakshatra]int8{
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

var NakshatraRajjuMap = map[models.Nakshatra]string{
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