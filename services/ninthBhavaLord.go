package services

func NinthBhavaLordEffect(lordPlacement int8) string {
	var bhavaLordEffect = ""

	// Add to the string based on the lord placement
	if lordPlacement == 1 {
		bhavaLordEffect = `The native will be fortunate, virtuous, honoured by those in leadership positions and learned and honoured by public`
	}
	if lordPlacement == 2 {
		bhavaLordEffect = `The native will be a scholar and endowed with happiness from wife and sons.`
	}
	if lordPlacement == 3 {
		bhavaLordEffect = `The native will be endowed with fraternal bliss, be wealthy and virtuous and charming`
	}
	if lordPlacement == 4 {
		bhavaLordEffect = `The native will enjoy houses, vehicales and happiness and be devoted to his mother.`
	}
	if lordPlacement == 5 {
		bhavaLordEffect = `The native will be endowed with sons and prosperity, devoted to elders, bold, charitable and leanred.`
	}
	if lordPlacement == 6 {
		bhavaLordEffect = `The native will enjoy meagre prosperity, be devoid of happiness from maternal relatives and be always troubled by enemies.`
	}
	if lordPlacement == 7 {
		bhavaLordEffect = `The native will beget happiness from marriage, be virtuous and famous.`
	}
	if lordPlacement == 8 {
		bhavaLordEffect = `The native will not be prosperous`
	}
	if lordPlacement == 9 {
		bhavaLordEffect = `The native will be endowed with abundant fortunes, virtues ad beauty and will enjoy much happiness from co-born`
	}
	if lordPlacement == 10 {
		bhavaLordEffect = `The native will obtain leadership positions and be virtuous and dear to all`
	}
	if lordPlacement == 11 {
		bhavaLordEffect = `The native will enjoy gains, be virtuous and meritorious in acts (NA for Mithuna lagna)`
	}
	if lordPlacement == 12 {
		bhavaLordEffect = `The native will incure loss of fortunes will spend money on auspicious accounts and thereby become poor.`
	}

	return bhavaLordEffect
}