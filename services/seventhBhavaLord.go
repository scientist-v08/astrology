package services

func SeventhBhavaLordEffect(lordPlacement int8) string {
	var bhavaLordEffect = ""

	// Add to the string based on the lord placement
	if lordPlacement == 1 {
		bhavaLordEffect = `The native will go to others spouses, be wicked, skillful, devoid of courage and not be firm in words and actions.
		Will impart courage to others although not have courage with oneself.`
	}
	if lordPlacement == 2 {
		bhavaLordEffect = `The native will have many spouses and gain through his spouse and be procrastinatnig in nature.`
	}
	if lordPlacement == 3 {
		bhavaLordEffect = `The native may face loss of sons but all the daughters will live`
	}
	if lordPlacement == 4 {
		bhavaLordEffect = `The spouse of the native will not be under control but will be devoted to the native. Will be fond of the truth and be intelligent.`
	}
	if lordPlacement == 5 {
		bhavaLordEffect = `The native will be honorouable, endowed with seven principal virtues and endowed with all kinds of wealth.
		But there will be delays and disappointment in married life and the conjugal life seldom proves happy.
		Either there will be unhappiness on account of children or loss of children.`
	}
	if lordPlacement == 6 {
		bhavaLordEffect = `The native will beget a sickly spouse and be short tempred and devoid of happiness.`
	}
	if lordPlacement == 7 {
		bhavaLordEffect = `The native will be endowed with happiness through his spouse and the spouse will be courageous, skillful and intelligent.`
	}
	if lordPlacement == 8 {
		bhavaLordEffect = `The native will be deprived of marital happiness and will not obey the native.`
	}
	if lordPlacement == 9 {
		bhavaLordEffect = `The native will unite with other then own spouse but be well disposed to own spouse`
	}
	if lordPlacement == 10 {
		bhavaLordEffect = `The native will beget a disobedient spouse and will be religious and endowed with wealth and sons.`
	}
	if lordPlacement == 11 {
		bhavaLordEffect = `The native will gain wealth through the spouse and be unhappy from own sons (but gaining a son is very unlikely).`
	}
	if lordPlacement == 12 {
		bhavaLordEffect = `The native may become poor, be a miser and earn livelihood through selling clothes. The spouse maybe a spendthrift.`
	}

	return bhavaLordEffect
}