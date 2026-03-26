package services

func FirstBhavaLordEffect(lordPlacement int8) string {
	var firstBhavaLordEffect = ""

	// Add to the string based on the lord placement
	if lordPlacement == 1 {
		firstBhavaLordEffect = `The native will have good health and prowess. Will be intelligent, fickle minded,
		will have 2 wives and have sex with women other then his wives. The native will be fearless and long lived`
	}
	if lordPlacement == 2 {
		firstBhavaLordEffect = `The native will be scholarly, happy, will possess good qualities, be religious, honourable and have many wives.
		Further, the native will face problems having children and will gain through his own efforts`
	}
	if lordPlacement == 3 {
		firstBhavaLordEffect = `The native will be equal to those in leadership positions, be respected by others and indulge in unnatural methods of sex. The native will be 
		equal to a lion in valour, be endowed with wealth, have 2 wives and be intelligent`
	}
	if lordPlacement == 4 {
		firstBhavaLordEffect = `The native will enjoy maternal and paternal happiness and have many brothers be lustful, virtuous and charming.
		The native will be of noble descent and will prosper through own efforts and will enjoy physical pleasure.`
	}
	if lordPlacement == 5 {
		firstBhavaLordEffect = `Happiness through children will be mediocre, likely to lose the first child, be honourable, be short tempered and dear to those in leadership positions`
	}
	if lordPlacement == 6 {
		firstBhavaLordEffect = `If the 1st lord is with a malefic with no benefic protection, then the native will be troubled by enemies and devoid of physical pleasure.
		This does not apply to scorpio or taurus ascendant. But this also ensures more then 1 marriage and gaining wealth and status`
	}
	if lordPlacement == 7 {
		firstBhavaLordEffect = `The native will become like those in leadership positions`
	}
	if lordPlacement == 8 {
		firstBhavaLordEffect = `The native will be an accomplished scholar, be sickly and thievish and have sex with other's spouse and be long lived.
		The negetive effects may not apply to Aries and Libra ascendants.`
	}
	if lordPlacement == 9 {
		firstBhavaLordEffect = `The native will be fortunate be a devotee of Mahavishnu, be eloquent in speech and have sons and wealth.
		Will gain from his father and be dear to co-borns.`
	}
	if lordPlacement == 10 {
		firstBhavaLordEffect = `The native will be honoured and have fame amongst men and have self earned wealth and will defintely have co-borns`
	}
	if lordPlacement == 11 {
		firstBhavaLordEffect = `The native will have gains, good qualities, fame and many wives.`
	}
	if lordPlacement == 12 {
		firstBhavaLordEffect = `If there is no benefic aspect/ conjunction with the 1st lord the native will spend on wasteful things,
		be short tempered and not have physical pleasure`
	}

	return firstBhavaLordEffect
}