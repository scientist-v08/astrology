package services

func ThirdBhavaLordEffect(lordPlacement int8) string {
	var bhavaLordEffect = ""

	// Add to the string based on the lord placement
	if lordPlacement == 1 {
		bhavaLordEffect = `The native will have self earned wealth be intelligent even if not gone to school.
		The native will have a unslakable lust`
	}
	if lordPlacement == 2 {
		bhavaLordEffect = `The native will not have valour and be lazy and have an eye on others wives and wealth. May resort to unnatural means for gratification.`
	}
	if lordPlacement == 3 {
		bhavaLordEffect = `The native will get happy through co-borns`
	}
	if lordPlacement == 4 {
		bhavaLordEffect = `The native will be wealthy, intelligent but acquire a wicked spouse`
	}
	if lordPlacement == 5 {
		bhavaLordEffect = `The native will have sons and if conjunct/ aspected by a malefic will have a formidable wife.`
	}
	if lordPlacement == 6 {
		bhavaLordEffect = `The native will be inimical to the co-born, be affluent and be dear to his maternal aunt`
	}
	if lordPlacement == 7 {
		bhavaLordEffect = `The native will be interested in serving those in leadership positions and will be happy during the end of his life. Do not make own business`
	}
	if lordPlacement == 8 {
		bhavaLordEffect = `The native will be a thief`
	}
	if lordPlacement == 9 {
		bhavaLordEffect = `The native will not get happiness from father but make fortune through his wife and enjoy happiness through his own children`
	}
	if lordPlacement == 10 {
		bhavaLordEffect = `The native will have all kinds of happiness and be self-made and interested in nurturing wicked females`
	}
	if lordPlacement == 11 {
		bhavaLordEffect = `The native will gain in trading be intelligent but may not be educated and may not be a worthy friend.`
	}
	if lordPlacement == 12 {
		bhavaLordEffect = `The native will spend on evil deeds, have a wicked father and will be fortunate through a female / wife.
		The native will be bestowed with all kinds of happiness but still feel miserable.`
	}

	return bhavaLordEffect
}