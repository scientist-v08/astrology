package services

func FifthBhavaLordEffect(lordPlacement int8) string {
	var bhavaLordEffect = ""

	// Add to the string based on the lord placement
	if lordPlacement == 1 {
		bhavaLordEffect = `The native will be scholarly, get progenic happiness but maybe put into an unpleasant situation with one of the sons.
		Be a miser, crooked and steal others wealth.`
	}
	if lordPlacement == 2 {
		bhavaLordEffect = `The native will have many sons and wealth, be honorouable, be attached to his spouse and be famous in the world.`
	}
	if lordPlacement == 3 {
		bhavaLordEffect = `The native will be attached to the co-born, be a liar, a miser and interested in only his own work.
		The native is likely to be a imposter and may not help anyone in any manner.`
	}
	if lordPlacement == 4 {
		bhavaLordEffect = `The native will be happy and obtain maternal happiness, wealth and intelligence and obtain leadership roles in life.
		The native will come to own a posh house and start acquiring prosperity from a young age and the mother will have a long life.`
	}
	if lordPlacement == 5 {
		bhavaLordEffect = `The native will have progeny if related to benefic or no children if related to malefics. Will make the native virtuous and dear to friends.`
	}
	if lordPlacement == 6 {
		bhavaLordEffect = `The children of the native will be inimical or may adopt children or lose own children.`
	}
	if lordPlacement == 7 {
		bhavaLordEffect = `The native will be honorouable, very religious, endowed with progenic happiness and be helpful to others.
		The native will serve his employer honestly.`
	}
	if lordPlacement == 8 {
		bhavaLordEffect = `The native will not have much progenic happiness, be troubled by cough and be short tempered`
	}
	if lordPlacement == 9 {
		bhavaLordEffect = `The native will author books and be famous and will shine in his race like Adi Shankaracharya.`
	}
	if lordPlacement == 10 {
		bhavaLordEffect = `The native will be famous and enjoy various pleasures`
	}
	if lordPlacement == 11 {
		bhavaLordEffect = `The native will be learned, dear to people, be skillful and endowed with many sons and wealth.`
	}
	if lordPlacement == 12 {
		bhavaLordEffect = `The native will be bereft of happiness from his own children and will have issues with adopting (if going for adoption)`
	}

	return bhavaLordEffect
}