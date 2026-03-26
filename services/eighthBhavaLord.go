package services

func EighthBhavaLordEffect(lordPlacement int8) string {
	var bhavaLordEffect = ""

	// Add to the string based on the lord placement
	if lordPlacement == 1 {
		bhavaLordEffect = `The native will be devoid of physical happiness and suffer from wounds. Be hostile to Gods and religious people.`
	}
	if lordPlacement == 2 {
		bhavaLordEffect = `The native will enjoy little wealth and will not regain lost wealth`
	}
	if lordPlacement == 3 {
		bhavaLordEffect = `The native will be devoid of fraternal happiness, be indolent and devoid of strength and servants.`
	}
	if lordPlacement == 4 {
		bhavaLordEffect = `The native may be deprived of own mother. Be devoid of house, lands and happiness and will betray friends`
	}
	if lordPlacement == 5 {
		bhavaLordEffect = `The native will be dull witted, will have limited number of children be long lived and wealthy.
		The wealth acquisition will be high but not steady. Will keep on changing his line of thinking.
		May not enjoy happiness through children.`
	}
	if lordPlacement == 6 {
		bhavaLordEffect = `Will win over enemies and have a long life even if not healthy.`
	}
	if lordPlacement == 7 {
		bhavaLordEffect = `The native will have more then 1 spouse. Will be an expert in stealing others things.`
	}
	if lordPlacement == 8 {
		bhavaLordEffect = `The native will be long lived, be a theif, be blame worthy and blame others as well`
	}
	if lordPlacement == 9 {
		bhavaLordEffect = `The native will betray his religion and be a heterodox will beget a wicked spouse and will steal others wealth.
		The native's father maybe short lived and may not understand his own father.
		The native will be addicted to others spouses while the spouse will go to others for progeny.`
	}
	if lordPlacement == 10 {
		bhavaLordEffect = `If the 10th house is bereft of benefic then the native will be a liar and be bereft of livelihood.`
	}
	if lordPlacement == 11 {
		bhavaLordEffect = `The native will be devoid of wealth if conjunct a malefic else if conjunct a benefic will be long lived.
		For taurus ascendant jupiter in 11th is not bad and for scorpio ascendant mercury in 11th is good for wealth and fame.`
	}
	if lordPlacement == 12 {
		bhavaLordEffect = `If there is a malefic in the 12th along with the 8th lord the native will spend money on vices and die young.`
	}

	return bhavaLordEffect
}