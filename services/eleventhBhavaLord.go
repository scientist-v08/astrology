package services

func EleventhBhavaLordEffect(lordPlacement int8) string {
	var bhavaLordEffect = ""

	// Add to the string based on the lord placement
	if lordPlacement == 1 {
		bhavaLordEffect = `The native will be rich, happy, be a poet, be eloquent in speech and be endowed with gains`
	}
	if lordPlacement == 2 {
		bhavaLordEffect = `The native will be endowed with all kinds of wealth and all kinds of accomplishments, be charitable, religious and always happy`
	}
	if lordPlacement == 3 {
		bhavaLordEffect = `The native will be skillful in jobs, wealthy, endowed with fraternal bliss and may sometimes incurr gout pains`
	}
	if lordPlacement == 4 {
		bhavaLordEffect = `The native will gain from maternal relatives, will undertake visits to shrines and will possess happiness of house and lands`
	}
	if lordPlacement == 5 {
		bhavaLordEffect = `The native's children will be happy, educated and virtuous. The native will be religious and happy`
	}
	if lordPlacement == 6 {
		bhavaLordEffect = `The native will be afflicted by diseases, be cruel, living in foreign places and troubled by enemies`
	}
	if lordPlacement == 7 {
		bhavaLordEffect = `The native will always gain through his wife's relatives be liberal, virtous, sensous and will remain in command of his spouse`
	}
	if lordPlacement == 8 {
		bhavaLordEffect = `The native will incur reversals in his undertakings and will live long while his wife will predecease him`
	}
	if lordPlacement == 9 {
		bhavaLordEffect = `The native will be fortunate, skillful, truthful, honoured by the those in leadership positions and be affluent`
	}
	if lordPlacement == 10 {
		bhavaLordEffect = `The native will be honoured by by the those in leadership positions, be virtuous, attached to his religion, intelligent, truthful and will subdue his senses`
	}
	if lordPlacement == 11 {
		bhavaLordEffect = `The native will gain in all his undertakings while his learning and happiness will increase gradually.`
	}
	if lordPlacement == 12 {
		bhavaLordEffect = `The native will expend on good deeds, be sensous, will have many wifes and will befriend foreigners`
	}

	return bhavaLordEffect
}