package services

func TwelfthBhavaLordEffect(lordPlacement int8) string {
	var bhavaLordEffect = ""

	// Add to the string based on the lord placement
	if lordPlacement == 1 {
		bhavaLordEffect = `The native will be a spendthrift, be weak in constitution, will suffer from diseases and be devoid of wealth and learning`
	}
	if lordPlacement == 2 {
		bhavaLordEffect = `The native will always spend on auspicious deeds be religious, will speak sweetly and be endowed with virtues and happiness`
	}
	if lordPlacement == 3 {
		bhavaLordEffect = `The native will be devoid of fraternal bliss, will hate others and be selfish`
	}
	if lordPlacement == 4 {
		bhavaLordEffect = `The native will be devoid of maternal happiness and will gradually accure losses in respect of lands, vehicles and houses`
	}
	if lordPlacement == 5 {
		bhavaLordEffect = `The native will be bereft of sons and learning`
	}
	if lordPlacement == 6 {
		bhavaLordEffect = `The native will incur enemit with his own men, be given to anger, be sinful, miserable and will go to others spouse`
	}
	if lordPlacement == 7 {
		bhavaLordEffect = `The native will incur expenses on account of his wife, will not enjoy conjugal bliss`
	}
	if lordPlacement == 8 {
		bhavaLordEffect = `The native will always gain, will speak affably and be endowed with good qualities`
	}
	if lordPlacement == 9 {
		bhavaLordEffect = `The native will dishonour his elders, be inimical even to his friends and be always intent on achieving his own ends`
	}
	if lordPlacement == 10 {
		bhavaLordEffect = `The native will incur expenses through royal persons and will enjoy moderate paternal bliss`
	}
	if lordPlacement == 11 {
		bhavaLordEffect = `The native will incur losses, be bought up by others and will sometimes gain through others`
	}
	if lordPlacement == 12 {
		bhavaLordEffect = `The native will only face heavy expenditures, will not have physical felicity, be iritable and spiteful`
	}

	return bhavaLordEffect
}