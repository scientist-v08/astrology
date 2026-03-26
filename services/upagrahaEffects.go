package services

/*
1. Dhuma own house makara exalted simha debilitated kumbha
2. Vyatipata Mithuna Vruschika Vrushabha
3. Parivesha Dhanu mithuna dhanu
4. Indrachapa karkataka dhanu mithuna
5. Upaketu Karkataka Kumbha simha
6. Gulika Kumbha
7. Yamaghantaka Dhanu
8. Ardhaprahara Mithuna
9. Mrityu Vruschika
10. Kala Makara
*/

func DhumaEffectsService(dhumaPlacement int16) []string {
	var dhumaEffects []string

	dhumaEffects = append(dhumaEffects, "The effects of dhuma will mature in the dasha periods of shani and he is exalted in Simha and Neech in Kumbha")

	switch dhumaPlacement {
	case 1:
		dhumaEffects = append(dhumaEffects, "The native will be valiant, endowed with beautiful eyes, unkind, wicked and highly short tempered")
	case 2:
		dhumaEffects = append(dhumaEffects, "The native will be sickly, wealthy, may have injuries on limbs or lose a limb, be humiliated, be dull witted and be a eunuch")
	case 3:
		dhumaEffects = append(dhumaEffects, "The native will be intelligent, very bold, delighted, eloquent and endowed with men and wealth")
	case 4:
		dhumaEffects = append(dhumaEffects, "The native will be devoid of happiness as a direct consequence of a female")
	case 5:
		dhumaEffects = append(dhumaEffects, "The native will have few children, be poor and be great")
	case 6:
		dhumaEffects = append(dhumaEffects, "The native will destroy enemies, be powerful, victorious, courageous and capable of overcoming obstacles")
	case 7:
		dhumaEffects = append(dhumaEffects, "The native will be horny and poor and fond of others' spouses and also skillful in pursuing others' wives")
	case 8:
		dhumaEffects = append(dhumaEffects, "The native will be bereft of courage although enthusiastic, be truthful and selfish")
	case 9:
		dhumaEffects = append(dhumaEffects, "The native will be endowed with sons and fortunes, be rich, honourable, kind and religious")
	case 10:
		dhumaEffects = append(dhumaEffects, "The native will be endowed with sons and fortunes, be delighted, intelligent, happy and truthful")
	case 11:
		dhumaEffects = append(dhumaEffects, "The native will gain wealth, be happy, be skillful in arts and be handsome or beautiful")
	case 12:
		dhumaEffects = append(dhumaEffects, "The native will be sinful, unkind and crafty")
	}

	return dhumaEffects
}

func VyatipatiEffectsService(vyatipatiPlacement int16) []string {
	var vyatipatiEffects []string

	vyatipatiEffects = append(vyatipatiEffects, "The effects of vyatipati will mature in the dasha periods of budha and he is exalted in Vruschika and Neech in Vrushabha")

	switch vyatipatiPlacement {
	case 1:
		vyatipatiEffects = append(vyatipatiEffects, "The native will be troubled by miseries, be cruel, will indulge in destructive acts, be foolish")
	case 2:
		vyatipatiEffects = append(vyatipatiEffects, "The native will be morally crooked, be bilious, enjoy pleasures, be unkind, be wicked and sinful")
	case 3:
		vyatipatiEffects = append(vyatipatiEffects, "The native will be a warrior, be liberal, very rich and head of an army")
	case 4:
		vyatipatiEffects = append(vyatipatiEffects, "The native will be endowed with relatives but not sons and fortune")
	case 5:
		vyatipatiEffects = append(vyatipatiEffects, "The native will be poor, charming in appearance, be hard hearted and shameless")
	case 6:
		vyatipatiEffects = append(vyatipatiEffects, "The native will destroy enemies, be powerful, victorious, courageous and capable of overcoming obstacles")
	case 7:
		vyatipatiEffects = append(vyatipatiEffects, "The native will be bereft of wealth, wife and sons and be horny")
	case 8:
		vyatipatiEffects = append(vyatipatiEffects, "The native will be unfortunate and be spiteful towards Brahmins")
	case 9:
		vyatipatiEffects = append(vyatipatiEffects, "The native will be learned, have many kinds of business")
	case 10:
		vyatipatiEffects = append(vyatipatiEffects, "The native will be rich, religious, peaceful, skillful in religious acts and far sighted")
	case 11:
		vyatipatiEffects = append(vyatipatiEffects, "The native will be opulent, be honourable, truthful, firm in policy and interested in singing")
	case 12:
		vyatipatiEffects = append(vyatipatiEffects, "The native will be given to anger, associated with many activities, disabled and irreligious")
	}

	return vyatipatiEffects
}

func ParidhiEffectsService(paridhiPlacement int16) []string {
	var paridhiEffects []string

	paridhiEffects = append(paridhiEffects, "The effects of paridhi will mature in the dasha periods of Guru and he is exalted in Mithuna and Neech in Dhanu")

	switch paridhiPlacement {
	case 1:
		paridhiEffects = append(paridhiEffects, "The native will be learned, truthful, peaceful, rich, endowed with sons, pure and charitable")
	case 2:
		paridhiEffects = append(paridhiEffects, "The native will be wealthy, charming, enjoy pleasures, be happy, very religious")
	case 3:
		paridhiEffects = append(paridhiEffects, "The native will be fond of his wife, be charming, and respectful to his elders")
	case 4:
		paridhiEffects = append(paridhiEffects, "The native will be kind, endowed with everything and be skillful in singing")
	case 5:
		paridhiEffects = append(paridhiEffects, "The native will be affluent, virtuous, affectionate, religious and dear to his wife")
	case 6:
		paridhiEffects = append(paridhiEffects, "The native will be wealthy, be endowed with sons and pleasures, be helpful to all and conquer his enemies")
	case 7:
		paridhiEffects = append(paridhiEffects, "The native will have limited children, be devoid of happiness and have a sickly wife")
	case 8:
		paridhiEffects = append(paridhiEffects, "The native will be spiritually disposed, peaceful, strong-bodied, firm in decision and religious")
	case 9:
		paridhiEffects = append(paridhiEffects, "The native will be endowed with sons, be happy, brilliant, very affluent, be honourable and be happy")
	case 10:
		paridhiEffects = append(paridhiEffects, "The native will be versed in arts, will enjoy pleasures, be strong-bodied and learned in all shastras")
	case 11:
		paridhiEffects = append(paridhiEffects, "The native will enjoy pleasures through women, be virtuous, intelligent and dear to his people")
	case 12:
		paridhiEffects = append(paridhiEffects, "The native will always be spendthrift, be miserable, firm and will dishonour elders")
	}

	return paridhiEffects
}

func IndraChapaEffectsService(indraChapaPlacement int16) []string {
	var indraChapaEffects []string

	indraChapaEffects = append(indraChapaEffects, "The effects of indrachapa will mature in the dasha periods of Chandra and he is exalted in Dhanu and Neech in Mithuna")

	switch indraChapaPlacement {
	case 1:
		indraChapaEffects = append(indraChapaEffects, "The native will be endowed with wealth, grains, gold and be grateful, and agreeable")
	case 2:
		indraChapaEffects = append(indraChapaEffects, "The native will speak well, be rich, leanred, charming and religious")
	case 3:
		indraChapaEffects = append(indraChapaEffects, "The native will be a miser, be versed in many arts, will indulge in thieving, be unfriendly")
	case 4:
		indraChapaEffects = append(indraChapaEffects, "The native will be happy, endowed with quadrupeds, wealth, grains etc and be honoured by the King")
	case 5:
		indraChapaEffects = append(indraChapaEffects, "The native will be far sighted, pious, affable and will acquire prosperity")
	case 6:
		indraChapaEffects = append(indraChapaEffects, "The native will destroy his enemies, be happy, affectionate, and pure")
	case 7:
		indraChapaEffects = append(indraChapaEffects, "The native will be wealthy, endowed with all virtues, learned in shastras, religious and agreeable")
	case 8:
		indraChapaEffects = append(indraChapaEffects, "The native will be interested in others' job, be cruel and interested in others' wives")
	case 9:
		indraChapaEffects = append(indraChapaEffects, "The native will perform penance, will take to religious observations, and be highly learned")
	case 10:
		indraChapaEffects = append(indraChapaEffects, "The native will be endowed with many sons, abundant wealth, cows, buffaloes etc")
	case 11:
		indraChapaEffects = append(indraChapaEffects, "The native will be gainful, free from diseases, affectionate to his spouse and possess knowledge of mantras and weapons")
	case 12:
		indraChapaEffects = append(indraChapaEffects, "The native will be wicked, shameless and go to others' spouses")
	}

	return indraChapaEffects
}

func UpaketuEffectsService(upaketuPlacement int16) []string {
	var upaketuEffects []string

	upaketuEffects = append(upaketuEffects, "The effects of Upaketu will mature during the dashas of the moon and he is exalted in Kumbha and debilitated in Simha")

	switch upaketuPlacement {
	case 1:
		upaketuEffects = append(upaketuEffects, "The native will be skillful in all branches of learning, be happy, efficient in speech, agreeable and be very affectionate")
	case 2:
		upaketuEffects = append(upaketuEffects, "The native will be a good speaker, write poetry, be scholarly, honourable, modest and endowed with vehicles")
	case 3:
		upaketuEffects = append(upaketuEffects, "The native will be miserly, cruel in acts, thin bodied, and poor")
	case 4:
		upaketuEffects = append(upaketuEffects, "The native will be charming, very virtuous, gentle and interested in the Vedas")
	case 5:
		upaketuEffects = append(upaketuEffects, "The native will be happy, will enjoy pleasures, be versed in arts, skilled in expedients, intelligent, eloquent and respect elders")
	case 6:
		upaketuEffects = append(upaketuEffects, "The native will be ominous for maternal relatives, will win over his enemies, be endowed with many relatives, valiant and skillful")
	case 7:
		upaketuEffects = append(upaketuEffects, "The native will be interested in gambling, be honry, enjoy pleasures and be-friend prostitutes")
	case 8:
		upaketuEffects = append(upaketuEffects, "The native will be interested in base acts, be sinful, shameless and will blame others and take others side")
	case 9:
		upaketuEffects = append(upaketuEffects, "The native will be delighted, helpfully disposed to all and be skillful in religious deeds")
	case 10:
		upaketuEffects = append(upaketuEffects, "The native will be endowed with happiness and fortunes, be fond of females, and be-friend Brahmins")
	case 11:
		upaketuEffects = append(upaketuEffects, "The native will be religious, acquire gains, fortunate, valiant and skilled in sacrificial rites")
	case 12:
		upaketuEffects = append(upaketuEffects, "The native will be interested in sinful acts, be untrustworthy, interested in others' wives and be short tempered")
	}

	return upaketuEffects
}

func GulikaEffectsService(gulikaPlacement int16) []string {
	var gulikaEffects []string

	gulikaEffects = append(gulikaEffects, "The effects of gulika will mature in the dasha of saturn")

	switch gulikaPlacement {
	case 1:
		gulikaEffects = append(gulikaEffects, "The native will be afflicted by diseases, be useful, be lustful, sinful, crafty and very miserable")
	case 2:
		gulikaEffects = append(gulikaEffects, "The native will be unsightly in appearance, miserable, mean, given to vices, shameless and penniless")
	case 3:
		gulikaEffects = append(gulikaEffects, "The native will be charming in appearance, will obtain leadership positions in life, be fond of virtuous men and be honoured by the King")
	case 4:
		gulikaEffects = append(gulikaEffects, "The native will be sickly, devoid of happiness, sinful and afflicted due to windy diseases")
	case 5:
		gulikaEffects = append(gulikaEffects, "The native will not be praise worthy, be poor, short lived, spiteful, mean, be a eunuch, be subdued by his wife and be a heterodox")
	case 6:
		gulikaEffects = append(gulikaEffects, "The native will be devoid of enemies, be strong bodied, liked by the spouse and helpful in disposition")
	case 7:
		gulikaEffects = append(gulikaEffects, "The native will subdue to his spouse, be sinful and go to others' spouses and live on the wives' wealth")
	case 8:
		gulikaEffects = append(gulikaEffects, "The native will be miserable, poor, cruel, short-tempered and bereft of good qualities")
	case 9:
		gulikaEffects = append(gulikaEffects, "The native will undergo many ordeals, will perform evil acts and be a tale bearer")
	case 10:
		gulikaEffects = append(gulikaEffects, "The native will be endowed with sons, be happy, will enjoy many things, be fond of worshipping Gods and practise meditation")
	case 11:
		gulikaEffects = append(gulikaEffects, "The native will enjoy women of class, possess leadership qualities, be short in height and obtain leadership positions")
	case 12:
		gulikaEffects = append(gulikaEffects, "The native will indulge in base deeds, be sinful, defective-limbed, unfortunate, indolent and join mean people")
	}

	return gulikaEffects
}

func PranapadaEffectsService(pranapadaPlacement int16) []string {
	var pranapadaEffects []string

	pranapadaEffects = append(pranapadaEffects, "The effects of the Pranapada will mature during the dasha of the dispositor of whichever sign it is placed in.")

	switch pranapadaPlacement {
	case 1:
		pranapadaEffects = append(pranapadaEffects, "The native will be weak, sickly, dumb, lunatic and miserable")
	case 2:
		pranapadaEffects = append(pranapadaEffects, "The native will be endowed with abundant grains, abundant wealth, children and fortune")
	case 3:
		pranapadaEffects = append(pranapadaEffects, "The native will be mischievous, proud, hard-hearted, very dirty and disrespect elders")
	case 4:
		pranapadaEffects = append(pranapadaEffects, "The native will be happy, friendly, attached to females and elders, soft and truthful")
	case 5:
		pranapadaEffects = append(pranapadaEffects, "The native will be happy, friendly, attached to females and elders and truthful")
	case 6:
		pranapadaEffects = append(pranapadaEffects, "The native will be sub-dued by his relatives and enemies, be sharp, be wicked and affluent")
	case 7:
		pranapadaEffects = append(pranapadaEffects, "The native will be green-eyed, horny, not worthy of respect and be ill-disposed")
	case 8:
		pranapadaEffects = append(pranapadaEffects, "The native will be afflicted by diseases, be troubled and will incur misery on account of those in leadership positions")
	case 9:
		pranapadaEffects = append(pranapadaEffects, "The native will be endowed with sons, be very rich, fortunate, charming, will serve others and be skillful")
	case 10:
		pranapadaEffects = append(pranapadaEffects, "The native will be heroic, intelligent, will follow the orders given by those in leadership positions and worship gods")
	case 11:
		pranapadaEffects = append(pranapadaEffects, "The native will be virtuous, learned, wealthy, fair-complexioned and attached to mother")
	case 12:
		pranapadaEffects = append(pranapadaEffects, "The native will be mean, wicked and suffer from eye diseases")
	}

	return pranapadaEffects
}