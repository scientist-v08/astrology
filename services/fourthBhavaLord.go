package services

func FourthBhavaLordEffect(lordPlacement int8) string {
	var bhavaLordEffect = ""

	// Add to the string based on the lord placement
	if lordPlacement == 1 {
		bhavaLordEffect = `The native will be endowed with learning, virtues, ornaments,lands, vehicles and maternal happiness.
		The native will acquire incomparable learning in various branches but has the risk of becoming a monk by renouncing everything.`
	}
	if lordPlacement == 2 {
		bhavaLordEffect = `The native will enjoy pleasures of all kinds, obtain wealth, family life and honour and be cunning.
		It will bring gains from mother and maternal relatives. The native will join evil company and get satified by stealing`
	}
	if lordPlacement == 3 {
		bhavaLordEffect = `The native will be valourous, have servants, be liberal and possess self earned wealth and be free from diseases`
	}
	if lordPlacement == 4 {
		bhavaLordEffect = `The native will obtain leadership positions, possess all kinds of wealth, be skillful, virtuous, leanred and happy and well disposed to the spouse.`
	}
	if lordPlacement == 5 {
		bhavaLordEffect = `The native will behappy and liked. Be devoted to MahaVishnu, be virtuous, honorouable and have self earned wealth.`
	}
	if lordPlacement == 6 {
		bhavaLordEffect = `The native will not have maternal happiness, be short-tempered, be a thief / magician, be independant in action.
		The natives mother will be sickly and a source of worry`
	}
	if lordPlacement == 7 {
		bhavaLordEffect = `The native will give up the property of his father and be edcated and have negetive skills in public speaking.`
	}
	if lordPlacement == 8 {
		bhavaLordEffect = `The native will not have motherly comforts and may lose the ability to have children as well.`
	}
	if lordPlacement == 9 {
		bhavaLordEffect = `The native will be dear to many, be devoted to God, be virtuous, honourable and endowed with happiness`
	}
	if lordPlacement == 10 {
		bhavaLordEffect = `The native will enjoy royal honours and will conquer the 5 senses. Will have abundant self-made properties.`
	}
	if lordPlacement == 11 {
		bhavaLordEffect = `The native will have fear of secret diseases, be liberal, virtuous, charitable and helpful to others`
	}
	if lordPlacement == 12 {
		bhavaLordEffect = `The native will be devoid of domestic and other comforts, will have vices and be foolish.`
	}

	return bhavaLordEffect
}