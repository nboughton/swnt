package place

import (
	"github.com/nboughton/rollt"
	"github.com/nboughton/swnt/gen/table"
)

// Reward roll reward
var Reward = rollt.List{
	Items: []string{
		"Large cache of credits",
		"Precious cultural artifact",
		"Vital data on the party’s goal",
		"Missing or kidnapped VIP",
		"Advanced pretech artifact",
		"Key to some guarded location",
		"Ancient treasure object",
		"Recently-stolen goods",
		"High-tech robotic servitor",
		"Token item of ruling legitimacy",
		"Juicy blackmail material",
		"History-rewriting evidence",
		"Alien artifact of great power",
		"Precious megacorp data files",
		"Map to some valuable thing",
		"Forbidden but precious drug",
		"Legal title to important land",
		"Awful secret of local government",
		"Cache of precious goods",
		"Stock of valuable weaponry",
	},
}

// Ongoings roll ongoings
var Ongoings = struct {
	Civilised  rollt.List
	Wilderness rollt.List
}{
	rollt.List{
		Items: []string{
			"Local festival going on",
			"Angry street protests",
			"Minor fire or other disorder",
			"Merchants and peddlers active",
			"Tourists from another country",
			"Building repair or maintenance",
			"Recent vehicle crash",
			"Public art performance",
			"Angry traffic jam",
			"Missionaries for a local religion",
			"Loud advertising campaign",
			"Memorial service ongoing",
			"Road work halting traffic",
			"Power outage in the area",
			"Police chasing criminals",
			"Annoying drunks being loud",
			"Beggars seeking alms",
			"Constructing a new building",
			"Local thugs swaggering around",
			"Aerial light display",
		},
	},
	rollt.List{
		Items: []string{
			"Bandits have moved in",
			"Flooding swept through",
			"Part of it has collapsed",
			"Refugees are hiding here",
			"Dangerous animals lair here",
			"A rebel cell uses it for a base",
			"Smugglers have landed here",
			"Foreign agents meet here",
			"A hermit has taken up residence",
			"A toxic plant is growing wild",
			"An artist seeks inspiration here",
			"An ancient structure was dug out",
			"The weather has turned savage",
			"A vehicle crashed nearby",
			"Some locals are badly lost",
			"Religious pilgrims come here",
			"Locals fight over control of it",
			"Nature threatens to wipe it out",
			"An old shrine was raised here",
			"A shell of a building remains",
		},
	},
}

// Hazard roll hazards
var Hazard = table.ThreePart{
	Headers: [3]string{"Hazard", "Specific Example", "Possible Danger"},
	Tables: []table.ThreePartSubTable{
		{
			Name: "Social",
			SubTable1: rollt.List{
				Items: []string{
					"An explosively temperamental VIP",
					"An unknown but critical social taboo",
					"A case of mistaken identity",
					"An expectation for specific PC action",
					"A frame job hung on the PCs",
				},
			},
			SubTable2: rollt.List{
				Items: []string{
					"An allied NPC breaks ties",
					"An enemy is alerted to them",
					"A new enemy is made",
					"Cads think the PCs are allies",
					"An opportunity is lost",
				},
			},
		},
		{
			Name: "Legal",
			SubTable1: rollt.List{
				Items: []string{
					"A regulation unknown to the PCs",
					"A tax or confiscation",
					"Vital gear is prohibited here",
					"Lawsuit from an aggrieved NPC",
					"A state agent conscripts PC help",
				},
			},
			SubTable2: rollt.List{
				Items: []string{
					"Substantial monetary fine",
					"Imprisonment for the party",
					"Confiscation of possessions",
					"Deportation from the place",
					"Loss of rights and protections",
				},
			},
		},
		{
			Name: "Environmental",
			SubTable1: rollt.List{
				Items: []string{
					"Heavy background radiation",
					"A planetary sickness foreigners get",
					"Strong or weak local gravity",
					"Gear-eating microbial life",
					"Unpredictable psychic power field",
				},
			},
			SubTable2: rollt.List{
				Items: []string{
					"Catch a lingering disease",
					"Suffer bodily harm",
					"Take a penalty on rolls",
					"Lose some equipment",
					"Psychic abilities are altered",
				},
			},
		},
		{
			Name: "Trap",
			SubTable1: rollt.List{
				Items: []string{
					"Alarm system attached to a trap",
					"Snare left for local animals",
					"Hermit’s self-defense measure",
					"Long-dead builder’s trapsmithing",
					"New occupant’s defensive trap",
				},
			},
			SubTable2: rollt.List{
				Items: []string{
					"Something set on fire",
					"Guards are summoned",
					"Fall to a new area",
					"Equipment is damaged",
					"Subject is injured",
				},
			},
		},
		{
			Name: "Animal",
			SubTable1: rollt.List{
				Items: []string{
					"Dangerous local swarm vermin",
					"A big predator lair",
					"Pack hunters haunt the area",
					"Flying threats pounce here",
					"Monstrous beast sleeps or is torpid",
				},
			},
			SubTable2: rollt.List{
				Items: []string{
					"They have a ranged attack",
					"They’re venomous",
					"Dangerously coordinated foe",
					"Killing them inflicts a fine",
					"Their deaths cause an effect",
				},
			},
		},
		{
			Name: "Sentient",
			SubTable1: rollt.List{
				Items: []string{
					"A group hostile to intruders",
					"Trickster thieves and con-men",
					"Hostile expert-system robots",
					"Secrecy-loving rebels or criminals",
					"Another area-clearing group",
				},
			},
			SubTable2: rollt.List{
				Items: []string{
					"Immediate combat",
					"Treacherous feigned friend",
					"Lead the PCs into a trap",
					"Demand payment or loot",
					"Activate other area defenses",
				},
			},
		},
		{
			Name: "Decay",
			SubTable1: rollt.List{
				Items: []string{
					"Crumbling floor or ceiling",
					"Waste or heating tubes rupture",
					"Dangerous standing liquid",
					"Maintenance robots gone haywire",
					"Power plant is unstable",
				},
			},
			SubTable2: rollt.List{
				Items: []string{
					"Ruptures to release a peril",
					"Toxic or radioactive debris",
					"Explosive decompression",
					"Invisible or slow-acting toxin",
					"Fires or explosions",
				},
			},
		},
		{
			Name: "PC-induced",
			SubTable1: rollt.List{
				Items: []string{
					"Activating a system causes a disaster",
					"Catastrophic plan proposed by NPCs",
					"Removing loot triggers defenses",
					"Handling an object ruins it",
					"Leaving a thing open brings calamity",
				},
			},
			SubTable2: rollt.List{
				Items: []string{
					"Horrible vermin are admitted",
					"Local system goes berserk",
					"Something ruptures violently",
					"Ancient defenses awaken",
					"The PC’s goal is imperiled",
				},
			},
		},
	},
}
