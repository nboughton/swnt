package religion

import (
	"github.com/nboughton/num"
	"github.com/nboughton/rollt"
)

// Evolution SWN Revised Free Edition p193
var Evolution = rollt.List{
	Name: "Evolution",
	Items: []string{
		"New holy book. Someone in the faith’s past penned or discovered a text that is now taken to be holy writ and the expressed will of the divine.",
		"New prophet. This faith reveres the words and example of a relatively recent prophet, esteeming him or her as the final word on the will of God. The prophet may or may not still be living.",
		"Syncretism. The faith has merged many of its beliefs with another religion. Roll again on the origin tradition table; this faith has reconciled the major elements of both beliefs into its tradition.",
		"Neofundamentalism. The faith is fiercely resistant to perceived innovations and deviations from their beliefs. Even extremely onerous traditions and restrictions will be observed to the letter.",
		"Quietism. The faith shuns the outside world and involvement with the affairs of nonbelievers. They prefer to keep to their own kind and avoid positions of wealth and power.",
		"Sacrifices. The faith finds it necessary to make substantial sacrifices to please God. Some faiths may go so far as to offer human sacrifices, while others insist on huge tithes offered to the building of religious edifices.",
		"Schism. The faith’s beliefs are actually almost identical to those of the majority of its origin tradition, save for a few minor points of vital interest to theologians and no practical difference whatsoever to believers. This does not prevent a burning resentment towards the parent faith.",
		"Holy family. God’s favor has been shown especially to members of a particular lineage. It may be that only men and women of this bloodline are permitted to become clergy, or they may serve as helpless figureheads for the real leaders of the faith",
	},
}

// OriginTradition SWN Revised Free Edition p193
var OriginTradition = rollt.List{
	Name: "Evolution",
	Items: []string{
		"Paganism",
		"Roman Catholicism",
		"Eastern Orthodox Christianity",
		"Protestant Christianity",
		"Buddhism",
		"Judaism",
		"Islam",
		"Taoism",
		"Hinduism",
		"Zoroastrianism",
		"Confucianism",
		"Ideology",
	},
}

// Leadership SWN Revised Free Edition p193
var Leadership = rollt.Table{
	Name: "Leadership",
	Dice: "1d6",
	Items: []rollt.Item{
		{Match: num.Set{1, 2}, Text: "Patriarch/Matriarch. A single leader determines doctrine for the entire religion, possibly in consultation with other clerics."},
		{Match: num.Set{3, 4}, Text: "Council. A group of the oldest and most revered clergy determine the course of the faith."},
		{Match: num.Set{5}, Text: "Democracy. Every member has an equal voice in matters of faith, with doctrine usually decided at regular church- wide councils."},
		{Match: num.Set{6}, Text: "No universal leadership. Roll again to determine how each region governs itself. If another 6 is rolled, this faith has no hierarchy."},
	},
}
