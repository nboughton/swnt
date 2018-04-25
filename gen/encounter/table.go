package encounter

import (
	"github.com/nboughton/rollt"
	"github.com/nboughton/swnt/gen/table"
)

// Urban represents the OneRoll tables for rolling quick encounters
var Urban = table.OneRoll{
	D4: rollt.List{
		Name: "What's the Conflict About?",
		Items: []string{
			"Money, extortion, payment due, debts",
			"Respect, submission to social authority",
			"Grudges, ethnic resentment, gang payback",
			"Politics, religion, or other ideology",
		},
	},
	D6: rollt.List{
		Name: "General Venue of the Event",
		Items: []string{
			"In the middle of the street",
			"In a public plaza",
			"Down a side alley",
			"Inside a local business",
			"Next to or in a public park",
			"At a mass-transit station",
		},
	},
	D8: rollt.List{
		Name: "Why are the PCs Involved?",
		Items: []string{
			"A sympathetic participant appeals to them",
			"Ways around it are all dangerous/blocked",
			"It happens immediately around them",
			"A valuable thing looks snatchable amid it",
			"A participant offers a reward for help",
			"Someone mistakenly involves the PCs in it",
			"The seeming way out just leads deeper in",
			"Responsibility is somehow pinned on them",
		},
	},
	D10: rollt.List{
		Name: "What's the Nature of the Event?",
		Items: []string{
			"A parade or festival is being disrupted",
			"Innocents are being assaulted",
			"An establishment is being robbed",
			"A disturbance over local politics happens",
			"Someone is being blamed for something",
			"Fires or building collapses are happening",
			"A medical emergency is happening",
			"Someone’s trying to cheat the PCs",
			"A vehicle accident is happening",
			"A religious ceremony is being disrupted",
		},
	},
	D12: rollt.List{
		Name: "What Antagonists are Involved?",
		Items: []string{
			"A local bully and their thugs",
			"A ruthless political boss and their zealots",
			"Violent criminals",
			"Religious fanatics",
			"A blisteringly obnoxious offworlder",
			"Corrupt or over-strict government official",
			"A mob of intoxicated locals",
			"A ranting demagogue and their followers",
			"A stupidly bull-headed local grandee",
			"A very capable assassin or strong-arm",
			"A self-centered local scion of power",
			"A confused foreigner or backwoodsman",
		},
	},
	D20: rollt.List{
		Name: "Relevant Urban Features",
		Items: []string{
			"Heavy traffic running through the place",
			"Music blaring at deafening volumes",
			"Two groups present that detest each other",
			"Large delivery taking place right there",
			"Swarm of schoolkids or feral youth",
			"Insistent soapbox preacher here",
			"Several pickpockets working the crowd",
			"A kiosk is tipping over and spilling things",
			"Streetlights are out or visibility is low",
			"A cop patrol is here and reluctant to act",
			"PC-hostile reporters are recording here",
			"Someone’s trying to sell something to PCs",
			"Feral dogs or other animals crowd here",
			"Unrelated activists are protesting here",
			"Street kids are trying to steal from the PCs",
			"GPS maps are dangerously wrong here",
			"Downed power lines are a danger here",
			"Numerous open manholes and utility holes",
			"The street’s blockaded by something",
			"Crowds so thick one can barely move",
		},
	},
}

// Wilderness represents the OneRoll tables for generating Wilderness Encounters
var Wilderness = table.OneRoll{
	D4: rollt.List{
		Name: "Initial Encounter Range",
		Items: []string{
			"Visible from a long distance away",
			"Noticed 1d4 hundred meters away",
			"Noticed only within 1d6 x 10 meters",
			"Noticed only when adjacent to the event",
		},
	},
	D6: rollt.List{
		Name: "Weather and Lighting",
		Items: []string{
			"Takes place in daylight and clear weather",
			"Daylight, but fog, mist, rain or the like",
			"Daylight, but harsh seasonal weather",
			"Night encounter, but clear weather",
			"Night, with rain or other obscuring effects",
			"Night, with terrible weather and wind",
		},
	},
	D8: rollt.List{
		Name: "Basic Nature of the Encounter",
		Items: []string{
			"Attack by pack of hostiles",
			"Ambush by single lone hostile",
			"Meet people who don’t want to be met",
			"Encounter people in need of aid",
			"Encounter hostile creatures",
			"Nearby feature is somehow dangerous",
			"Nearby feature promises useful loot",
			"Meet hostiles that aren’t immediately so",
		},
	},
	D10: rollt.List{
		Name: "Types of Friendly Creatures",
		Items: []string{
			"Affable but reclusive hermit",
			"Local herd animal let loose to graze",
			"Government ranger or circuit judge",
			"Curious local animal",
			"Remote homesteader and family",
			"Working trapper or hunter",
			"Back-country villager or native",
			"Hiker or wilderness tourist",
			"Religious recluse or holy person",
			"Impoverished social exile",
		},
	},
	D12: rollt.List{
		Name: "Types of Hostile Creatures",
		Items: []string{
			"Bandits in their wilderness hideout",
			"Dangerous locals looking for easy marks",
			"Rabid or diseased large predator",
			"Pack of hungry hunting beasts",
			"Herd of potentially dangerous prey animals",
			"Swarm of dangerous vermin",
			"Criminal seeking to evade the law",
			"Brutal local landowner and their men",
			"Crazed hermit seeking enforced solitude",
			"Friendly-seeming guide into lethal danger",
			"Harmless-looking but dangerous beast",
			"Confidence man seeking to gull the PCs",
		},
	},
	D20: rollt.List{
		Name: "Specific Nearby Feature of Relevance",
		Items: []string{
			"Overgrown homestead",
			"Stream prone to flash-flooding",
			"Narrow bridge or beam over deep cleft",
			"Box canyon with steep sides",
			"Unstable hillside that slides if disturbed",
			"Long-lost crash site of a gravflyer",
			"Once-inhabited cave or tunnel",
			"Steep and dangerous cliff",
			"Quicksand-laden swamp or dust pit",
			"Ruins of a ghost town or lost hamlet",
			"Hunting cabin with necessities",
			"Ill-tended graveyard of a lost family stead",
			"Narrow pass that’s easily blocked",
			"Dilapidated resort building",
			"Remote government monitoring outpost",
			"Illicit substance farm or processing center",
			"Old and forgotten battleground",
			"Zone overrun by dangerous plants",
			"Thick growth that lights up at a spark",
			"Abandoned vehicle",
		},
	},
}
