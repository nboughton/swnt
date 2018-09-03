package content

import (
	"bytes"
	"fmt"
	"math/rand"

	"github.com/nboughton/rollt"
	"github.com/nboughton/swnt/content/culture"
	"github.com/nboughton/swnt/content/format"
	"github.com/nboughton/swnt/content/gender"
	"github.com/nboughton/swnt/content/name"
	"github.com/nboughton/swnt/content/table"
)

// Patron is a Patron
type Patron struct {
	Fields [][]string
}

// NewPatron just roll Patron details
func NewPatron() Patron {
	return Patron{Fields: patronTable.Roll()}
}

// Format patron data as type t
func (p Patron) Format(t format.OutputType) string {
	return format.Table(t, []string{"Patron", ""}, p.Fields)
}

// NPC represents an NPC
type NPC struct {
	Name     string
	Gender   gender.Gender
	Culture  culture.Culture
	Fields   [][]string
	Hooks    NPCHooks
	Patron   Patron
	Reaction string
}

// NPCHooks character hooks, wants etc
type NPCHooks struct {
	Manner     string
	Outcome    string
	Motivation string
	Want       string
	Power      string
	Hook       string
}

// NewNPC roll a new NPC
func NewNPC(ctr culture.Culture, g gender.Gender, isPatron bool) NPC {
	n := NPC{
		Gender:  g,
		Culture: ctr,
		Fields:  npcTable.Roll(),
		Hooks: NPCHooks{
			Manner:     npcHooksTable.manner.Roll(),
			Outcome:    npcHooksTable.outcome.Roll(),
			Motivation: npcHooksTable.motivation.Roll(),
			Want:       npcHooksTable.want.Roll(),
			Power:      npcHooksTable.power.Roll(),
			Hook:       npcHooksTable.hook.Roll(),
		},
		Reaction: Reaction.Roll(),
	}

	if isPatron {
		n.Patron = NewPatron()
	}

	nm := name.Person.ByCulture(ctr)
	switch g {
	case gender.Male:
		n.Name = fmt.Sprintf("%s %s", nm.Male.Roll(), nm.Surname.Roll())
	case gender.Female:
		n.Name = fmt.Sprintf("%s %s", nm.Female.Roll(), nm.Surname.Roll())
	case gender.Other, gender.Any:
		switch rand.Intn(2) {
		case 0:
			n.Name = fmt.Sprintf("%s %s", nm.Male.Roll(), nm.Surname.Roll())
		case 1:
			n.Name = fmt.Sprintf("%s %s", nm.Female.Roll(), nm.Surname.Roll())
		}
	}

	return n
}

// Format returns a string output in the specified format t
func (n NPC) Format(t format.OutputType) string {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, format.Table(t, []string{n.Name, ""}, [][]string{
		{"Culture", n.Culture.String()},
		{"Gender", n.Gender.String()},
	}))
	fmt.Fprintf(buf, format.Table(t, []string{}, n.Fields))
	fmt.Fprintf(buf, format.Table(t, []string{}, [][]string{
		{"Hooks", ""},
		{npcHooksTable.manner.Name, n.Hooks.Manner},
		{npcHooksTable.outcome.Name, n.Hooks.Outcome},
		{npcHooksTable.motivation.Name, n.Hooks.Motivation},
		{npcHooksTable.want.Name, n.Hooks.Want},
		{npcHooksTable.power.Name, n.Hooks.Power},
		{npcHooksTable.hook.Name, n.Hooks.Hook},
		{Reaction.Name, n.Reaction},
	}))

	if len(n.Patron.Fields) > 0 {
		fmt.Fprintln(buf)
		fmt.Fprintf(buf, n.Patron.Format(t))
	}

	return buf.String()
}

func (n NPC) String() string {
	return n.Format(format.TEXT)
}

// Reaction of possible reaction rolls for NPCs
var Reaction = rollt.Table{
	Name: "Reaction Roll Results",
	Dice: "2d6",
	Items: []rollt.Item{
		{Match: []int{2}, Text: "Hostile, reacting negatively as is plausible"},
		{Match: []int{3, 4, 5}, Text: "Negative, unfriendly and unhelpful"},
		{Match: []int{6, 7, 8}, Text: "Neutral, reacting predictably or warily"},
		{Match: []int{9, 10, 11}, Text: "Positive, potentially cooperative with PCs"},
		{Match: []int{12}, Text: "Friendly, helpful as is plausible to be"},
	},
}

var npcHooksTable = struct {
	manner     rollt.List
	outcome    rollt.List
	motivation rollt.List
	want       rollt.List
	power      rollt.List
	hook       rollt.List
}{
	// Manner table
	rollt.List{
		Name: "Initial Manner",
		Items: []string{
			"Ingratiating and cloying",
			"Grim suspicion of the PCs or their backers",
			"Xenophilic interest in the novelty of the PCs",
			"Pragmatic and businesslike",
			"Romantically interested in one or more PCs",
			"A slimy used-gravcar dealer’s approach",
			"Wide-eyed awe at the PCs",
			"Cool and superior attitude toward PC “hirelings”",
			"Benevolently patronizing toward outsiders",
			"Sweaty-palmed need or desperation",
			"Xenophobic mistrust of the PCs",
			"Idealistic enthusiasm for a potentially shared cause",
			"Somewhat intoxicated by recent indulgence",
			"Smoothly persuasive and reasonable",
			"Visibly uncomfortable with the PCs",
			"Grossly overconfident in PC abilities",
			"Somewhat frightened by the PCs",
			"Deeply misunderstanding the PCs’ culture",
			"Extremely well-informed about the PCs’ past",
			"Distracted by their current situation",
		},
	},

	// Outcome table
	rollt.List{
		Name: "Default Deal Outcome",
		Items: []string{
			"They’ll screw the PCs over even at their own cost",
			"They firmly intend to actively betray the PCs",
			"They won’t keep the deal unless driven to it",
			"They plan to twist the deal to their own advantage",
			"They won’t keep their word unless it’s profitable",
			"They’ll flinch from paying up when the time comes",
			"They mean to keep the deal, but are reluctant",
			"They’ll keep most of the deal, but not all of it",
			"They’ll keep the deal slowly and grudgingly",
			"They’ll keep the deal but won’t go out of their way",
			"They’ll be reasonably punctual about the deal",
			"They’ll want a further small favor to pay up on it",
			"They’ll keep the deal in a way that helps them",
			"They’ll keep the deal if it’s still good for them",
			"They’ll offer a bonus for an additional favor",
			"Trustworthy as long as the deal won’t hurt them",
			"Trustworthy, with the NPC following through",
			"They’ll be very fair in keeping to their agreements",
			"They’ll keep bargains even to their own cost",
			"Complete and righteous integrity to the bitter end",
		},
	},

	// Motivation table
	rollt.List{
		Name: "Motivation",
		Items: []string{
			"An ambition for greater social status",
			"Greed for wealth and indulgent riches",
			"Protect a loved one who is somehow imperiled",
			"A sheer sadistic love of inflicting pain and suffering",
			"Hedonistic enjoyment of pleasing company",
			"Searching out hidden knowledge or science",
			"Establishing or promoting a cultural institution",
			"Avenging a grievous wrong to them or a loved one",
			"Promoting their religion and living out their faith",
			"Winning the love of a particular person",
			"Winning glory and fame in their profession",
			"Dodging an enemy who is pursuing them",
			"Driving out or killing an enemy group",
			"Deposing a rival to them in their line of work",
			"Getting away from this world or society",
			"Promote a friend or offspring’s career or future",
			"Taking control of a property or piece of land",
			"Building a structure or a complex prototype tech",
			"Perform or create their art to vast acclaim",
			"Redeem themselves from a prior failure",
		},
	},

	// Want table
	rollt.List{
		Name: "Want",
		Items: []string{
			"Bring them an exotic piece of tech",
			"Convince someone to meet with the NPC",
			"Kill a particular NPC",
			"Kidnap or non-fatally eliminate a particular NPC",
			"Pay them a large amount of money",
			"Take a message to someone hard to reach",
			"Acquire a tech component that’s hard to get",
			"Find proof of a particular NPC’s malfeasance",
			"Locate a missing NPC",
			"Bring someone to a destination via dangerous travel",
			"Retrieve a lost or stolen object",
			"Defend someone from an impending attack",
			"Burn down or destroy a particular structure",
			"Explore a dangerous or remote location",
			"Steal something from a rival NPC or group",
			"Intimidate a rival into ceasing their course of action",
			"Commit a minor crime to aid the NPC",
			"Trick a rival into doing something",
			"Rescue an NPC from a dire situation",
			"Force a person or group to leave an area",
		},
	},

	// Power table
	rollt.List{
		Name: "Power",
		Items: []string{
			"They’re just really appealing and sympathetic to PCs",
			"They have considerable liquid funds",
			"They control the use of large amounts of violence",
			"They have a position of great social status",
			"They’re a good friend of an important local leader",
			"They have blackmail info on the PCs",
			"They have considerable legal influence here",
			"They have tech the PCs might reasonably want",
			"They can get the PCs into a place they want to go",
			"They know where significant wealth can be found",
			"They have information about the PCs’ current goal",
			"An NPC the PCs need has implicit trust in them",
			"The NPC can threaten someone the PCs like",
			"They control a business relevant to PC needs",
			"They have considerable criminal contacts",
			"They have pull with the local religion",
			"They know a great many corrupt politicians",
			"They can alert the PCs to an unexpected peril",
			"They’re able to push a goal the PCs currently have",
			"They can get the PCs useful permits and rights",
		},
	},

	// Hook table
	rollt.List{
		Name: "Hook",
		Items: []string{
			"A particular odd style of dress",
			"An amputation or other maiming",
			"Visible cyberware or prosthetics",
			"Unusual hair, skin, or eye colors",
			"Scarring, either intentional or from old injuries",
			"Tic-like overuse of a particular word or phrase",
			"Specific unusual fragrance or cologne",
			"Constant fiddling with a particular item",
			"Visible signs of drug use",
			"Always seems to be in one particular mood",
			"Wears badges or marks of allegiance to a cause",
			"Extremely slow or fast pace of speech",
			"Wheezes, shakes, or other signs of infirmity",
			"Constantly with a drink to hand",
			"Always complaining about a group or organization",
			"Paranoid, possibly for justifiable reasons",
			"Insists on a particular location for all meetings",
			"Communicates strictly through a third party",
			"Abnormally obese, emaciated, tall, or short",
			"Always found with henchmen or friends",
		},
	},
}

// PatronTable represents the tables to roll on to create an adventure Patron
var patronTable = table.OneRoll{
	D4: rollt.List{
		Name: "Patron Eagerness to Hire",
		Items: []string{
			"Cautious, but can be convinced to hire",
			"Willing to promise standard rates",
			"Eager, willing to offer a bonus",
			"Desperate, might offer what they can’t pay",
		},
	},
	D6: rollt.List{
		Name: "Patron Trustworthiness",
		Items: []string{
			"They intend to totally screw the PCs",
			"They won’t pay unless forced to do so",
			"They’ll pay slowly or reluctantly",
			"They’ll pay, but discount for mistakes",
			"They’ll pay without quibbling",
			"They’ll pay more than they promised",
		},
	},
	D8: rollt.List{
		Name: "Basic Challenge of the Job",
		Items: []string{
			"Kill somebody who might deserve it",
			"Kidnap someone dangerous",
			"Steal a well-guarded object",
			"Arson or sabotage on a place",
			"Get proof of some misdeed",
			"Protect someone from an immediate threat",
			"Transport someone through danger",
			"Guard an object being transported",
		},
	},
	D10: rollt.List{
		Name: "Main Countervailing Force",
		Items: []string{
			"A treacherous employer or subordinate",
			"An open and known enemy of the patron",
			"Official governmental meddling",
			"An unknown rival of the patron",
			"The macguffin itself opposes them",
			"Very short time frame allowed",
			"The job is spectacularly illegal",
			"A participant would profit by their failure",
			"The patron is badly wrong about a thing",
			"The locals are against the patron",
		},
	},
	D12: rollt.List{
		Name: "Potential Non-Cash Rewards",
		Items: []string{
			"Government official favors owed",
			"Property in the area",
			"An item very valuable on another world",
			"Pretech mod components",
			"Useful pretech artifact",
			"Information the PCs need",
			"Membership in a powerful group",
			"Black market access",
			"Use of restricted facilities or shipyards",
			"Shares in a profitable business",
			"Maps to a hidden or guarded treasure",
			"Illegal but valuable weapons or gear",
		},
	},
	D20: rollt.List{
		Name: "Complication to the Job",
		Items: []string{
			"An ambush is laid somewhere",
			"PC involvement is leaked to the enemy",
			"The patron gives faulty aid somehow",
			"Failing would be extremely unhealthy",
			"The job IDs them as allies of a local faction",
			"The macguffin is physically dangerous",
			"An important location is hard to get into",
			"Succeeding would be morally distasteful",
			"A supposed ally is very unhelpful or stupid",
			"The patron badly misunderstood the PCs",
			"The job changes suddenly partway through",
			"An unexpected troublemaker is involved",
			"Critical gear will fail partway through",
			"An unrelated accident complicates things",
			"Payment comes in a hard-to-handle form",
			"Someone is turning traitor on the patron",
			"A critical element has suddenly moved",
			"Payment is in avidly-pursued hot goods",
			"The true goal is a subsidiary part of the job",
			"No complications; it’s just as it seems to be",
		},
	},
}

// NPCTable represents the tables to roll on to create an NPC
var npcTable = table.OneRoll{
	D4: rollt.List{
		Name: "Age",
		Items: []string{
			"Unusually young or old for their role",
			"Young adult",
			"Mature prime",
			"Middle-aged or elderly",
		},
	},
	D6: rollt.Table{
		Name: "Background",
		Dice: "1d6",
		Items: []rollt.Item{
			{Match: []int{1}, Text: "The local underclass or poorest natives"},
			{Match: []int{2}, Text: "Common laborers or cube workers"},
			{Match: []int{3}, Text: "Aspiring bourgeoise or upper class"},
			{Match: []int{4}, Text: "The elite of this society"},
			{Match: []int{5}, Text: "Minority or foreigners"},
			{Match: []int{6}, Text: "Offworlders or exotics"},
		},
		Reroll: rollt.Reroll{
			Dice:  "1d4",
			Match: []int{5, 6},
		},
	},
	D8: rollt.List{
		Name: "Role in Society",
		Items: []string{
			"Criminal, thug, thief, swindler",
			"Menial, cleaner, retail worker, servant",
			"Unskilled heavy labor, porter, construction",
			"Skilled trade, electrician, mechanic, pilot",
			"Idea worker, programmer, writer",
			"Merchant, business owner, trader, banker",
			"Official, bureaucrat, courtier, clerk",
			"Military, soldier, enforcer, law officer",
		},
	},
	D10: rollt.List{
		Name: "Biggest Problem",
		Items: []string{
			"They have significant debt or money woes",
			"A loved one is in trouble; reroll for it",
			"Romantic failure with a desired person",
			"Drug or behavioral addiction",
			"Their superior dislikes or resents them",
			"They have a persistent sickness",
			"They hate their job or life situation",
			"Someone dangerous is targeting them",
			"They’re pursuing a disastrous purpose",
			"They have no problems worth mentioning",
		},
	},
	D12: rollt.List{
		Name: "Greatest Desire",
		Items: []string{
			"They want a particular romantic partner",
			"They want money for them or a loved one",
			"They want a promotion in their job",
			"They want answers about a past trauma",
			"They want revenge on an enemy",
			"They want to help a beleaguered friend",
			"They want an entirely different job",
			"They want protection from an enemy",
			"They want to leave their current life",
			"They want fame and glory",
			"They want power over those around them",
			"They have everything they want from life",
		},
	},
	D20: rollt.List{
		Name: "Most Obvious Character Trait",
		Items: []string{
			"Ambition",
			"Avarice",
			"Bitterness",
			"Courage",
			"Cowardice",
			"Curiosity",
			"Deceitfulness",
			"Determination",
			"Devotion to a cause",
			"Filiality",
			"Hatred",
			"Honesty",
			"Hopefulness",
			"Love of a person",
			"Nihilism",
			"Paternalism",
			"Pessimism",
			"Protectiveness",
			"Resentment",
			"Shame",
		},
	},
}
