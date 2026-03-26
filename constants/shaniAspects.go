package constants

import "github.com/scientist-v08/bphs/models"

var ShaniAspectsStore = models.SpecialAspectsMap{
		Mesha: {
			{House: Tula}, // 7th
			{House: Mithuna}, // 3rd
			{House: Makara}, // 10th
		},
		Vrushabha: {
			{House: Vruschika}, // 7th
			{House: Karkataka},      // 3rd
			{House: Kumbha},     // 10th
		},
		Mithuna: {
			{House: Dhanassu},    // 7th
			{House: Simha},    // 3rd
			{House: Meena},    // 10th
		},
		Karkataka: {
			{House: Makara},   // 7th
			{House: Kanya},    // 3rd
			{House: Mesha},    // 10th
		},
		Simha: {
			{House: Kumbha},   // 7th
			{House: Tula},     // 3rd
			{House: Vrushabha}, // 10th
		},
		Kanya: {
			{House: Meena},    // 7th
			{House: Vruschika}, // 3rd
			{House: Mithuna},  // 10th
		},
		Tula: {
			{House: Mesha},    // 7th
			{House: Dhanassu},    // 3rd
			{House: Karkataka},    // 10th
		},
		Vruschika: {
			{House: Vrushabha}, // 7th
			{House: Makara},   // 3rd
			{House: Simha},    // 10th
		},
		Dhanassu: {
			{House: Mithuna},  // 7th
			{House: Kumbha},   // 3rd
			{House: Kanya},    // 10th
		},
		Makara: {
			{House: Karkataka},    // 7th
			{House: Meena},    // 3rd
			{House: Tula},     // 10th
		},
		Kumbha: {
			{House: Simha},    // 7th
			{House: Mesha},    // 3rd
			{House: Vruschika}, // 10th
		},
		Meena: {
			{House: Kanya},    // 7th
			{House: Vrushabha}, // 3rd
			{House: Dhanassu},    // 10th
		},
}