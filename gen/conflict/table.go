package conflict

import (
	"github.com/nboughton/rollt"
	"github.com/nboughton/swnt/gen/table"
)

var Restraint = rollt.Table{
	Dice: "1d20",
	Items: []rollt.Item{
		{Match: []int{1}, Text: "The government is cracking down on the conflict"},
		{Match: []int{2}, Text: "One side seems invincibly stronger to the other"},
		{Match: []int{3}, Text: "Both sides have “doomsday” info or devices"},
		{Match: []int{4}, Text: "A prior conflict ended horribly for both of them"},
		{Match: []int{5}, Text: "Foreign participants are keeping things tamped"},
		{Match: []int{6}, Text: "Elements of both sides seek accommodation"},
		{Match: []int{7}, Text: "The conflict is only viable in a narrow location"},
		{Match: []int{8}, Text: "Catastrophic cost of losing a direct showdown"},
		{Match: []int{9}, Text: "Each thinks they’ll win without further exertion"},
		{Match: []int{10}, Text: "They expect a better opening to appear soon"},
		{Match: []int{11}, Text: "Former ties of friendship or family restrain them"},
		{Match: []int{12}, Text: "Religious principles are constraining them"},
		{Match: []int{13}, Text: "One side’s still licking their wounds after a failure"},
		{Match: []int{14}, Text: "They’re building up force to make sure they win"},
		{Match: []int{15}, Text: "Their cultural context makes open struggle hard"},
		{Match: []int{16}, Text: "They expect an outside power to hand them a win"},
		{Match: []int{17}, Text: "They’re still searching for a way to get at their goal"},
		{Match: []int{18}, Text: "One side mistakenly thinks they’ve already won"},
		{Match: []int{19}, Text: "A side is busy integrating a recent success"},
		{Match: []int{20}, Text: "An outside power threatens both sides"},
	},
}

var Twist = rollt.Table{
	Dice: "1d20",
	Items: []rollt.Item{
		{Match: []int{1}, Text: "There’s a very sharp time limit for any resolution"},
		{Match: []int{2}, Text: "The sympathetic side is actually a bunch of bastards"},
		{Match: []int{3}, Text: "There’s an easy but very repugnant solution to hand"},
		{Match: []int{4}, Text: "PC success means a big benefit to a hostile group"},
		{Match: []int{5}, Text: "The real bone of contention is hidden from most"},
		{Match: []int{6}, Text: "A sympathetic figure’s on an unsympathetic side"},
		{Match: []int{7}, Text: "There’s a profitable chance for PCs to turn traitor"},
		{Match: []int{8}, Text: "The “winner” will actually get in terrible trouble"},
		{Match: []int{9}, Text: "There’s a very appealing third party in the mix"},
		{Match: []int{10}, Text: "The PCs could really profit off the focus of the strife"},
		{Match: []int{11}, Text: "The PCs are mistaken for an involved group"},
		{Match: []int{12}, Text: "Somebody plans on screwing over the PCs"},
		{Match: []int{13}, Text: "Both sides think the PCs are working for them"},
		{Match: []int{14}, Text: "A side wants to use the PCs as a distraction for foes"},
		{Match: []int{15}, Text: "The PCs’ main contact is mistrusted by their allies"},
		{Match: []int{16}, Text: "If the other side can’t get it, they’ll destroy it"},
		{Match: []int{17}, Text: "The focus isn’t nearly as valuable as both sides think"},
		{Match: []int{18}, Text: "The focus somehow has its own will and goals"},
		{Match: []int{19}, Text: "Victory will drastically change one of the sides"},
		{Match: []int{20}, Text: "Actually, there is no twist. It’s all exactly as it seems."},
	},
}

var Problem = table.ThreePart{
	Headers: [3]string{"Conflict Type", "Overall Situation", "Specific Focus"},
	Tables: []table.ThreePartSubTable{
		{
			Name: "Money",
			SubTable1: rollt.List{
				"Money is owed to a ruthless creditor",
				"Money was stolen from someone",
				"A sudden profit opportunity arises",
				"There’s a hidden stash of wealth",
				"Money is offered from an evil source",
			},
			SubTable2: rollt.List{
				"Organized crime wants it",
				"Corrupt officials want it",
				"A sympathetic NPC needs it",
				"The PCs are owed it",
				"It will disappear very soon",
			},
		},
		{
			Name: "Revenge",
			SubTable1: rollt.List{
				"Someone was murdered",
				"Someone was stripped of rank",
				"Someone lost all their wealth",
				"Someone lost someone’s love",
				"Someone was framed for a crime",
			},
			SubTable2: rollt.List{
				"It was wholly justified",
				"The wrong person is targeted",
				"The reaction is excessive",
				"The PCs are somehow blamed",
				"Both sides were wronged",
			},
		},
		{
			Name: "Power",
			SubTable1: rollt.List{
				"An influential political leader",
				"A stern community elder",
				"A ruling patriarch of a large family",
				"A star expert in a particular industry",
				"A criminal boss or outcast leader",
			},
			SubTable2: rollt.List{
				"They’ve betrayed their own",
				"Someone’s gunning for them",
				"They made a terrible choice",
				"They usurped their position",
				"They’re oppressing their own",
			},
		},
		{
			Name: "Natural Danger",
			SubTable1: rollt.List{
				"A cyclical planetary phenomenon",
				"A sudden natural disaster",
				"Sudden loss of vital infrastructure",
				"Catastrophe from outside meddling",
				"Formerly-unknown planetary peril",
			},
			SubTable2: rollt.List{
				"Anti-helpful bureaucrats",
				"Religious zealots panic",
				"Bandits and looters strike",
				"The government hushes it up",
				"There’s money in exploiting it",
			},
		},
		{
			Name: "Religion",
			SubTable1: rollt.List{
				"Sects that hate each other bitterly",
				"Zealot reformers forcing new things",
				"Radical traditionalists fighting back",
				"Ethnic religious divisions",
				"Corrupt and decadent institutions",
			},
			SubTable2: rollt.List{
				"Charismatic new leader",
				"Mandatory state religion",
				"Heavy foreign influence",
				"Religious purging underway",
				"Fighting for holy ground",
			},
		},
		{
			Name: "Ideology",
			SubTable1: rollt.List{
				"A universally-despised fringe group",
				"Terrorists with widespread support",
				"A political party’s goon squads",
				"Dead-end former regime supporters",
				"Ruthless ascendant political group",
			},
			SubTable2: rollt.List{
				"Terrorist attack",
				"Street rioting",
				"Police state crackdown",
				"Forced expulsions",
				"Territory under hostile rule",
			},
		},
		{
			Name: "Ethnicity",
			SubTable1: rollt.List{
				"A traditionally subordinate group",
				"An ethnic group from offworld",
				"A dominant caste or ethnicity",
				"An alien or transhuman group",
				"Two groups that hate each other",
			},
			SubTable2: rollt.List{
				"Forced immigration",
				"Official ethnic ghettos",
				"Rigid separation of groups",
				"Group statuses have changed",
				"Rising ethnic violence",
			},
		},
		{
			Name: "Resources",
			SubTable1: rollt.List{
				"There’s a cache of illegal materials",
				"A hidden strike of rare resources",
				"Cargo has been abandoned as lost",
				"Land ownership is disputed",
				"A resource is desperately necessary",
			},
			SubTable2: rollt.List{
				"Someone thinks they own it",
				"The state is looking for it",
				"It has its own protectors",
				"Rights to it were stolen",
				"Offworlders want it badly",
			},
		},
	},
}
