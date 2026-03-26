package services

func TenthBhavaLordEffect(lordPlacement int8) string {
	var bhavaLordEffect = ""

	// Add to the string based on the lord placement
	if lordPlacement == 1 {
		bhavaLordEffect = `The native will be scholarly, famous, be a poet and increase his wealth slowly`
	}
	if lordPlacement == 2 {
		bhavaLordEffect = `The native will be wealthy, virtuous, honoroued by those in leadership positions, be charitable and enjoy happiness from father`
	}
	if lordPlacement == 3 {
		bhavaLordEffect = `The native will enjoy happiness from co-born and servants and be valorous, virtuous, eloquent and beautiful`
	}
	if lordPlacement == 4 {
		bhavaLordEffect = `The native will happy, interested in mothers welfare, possess vehicles, lands and houses be virtous and wealthy`
	}
	if lordPlacement == 5 {
		bhavaLordEffect = `The native will be endowed with all kinds of learning, be always delighted and be wealthy and endowed with sons.`
	}
	if lordPlacement == 6 {
		bhavaLordEffect = `The native will be bereft of paternal happiness, be skillful, be bereft of we4alth and troubled by enemies`
	}
	if lordPlacement == 7 {
		bhavaLordEffect = `The native will be endowed with happiness through wife, be intelligent, virtous, eloquent, truthful and religious`
	}
	if lordPlacement == 8 {
		bhavaLordEffect = `The native will be devoid of good acts, long lived and intent on blaming others`
	}
	if lordPlacement == 9 {
		bhavaLordEffect = `The native who be born of royal scion will obtain leadership positions and the ordinary individual will be equal to those in leadership positions.`
	}
	if lordPlacement == 10 {
		bhavaLordEffect = `The native will be skillful in all jobs, be valouruous, truthful and devoted to elders`
	}
	if lordPlacement == 11 {
		bhavaLordEffect = `The native will be endowed with wealth, happiness and sons.`
	}
	if lordPlacement == 12 {
		bhavaLordEffect = `The native will spend on royal abodes will have fear from enemies and will be worried in spite of being skillful`
	}

	return bhavaLordEffect
}