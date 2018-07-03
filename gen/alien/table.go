package alien

import (
	"github.com/nboughton/num"
	"github.com/nboughton/rollt"
)

// Body SWN Revised Free Edition p203
var Body = rollt.Table{
	Name: "Body",
	Dice: "1d6",
	Items: []rollt.Item{
		{Match: num.Set{1}, Text: "Avian, bat-like, pterodactylian"},
		{Match: num.Set{2}, Text: "Reptilian, amphibian, draconic"},
		{Match: num.Set{3}, Text: "Insectile, beetle-like, spiderish, wasp-like"},
		{Match: num.Set{4}, Text: "Mammalian, furred or bare-skinned"},
		{Match: num.Set{5}, Text: "Exotic, composed of some novel substance"},
		{Match: num.Set{6}, Text: "Hybrid of two or more types"},
	},
}

// Lense from SWN Revised Free Edition p205
var Lense = rollt.Table{
	Name: "Lense",
	Dice: "1d20",
	Items: []rollt.Item{
		{Match: num.Set{1}, Text: "Collectivity"},
		{Match: num.Set{2}, Text: "Curiosity"},
		{Match: num.Set{3}, Text: "Despair"},
		{Match: num.Set{4}, Text: "Dominion"},
		{Match: num.Set{5}, Text: "Faith"},
		{Match: num.Set{6}, Text: "Fear"},
		{Match: num.Set{7}, Text: "Gluttony"},
		{Match: num.Set{8}, Text: "Greed"},
		{Match: num.Set{9}, Text: "Hate"},
		{Match: num.Set{10}, Text: "Honor"},
		{Match: num.Set{11}, Text: "Journeying"},
		{Match: num.Set{12}, Text: "Joy"},
		{Match: num.Set{13}, Text: "Pacifism"},
		{Match: num.Set{14}, Text: "Pride"},
		{Match: num.Set{15}, Text: "Sagacity"},
		{Match: num.Set{16}, Text: "Subtlety"},
		{Match: num.Set{17}, Text: "Tradition"},
		{Match: num.Set{18}, Text: "Treachery"},
		{Match: num.Set{19}, Text: "Tribalism"},
		{Match: num.Set{20}, Text: "Wrath"},
	},
}

// SocialStructure from SWN Revised Free Edition p207
var SocialStructure = rollt.Table{
	Name: "Social Structure",
	Dice: "1d8",
	Items: []rollt.Item{
		{Match: num.Set{1}, Text: "Democratic"},
		{Match: num.Set{2}, Text: "Monarchic"},
		{Match: num.Set{3}, Text: "Tribal"},
		{Match: num.Set{4}, Text: "Oligarchic"},
		{Match: num.Set{5, 6}, Text: "Multipolar Competitive"}, // TODO: Roll to expand
		{Match: num.Set{7, 8}, Text: "Multipolar Cooperative"}, // TODO: Roll to expand
	},
}
