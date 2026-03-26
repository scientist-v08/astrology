package services

func SixthBhavaLordEffect(lordPlacement int8) string {
	var bhavaLordEffect = ""

	// Add to the string based on the lord placement
	if lordPlacement == 1 {
		bhavaLordEffect = `The native will be sickly, famous, inimical to his own men, rich, honorouable, adventoruos and virtuous`
	}
	if lordPlacement == 2 {
		bhavaLordEffect = `Adventoruous, famous among his race men, will live in alien places, be a skillful speaker and interested in his own work.`
	}
	if lordPlacement == 3 {
		bhavaLordEffect = `Be short tempered, bereft of courage, inimical to his co-born and have disobedient servants.`
	}
	if lordPlacement == 4 {
		bhavaLordEffect = `Be devoid of maternal happiness. Be intelligent, be a liar, be jealous and rich`
	}
	if lordPlacement == 5 {
		bhavaLordEffect = `The native will have fluctuating finances be inimical to his children and friends. Be happy selfish and kind.`
	}
	if lordPlacement == 6 {
		bhavaLordEffect = `The native will have enimity with his own kinsmen and be happy with others. The native will have a long life`
	}
	if lordPlacement == 7 {
		bhavaLordEffect = `The native will not derive happiness through wedlock. Be famous, virtuous, honorouable, adventorous and wealthy.
		The native's spouse may be an enemy to the native`
	}
	if lordPlacement == 8 {
		bhavaLordEffect = `Be sickly, desire others wealth, be interested in others spouses and be impure. The native will keep on incurring enmity with others and be sad.`
	}
	if lordPlacement == 9 {
		bhavaLordEffect = `Maybe involved in real estate business and face ups and downs in finances`
	}
	if lordPlacement == 10 {
		bhavaLordEffect = `Will not be disposed to his father and be happy in foreign countries and be a gifted speaker.
		Will be greatly valourous and litigation on account of ancestral property.`
	}
	if lordPlacement == 11 {
		bhavaLordEffect = `The native will gain wealth through enemies and have progeic issues.`
	}
	if lordPlacement == 12 {
		bhavaLordEffect = `The native will spend on vices, be hostile to learned people and will torture living beings. Will have questionable character and 
		on derving pleasure from other females.`
	}

	return bhavaLordEffect
}