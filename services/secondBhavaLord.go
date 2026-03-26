package services

func SecondBhavaLordEffect(lordPlacement int8) string {
	var secondBhavaLordEffect = ""

	// Add to the string based on the lord placement
	if lordPlacement == 1 {
		secondBhavaLordEffect = `The native will obtain sons and wealth, be inimical to his own family, lustful and will do others job and be a fraud.
		The native would face financial upheavals and this may not apply completely to a capricorn ascendant.`
	}
	if lordPlacement == 2 {
		secondBhavaLordEffect = `The native will be wealthy and proud and have many spouses and be bereft of progeny.`
	}
	if lordPlacement == 3 {
		secondBhavaLordEffect = `The native will be valorous, wise, lustful and virtuous. All of these will happen if the lord is related to a benefic.
		If related to a malefic the native will be a heterodox, will not fear God and have dirty conduct. The native may be a pimp`
	}
	if lordPlacement == 4 {
		secondBhavaLordEffect = `The native will acquire all kinds of wealth be a heterodox and of questionable character.`
	}
	if lordPlacement == 5 {
		secondBhavaLordEffect = `The native will be wealthy and his sons will also earn money on their own. Resort to trickery and not have a happy family life.
		Will not be kind to others and be lustful and may be prone to losing a child prematurely`
	}
	if lordPlacement == 6 {
		secondBhavaLordEffect = `If the 2nd lord is with a benefic the native will gain through his enemies and lose if with malefics along with getting hit by them.
		Loss of wealth through theives and servants stealing wealth.`
	}
	if lordPlacement == 7 {
		secondBhavaLordEffect = `The native will be addicted to others wives and be a doctor. If a malefic is associated with the second lord,
		then even the wife may have questionable character.`
	}
	if lordPlacement == 8 {
		secondBhavaLordEffect = `The native will have land and wealth but limited marital felicity and be bereft of happiness from co-borns`
	}
	if lordPlacement == 9 {
		secondBhavaLordEffect = `The native will be wealthy, diligent, skillful, visit shrines and observing religious code`
	}
	if lordPlacement == 10 {
		secondBhavaLordEffect = `The native will be horny, honourable, learned, will have many wives and much wealth`
	}
	if lordPlacement == 11 {
		secondBhavaLordEffect = `The native will have all kinds of wealth, be ever diligent, honourable and famous`
	}
	if lordPlacement == 12 {
		secondBhavaLordEffect = `The native will be adventurous, be devoid of wealth and interested in others wealth and the first born will make the native sad.
		Except Aries ascendant or if the second lord is with 2 or more benefics then the native instead gain from this configuration.`
	}

	return secondBhavaLordEffect
}