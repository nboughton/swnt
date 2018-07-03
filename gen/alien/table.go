package alien

import (
	"github.com/nboughton/rollt"
)

// Body SWN Revised Free Edition p203
var Body = rollt.Table{
	Name: "Body",
	Dice: "1d6",
	Items: []rollt.Item{
		{Match: []int{1}, Text: "Avian, bat-like, pterodactylian"},
		{Match: []int{2}, Text: "Reptilian, amphibian, draconic"},
		{Match: []int{3}, Text: "Insectile, beetle-like, spiderish, wasp-like"},
		{Match: []int{4}, Text: "Mammalian, furred or bare-skinned"},
		{Match: []int{5}, Text: "Exotic, composed of some novel substance"},
		{Match: []int{6}, Text: "Hybrid of two or more types"},
	},
}

// Lense from SWN Revised Free Edition p205
var Lense = rollt.Table{
	Name: "Lense",
	Dice: "1d20",
	Items: []rollt.Item{
		{Match: []int{1}, Text: "Collectivity"},
		{Match: []int{2}, Text: "Curiosity"},
		{Match: []int{3}, Text: "Despair"},
		{Match: []int{4}, Text: "Dominion"},
		{Match: []int{5}, Text: "Faith"},
		{Match: []int{6}, Text: "Fear"},
		{Match: []int{7}, Text: "Gluttony"},
		{Match: []int{8}, Text: "Greed"},
		{Match: []int{9}, Text: "Hate"},
		{Match: []int{10}, Text: "Honor"},
		{Match: []int{11}, Text: "Journeying"},
		{Match: []int{12}, Text: "Joy"},
		{Match: []int{13}, Text: "Pacifism"},
		{Match: []int{14}, Text: "Pride"},
		{Match: []int{15}, Text: "Sagacity"},
		{Match: []int{16}, Text: "Subtlety"},
		{Match: []int{17}, Text: "Tradition"},
		{Match: []int{18}, Text: "Treachery"},
		{Match: []int{19}, Text: "Tribalism"},
		{Match: []int{20}, Text: "Wrath"},
	},
}

// SocialStructure from SWN Revised Free Edition p207
var SocialStructure = rollt.Table{
	Name: "Social Structure",
	Dice: "1d8",
	Items: []rollt.Item{
		{Match: []int{1}, Text: "Democratic"},
		{Match: []int{2}, Text: "Monarchic"},
		{Match: []int{3}, Text: "Tribal"},
		{Match: []int{4}, Text: "Oligarchic"},
		{Match: []int{5, 6}, Text: "Multipolar Competitive"}, // TODO: Roll to expand
		{Match: []int{7, 8}, Text: "Multipolar Cooperative"}, // TODO: Roll to expand
	},
}
