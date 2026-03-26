package constants

import "github.com/scientist-v08/bphs/models"

var BrihaspatiAspectsStore = models.SpecialAspectsMap{
	// Mesha
	Mesha: {
		{House: Tula},     // 7th
		{House: Simha},    // 5th
		{House: Dhanassu}, // 9th
	},
	// Vrushabha
	Vrushabha: {
		{House: Vruschika}, // 7th
		{House: Kanya},     // 5th
		{House: Makara},    // 9th
	},
	// Mithuna
	Mithuna: {
		{House: Dhanassu}, // 7th
		{House: Tula},     // 5th
		{House: Kumbha},   // 9th
	},
	// Karkataka
	Karkataka: {
		{House: Makara},   // 7th
		{House: Vruschika}, // 5th
		{House: Meena},    // 9th
	},
	// Simha
	Simha: {
		{House: Kumbha},   // 7th
		{House: Dhanassu}, // 5th
		{House: Mesha},    // 9th
	},
	// Kanya
	Kanya: {
		{House: Meena},    // 7th
		{House: Makara},   // 5th
		{House: Vrushabha}, // 9th
	},
	// Tula
	Tula: {
		{House: Mesha},    // 7th
		{House: Kumbha},   // 5th
		{House: Mithuna},  // 9th
	},
	// Vruschika
	Vruschika: {
		{House: Vrushabha}, // 7th
		{House: Meena},     // 5th
		{House: Karkataka}, // 9th
	},
	// Dhanassu
	Dhanassu: {
		{House: Mithuna},  // 7th
		{House: Mesha},    // 5th
		{House: Simha},    // 9th
	},
	// Makara
	Makara: {
		{House: Karkataka}, // 7th
		{House: Vrushabha}, // 5th
		{House: Tula},      // 9th
	},
	// Kumbha
	Kumbha: {
		{House: Simha},    // 7th
		{House: Mithuna},  // 5th
		{House: Tula}, // 9th
	},
	// Meena
	Meena: {
		{House: Kanya},     // 7th
		{House: Karkataka}, // 5th
		{House: Vruschika}, // 9th
	},
}