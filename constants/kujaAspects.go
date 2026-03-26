package constants

import "github.com/scientist-v08/bphs/models"

var KujaAspectsStore = models.SpecialAspectsMap{
	// Mesha
	Mesha: {
		{House: Tula},      // 7th
		{House: Karkataka}, // 4th
		{House: Vruschika}, // 8th
	},
	// Vrushabha
	Vrushabha: {
		{House: Vruschika}, // 7th
		{House: Simha},     // 4th
		{House: Dhanassu},  // 8th
	},
	// Mithuna
	Mithuna: {
		{House: Dhanassu},  // 7th
		{House: Kanya},     // 4th
		{House: Makara},    // 8th
	},
	// Karkataka
	Karkataka: {
		{House: Makara},    // 7th
		{House: Tula},      // 4th
		{House: Kumbha},    // 8th
	},
	// Simha
	Simha: {
		{House: Kumbha},    // 7th
		{House: Vruschika}, // 4th
		{House: Meena},     // 8th
	},
	// Kanya
	Kanya: {
		{House: Meena},     // 7th
		{House: Dhanassu},  // 4th
		{House: Mesha},     // 8th
	},
	// Tula
	Tula: {
		{House: Mesha},     // 7th
		{House: Makara},    // 4th
		{House: Vrushabha}, // 8th
	},
	// Vruschika
	Vruschika: {
		{House: Vrushabha}, // 7th
		{House: Kumbha},    // 4th
		{House: Mithuna},   // 8th
	},
	// Dhanassu
	Dhanassu: {
		{House: Mithuna},   // 7th
		{House: Meena},     // 4th
		{House: Karkataka}, // 8th
	},
	// Makara
	Makara: {
		{House: Karkataka}, // 7th
		{House: Mesha},     // 4th
		{House: Simha},     // 8th
	},
	// Kumbha
	Kumbha: {
		{House: Simha},     // 7th
		{House: Vrushabha}, // 4th
		{House: Kanya},     // 8th
	},
	// Meena
	Meena: {
		{House: Kanya},     // 7th
		{House: Mithuna},   // 4th
		{House: Tula},      // 8th
	},
}