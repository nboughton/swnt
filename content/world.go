package content

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"

	"github.com/nboughton/go-roll"
	"github.com/nboughton/swnt/content/culture"
	"github.com/nboughton/swnt/content/format"
	"github.com/nboughton/swnt/content/name"
)

// TagsTable represents the collection of Tags
type TagsTable []Tag

// Roll selects a random Tag
func (t TagsTable) Roll() string {
	return fmt.Sprint(t[rand.Intn(len(t))])
}

// Random selects a random tag (used in Adventure seed generation)
func (t TagsTable) Random() string {
	return Tags[rand.Intn(len(Tags))].Name
}

// Find returns the tag specified. The search is case insensitive for convenience
func (t TagsTable) Find(name string) (Tag, error) {
	for _, tag := range t {
		if strings.ToLower(tag.Name) == strings.ToLower(name) {
			return tag, nil
		}
	}

	return Tag{}, fmt.Errorf("no tag with name \"%s\"", name)
}

func selectTags(exclude []string) (Tag, Tag) {
	var t TagsTable
	for _, tag := range Tags {
		if !tag.match(exclude) {
			t = append(t, tag)
		}
	}

	t1Idx, t2Idx := rand.Intn(len(t)), rand.Intn(len(t))
	for t1Idx == t2Idx { // Ensure the same tag isn't selected twice
		t2Idx = rand.Intn(len(t))
	}

	return t[t1Idx], t[t2Idx]
}

func (t Tag) match(s []string) bool {
	for _, x := range s {
		if strings.ToLower(t.Name) == strings.ToLower(x) {
			return true
		}
	}

	return false
}

// Tag represents a complete World Tag structure as extracted from the Stars Without Number core book.
type Tag struct {
	Name          string
	Desc          string
	Enemies       roll.List
	Friends       roll.List
	Complications roll.List
	Things        roll.List
	Places        roll.List
}

// Format tag as type o
func (t Tag) Format(o format.OutputType) string {
	return format.Table(o, []string{t.Name, ""}, [][]string{
		{"Description", t.Desc},
		{"Enemies", t.Enemies.String()},
		{"Friends", t.Friends.String()},
		{"Complications", t.Complications.String()},
		{"Things", t.Things.String()},
		{"Places", t.Places.String()},
	})
}

func (t Tag) String() string {
	return t.Format(format.TEXT)
}

// World represents a generated world
type World struct {
	Primary      bool
	FullTags     bool
	Name         string
	Culture      culture.Culture
	Tags         [2]Tag
	Atmosphere   string
	Temperature  string
	Population   string
	Biosphere    string
	TechLevel    string
	Origin       string
	Relationship string
	Contact      string
}

// NewWorld creates a new world. Set culture to culture.Any for a random culture and primary to false
// to include relationship information. If tagNamesOnly is true then format output will not include full
// tag text
func NewWorld(primary bool, c culture.Culture, fullTags bool, excludeTags []string) World {
	t1, t2 := selectTags(excludeTags)

	w := World{
		Primary:     primary,
		FullTags:    fullTags,
		Name:        name.Table.ByCulture(c).Place.Roll(),
		Culture:     c,
		Tags:        [2]Tag{t1, t2},
		Atmosphere:  worldTable.atmosphere.Roll(),
		Temperature: worldTable.temperature.Roll(),
		Population:  worldTable.population.Roll(),
		Biosphere:   worldTable.biosphere.Roll(),
		TechLevel:   worldTable.techLevel.Roll(),
	}

	if !w.Primary {
		w.Origin = otherWorldTable.origin.Roll()
		w.Relationship = otherWorldTable.relationship.Roll()
		w.Contact = otherWorldTable.contact.Roll()
	}

	return w
}

// Format returns the content of World w in format t
func (w World) Format(t format.OutputType) string {
	var buf = new(bytes.Buffer)

	fmt.Fprintf(buf, format.Table(t, []string{w.Name, ""}, [][]string{
		{"Atmosphere", w.Atmosphere},
		{"Temperature", w.Temperature},
		{"Biosphere", w.Biosphere},
		{"Population", w.Population},
		{"Culture", w.Culture.String()},
		{"Tech Level", w.TechLevel},
	}))

	if !w.FullTags {
		fmt.Fprintf(buf, format.Table(t, []string{}, [][]string{
			{"Tags", fmt.Sprintf("%s, %s", w.Tags[0].Name, w.Tags[1].Name)},
		}))
	} else {
		fmt.Fprintf(buf, format.Table(t, []string{}, [][]string{
			{"Tags", ""},
			{w.Tags[0].Name, w.Tags[0].Desc},
			{w.Tags[1].Name, w.Tags[1].Desc},
		}))
	}

	if !w.Primary {
		fmt.Fprintf(buf, format.Table(t, []string{}, [][]string{
			{"Origins", ""},
			{otherWorldTable.origin.Name, w.Origin},
			{otherWorldTable.relationship.Name, w.Relationship},
			{otherWorldTable.contact.Name, w.Contact},
		}))
	}

	return buf.String()
}

func (w World) String() string {
	return w.Format(format.TEXT)
}

// Other represents origins of secondary population centers in a System
var otherWorldTable = struct {
	origin       roll.List
	relationship roll.List
	contact      roll.List
}{
	roll.List{
		Name: "Origin of the World",
		Items: []string{
			"Recent colony from the primary world",
			"Refuge for exiles from primary",
			"Founded ages ago by a different group",
			"Founded long before the primary world",
			"Lost ancient colony of the primary",
			"Colony recently torn free of the primary",
			"Long-standing cooperative colony world",
			"Recent interstellar colony from elsewhere",
		},
	},
	roll.List{
		Name: "Current Relationship",
		Items: []string{
			"Confirmed hatred of each other",
			"Active cold war between them",
			"Old grudges or resentments",
			"Cultural disgust and avoidance",
			"Polite interchange and trade",
			"Cultural admiration for primary",
			"Long-standing friendship",
			"Unflinching mutual loyalty",
		},
	},
	roll.List{
		Name: "Contact Point",
		Items: []string{
			"Trade in vital goods",
			"Shared religion",
			"Mutual language",
			"Entertainment content",
			"Shared research",
			"Threat to both of them",
			"Shared elite families",
			"Exploiting shared resource",
		},
	},
}

var worldTable = struct {
	atmosphere  roll.Table
	biosphere   roll.Table
	temperature roll.Table
	techLevel   roll.Table
	population  roll.Table
}{
	// Atmosphere List
	roll.Table{
		Dice: "2d6",
		Items: []roll.TableItem{
			{Match: []int{2}, Text: "Corrosive, damaging to foreign objects"},
			{Match: []int{3}, Text: "Inert gas, useless for respiration"},
			{Match: []int{4}, Text: "Airless or thin to the point of suffocation"},
			{Match: []int{5, 6, 7, 8, 9}, Text: "Breathable mix"},
			{Match: []int{10}, Text: "Thick, but breathable with a pressure mask"},
			{Match: []int{11}, Text: "Invasive, penetrating suit seals"},
			{Match: []int{12}, Text: "Both corrosive and invasive in its effect"},
		},
	},

	// Biosphere List
	roll.Table{
		Dice: "2d6",
		Items: []roll.TableItem{
			{Match: []int{2}, Text: "Remnant biosphere"},
			{Match: []int{3}, Text: "Microbial life forms exist"},
			{Match: []int{4, 5}, Text: "No native biosphere"},
			{Match: []int{6, 7, 8}, Text: "Human-miscible biosphere"},
			{Match: []int{9, 10}, Text: "Immiscible biosphere"},
			{Match: []int{11}, Text: "Hybrid biosphere"},
			{Match: []int{12}, Text: "Engineered biosphere"},
		},
	},

	// Temperature List
	roll.Table{
		Dice: "2d6",
		Items: []roll.TableItem{
			{Match: []int{2}, Text: "Frozen, locked in perpetual ice"},
			{Match: []int{3}, Text: "Cold, dominated by glaciers and tundra"},
			{Match: []int{4, 5}, Text: "Variable cold with temperate places"},
			{Match: []int{6, 7, 8}, Text: "Temperate, Earthlike in its ranges"},
			{Match: []int{9, 10}, Text: "Variable warm, with temperate places"},
			{Match: []int{11}, Text: "Warm, tropical and hotter in places"},
			{Match: []int{12}, Text: "Burning, intolerably hot on its surface"},
		},
	},

	// TechLevel List
	roll.Table{
		Dice: "2d6",
		Items: []roll.TableItem{
			{Match: []int{2}, Text: "TL0, neolithic-level technology"},
			{Match: []int{3}, Text: "TL1, medieval technology"},
			{Match: []int{4, 5}, Text: "TL2, early Industrial Age tech"},
			{Match: []int{6, 7, 8}, Text: "TL4, modern postech"},
			{Match: []int{9, 10}, Text: "TL3, tech like that of present-day Earth"},
			{Match: []int{11}, Text: "TL4+, postech with specialties"},
			{Match: []int{12}, Text: "TL5, pretech with surviving infrastructure"},
		},
	},

	// Population List
	roll.Table{
		Dice: "2d6",
		Items: []roll.TableItem{
			{Match: []int{2}, Text: "Failed colony"},
			{Match: []int{3}, Text: "Outpost"},
			{Match: []int{4, 5}, Text: "Fewer than a million inhabitants"},
			{Match: []int{6, 7, 8}, Text: "Several million inhabitants"},
			{Match: []int{9, 10}, Text: "Hundreds of millions of inhabitants"},
			{Match: []int{11}, Text: "Billions of inhabitants"},
			{Match: []int{12}, Text: "Alien inhabitants"},
		},
	},
}

// Tags List
var Tags = TagsTable{
	{
		Name: "Abandoned Colony",
		Desc: "The world once hosted a colony, whether human or otherwise, until some crisis or natural disaster drove the inhabitants away or killed them off. The colony might have been mercantile in nature, an expedition to extract valuable local resources, or it might have been a reclusive cabal of zealots. The remains of the colony are usually in ruins, and might still be dangerous from the aftermath of whatever destroyed it in the first place.",
		Enemies: roll.List{
			Items: []string{
				"Crazed survivors",
				"Ruthless plunderers of the ruins",
				"Automated defense system",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Inquisitive stellar archaeologist",
				"Heir to the colony’s property",
				"Local wanting the place cleaned out and made safe",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The local government wants the ruins to remain a secret",
				"The locals claim ownership of it",
				"The colony is crumbling and dangerous to navigate",
			},
		},
		Things: roll.List{
			Items: []string{
				"Long-lost property deeds",
				"Relic stolen by the colonists when they left",
				"Historical record of the colonization attempt",
			},
		},
		Places: roll.List{
			Items: []string{
				"Decaying habitation block",
				"Vine-covered town square",
				"Structure buried by an ancient landslide",
			},
		},
	},
	{
		Name: "Alien Ruins",
		Desc: "The world has significant alien ruins present. The locals may or may not permit others to investigate the ruins, and may make it difficult to remove any objects of value without substantial payment. Any surviving ruins with worthwhile salvage almost certainly have some defense or hazard to explain their unplundered state.",
		Enemies: roll.List{
			Items: []string{
				"Customs inspector",
				"Worshipper of the ruins",
				"Hidden alien survivor",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Curious scholar",
				"Avaricious local resident",
				"Interstellar smuggler",
			},
		},
		Complications: roll.List{
			Items: []string{
				"Traps in the ruins",
				"Remote location",
				"Paranoid customs officials",
			},
		},
		Things: roll.List{
			Items: []string{
				"Precious alien artifacts",
				"Objects left with the remains of a prior unsuccessful expedition",
				"Untranslated alien texts",
				"Untouched hidden ruins",
			},
		},
		Places: roll.List{
			Items: []string{
				"Undersea ruin",
				"Orbital ruin",
				"Perfectly preserved",
			},
		},
	},
	{
		Name: "Altered Humanity",
		Desc: "The humans on this world are visibly and drastically different from normal humanity. They may have additional limbs, new sensory organs, or other significant changes. Were these from ancestral eugenic manipulation, strange stellar mutations, or from an environmental toxin unique to this world?",
		Enemies: roll.List{
			Items: []string{
				"Biochauvinist local",
				"Local experimenter",
				"Mentally unstable mutant",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Local seeking a “cure”",
				"Curious xenophiliac",
				"Anthropological researcher",
			},
		},
		Complications: roll.List{
			Items: []string{
				"Alteration is contagious",
				"Alteration is necessary for long-term survival",
				"Locals fear and mistrust non-local humans",
			},
		},
		Things: roll.List{
			Items: []string{
				"Original pretech mutagenic equipment",
				"Valuable biological byproduct from the mutants",
				"“Cure” for the altered genes",
				"Record of the original colonial genotypes",
			},
		},
		Places: roll.List{
			Items: []string{
				"Abandoned eugenics laboratory",
				"An environment requiring the mutation for survival",
				"A sacred site where the first local was transformed",
			},
		},
	},
	{
		Name: "Anarchists",
		Desc: "Rather than being an incidental anarchy of struggling tribes and warring factions, this world actually has a functional society with no centralized authority. Authority might be hyper-localized to extended families, specific religious parishes, or voluntary associations. Some force is preventing an outside group or internal malcontents from coalescing into a power capable of imposing its rule on the locals; this force might be an ancient pretech defense system, a benevolent military AI, or the sheer obscurity and isolation of the culture.",
		Enemies: roll.List{
			Items: []string{
				"Offworlder imperialist",
				"Reformer seeking to impose “good government”",
				"Exploiter taking advantage of the lack of centralized resistance",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Proud missionary for anarchy",
				"Casual local free spirit",
				"Curious offworlder political scientist",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The anarchistic structure is compelled by an external power",
				"The anarchy is enabled by currently abundant resources",
				"The protecting force that shelters the anarchy is waning",
			},
		},
		Things: roll.List{
			Items: []string{
				"A macguffin that would let the possessor enforce their rule on others",
				"A vital resource needed to preserve general liberty",
				"Tech forbidden as disruptive to the social order",
			},
		},
		Places: roll.List{
			Items: []string{
				"Community of similar-sized homes",
				"Isolated clan homestead",
				"Automated mining site",
			},
		},
	},
	{
		Name: "Anthropomorphs",
		Desc: "The locals were originally human, but at some point became anthropomorphic, either as an ancient furry colony, a group of animal-worshiping sectarians, or gengineers who just happened to find animal elements most convenient for surviving on the world. Depending on the skill of the original gengineers, their feral forms may or may not work as well as their original human bodies, or may come with drawbacks inherited from their animal elements.",
		Enemies: roll.List{
			Items: []string{
				"Anthro-supremacist local",
				"Native driven by feral urges",
				"Outside exploiter who sees the locals as subhuman creatures",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Fascinated genetic researcher",
				"Diplomat trained to deal with normals",
				"Local needing outside help",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The locals consider their shapes a curse from their foolish ancestors",
				"Society is ordered according to animal forms",
				"The locals view normal humans as repulsive or inferior",
			},
		},
		Things: roll.List{
			Items: []string{
				"Pretech gengineering tech",
				"A “cure” that may not be wanted",
				"Sacred feral totem",
			},
		},
		Places: roll.List{
			Items: []string{
				"Shrine to a feral deity",
				"Nature preserve suited to an animal type",
				"Living site built to take advantage of animal traits",
			},
		},
	},
	{
		Name: "Area 51",
		Desc: "The world’s government is fully aware of their local stellar neighbors, but the common populace has no idea about it- and the government means to keep it that way. Trade with government officials in remote locations is possible, but any attempt to clue the commoners in on the truth will be met with lethal reprisals.",
		Enemies: roll.List{
			Items: []string{
				"Suspicious government minder",
				"Free merchant who likes his local monopoly",
				"Local who wants a specimen for dissection",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Crusading offworld investigator",
				"Conspiracy-theorist local",
				"Idealistic government reformer",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The government has a good reason to keep the truth concealed",
				"The government ruthlessly oppresses the natives",
				"The government is actually composed of offworlders",
			},
		},
		Things: roll.List{
			Items: []string{
				"Elaborate spy devices",
				"Memory erasure tech",
				"Possessions of the last offworlder who decided to spread the truth",
			},
		},
		Places: roll.List{
			Items: []string{
				"Desert airfield",
				"Deep subterranean bunker",
				"Hidden mountain valley",
			},
		},
	},
	{
		Name: "Badlands World",
		Desc: "Whatever the original climate and atmosphere type, something horrible happened to this world. Biological, chemical, or nanotechnical weaponry has reduced it to a wretched hellscape. Some local life might still be able to survive on its blasted surface, usually at some dire cost in health or humanity.",
		Enemies: roll.List{
			Items: []string{
				"Mutated badlands fauna",
				"Desperate local",
				"Badlands raider chief",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Native desperately wishing to escape the world",
				"Scientist researching ecological repair methods",
				"Ruin scavenger",
			},
		},
		Complications: roll.List{
			Items: []string{
				"Radioactivity",
				"Bioweapon traces",
				"Broken terrain",
				"Sudden local plague",
			},
		},
		Things: roll.List{
			Items: []string{
				"Maltech research core",
				"Functional pretech Maltech research core",
				"Functional pretech weaponry",
				"An uncontaminated well",
			},
		},
		Places: roll.List{
			Items: []string{
				"Untouched oasis",
				"Ruined city",
				"Salt flat",
			},
		},
	},
	{
		Name: "Battleground",
		Desc: "The world is a battleground for two or more outside powers. They may be interstellar rivals, or groups operating out of orbitals or other system bodies. Something about the planet is valuable enough for them to fight over, but the natives are too weak to be anything but animate obstacles to the fight.",
		Enemies: roll.List{
			Items: []string{
				"Ruthless military commander",
				"Looter pack chieftain",
				"Traitorous collaborator",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Native desperately seeking protection",
				"Pragmatic military officer",
				"Hapless war orphan",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The war just ended as both sides are leaving",
				"The natives somehow brought this on themselves",
				"A small group of natives profit tremendously from the fighting",
			},
		},
		Things: roll.List{
			Items: []string{
				"A cache of the resource the invaders seek",
				"Abandoned prototype military gear",
				"Precious spy intelligence lost by someone",
			},
		},
		Places: roll.List{
			Items: []string{
				"Artillery-pocked wasteland",
				"Reeking refugee camp",
				"Burnt-out shell of a city",
			},
		},
	},
	{
		Name: "Beastmasters",
		Desc: "The natives have extremely close bonds with the local fauna, possibly having special means of communication and control through tech or gengineering. Local animal life plays a major role in their society, industry, or warfare, and new kinds of beasts may be bred to suit their purposes.",
		Enemies: roll.List{
			Items: []string{
				"Half-feral warlord of a beast swarm",
				"Coldly inhuman scientist",
				"Altered beast with human intellect and furious malice",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Native bonded with an adorable animal",
				"Herder of very useful beasts",
				"Animal-revering mystic",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The “animals” are very heavily gengineered humans",
				"The animals actually run the society",
				"The animals have the same rights as humans",
			},
		},
		Things: roll.List{
			Items: []string{
				"Tech used to alter animal life",
				"A plague vial that could wipe out the animals",
				"A pretech device that can perform a wonder if operated by a beast",
			},
		},
		Places: roll.List{
			Items: []string{
				"Park designed as a comfortable home for beasts",
				"Public plaza designed to accommodate animal companions",
				"Factory full of animal workers",
			},
		},
	},
	{
		Name: "Bubble Cities",
		Desc: "Whether due to a lack of atmosphere or an uninhabitable climate, the world’s cities exist within domes or pressurized buildings. In such sealed environments, techniques of surveillance and control can grow baroque and extreme.",
		Enemies: roll.List{
			Items: []string{
				"Native dreading outsider contamination",
				"Saboteur from another bubble city",
				"Local official hostile to outsider ignorance of laws",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Local rebel against the city officials",
				"Maintenance chief in need of help",
				"Surveyor seeking new building sites",
			},
		},
		Complications: roll.List{
			Items: []string{
				"Bubble rupture",
				"Failing atmosphere reprocessor",
				"Native revolt against officials",
				"All-seeing surveillance cameras",
			},
		},
		Things: roll.List{
			Items: []string{
				"Pretech habitat technology",
				"Valuable industrial products",
				"Master key codes to a city’s security system",
			},
		},
		Places: roll.List{
			Items: []string{
				"City power core",
				"Surface of the bubble",
				"Hydroponics complex",
				"Warren-like hab block",
			},
		},
	},
	{
		Name: "Cheap Life",
		Desc: "Human life is near-worthless on this world. Ubiquitous cloning, local conditions that ensure early death, a culture that reveres murder, or a social structure that utterly discounts the value of most human lives ensures that death is the likely outcome for any action that irritates someone consequential. ",
		Enemies: roll.List{
			Items: []string{
				"Master assassin",
				"Bloody-handed judge",
				"Overseer of disposable clones",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Endearing local whose life the PCs accidentally bought",
				"Escapee from death seeking outside help",
				"Reformer trying to change local mores",
			},
		},
		Complications: roll.List{
			Items: []string{
				"Radiation or local diseases ensure all locals die before twenty-five years of age",
				"Tech ensures that death is just an annoyance",
				"Locals are totally convinced of a blissful afterlife",
			},
		},
		Things: roll.List{
			Items: []string{
				"Device that revives or re-embodies the dead",
				"Maltech engine fueled by human life",
				"Priceless treasure held by a now-dead owner",
			},
		},
		Places: roll.List{
			Items: []string{
				"Thronging execution ground",
				"extremely cursory cemetery",
				"Factory full of lethal dangers that could be corrected easily",
			},
		},
	},
	{
		Name: "Civil War",
		Desc: "The world is currently torn between at least two opposing factions, all of which claim legitimacy. The war may be the result of a successful rebel uprising against tyranny, or it might just be the result of schemers who plan to be the new masters once the revolution is complete.",
		Enemies: roll.List{
			Items: []string{
				"Faction commissar",
				"Angry native",
				"Conspiracy theorist who blames offworlders for the war",
				"Deserter looking out for himself",
				"Guerrilla bandit chieftain",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Faction loyalist seeking aid",
				"Native caught in the crossfire",
				"Offworlder seeking passage off the planet",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The front rolls over the group",
				"Famine strikes",
				"Bandit infestations are in the way",
			},
		},
		Things: roll.List{
			Items: []string{
				"Ammo dump",
				"Military cache",
				"Treasure buried for after the war",
				"Secret war plans",
			},
		},
		Places: roll.List{
			Items: []string{
				"Battle front",
				"Bombed-out town",
				"Rear-area red light zone",
				"Propaganda broadcast tower",
			},
		},
	},
	{
		Name: "Cold War",
		Desc: "Two or more great powers control the planet, and they have a hostility to each other that’s just barely less than open warfare. The hostility might be ideological in nature, or it might revolve around control of some local resource.",
		Enemies: roll.List{
			Items: []string{
				"Suspicious chief of intelligence",
				"Native who thinks the outworlders are with the other side",
				"Femme fatale",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Apolitical information broker",
				"Spy for the other side",
				"Unjustly accused innocent",
				"“He’s a bastard",
				"but he’s our bastard” official",
			},
		},
		Complications: roll.List{
			Items: []string{
				"Police sweep",
				"Low-level skirmishing",
				"“Red scare”",
			},
		},
		Things: roll.List{
			Items: []string{
				"List of traitors in government",
				"secret military plans",
				"Huge cache of weapons built up in preparation for war",
			},
		},
		Places: roll.List{
			Items: []string{
				"Seedy bar in a neutral area",
				"Political rally",
				"Isolated area where fighting is underway",
			},
		},
	},
	{
		Name: "Colonized Population",
		Desc: "A neighboring world has successfully colonized this less-advanced or less-organized planet, and the natives aren’t happy about it. A puppet government may exist, but all real decisions are made by the local viceroy.",
		Enemies: roll.List{
			Items: []string{
				"Suspicious security personnel",
				"Offworlder-hating natives",
				"Local crime boss preying on rich offworlders",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Native resistance leader",
				"Colonial official seeking help",
				"Native caught between the two sides",
			},
		},
		Complications: roll.List{
			Items: []string{
				"Natives won’t talk to offworlders",
				"Colonial repression",
				"Misunderstood local customs",
			},
		},
		Things: roll.List{
			Items: []string{
				"Relic of the resistance movement",
				"List of collaborators",
				"Precious substance extracted by colonial labor",
			},
		},
		Places: roll.List{
			Items: []string{
				"Deep wilderness resistance camp",
				"City district off-limits to natives",
				"Colonial labor site",
			},
		},
	},
	{
		Name: "Cultural Power",
		Desc: "The world is a considerable cultural power in the sector, producing music, art, philosophy, or some similar intangible that their neighbors find irresistibly attractive. Other worlds might have a profound degree of cultural cachet as the inheritor of some venerable artistic tradition.",
		Enemies: roll.List{
			Items: []string{
				"Murderously eccentric artist",
				"Crazed fan",
				"Failed artist with an obsessive grudge",
				"Critic with a crusade to enact",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Struggling young artist",
				"Pupil of the artistic tradition",
				"Scholar of the art",
				"Offworlder hating the source of corrupting alien ways",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The art is slowly lethal to its masters",
				"The art is mentally or physically addictive",
				"The art is a fragment of ancient technical or military science",
			},
		},
		Things: roll.List{
			Items: []string{
				"The instrument of a legendary master",
				"The only copy of a dead master’s opus",
				"Proof of intellectual property ownership",
			},
		},
		Places: roll.List{
			Items: []string{
				"Recording or performance studio",
				"Public festival choked with tourists",
				"Monument to a dead master of the art",
			},
		},
	},
	{
		Name: "Cybercommunists",
		Desc: "On this world communism actually works, thanks to pretech computing devices and greater or lesser amounts of psychic precognition. Central planning nodes direct all production and employment on the world. Citizens in good standing have access to ample amounts of material goods for all needs and many wants. Instead of strife over wealth, conflicts erupt over political controls, cultural precepts, or control over the planning nodes. Many cybercommunist worlds show a considerable bias toward the private interests of those who run the planning nodes.",
		Enemies: roll.List{
			Items: []string{
				"Embittered rebel against perceived unfairness",
				"Offworlder saboteur",
				"Aspiring Stalin-figure",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Idealistic planning node tech",
				"Cynical anti-corruption cop",
				"Precognitive economist",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The pretech planning computers are breaking down",
				"The planning only works because the locals have been mentally or physically altered",
				"The planning computers can’t handle the increasing population within the system",
			},
		},
		Things: roll.List{
			Items: []string{
				"Planning node computer",
				"Wildly destabilizing commodity that can’t be factored into plans",
				"A tremendous store of valuables made by accident",
			},
		},
		Places: roll.List{
			Items: []string{
				"Humming factory",
				"Apartment block of perfectly equal flats",
				"Mass demonstration of unity",
			},
		},
	},
	{
		Name: "Cyborgs",
		Desc: "The planet’s population makes heavy use of cybernetics, with many of the inhabitants possessing at least a cosmetic amount of chrome. This may be the result of a strong local cyber tech base, a religious injunction, or simply a necessary measure to survive the local conditions.",
		Enemies: roll.List{
			Items: []string{
				"Ambitious hacker of cyber implants",
				"Cybertech oligarch",
				"Researcher craving fresh offworlders",
				"Cybered-up gang boss",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Charity-working implant physician",
				"Idealistic young cyber researcher",
				"Avant-garde activist",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The powerful and dangerous come here often for cutting-edge implants",
				"The cyber has some universal negative side-effect",
				"Cyber and those implanted with it are forbidden to leave the planet as a tech security measure",
			},
		},
		Things: roll.List{
			Items: []string{
				"Unique prototype cyber implant",
				"Secret research files",
				"A virus that debilitates cyborgs",
				"A cache of critically-needed therapeutic cyber",
			},
		},
		Places: roll.List{
			Items: []string{
				"Grimy slum chop-shop",
				"Bloody lair of implant rippers",
				"Stark plaza where everyone is seeing things through their augmented-reality cyber",
			},
		},
	},
	{
		Name: "Cyclical Doom",
		Desc: "The world regularly suffers some apocalyptic catastrophe that wipes out organized civilization on it. The local culture is aware of this cycle and has traditions to ensure a fragment of civilization survives into the next era, but these traditions don’t always work properly, and sometimes dangerous fragments of the past emerge.",
		Enemies: roll.List{
			Items: []string{
				"Offworlder seeking to trigger the apocalypse early for profit",
				"Local recklessly taking advantage of preparation stores",
				"Demagogue claiming the cycle is merely a myth of the authorities",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Harried official working to prepare",
				"Offworlder studying the cycles",
				"Local threatened by perils of the cycle’s initial stages",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The cycles really are a myth of the authorities",
				"The cycles are controlled by alien constructs",
				"An outside power is interfering with preparation",
			},
		},
		Things: roll.List{
			Items: []string{
				"A lost cache of ancient treasures",
				"Tech or archives that will pinpoint the cycle’s timing",
				"Keycodes to bypass an ancient vault’s security",
			},
		},
		Places: roll.List{
			Items: []string{
				"Lethally-defended vault of forgotten secrets",
				"Starport crowded with panicked refugees",
				"Town existing in the shadow of some monstrous monument to a former upheaval",
			},
		},
	},
	{
		Name: "Desert World",
		Desc: "The world may have a breathable atmosphere and a human-tolerable temperature range, but it is an arid, stony waste outside of a few places made habitable by human effort. The deep wastes are largely unexplored and inhabited by outcasts and worse.",
		Enemies: roll.List{
			Items: []string{
				"Raider chieftain",
				"Crazed hermit",
				"Angry isolationists",
				"Paranoid mineral prospector",
				"Strange desert beast",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Native guide",
				"Research biologist",
				"Aspiring terraformer",
			},
		},
		Complications: roll.List{
			Items: []string{
				"Sandstorms",
				"Water supply failure",
				"Native warfare over water rights",
			},
		},
		Things: roll.List{
			Items: []string{
				"Enormous water reservoir",
				"Map of hidden wells",
				"Pretech rainmaking equipment",
			},
		},
		Places: roll.List{
			Items: []string{
				"Oasis",
				"“The Empty Quarter” of the desert",
				"Hidden underground cistern",
			},
		},
	},
	{
		Name: "Doomed World",
		Desc: "The world is doomed, and the locals may or may not know it. Some cosmic catastrophe looms before them, and the locals have no realistic way to get everyone to safety. To the extent that the public is aware, society is disintegrating into a combination of religious fervor, abject hedonism, and savage violence.",
		Enemies: roll.List{
			Items: []string{
				"Crazed prophet of a false salvation",
				"Ruthless leader seeking to flee with their treasures",
				"Cynical ship captain selling a one-way trip into hard vacuum as escape to another world",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Appealing waif or family head seeking escape",
				"Offworld relief coordinator",
				"Harried law officer",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The doom is false or won’t actually kill everyone",
				"The doom was intentionally triggered by someone",
				"Mass escape is possible if warring groups can somehow be brought to cooperate",
			},
		},
		Things: roll.List{
			Items: []string{
				"Clearance for a ship to leave the planet",
				"A cache of priceless cultural artifacts",
				"The life savings of someone trying to buy passage out",
				"Data that would prove to the public the end is nigh",
			},
		},
		Places: roll.List{
			Items: []string{
				"Open square beneath a sky angry with a foretaste of th impending ruin",
				"Orgiastic celebration involving sex and murder in equal parts",
				"Holy site full of desperate petitioners to the divine",
			},
		},
	},
	{
		Name: "Dying Race",
		Desc: "The inhabitants of this world are dying out, and they know it. Through environmental toxins, hostile bio-weapons, or sheer societal despair, the culture cannot replenish its numbers. Members seek meaning in their own strange goals or peculiar faiths, though a few might struggle to find some way to reverse their slow yet inevitable doom.",
		Enemies: roll.List{
			Items: []string{
				"Hostile outsider who wants the locals dead",
				"Offworlder seeking to take advantage of their weakened state",
				"Invaders eager to push the locals out of their former lands",
			},
		},
		Friends: roll.List{
			Items: []string{
				"One of the few youth among the population",
				"Determined and hopeful reformer",
				"Researcher seeking a new method of reproduction",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The dying culture’s values were monstrous",
				"The race’s death is somehow necessary to prevent some grand catastrophe",
				"The race is somehow convinced they deserve this fate",
			},
		},
		Things: roll.List{
			Items: []string{
				"Extremely valuable reproductive tech",
				"Treasured artifacts of the former age",
				"Bioweapon used on the race",
			},
		},
		Places: roll.List{
			Items: []string{
				"City streets devoid of pedestrians",
				"Mighty edifice now crumbling with disrepair",
				"Small dwelling full of people in a town now otherwise empty",
			},
		},
	},
	{
		Name: "Eugenic Cult",
		Desc: "Even in the days before the Silence, major improvement of the human genome always seemed to come with unacceptable side-effects. Some worlds host secret cults that perpetuate these improvements regardless of the cost, and a few planets have been taken over entirely by the cults.",
		Enemies: roll.List{
			Items: []string{
				"Eugenic superiority fanatic",
				"Mentally unstable homo superior",
				"Mad eugenic scientist",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Eugenic propagandist",
				"Biotechnical investigator",
				"Local seeking revenge on cult",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The altered cultists look human",
				"The locals are terrified of any unusual physical appearance",
				"The genetic modifications- and drawbacks- are contagious with long exposure",
			},
		},
		Things: roll.List{
			Items: []string{
				"Serum that induces the alteration",
				"Elixir that reverses the alteration",
				"Pretech biotechnical databanks",
				"List of secret cult sympathizers",
			},
		},
		Places: roll.List{
			Items: []string{
				"Eugenic breeding pit",
				"Isolated settlement of altered humans",
				"Public place infiltrated by cult sympathizers",
			},
		},
	},
	{
		Name: "Exchange Consulate",
		Desc: "The Exchange of Light once served as the largest, most trusted banking and diplomatic service in human space. Even after the Silence, some worlds retain a functioning Exchange Consulate where banking services and arbitration can be arranged.",
		Enemies: roll.List{
			Items: []string{
				"Corrupt Exchange official",
				"Indebted native who thinks the players are Exchange agents",
				"Exchange official dunning the players for debts incurred",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Consul in need of offworld help",
				"Local banker seeking to hurt his competition",
				"Exchange diplomat",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The local Consulate has been corrupted",
				"the Consulate is cut off from its funds",
				"A powerful debtor refuses to pay",
			},
		},
		Things: roll.List{
			Items: []string{
				"Exchange vault codes",
				"Wealth hidden to conceal it from a bankruptcy judgment",
				"Location of forgotten vault",
			},
		},
		Places: roll.List{
			Items: []string{
				"Consulate meeting chamber",
				"Meeting site between fractious disputants",
				"Exchange vault",
			},
		},
	},
	{
		Name: "Fallen Hegemon",
		Desc: "At some point in the past, this world was a hegemonic power over some or all of the sector, thanks to superior tech, expert diplomacy, the weakness of their neighbors, or inherited Mandate legitimacy. Some kind of crash or revolt broke their power, however, and now the world is littered with the wreckage of former glory.",
		Enemies: roll.List{
			Items: []string{
				"Bitter pretender to a meaningless throne",
				"Resentful official dreaming of empire",
				"Vengeful offworlder seeking to punish their old rulers",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Realistic local leader trying to hold things together",
				"Scholar of past glories",
				"Refugee from an overthrown colonial enclave",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The hegemon’s rule was enlightened and fair",
				"It collapsed due to its own internal strife rather than external resistance",
				"It pretends that nothing has happened to its power",
				"It’s been counter-colonized by vengeful outsiders",
			},
		},
		Things: roll.List{
			Items: []string{
				"Precious insignia of former rule",
				"Relic tech important to its power",
				"Plundered colonial artifact",
			},
		},
		Places: roll.List{
			Items: []string{
				"Palace far too grand for its current occupant",
				"Oversized spaceport now in disrepair",
				"Boulevard lined with monuments to past glories",
			},
		},
	},
	{
		Name: "Feral World",
		Desc: "In the long, isolated night of the Silence, some worlds have experienced total moral and cultural collapse. Whatever remains has been twisted beyond recognition into assorted death cults, xenophobic fanaticism, horrific cultural practices, or other behavior unacceptable on more enlightened worlds. These worlds are almost invariably quarantined by other planets.",
		Enemies: roll.List{
			Items: []string{
				"Decadent noble",
				"Mad cultist",
				"Xenophobic local",
				"Cannibal chief",
				"Maltech researcher",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Trapped outworlder",
				"Aspiring reformer",
				"Native wanting to avoid traditional flensing",
			},
		},
		Complications: roll.List{
			Items: []string{
				"Horrific local “celebration”",
				"Inexplicable and repugnant social rules",
				"Taboo zones and people",
			},
		},
		Things: roll.List{
			Items: []string{
				"Terribly misused piece of pretech",
				"Wealth accumulated through brutal evildoing",
				"Valuable possession owned by luckless outworlder victim",
			},
		},
		Places: roll.List{
			Items: []string{
				"Atrocity amphitheater",
				"Traditional torture parlor",
				"Ordinary location twisted into something terrible.",
			},
		},
	},
	{
		Name: "Flying Cities",
		Desc: "Perhaps the world is a gas giant, or plagued with unendurable storms at lower levels of the atmosphere. For whatever reason, the cities of this world fly above the surface of the planet. Perhaps they remain stationary, or perhaps they move from point to point in search of resources.",
		Enemies: roll.List{
			Items: []string{
				"Rival city pilot",
				"Tech thief attempting to steal outworld gear",
				"Saboteur or scavenger plundering the city’s tech",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Maintenance tech in need of help",
				"City defense force pilot",
				"Meteorological researcher",
			},
		},
		Complications: roll.List{
			Items: []string{
				"Sudden storms",
				"Drastic altitude loss",
				"Rival city attacks",
				"Vital machinery breaks down",
			},
		},
		Things: roll.List{
			Items: []string{
				"Precious refined atmospheric gases",
				"Pretech grav engine plans",
				"Meteorological codex predicting future storms",
			},
		},
		Places: roll.List{
			Items: []string{
				"Underside of the city",
				"The one calm place on the planet’s surface",
				"Catwalks stretching over unimaginable gulfs below.",
			},
		},
	},
	{
		Name: "Forbidden Tech",
		Desc: "Some group on this planet fabricates or uses maltech. Unbraked AIs doomed to metastasize into insanity, nation-destroying nanowarfare particles, slow-burn DNA corruptives, genetically engineered slaves, or something worse still. The planet’s larger population may or may not be aware of the danger in their midst.",
		Enemies: roll.List{
			Items: []string{
				"Mad scientist",
				"Maltech buyer from offworld",
				"Security enforcer",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Victim of maltech",
				"Perimeter agent",
				"Investigative reporter",
				"Conventional arms merchant",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The maltech is being fabricated by an unbraked AI",
				"The government depends on revenue from maltech sales to offworlders",
				"Citizens insist that it’s not really maltech",
			},
		},
		Things: roll.List{
			Items: []string{
				"Maltech research data",
				"The maltech itself",
				"Precious pretech equipment used to create it",
			},
		},
		Places: roll.List{
			Items: []string{
				"Horrific laboratory",
				"Hellscape sculpted by the maltech’s use",
				"Government building meeting room",
			},
		},
	},
	{
		Name: "Former Warriors",
		Desc: "The locals of this world were once famed for their martial prowess. They may have simply had a very militaristic culture, or were genetically engineered for combat, or developed high-tech weaponry, or had brilliant leadership. Those days are past, however, either due to crushing defeat, external restrictions, or a cultural turn toward peace.",
		Enemies: roll.List{
			Items: []string{
				"Unreformed warlord leader",
				"Bitter mercenary chief",
				"Victim of their warfare seeking revenge",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Partisan of the new peaceful ways",
				"Outsider desperate for military aid",
				"Martial genius repressed by the new dispensation",
			},
		},
		Complications: roll.List{
			Items: []string{
				"Neighboring worlds want them pacified or dead",
				"They only ever used their arts in self-defense",
				"The source of their gifts has been “turned off” in a reversible way",
			},
		},
		Things: roll.List{
			Items: []string{
				"War trophy taken from a defeated foe",
				"Key to re-activating their martial ways",
				"Secret cache of high-tech military gear",
			},
		},
		Places: roll.List{
			Items: []string{
				"Cemetery of dead heroes",
				"Memorial hall now left to dust and silence",
				"Monument plaza dedicated to the new culture",
			},
		},
	},
	{
		Name: "Freak Geology",
		Desc: "The geology or geography of this world is simply freakish. Perhaps it’s composed entirely of enormous mountain ranges, or regular bands of land and sea, or the mineral structures all fragment into perfect cubes. The locals have learned to deal with it and their culture will be shaped by its requirements.",
		Enemies: roll.List{
			Items: []string{
				"Crank xenogeologist",
				"Cultist who believes it the work of aliens",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Research scientist",
				"Prospector",
				"Artist",
			},
		},
		Complications: roll.List{
			Items: []string{
				"Local conditions that no one remembers to tell outworlders about",
				"Lethal weather",
				"Seismic activity",
			},
		},
		Things: roll.List{
			Items: []string{
				"Unique crystal formations",
				"Hidden veins of a major precious mineral strike",
				"Deed to a location of great natural beauty",
			},
		},
		Places: roll.List{
			Items: []string{
				"Atop a bizarre geological formation",
				"Tourist resort catering to offworlders",
			},
		},
	},
	{
		Name: "Freak Weather",
		Desc: "The planet is plagued with some sort of bizarre or hazardous weather pattern. Perhaps city-flattening storms regularly scourge the surface, or the world’s sun never pierces its thick banks of clouds.",
		Enemies: roll.List{
			Items: []string{
				"Criminal using the weather as a cover",
				"Weather cultists convinced the offworlders are responsible for some disaster",
				"Native predators dependent on the weather",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Meteorological researcher",
				"Holodoc crew wanting shots of the weather",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The weather itself",
				"Malfunctioning pretech terraforming engines that cause the weather",
			},
		},
		Things: roll.List{
			Items: []string{
				"Wind-scoured deposits of precious minerals",
				"Holorecords of a spectacularly and rare weather pattern",
				"Naturally-sculpted objects of intricate beauty",
			},
		},
		Places: roll.List{
			Items: []string{
				"Eye of the storm",
				"The one sunlit place",
				"Terraforming control room",
			},
		},
	},
	{
		Name: "Friendly Foe",
		Desc: "Some hostile alien race or malevolent cabal has a branch or sect on this world that is actually quite friendly toward outsiders. For whatever internal reason, they are willing to negotiate and deal honestly with strangers, and appear to lack the worst impulses of their fellows.",
		Enemies: roll.List{
			Items: []string{
				"Driven hater of all their kind",
				"Internal malcontent bent on creating conflict",
				"Secret master who seeks to lure trust",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Well-meaning bug-eyed monster",
				"Principled eugenics cultist",
				"Suspicious investigator",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The group actually is as harmless and benevolent as they seem",
				"The group offers a vital service at the cost of moral compromise",
				"The group still feels bonds of affiliation with their hostile brethren",
			},
		},
		Things: roll.List{
			Items: []string{
				"Forbidden xenotech",
				"Eugenic biotech template",
				"Evidence to convince others of their kind that they are right",
			},
		},
		Places: roll.List{
			Items: []string{
				"Repurposed maltech laboratory",
				"Alien conclave building",
				"Widely-feared starship interior",
			},
		},
	},
	{
		Name: "Gold Rush",
		Desc: "Gold, silver, and other conventional precious minerals are common and cheap now that asteroid mining is practical for most worlds. But some minerals and compounds remain precious and rare, and this world has recently been discovered to have a supply of them. People from across the sector have come to strike it rich.",
		Enemies: roll.List{
			Items: []string{
				"Paranoid prospector",
				"Aspiring mining tycoon",
				"Rapacious merchant",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Claim-jumped miner",
				"Native alien",
				"Curious tourist",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The strike is a hoax",
				"The strike is of a dangerous toxic substance",
				"Export of the mineral is prohibited by the planetary government",
				"The native aliens live around the strike’s location",
			},
		},
		Things: roll.List{
			Items: []string{
				"Cases of the refined element",
				"Pretech mining equipment",
				"A dead prospector’s claim deed",
			},
		},
		Places: roll.List{
			Items: []string{
				"Secret mine",
				"Native alien village",
				"Processing plant",
				"Boom town",
			},
		},
	},
	{
		Name: "Great Work",
		Desc: "The locals are obsessed with completing a massive project, one that has consumed them for generations. It might be the completion of a functioning spaceyard, a massive solar power array, a network of terraforming engines, or the universal conversion of their neighbors to their own faith. The purpose of their entire civilization is to progress and some day complete the work.",
		Enemies: roll.List{
			Items: []string{
				"Local planning to sacrifice the PCs for the work",
				"Local who thinks the PCs threaten the work",
				"Obsessive zealot ready to destroy someone or something important to the PCs for the sake of the work",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Outsider studying the work",
				"Local with a more temperate attitude",
				"Supplier of work materials",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The work is totally hopeless",
				"Different factions disagree on what the work is",
				"An outside power is determined to thwart the work",
			},
		},
		Things: roll.List{
			Items: []string{
				"Vital supplies for the work",
				"Plans that have been lost",
				"Tech that greatly speeds the work",
			},
		},
		Places: roll.List{
			Items: []string{
				"A bustling work site",
				"Ancestral worker housing",
				"Local community made only semi-livable by the demands of the work",
			},
		},
	},
	{
		Name: "Hatred",
		Desc: "For whatever reason, this world’s populace has a burning hatred for the inhabitants of a neighboring system. Perhaps this world was colonized by exiles, or there was a recent interstellar war, or ideas of racial or religious superiority have fanned the hatred. Regardless of the cause, the locals view their neighbor and any sympathizers with loathing.",
		Enemies: roll.List{
			Items: []string{
				"Native convinced that the offworlders are agents of Them",
				"Cynical politician in need of scapegoats",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Intelligence agent needing catspaws",
				"Holodoc producers needing “an inside look”",
				"Unlucky offworlder from the hated system",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The characters are wearing or using items from the hated world",
				"The characters are known to have done business there",
				"The characters “look like” the hated others",
			},
		},
		Things: roll.List{
			Items: []string{
				"Proof of Their evildoing",
				"Reward for turning in enemy agents",
				"Relic stolen by Them years ago",
			},
		},
		Places: roll.List{
			Items: []string{
				"War crimes museum",
				"Atrocity site",
				"Captured and decommissioned spaceship kept as a trophy",
			},
		},
	},
	{
		Name: "Heavy Industry",
		Desc: "With interstellar transport so limited in the bulk it can move, worlds have to be largely self-sufficient in industry. Some worlds are more sufficient than others, however, and this planet has a thriving manufacturing sector capable of producing large amounts of goods appropriate to its tech level. The locals may enjoy a correspondingly higher lifestyle, or the products might be devoted towards vast projects for the aggrandizement of the rulers.",
		Enemies: roll.List{
			Items: []string{
				"Tycoon monopolist",
				"Industrial spy",
				"Malcontent revolutionary",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Aspiring entrepreneur",
				"Worker union leader",
				"Ambitious inventor",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The factories are toxic",
				"The resources extractable at their tech level are running out",
				"The masses require the factory output for survival",
				"The industries’ major output is being obsoleted by offworld tech",
			},
		},
		Things: roll.List{
			Items: []string{
				"Confidential industrial data",
				"Secret union membership lists",
				"Ownership shares in an industrial complex",
			},
		},
		Places: roll.List{
			Items: []string{
				"Factory floor",
				"Union meeting hall",
				"Toxic waste dump",
				"R&D complex",
			},
		},
	},
	{
		Name: "Heavy Mining",
		Desc: "This world has large stocks of valuable minerals, usually necessary for local industry, life support, or refinement into loads small enough to export offworld. Major mining efforts are necessary to extract the minerals, and many natives work in the industry.",
		Enemies: roll.List{
			Items: []string{
				"Mine boss",
				"Tunnel saboteur",
				"Subterranean predators",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Hermit prospector",
				"Offworld investor",
				"Miner’s union representative",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The refinery equipment breaks down",
				"Tunnel collapse",
				"Silicate life forms growing in the miners’ lungs",
			},
		},
		Things: roll.List{
			Items: []string{
				"The mother lode",
				"Smuggled case of refined mineral",
				"Faked crystalline mineral samples",
			},
		},
		Places: roll.List{
			Items: []string{
				"Vertical mine face",
				"Tailing piles",
				"Roaring smelting complex",
			},
		},
	},
	{
		Name: "Hivemind",
		Desc: "Natives of this world exist in a kind of mental gestalt, sharing thoughts and partaking of a single identity. Powerful pretech, exotic psionics, alien influence, or some other cause has left the world sharing one identity. Individual members may have greater or lesser degrees of effective coordination with the whole.",
		Enemies: roll.List{
			Items: []string{
				"A hivemind that wants to assimilate outsiders",
				"A hivemind that has no respect for unjoined life",
				"A hivemind that fears and hates unjoined life",
			},
		},
		Friends: roll.List{
			Items: []string{
				"A scholar studying the hivemind",
				"A person severed from the gestalt",
				"A relative of someone who has been assimilated",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The hivemind only functions on this world",
				"The hivemind has strict range limits",
				"The hivemind has different personality factions",
				"The hivemind only happens at particular times",
				"The world is made of semi-sentient drones and a single AI",
			},
		},
		Things: roll.List{
			Items: []string{
				"Vital tech for maintaining the mind",
				"Precious treasure held by now-assimilated outsider",
				"Tech that “blinds” the hivemind to the tech’s users",
			},
		},
		Places: roll.List{
			Items: []string{
				"Barely tolerable living cells for individuals",
				"Workside where individuals casually die in their labors",
				"Community with absolutely no social or group-gathering facilities",
			},
		},
	},
	{
		Name: "Holy War",
		Desc: "A savage holy war is raging on this world, either between factions of locals or as a united effort against the pagans of some neighboring world. This war might involve a conventional religion, or it might be the result of a branding campaign, political ideology, artistic movement, or any other cause that people use as a substitute for traditional religion.",
		Enemies: roll.List{
			Items: []string{
				"Blood-mad pontiff",
				"Coldly cynical secular leader",
				"Totalitarian political demagogue",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Desperate peacemaker",
				"Hard-pressed refugee of the fighting",
				"Peaceful religious leader who lost the internal debate",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The targets of the war really are doing something diabolically horrible",
				"The holy war is just a mask for a very traditional casus belli",
				"The leaders don’t want the war won but only prolonged",
				"Both this world and the target of the war are religion-obsessed",
			},
		},
		Things: roll.List{
			Items: []string{
				"Sacred relic of the faith",
				"A captured blasphemer under a death sentence",
				"Plunder seized in battle",
			},
		},
		Places: roll.List{
			Items: []string{
				"Massive holy structure",
				"Razed community of infidels",
				"Vast shrine to the martyrs dead in war",
			},
		},
	},
	{
		Name: "Hostile Biosphere",
		Desc: "The world is teeming with life, and it hates humans. Perhaps the life is xenoallergenic, forcing filter masks and tailored antiallergens for survival. It could be the native predators are huge and fearless, or the toxic flora ruthlessly outcompetes earth crops.",
		Enemies: roll.List{
			Items: []string{
				"Local fauna",
				"Nature cultist",
				"Native aliens",
				"Callous labor overseer",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Xenobiologist",
				"Tourist on safari",
				"Grizzled local guide",
			},
		},
		Complications: roll.List{
			Items: []string{
				"Filter masks fail",
				"Parasitic alien infestation",
				"Crop greenhouses lose bio-integrity",
			},
		},
		Things: roll.List{
			Items: []string{
				"Valuable native biological extract",
				"Abandoned colony vault",
				"Remains of an unsuccessful expedition",
			},
		},
		Places: roll.List{
			Items: []string{
				"Deceptively peaceful glade",
				"Steaming polychrome jungle",
				"Nightfall when surrounded by Things",
			},
		},
	},
	{
		Name: "Hostile Space",
		Desc: "The system in which the world exists is a dangerous neighborhood. Something about the system is perilous to inhabitants, either through meteor swarms, stellar radiation, hostile aliens in the asteroid belt, or periodic comet clouds.",
		Enemies: roll.List{
			Items: []string{
				"Alien raid leader",
				"Meteor-launching terrorists",
				"Paranoid local leader",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Astronomic researcher",
				"Local defense commander",
				"Early warning monitor agent",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The natives believe the danger is divine chastisement",
				"The natives blame outworlders for the danger",
				"The native elite profit from the danger in some way",
			},
		},
		Things: roll.List{
			Items: []string{
				"Early warning of a raid or impact",
				"Abandoned riches in a disaster zone",
				"Key to a secure bunker",
			},
		},
		Places: roll.List{
			Items: []string{
				"City watching an approaching asteroid",
				"Village burnt in an alien raid",
				"Massive ancient crater",
			},
		},
	},
	{
		Name: "Immortals",
		Desc: "Natives of this world are effectively immortal. They may have been gengineered for tremendous lifespans, or have found some local anagathic, or be cyborg life forms, or be so totally convinced of reincarnation that death is a cultural irrelevance. Any immortality technique is likely applicable only to them, or else it’s apt to be a massive draw to outside imperialists.",
		Enemies: roll.List{
			Items: []string{
				"Outsider determined to steal immortality",
				"Smug local convinced of their immortal wisdom to rule all",
				"Offworlder seeking the world’s ruin before it becomes a threat to all",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Curious longevity researcher",
				"Thrill-seeking local",
			},
		},
		Complications: roll.List{
			Items: []string{
				"Immortality requires doing something that outsiders can’t or won’t willingly do",
				"The immortality ends if they leave the world",
				"Death is the punishment for even minor crimes",
				"Immortals must die or go offworld after a certain span",
				"Immortality has brutal side-effects",
			},
		},
		Things: roll.List{
			Items: []string{
				"Immortality drug",
				"Masterwork of an ageless artisan",
				"Toxin that only affects immortals",
			},
		},
		Places: roll.List{
			Items: []string{
				"Community with no visible children",
				"Unchanging structure of obvious ancient age",
				"Cultural performance relying on a century of in-jokes",
			},
		},
	},
	{
		Name: "Local Specialty",
		Desc: "The world may be sophisticated or barely capable of steam engines, but either way it produces something rare and precious to the wider galaxy. It might be some pharmaceutical extract produced by a secret recipe, a remarkably popular cultural product, or even gengineered humans uniquely suited for certain work.",
		Enemies: roll.List{
			Items: []string{
				"Monopolist",
				"Offworlder seeking prohibition of the specialty",
				"Native who views the specialty as sacred",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Spy searching for the source",
				"Artisan seeking protection",
				"Exporter with problems",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The specialty is repugnant in nature",
				"The crafters refuse to sell to offworlders",
				"The specialty is made in a remote",
				"dangerous place",
				"The crafters don’t want to make the specialty any more",
			},
		},
		Things: roll.List{
			Items: []string{
				"The specialty itself",
				"The secret recipe",
				"Sample of a new improved variety",
			},
		},
		Places: roll.List{
			Items: []string{
				"Secret manufactory",
				"Hidden cache",
				"Artistic competition for best artisan",
			},
		},
	},
	{
		Name: "Local Tech",
		Desc: "The locals can create a particular example of extremely high tech, possibly even something that exceeds pretech standards. They may use unique local resources to do so, or have stumbled on a narrow scientific breakthrough, or still have a functional experimental manufactory.",
		Enemies: roll.List{
			Items: []string{
				"Keeper of the tech",
				"Offworld industrialist",
				"Automated defenses that suddenly come alive",
				"Native alien mentors",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Curious offworld scientist",
				"Eager tech buyer",
				"Native in need of technical help",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The tech is unreliable",
				"The tech only works on this world",
				"The tech has poorly-understood side effects",
				"The tech is alien in nature.",
			},
		},
		Things: roll.List{
			Items: []string{
				"The tech itself",
				"An unclaimed payment for a large shipment",
				"The secret blueprints for its construction",
				"An ancient alien R&D database",
			},
		},
		Places: roll.List{
			Items: []string{
				"Alien factory",
				"Lethal R&D center",
				"Tech brokerage vault",
			},
		},
	},
	{
		Name: "Major Spaceyard",
		Desc: "Most worlds of tech level 4 or greater have the necessary tech and orbital facilities to build spike drives and starships. This world is blessed with a major spaceyard facility, either inherited from before the Silence or painstakingly constructed in more recent decades. It can build even capital-class hulls, and do so more quickly and cheaply than its neighbors.",
		Enemies: roll.List{
			Items: []string{
				"Enemy saboteur",
				"Industrial spy",
				"Scheming construction tycoon",
				"Aspiring ship hijacker",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Captain stuck in drydock",
				"Maintenance chief",
				"Mad innovator",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The spaceyard is an alien relic",
				"The spaceyard is burning out from overuse",
				"The spaceyard is alive",
				"The spaceyard relies on maltech to function",
			},
		},
		Things: roll.List{
			Items: []string{
				"Intellectual property-locked pretech blueprints",
				"Override keys for activating old pretech facilities",
				"A purchased but unclaimed spaceship.",
			},
		},
		Places: roll.List{
			Items: []string{
				"Hidden shipyard bay",
				"Surface of a partially-completed ship",
				"Ship scrap graveyard",
			},
		},
	},
	{
		Name: "Mandarinate",
		Desc: "The planet is ruled by an intellectual elite chosen via ostensibly neutral examinations or tests. The values this system selects for may or may not have anything to do with actual practical leadership skills, and the examinations may be more or less corruptible.",
		Enemies: roll.List{
			Items: []string{
				"Corrupt test administrator",
				"Incompetent but highly-rated graduate",
				"Ruthless leader of a clan of high-testing relations",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Crusader for test reform",
				"Talented but poorly-connected graduate",
				"Genius who tests badly",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The test is totally unrelated to necessary governing skills",
				"The test was very pertinent in the past but tech or culture has changed",
				"The test is for a skill that is vital to maintaining society but irrelevant to day-to-day governance",
				"The test is a sham and passage is based on wealth or influence",
			},
		},
		Things: roll.List{
			Items: []string{
				"Answer key to the next test",
				"Lost essay of incredible merit",
				"Proof of cheating",
			},
		},
		Places: roll.List{
			Items: []string{
				"Massive structure full of test-taking cubicles",
				"School filled with desperate students",
				"Ornate government building decorated with scholarly quotes and academic images",
			},
		},
	},
	{
		Name: "Mandate Base",
		Desc: "The Terran Mandate retained its control over this world for much longer than usual, and the world may still consider itself a true inheritor of Mandate legitimacy. Most of these worlds have or had superior technology, but they may still labor under the burden of ancient restrictive tech or monitoring systems designed to prevent them from rebelling.",
		Enemies: roll.List{
			Items: []string{
				"Deranged Mandate monitoring AI",
				"Aspiring sector ruler",
				"Demagogue preaching local superiority over “traitorous rebel worlds”.",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Idealistic do-gooder local",
				"Missionary for advanced Mandate tech",
				"Outsider seeking lost data from Mandate records",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The monitoring system forces the locals to behave in aggressive ways toward “rebel” worlds",
				"The monitoring system severely hinders offworld use of their tech",
				"The original colonists are all dead and have been replaced by outsiders who don’t understand all the details",
			},
		},
		Things: roll.List{
			Items: []string{
				"Ultra-advanced pretech",
				"Mandate military gear",
				"Databank containing precious tech schematics",
			},
		},
		Places: roll.List{
			Items: []string{
				"Faded Mandate offices still in use",
				"Vault containing ancient pretech",
				"Carefully-maintained monument to Mandate glory",
			},
		},
	},
	{
		Name: "Maneaters",
		Desc: "The locals are cannibals, either out of necessity or out of cultural preference. Some worlds may actually eat human flesh, while others simply require the rendering of humans into important chemicals or pharmaceutical compounds, perhaps to prolong the lives of ghoul overlords. This cannibalism plays a major role in their society.",
		Enemies: roll.List{
			Items: []string{
				"Ruthless ghoul leader",
				"Chieftain of a ravenous tribe",
				"Sophisticated degenerate preaching the splendid authenticity of cannibalism",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Sympathetic local fleeing the fork",
				"Escapee from a pharmaceutical rendering plant",
				"Outsider chosen for dinner",
				"Reformer seeking to break the custom or its necessity",
			},
		},
		Complications: roll.List{
			Items: []string{
				"Local food or environmental conditions make human consumption grimly necessary",
				"The locals farm human beings",
				"Outsiders are expected to join in the custom",
				"The custom is totally unnecessary but jealously maintained by the people",
			},
		},
		Things: roll.List{
			Items: []string{
				"Belongings of a recent meal",
				"An offworlder VIP due for the menu",
				"A toxin that makes human flesh lethal to consumers",
			},
		},
		Places: roll.List{
			Items: []string{
				"Hideous human abattoir",
				"Extremely civilized restaurant",
				"Funeral-home-cum-kitchen",
			},
		},
	},
	{
		Name: "Megacorps",
		Desc: "The world is dominated by classic cyberpunk-esque megacorporations, each one far more important than the vestigial national remnants that encompass them. These megacorps are usually locked in a cold war, trading and dealing with each other even as they try to strike in deniable ways. An over-council of corporations usually acts to bring into line any that get excessively overt in their activities.",
		Enemies: roll.List{
			Items: []string{
				"Megalomaniacal executive",
				"Underling looking to use the PCs as catspaws",
				"Ruthless mercenary who wants what the PCs have",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Victim of megacorp scheming",
				"Offworlder merchant in far over their head",
				"Local reformer struggling to cope with megacorp indifference",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The megacorps are the only source of something vital to life on this world",
				"An autonomous Mandate system acts to punish excessively overt violence",
				"The megacorps are struggling against much more horrible national governments",
			},
		},
		Things: roll.List{
			Items: []string{
				"Blackmail on a megacorp exec",
				"Keycodes to critical corp secrets",
				"Proof of corp responsibility for a heinously unacceptable public atrocity",
				"Data on a vital new product line coming out soon",
			},
		},
		Places: roll.List{
			Items: []string{
				"A place plastered in megacorp ads",
				"A public plaza discreetly branded",
				"Private corp military base",
			},
		},
	},
	{
		Name: "Mercenaries",
		Desc: "The world is either famous for its mercenary bands or it is plagued by countless groups of condottieri in service to whatever magnate can afford to pay or bribe them adequately.",
		Enemies: roll.List{
			Items: []string{
				"Amoral mercenary leader",
				"Rich offworlder trying to buy rule of the world",
				"Mercenary press gang chief forcing locals into service",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Young and idealistic mercenary chief",
				"Harried leader of enfeebled national army",
				"Offworlder trying to hire help for a noble cause",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The mercenaries are all that stand between the locals and a hungry imperial power",
				"The mercenaries are remnants of a former official army",
				"The mercenaries hardly ever actually fight as compared to taking bribes to walk away",
			},
		},
		Things: roll.List{
			Items: []string{
				"Lost mercenary payroll shipment",
				"Forbidden military tech",
				"Proof of a band’s impending treachery against their employers",
			},
		},
		Places: roll.List{
			Items: []string{
				"Shabby camp of undisciplined mercs",
				"Burnt-out village occupied by mercenaries",
				"Luxurious and exceedingly well-defended merc leader villa",
			},
		},
	},
	{
		Name: "Minimal Contact",
		Desc: "The locals refuse most contact with offworlders. Only a small, quarantined treaty port is provided for offworld trade, and ships can expect an exhaustive search for contraband. Local governments may be trying to keep the very existence of interstellar trade a secret from their populations, or they may simply consider offworlders too dangerous or repugnant to be allowed among the population.",
		Enemies: roll.List{
			Items: []string{
				"Customs official",
				"Xenophobic natives",
				"Existing merchant who doesn’t like competition",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Aspiring tourist",
				"Anthropological researcher",
				"Offworld thief",
				"Religious missionary",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The locals carry a disease harmless to them and lethal to outsiders",
				"The locals hide dark purposes from offworlders",
				"The locals have something desperately needed but won’t bring it into the treaty port",
			},
		},
		Things: roll.List{
			Items: []string{
				"Contraband trade goods",
				"Security perimeter codes",
				"Black market local products",
			},
		},
		Places: roll.List{
			Items: []string{
				"Treaty port bar",
				"Black market zone",
				"Secret smuggler landing site",
			},
		},
	},
	{
		Name: "Misandry/Misogyny",
		Desc: "The culture on this world holds a particular gender in contempt. Members of that gender are not permitted positions of formal power, and may be restricted in their movements and activities. Some worlds may go so far as to scorn both traditional genders, using gengineering techniques to hybridize or alter conventional human biology.",
		Enemies: roll.List{
			Items: []string{
				"Cultural fundamentalist",
				"Cultural missionary to outworlders",
				"Local rebel driven to pointless and meaningless violence",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Oppressed native",
				"Research scientist",
				"Offworld emancipationist",
				"Local reformer",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The oppressed gender is restive against the customs",
				"The oppressed gender largely supports the customs",
				"The customs relate to some physical quality of the world",
				"The oppressed gender has had maltech gengineering done to “tame” them.",
			},
		},
		Things: roll.List{
			Items: []string{
				"Aerosol reversion formula for undoing gengineered docility",
				"Hidden history of the world",
				"Pretech gengineering equipment",
			},
		},
		Places: roll.List{
			Items: []string{
				"Shrine to the virtues of the favored gender",
				"Security center for controlling the oppressed",
				"Gengineering lab",
			},
		},
	},
	{
		Name: "Night World",
		Desc: "The world is plunged into eternal darkness. The only life on this planet derives its energy from other sources, such as geothermal heat, extremely volatile chemical reactions in the planet’s soil, or light in a non-visible spectrum. Most flora and fauna is voraciously eager to consume other life.",
		Enemies: roll.List{
			Items: []string{
				"Monstrous thing from the night",
				"Offworlder finding the obscurity of the world convenient for dark purposes",
				"Mad scientist experimenting with local life",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Curious offworlder researcher",
				"Hard-pressed colony leader",
				"High priest of a sect that finds religious significance in the night",
			},
		},
		Complications: roll.List{
			Items: []string{
				"Daylight comes as a cataclysmic event at very long intervals",
				"Light causes very dangerous reactions in native life or chemicals here",
				"The locals have been gengineered to exist without sight",
			},
		},
		Things: roll.List{
			Items: []string{
				"Rare chemicals created in the darkness",
				"Light source usable on this world",
				"Smuggler cache hidden here in ages pastP Formlessly pitch-black wilderness",
				"Sea without a sun",
				"Location defined by sounds or smells",
			},
		},
		Places: roll.List{
			Items: []string{
				"Formlessly pitch-black wilderness",
				"Sea without a sun",
				"Location defined by sounds or smells",
			},
		},
	},
	{
		Name: "Nomads",
		Desc: "Most of the natives of this world are nomadic, usually following a traditional cycle of movement through the lands they possess. Promises of rich plunder or local environmental perils can force these groups to strike out against neighbors. Other groups are forced to move constantly due to unpredictable dangers that crop up on the planet.",
		Enemies: roll.List{
			Items: []string{
				"Desperate tribal leader who needs what the PCs have",
				"Ruthless raider chieftain",
				"Leader seeking to weld the nomads into an army",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Free-spirited young nomad",
				"Dreamer imagining a stable life",
				"Offworlder enamored of the life",
			},
		},
		Complications: roll.List{
			Items: []string{
				"An irresistibly lethal swarm of native life forces locals to move regularly",
				"Ancient defense systems destroy too-long-stationary communities",
				"Local chemical patches require careful balancing of exposure times to avoid side effects",
			},
		},
		Things: roll.List{
			Items: []string{
				"Cache of rare and precious resource",
				"Plunder seized by a tribal raid",
				"Tech that makes a place safe for long-term inhabitation",
			},
		},
		Places: roll.List{
			Items: []string{
				"Temporary nomad camp",
				"Oasis or resource reserve",
				"Trackless waste that kills the unprepared",
			},
		},
	},
	{
		Name: "Oceanic World",
		Desc: "The world is entirely or almost entirely covered with liquid water. Habitations might be floating cities, or might cling precariously to the few rocky atolls jutting up from the waves, or are planted as bubbles on promontories deep beneath the stormy surface. Survival depends on aquaculture. Planets with inedible alien life rely on gengineered Terran sea crops.",
		Enemies: roll.List{
			Items: []string{
				"Pirate raider",
				"Violent “salvager” gang",
				"Tentacled sea monster",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Daredevil fisherman",
				"Sea hermit",
				"Sapient native life",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The liquid flux confuses grav engines too badly for them to function on this world",
				"Sea is corrosive or toxic",
				"The seas are wracked by regular storms",
			},
		},
		Things: roll.List{
			Items: []string{
				"Buried pirate treasure",
				"Location of enormous schools of fish",
				"Pretech water purification equipment",
			},
		},
		Places: roll.List{
			Items: []string{
				"The only island on the planet",
				"Floating spaceport",
				"Deck of a storm-swept ship",
				"Undersea bubble city",
			},
		},
	},
	{
		Name: "Out of Contact",
		Desc: "The natives have been entirely out of contact with the greater galaxy for centuries or longer. Perhaps the original colonists were seeking to hide from the rest of the universe, or the Silence destroyed any means of communication. It may have been so long that human origins on other worlds have regressed into a topic for legends. The players might be on the first offworld ship to land since the First Wave of colonization a thousand years ago.",
		Enemies: roll.List{
			Items: []string{
				"Fearful local ruler",
				"Zealous native cleric",
				"Sinister power that has kept the world isolated",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Scheming native noble",
				"Heretical theologian",
				"UFO cultist native",
			},
		},
		Complications: roll.List{
			Items: []string{
				"Automatic defenses fire on ships that try to take off",
				"The natives want to stay out of contact",
				"The natives are highly vulnerable to offworld diseases",
				"The native language is completely unlike any known to the group",
			},
		},
		Things: roll.List{
			Items: []string{
				"Ancient pretech equipment",
				"Terran relic brought from Earth",
				"Logs of the original colonists",
			},
		},
		Places: roll.List{
			Items: []string{
				"Long-lost colonial landing site",
				"Court of the local ruler",
				"Ancient defense battery controls",
			},
		},
	},
	{
		Name: "Outpost World",
		Desc: "The world is only a tiny outpost of human habitation planted by an offworld corporation or government. Perhaps the staff is there to serve as a refueling and repair stop for passing ships, or to oversee an automated mining and refinery complex. They might be there to study ancient ruins, or simply serve as a listening and monitoring post for traffic through the system. The outpost is likely well-equipped with defenses against casual piracy.",
		Enemies: roll.List{
			Items: []string{
				"Space-mad outpost staffer",
				"Outpost commander who wants it to stay undiscovered",
				"Undercover saboteur",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Lonely staffer",
				"Fixated researcher",
				"Overtaxed maintenance chief",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The alien ruin defense systems are waking up",
				"Atmospheric disturbances trap the group inside the outpost for a month",
				"Pirates raid the outpost",
				"The crew have become converts to a strange set of beliefs",
			},
		},
		Things: roll.List{
			Items: []string{
				"Alien relics",
				"Vital scientific data",
				"Secret corporate exploitation plans",
			},
		},
		Places: roll.List{
			Items: []string{
				"Grimy recreation room",
				"Refueling station",
				"The only building on the planet",
				"A “starport” of swept bare rock.",
			},
		},
	},
	{
		Name: "Perimeter Agency",
		Desc: "Before the Silence, the Perimeter was a Terran-sponsored organization charged with rooting out use of maltech, technology banned in human space as too dangerous for use or experimentation. Unbraked AIs, gengineered slave species, nanotech replicators, weapons of planetary destruction… the Perimeter hunted down experimenters with a great indifference to planetary laws. Most Perimeter Agencies collapsed during the Silence, but a few managed to hold on to their mission, though modern Perimeter agents often find more work as conventional spies.",
		Enemies: roll.List{
			Items: []string{
				"Renegade Agency Director",
				"Maltech researcher",
				"Paranoid intelligence chief",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Agent in need of help",
				"Support staffer",
				"“Unjustly” targeted researcher",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The local Agency has gone rogue and now uses maltech",
				"The Agency archives have been compromised",
				"The Agency has been targeted by a maltech-using organization",
				"The Agency’s existence is unknown to the locals",
			},
		},
		Things: roll.List{
			Items: []string{
				"Agency maltech research archives",
				"Agency pretech spec-ops gear",
				"File of blackmail on local politicians",
			},
		},
		Places: roll.List{
			Items: []string{
				"Interrogation room",
				"Smoky bar",
				"Maltech laboratory",
				"Secret Agency base",
			},
		},
	},
	{
		Name: "Pilgrimage Site",
		Desc: "The world is noted for an important spiritual or historical location, and might be the sector headquarters for a widespread religion or political movement. The site attracts wealthy pilgrims from throughout nearby space, and those with the money necessary to manage interstellar travel can be quite generous to the site and its keepers. The locals tend to be fiercely protective of the place and its reputation, and some places may forbid the entrance of those not suitably pious or devout.",
		Enemies: roll.List{
			Items: []string{
				"Saboteur devoted to a rival belief",
				"Bitter reformer who resents the current leadership",
				"Swindler conning the pilgrims",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Protector of the holy site",
				"Naive offworlder pilgrim",
				"Outsider wanting to learn the sanctum’s inner secrets",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The site is actually a fake",
				"The site is run by corrupt and venal keepers",
				"A natural disaster threatens the site",
			},
		},
		Things: roll.List{
			Items: []string{
				"Ancient relic guarded at the site",
				"Proof of the site’s inauthenticity",
				"Precious offering from a pilgrim",
			},
		},
		Places: roll.List{
			Items: []string{
				"Incense-scented sanctum",
				"Teeming crowd of pilgrims",
				"Imposing holy structure",
			},
		},
	},
	{
		Name: "Pleasure World",
		Desc: "This world provides delights either rare or impermissible elsewhere. Matchless local beauty, stunningly gengineered natives, a wide variety of local drugs, carnal pleasures unacceptable on other worlds, or some other rare delight is readily available here. Most worlds are fully aware of the value of their offerings, and the prices they demand can be in credits or in less tangible recompense.",
		Enemies: roll.List{
			Items: []string{
				"Purveyor of evil delights",
				"Local seeking to control others with addictions",
				"Offworlder exploiter of native resources",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Tourist who’s in too deep",
				"Native seeking a more meaningful life elsewhere",
				"Offworld entertainer looking for training here",
			},
		},
		Complications: roll.List{
			Items: []string{
				"A deeply repugnant pleasure is offered here by a culture that sees nothing wrong with it",
				"Certain pleasures here are dangerously addictive",
				"The prices here can involve enslavement or death",
				"The world has been seized and exploited by an imperial power",
			},
		},
		Things: roll.List{
			Items: []string{
				"Forbidden drug",
				"A contract for some unspeakable payment",
				"Powerful tech repurposed for hedonistic ends",
			},
		},
		Places: roll.List{
			Items: []string{
				"Breathtaking natural feature",
				"Artful but decadent salon",
				"Grimy den of desperate vice",
			},
		},
	},
	{
		Name: "Police State",
		Desc: "The world is a totalitarian police state. Any sign of disloyalty to the planet’s rulers is punished severely, and suspicion riddles society. Some worlds might operate by Soviet-style informers and indoctrination, while more technically sophisticated worlds might rely on omnipresent cameras or braked AI “guardian angels”. Outworlders are apt to be treated as a necessary evil at best, and “disappeared” if they become troublesome.",
		Enemies: roll.List{
			Items: []string{
				"Secret police chief",
				"Scapegoating official",
				"Treacherous native informer",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Rebel leader",
				"Offworld agitator",
				"Imprisoned victim",
				"Crime boss",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The natives largely believe in the righteousness of the state",
				"The police state is automated and its “rulers” can’t shut it off",
				"The leaders foment a pogrom against “offworlder spies”.",
			},
		},
		Things: roll.List{
			Items: []string{
				"List of police informers",
				"Wealth taken from “enemies of the state”",
				"Dear Leader’s private stash",
			},
		},
		Places: roll.List{
			Items: []string{
				"Military parade",
				"Gulag",
				"Gray concrete housing block",
				"Surveillance center",
			},
		},
	},
	{
		Name: "Post-Scarcity",
		Desc: "The locals have maintained sufficient Mandate-era tech to be effectively post-scarcity in their economic structure. Everyone has all the necessities and most of the desires they can imagine. Conflict now exists over the apportionment of services and terrestrial space, since anything else can be had in abundance. Military goods and items of mass destruction may still be restricted, and there is probably some reason that the locals do not export their vast wealth.",
		Enemies: roll.List{
			Items: []string{
				"Frenzied ideologue fighting over an idea",
				"Paranoid local fearing offworlder influence",
				"Grim reformer seeking the destruction of the “enfeebling” productive tech",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Offworlder seeking something available only here",
				"Local struggling to maintain the production tech",
				"Native missionary seeking to bring abundance to other worlds",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The tech causes serious side-effects on those who take advantage of it",
				"The tech is breaking down",
				"The population is growing too large",
				"The tech produces only certain things in abundance",
			},
		},
		Things: roll.List{
			Items: []string{
				"A cornucopia device",
				"A rare commodity that cannot be duplicated",
				"Contract for services",
			},
		},
		Places: roll.List{
			Items: []string{
				"Tiny but richly-appointed private quarters",
				"Market for services",
				"Hushed non-duped art salon",
			},
		},
	},
	{
		Name: "Preceptor Archive",
		Desc: "The Preceptors of the Great Archive were a pre-Silence organization devoted to ensuring the dissemination of human culture, history, and basic technology to frontier worlds that risked losing this information during the human expansion. Most frontier planets had an Archive where natives could learn useful technical skills in addition to human history and art. Those Archives that managed to survive the Silence now strive to send their missionaries of knowledge to new worlds in need of their lore.",
		Enemies: roll.List{
			Items: []string{
				"Luddite native",
				"Offworld merchant who wants the natives kept ignorant",
				"Religious zealot",
				"Corrupted First Speaker who wants to keep a monopoly on learning",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Preceptor Adept missionary",
				"Offworld scholar",
				"Reluctant student",
				"Roving Preceptor Adept",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The local Archive has taken a very religious and mystical attitude toward their teaching",
				"The Archive has maintained some replicable pretech science",
				"The Archive has been corrupted and their teaching is incorrect",
			},
		},
		Things: roll.List{
			Items: []string{
				"Lost Archive database",
				"Ancient pretech teaching equipment",
				"Hidden cache of unacceptable tech",
			},
		},
		Places: roll.List{
			Items: []string{
				"Archive lecture hall",
				"Experimental laboratory",
				"Student-local riot",
			},
		},
	},
	{
		Name: "Pretech Cultists",
		Desc: "The capacities of human science before the Silence vastly outmatch the technology available since the Scream. The Jump Gates alone were capable of crossing hundreds of light years in a moment, and they were just one example of the results won by blending psychic artifice with pretech science. Some worlds outright worship the artifacts of their ancestors, seeing in them the work of more enlightened and perfect humanity. These cultists may or may not understand the operation or replication of these devices, but they seek and guard them jealously.",
		Enemies: roll.List{
			Items: []string{
				"Cult leader",
				"Artifact supplier",
				"Pretech smuggler",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Offworld scientist",
				"Robbed collector",
				"Cult heretic",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The cultists can actually replicate certain forms of pretech",
				"The cultists abhor use of the devices as “presumption on the holy”",
				"The cultists mistake the party’s belongings for pretech",
			},
		},
		Things: roll.List{
			Items: []string{
				"Pretech artifacts both functional and broken",
				"Religious-jargon laced pretech replication techniques",
				"Waylaid payment for pretech artifacts",
			},
		},
		Places: roll.List{
			Items: []string{
				"Shrine to nonfunctional pretech",
				"Smuggler’s den",
				"Public procession showing a prized artifact",
			},
		},
	},
	{
		Name: "Primitive Aliens",
		Desc: "The world is populated by a large number of sapient aliens that have yet to develop advanced technology. The human colonists may have a friendly or hostile relationship with the aliens, but a certain intrinsic tension is likely. Small human colonies might have been enslaved or otherwise subjugated.",
		Enemies: roll.List{
			Items: []string{
				"Hostile alien chief",
				"Human firebrand",
				"Dangerous local predator",
				"Alien religious zealot",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Colonist leader",
				"Peace-faction alien chief",
				"Planetary frontiersman",
				"Xenoresearcher",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The alien numbers are huge and can overwhelm the humans whenever they so choose",
				"One group is trying to use the other to kill their political opponents",
				"The aliens are incomprehensibly strange",
				"One side commits an atrocity",
			},
		},
		Things: roll.List{
			Items: []string{
				"Alien religious icon",
				"Ancient alien-human treaty",
				"Alien technology",
			},
		},
		Places: roll.List{
			Items: []string{
				"Alien village",
				"Fortified human settlement",
				"Massacre site",
			},
		},
	},
	{
		Name: "Prison Planet",
		Desc: "This planet is or was intended as a prison. Some such prisons were meant for specific malefactors of the Terran Mandate, while others were to contain entire “dangerous” ethnic groups or alien races. Some may still have warden AIs or automatic systems to prevent any unauthorized person from leaving, and any authorization permits have long since expired.",
		Enemies: roll.List{
			Items: []string{
				"Crazed warden AI",
				"Brutal heir to gang leadership",
				"Offworlder who’s somehow acquired warden powers and exploits the locals",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Innocent local born here",
				"Native technician forced to maintain the very tech that imprisons them",
				"Offworlder trapped here by accident",
			},
		},
		Complications: roll.List{
			Items: []string{
				"Departure permits are a precious currency",
				"The prison industry still makes valuable pretech devices",
				"Gangs have metamorphosed into governments",
				"The local nobility descended from the prison staff",
			},
		},
		Things: roll.List{
			Items: []string{
				"A pass to get offworld",
				"A key to bypass ancient security devices",
				"Contraband forbidden by the security scanners",
			},
		},
		Places: roll.List{
			Items: []string{
				"Mandate-era prison block converted to government building",
				"Industrial facility manned by mandatory numbers of prisoners",
				"Makeshift shop where contraband is assembled",
			},
		},
	},
	{
		Name: "Psionics Academy",
		Desc: "This world is one of the few that have managed to redevelop the basics of psychic training. Without this education, a potential psychic is doomed to either madness or death unless they refrain from using their abilities. Psionic academies are rare enough that offworlders are often sent there to study by wealthy patrons. The secrets of psychic mentorship, the protocols and techniques that allow a psychic to successfully train another, are carefully guarded at these academies. Most are closely affiliated with the planetary government.",
		Enemies: roll.List{
			Items: []string{
				"Corrupt psychic instructor",
				"Renegade student",
				"Mad psychic researcher",
				"Resentful townie",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Offworld researcher",
				"Aspiring student",
				"Wealthy tourist",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The academy curriculum kills a significant percentage of students",
				"The faculty use students as research subjects",
				"The students are indoctrinated as sleeper agents",
				"The local natives hate the academy",
				"The academy is part of a religion.",
			},
		},
		Things: roll.List{
			Items: []string{
				"Secretly developed psitech",
				"A runaway psychic mentor",
				"Psychic research prize",
			},
		},
		Places: roll.List{
			Items: []string{
				"Training grounds",
				"Experimental laboratory",
				"School library",
				"Campus hangout",
			},
		},
	},
	{
		Name: "Psionics Fear",
		Desc: "The locals are terrified of psychics. Perhaps their history is studded with feral psychics who went on murderous rampages, or perhaps they simply nurse an unreasoning terror of those “mutant freaks”. Psychics demonstrate their powers at risk of their lives.",
		Enemies: roll.List{
			Items: []string{
				"Mental purity investigator",
				"Suspicious zealot",
				"Witch-finder",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Hidden psychic",
				"Offworlder psychic trapped here",
				"Offworld educator",
			},
		},
		Complications: roll.List{
			Items: []string{
				"Psychic potential is much more common here",
				"Some tech is mistaken as psitech",
				"Natives believe certain rituals and customs can protect them from psychic powers",
			},
		},
		Things: roll.List{
			Items: []string{
				"Hidden psitech cache",
				"Possessions of convicted psychics",
				"Reward for turning in a psychic",
			},
		},
		Places: roll.List{
			Items: []string{
				"Inquisitorial chamber",
				"Lynching site",
				"Museum of psychic atrocities",
			},
		},
	},
	{
		Name: "Psionics Worship",
		Desc: "These natives view psionic powers as a visible gift of god or sign of superiority. If the world has a functional psychic training academy, psychics occupy almost all major positions of power and are considered the natural and proper rulers of the world. If the world lacks training facilities, it is likely a hodgepodge of demented cults, with each one dedicated to a marginally-coherent feral prophet and their psychopathic ravings.",
		Enemies: roll.List{
			Items: []string{
				"Psychic inquisitor",
				"Haughty mind-noble",
				"Psychic slaver",
				"Feral prophet",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Offworlder psychic researcher",
				"Native rebel",
				"Offworld employer seeking psychics",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The psychic training is imperfect",
				"and the psychics all show significant mental illness",
				"The psychics have developed a unique discipline",
				"The will of a psychic is law",
				"Psychics in the party are forcibly kidnapped for “enlightening”.",
			},
		},
		Things: roll.List{
			Items: []string{
				"Ancient psitech",
				"Valuable psychic research records",
				"Permission for psychic training",
			},
		},
		Places: roll.List{
			Items: []string{
				"Psitech-imbued council chamber",
				"Temple to the mind",
				"Sanitarium-prison for feral psychics",
			},
		},
	},
	{
		Name: "Quarantined World",
		Desc: "The world is under a quarantine, and space travel to and from it is strictly forbidden. This may be enforced by massive ground batteries that burn any interlopers from the planet’s sky, or it may be that a neighboring world runs a persistent blockade.",
		Enemies: roll.List{
			Items: []string{
				"Defense installation commander",
				"Suspicious patrol leader",
				"Crazed asteroid hermit",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Relative of a person trapped on the world",
				"Humanitarian relief official",
				"Treasure hunter",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The natives want to remain isolated",
				"The quarantine is enforced by an ancient alien installation",
				"The world is rife with maltech abominations",
				"The blockade is meant to starve everyone on the barren world.",
			},
		},
		Things: roll.List{
			Items: []string{
				"Defense grid key",
				"Bribe for getting someone out",
				"Abandoned alien tech",
			},
		},
		Places: roll.List{
			Items: []string{
				"Bridge of a blockading ship",
				"Defense installation control room",
				"Refugee camp",
			},
		},
	},
	{
		Name: "Radioactive World",
		Desc: "Whether due to a legacy of atomic warfare unhindered by nuke snuffers or a simple profusion of radioactive elements, this world glows in the dark. Even heavy vacc suits can filter only so much of the radiation, and most natives suffer a wide variety of cancers, mutations and other illnesses without the protection of advanced medical treatments.",
		Enemies: roll.List{
			Items: []string{
				"Bitter mutant",
				"Relic warlord",
				"Desperate wouldbe escapee",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Reckless prospector",
				"Offworld scavenger",
				"Biogenetic variety seeker",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The radioactivity is steadily growing worse",
				"The planet’s medical resources break down",
				"The radioactivity has inexplicable effects on living creatures",
				"The radioactivity is the product of a malfunctioning pretech manufactory.",
			},
		},
		Things: roll.List{
			Items: []string{
				"Ancient atomic weaponry",
				"Pretech anti-radioactivity drugs",
				"Untainted water supply",
			},
		},
		Places: roll.List{
			Items: []string{
				"Mutant-infested ruins",
				"Scorched glass plain",
				"Wilderness of bizarre native life",
				"Glowing barrens",
			},
		},
	},
	{
		Name: "Refugees",
		Desc: "The world teems with refugees, either exiles from another planet who managed to get here, or the human detritus of some local conflict that have fled to the remaining stable states. The natives usually regard the refugees with hostility, an attitude returned by many among their unwilling guests.",
		Enemies: roll.List{
			Items: []string{
				"Xenophobic native leader",
				"Refugee chief aspiring to seize the host nation",
				"Politician seeking to use the refugees as a weapon",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Sympathetic refugee waif",
				"Local hard-pressed by refugee gangs",
				"Clergy seeking peace",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The xenophobes are right that the refugees are taking over",
				"The refugees are right that the xenophobes want them out or dead",
				"Both are right",
				"Outside powers are using the refugees to destabilize an enemy government",
				"Refugee and local cultures are extremely incompatible",
			},
		},
		Things: roll.List{
			Items: []string{
				"Treasures brought out by fleeing refugees",
				"Citizenship papers",
				"Cache of vital refugee supplies",
				"Hidden arms for terrorists",
			},
		},
		Places: roll.List{
			Items: []string{
				"Hopeless refugee camp",
				"City swarming with confused strangers",
				"Festival full of angry locals",
			},
		},
	},
	{
		Name: "Regional Hegemon",
		Desc: "This world has the technological sophistication, natural resources, and determined polity necessary to be a regional hegemon for the sector. Nearby worlds are likely either directly subservient to it or tack carefully to avoid its anger. It may even be the capital of a small stellar empire.",
		Enemies: roll.List{
			Items: []string{
				"Ambitious general",
				"Colonial official",
				"Contemptuous noble",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Diplomat",
				"Offworld ambassador",
				"Foreign spy",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The hegemon’s influence is all that’s keeping a murderous war from breaking out on nearby worlds",
				"The hegemon is decaying and losing its control",
				"The government is riddled with spies",
				"The hegemon is genuinely benign",
			},
		},
		Things: roll.List{
			Items: []string{
				"Diplomatic carte blanche",
				"Deed to an offworld estate",
				"Foreign aid grant",
			},
		},
		Places: roll.List{
			Items: []string{
				"Palace or seat of government",
				"Salon teeming with spies",
				"Protest rally",
				"Military base",
			},
		},
	},
	{
		Name: "Restrictive Laws",
		Desc: "A myriad of laws, customs, and rules constrain the inhabitants of this world, and even acts that are completely permissible elsewhere are punished severely here. The locals may provide lists of these laws to offworlders, but few non-natives can hope to master all the important intricacies.",
		Enemies: roll.List{
			Items: []string{
				"Law enforcement officer",
				"Outraged native",
				"Native lawyer specializing in peeling offworlders",
				"Paid snitch",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Frustrated offworlder",
				"Repressed native",
				"Reforming crusader",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The laws change regularly in patterns only natives understand",
				"The laws forbid some action vital to the party",
				"The laws forbid the simple existence of some party members",
				"The laws are secret to offworlders",
			},
		},
		Things: roll.List{
			Items: []string{
				"Complete legal codex",
				"Writ of diplomatic immunity",
				"Fine collection vault contents",
			},
		},
		Places: roll.List{
			Items: []string{
				"Courtroom",
				"Mob scene of outraged locals",
				"Legislative chamber",
				"Police station",
			},
		},
	},
	{
		Name: "Revanchists",
		Desc: "The locals formerly owned another world, or a major nation on the planet formerly owned an additional region of land. Something happened to take away this control or drive out the former rulers, and they’ve never forgotten it. The locals are obsessed with reclaiming their lost lands, and will allow no questions of practicality to interfere with their cause.",
		Enemies: roll.List{
			Items: []string{
				"Demagogue whipping the locals on to a hopeless war",
				"Politician seeking to use the resentment for their own ends",
				"Local convinced the PCs are agents of the “thieving” power",
				"Refugee from the land bitterly demanding it be reclaimed",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Realist local clergy seeking peace",
				"Politician trying to calm the public",
				"Third-party diplomat trying to stamp out the fire",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The revanchists’ claim is completely just and reasonable",
				"The land is now occupied entirely by heirs of the conquerors",
				"Both sides have seized lands the other thinks are theirs",
			},
		},
		Things: roll.List{
			Items: []string{
				"Stock of vital resource produced by the taken land",
				"Relic carried out of it",
				"Proof that the land claim is justified or unjustified",
			},
		},
		Places: roll.List{
			Items: []string{
				"Memorial monument to the loss",
				"Cemetery of those who died in the conquest",
				"Public ceremony commemorating the disaster",
			},
		},
	},
	{
		Name: "Revolutionaries",
		Desc: "The world is convulsed by one or more bands of revolutionaries, with some nations perhaps in the grip of a current revolution. Most of these upheavals can be expected only to change the general flavor of problems in the polity, but the process of getting there usually produces a tremendous amount of suffering.",
		Enemies: roll.List{
			Items: []string{
				"Blood-drenched revolutionary leader",
				"Blooddrenched secret police chief",
				"Hostile foreign agent seeking further turmoil",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Sympathetic victim accused of revolutionary sympathies or government collaboration",
				"Revolutionary or state agent who now repents",
				"Agent of a neutral power that wants peace",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The revolutionaries actually do seem likely to put in better rulers",
				"The revolutionaries are client groups that got out of hand",
				"The revolutionaries are clearly much worse than the government",
				"The revolutionaries have no real ideals beyond power and merely pretend to ideology",
			},
		},
		Things: roll.List{
			Items: []string{
				"List of secret revolutionary sympathizers",
				"Proof of rebel hypocrisy",
				"Confiscated wealth",
			},
		},
		Places: roll.List{
			Items: []string{
				"Festival that explodes into violence",
				"Heavily-fortified police station",
				"Revolutionary base hidden in the wilderness",
			},
		},
	},
	{
		Name: "Rigid Culture",
		Desc: "The local culture is extremely rigid. Certain forms of behavior and belief are absolutely mandated, and any deviation from these principles is punished, or else society may be strongly stratified by birth with limited prospects for change. Anything which threatens the existing social order is feared and shunned.",
		Enemies: roll.List{
			Items: []string{
				"Rigid reactionary",
				"Wary ruler",
				"Regime ideologue",
				"Offended potentate",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Revolutionary agitator",
				"Ambitious peasant",
				"Frustrated merchant",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The cultural patterns are enforced by technological aids",
				"The culture is run by a secret cabal of manipulators",
				"The culture has explicit religious sanction",
				"The culture evolved due to important necessities that have since been forgotten",
			},
		},
		Things: roll.List{
			Items: []string{
				"Precious traditional regalia",
				"Peasant tribute",
				"Opulent treasures of the ruling class",
			},
		},
		Places: roll.List{
			Items: []string{
				"Time-worn palace",
				"Low-caste slums",
				"Bandit den",
				"Reformist temple",
			},
		},
	},
	{
		Name: "Rising Hegemon",
		Desc: "This world is not yet a dominant power in the sector, but it’s well on its way there. Whether through newly-blossoming economic, military, or cultural power, they’re extending their influence over their neighbors and forging new arrangements between their government and the rulers of nearby worlds.",
		Enemies: roll.List{
			Items: []string{
				"Jingoistic supremacist",
				"Official bent on glorious success",
				"Foreign agent saboteur",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Friendly emissary to the benighted",
				"Hardscrabble local turned great success",
				"Foreign visitor seeking contacts or knowledge",
			},
		},
		Complications: roll.List{
			Items: []string{
				"They’re only strong because their neighbors have been weakened",
				"Their success is based on a fluke resource or pretech find",
				"They bitterly resent their neighbors as former oppressors",
			},
		},
		Things: roll.List{
			Items: []string{
				"Tribute shipment",
				"Factory or barracks emblematic of their power source",
				"Tech or data that will deal a blow to their rise",
			},
		},
		Places: roll.List{
			Items: []string{
				"Rustic town being hurled into prosperity",
				"Government building being expanded",
				"Starport struggling under the flow of new ships",
			},
		},
	},
	{
		Name: "Ritual Combat",
		Desc: "The locals favor some form of stylized combat to resolve disputes, provide entertainment, or settle religious differences. This combat is probably not normally lethal unless it’s reserved for a specific disposable class of slaves or professionals. Some combat may involve mastery of esoteric weapons and complex arenas, while other forms might require nothing more than a declaration in the street and a drawn gun.",
		Enemies: roll.List{
			Items: []string{
				"Bloodthirsty local champion",
				"Ambitious gladiator stable owner",
				"Xenophobic master fighter",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Peace-minded foreign missionary",
				"Temperate defender of the weak",
				"Local eager to learn of offworld fighting styles",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The required weapons are strange pretech artifacts",
				"Certain classes are forbidden from fighting and require champions",
				"Loss doesn’t mean death but it does mean ritual scarring or property loss",
			},
		},
		Things: roll.List{
			Items: []string{
				"Magnificent weapon",
				"Secret book of martial techniques",
				"Token signifying immunity to ritual combat challenges",
				"Prize won in bloody battle",
			},
		},
		Places: roll.List{
			Items: []string{
				"Area full of cheering spectators",
				"Dusty street outside a saloon",
				"Memorial for fallen warriors",
			},
		},
	},
	{
		Name: "Robots",
		Desc: "The world has a great many robots on it. Most bots are going to be non-sentient expert systems, though an AI with enough computing resources can control many bots at once, and some worlds may have developed VIs to a degree that individual bots can seem (or be) sentient. Some worlds might even be ruled by metal overlords, ones which do not need to be sentient so long as they have overwhelming force.",
		Enemies: roll.List{
			Items: []string{
				"Hostile robot master",
				"Robot greedy to seize offworld tech",
				"Robot fallen in love with the PC’s ship",
				"Oligarch whose factories build robots",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Data-seeking robot",
				"Plucky young robot tech",
				"Local being pushed out of a job by robots",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The robots are only partially controlled",
				"The robots are salvaged and originally meant for a much darker use",
				"The robots require a rare material that the locals fight over",
				"The robots require the planet’s specific infrastructure so cannot be exported",
			},
		},
		Things: roll.List{
			Items: []string{
				"Prototype robot",
				"Secret robot override codes",
				"Vast cache of robot-made goods",
				"Robot-destroying pretech weapon",
			},
		},
		Places: roll.List{
			Items: []string{
				"Humming robotic factory",
				"Stark robotic “barracks”",
				"House crowded with robot servants and only one human owner",
			},
		},
	},
	{
		Name: "Seagoing Cities",
		Desc: "Either the world is entirely water or else the land is simply too dangerous for most humans. Human settlement on this world consists of a number of floating cities that follow the currents and the fish. These city-ships might have been purpose-built for their task, or they could be jury-rigged conglomerations of ships and structures thrown together when the need for seagoing life become apparent to the locals.",
		Enemies: roll.List{
			Items: []string{
				"Pirate city lord",
				"Mer-human raider chieftain",
				"Hostile landsman noble",
				"Enemy city saboteur",
			},
		},
		Friends: roll.List{
			Items: []string{
				"City navigator",
				"Scout captain",
				"Curious mer-human",
				"Hard-pressed ship-city engineer",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The seas are not water",
				"The fish schools have vanished and the city faces starvation",
				"Terrible storms drive the city into the glacial regions",
				"Suicide ships ram the city’s hull",
			},
		},
		Things: roll.List{
			Items: []string{
				"Giant pearls with mysterious chemical properties",
				"Buried treasure",
				"Vital repair materials",
			},
		},
		Places: roll.List{
			Items: []string{
				"Bridge of the city",
				"Storm-tossed sea",
				"A bridge fashioned of many small boats.",
			},
		},
	},
	{
		Name: "Sealed Menace",
		Desc: "Something on this planet has the potential to create enormous havoc for the inhabitants if it is not kept safely contained by its keepers. Whether a massive seismic fault line suppressed by pretech terraforming technology, a disease that has to be quarantined within hours of discovery, or an ancient alien relic that requires regular upkeep in order to prevent planetary catastrophe, the menace is a constant shadow on the fearful populace.",
		Enemies: roll.List{
			Items: []string{
				"Hostile outsider bent on freeing the menace",
				"Misguided fool who thinks he can use it",
				"Reckless researcher who thinks he can fix it",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Keeper of the menace",
				"Student of its nature",
				"Victim of the menace",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The menace would bring great wealth along with destruction",
				"The menace is intelligent",
				"The natives don’t all believe in the menace",
			},
		},
		Things: roll.List{
			Items: []string{
				"A key to unlock the menace",
				"A precious byproduct of the menace",
				"The secret of the menace’s true nature",
			},
		},
		Places: roll.List{
			Items: []string{
				"Guarded fortress containing the menace",
				"Monitoring station",
				"Scene of a prior outbreak of the menace",
			},
		},
	},
	{
		Name: "Secret Masters",
		Desc: "The world is actually run by a hidden cabal, acting through their catspaws in the visible government. For one reason or another, this group finds it imperative that they not be identified by outsiders, and in some cases even the planet’s own government may not realize that they’re actually being manipulated by hidden masters.",
		Enemies: roll.List{
			Items: []string{
				"An agent of the cabal",
				"Government official who wants no questions asked",
				"Willfully blinded local",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Paranoid conspiracy theorist",
				"Machiavellian gamesman within the cabal",
				"Interstellar investigator",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The secret masters have a benign reason for wanting secrecy",
				"The cabal fights openly amongst itself",
				"The cabal is recruiting new members",
			},
		},
		Things: roll.List{
			Items: []string{
				"A dossier of secrets on a government official",
				"A briefcase of unmarked credit notes",
				"The identity of a cabal member",
			},
		},
		Places: roll.List{
			Items: []string{
				"Smoke-filled room",
				"Shadowy alleyway",
				"Secret underground bunker",
			},
		},
	},
	{
		Name: "Sectarians",
		Desc: "The world is torn by violent disagreement between sectarians of a particular faith. Each views the other as a damnable heresy in need of extirpation. Local government may be able to keep open war from breaking out, but the poisonous hatred divides communities. The nature of the faith may be religious, or it may be based on some secular ideology.",
		Enemies: roll.List{
			Items: []string{
				"Paranoid believer",
				"Native convinced the party is working for the other side",
				"Absolutist ruler",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Reformist clergy",
				"Local peacekeeping official",
				"Offworld missionary",
				"Exhausted ruler",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The conflict has more than two sides",
				"The sectarians hate each other for multiple reasons",
				"The sectarians must cooperate or else life on this world is imperiled",
				"The sectarians hate outsiders more than they hate each other",
				"The differences in sects are incomprehensible to an outsider",
			},
		},
		Things: roll.List{
			Items: []string{
				"Ancient holy book",
				"Incontrovertible proof",
				"Offering to a local holy man",
			},
		},
		Places: roll.List{
			Items: []string{
				"Sectarian battlefield",
				"Crusading temple",
				"Philosopher’s salon",
				"Bitterly divided village",
			},
		},
	},
	{
		Name: "Seismic Instability",
		Desc: "The local land masses are remarkably unstable, and regular earthquakes rack the surface. Local construction is either advanced enough to sway and move with the vibrations or primitive enough that it is easily rebuilt. Severe volcanic activity may be part of the instability.",
		Enemies: roll.List{
			Items: []string{
				"Earthquake cultist",
				"Hermit seismologist",
				"Burrowing native life form",
				"Earthquake-inducing saboteur",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Experimental construction firm owner",
				"Adventurous volcanologist",
				"Geothermal prospector",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The earthquakes are caused by malfunctioning pretech terraformers",
				"They’re caused by alien technology",
				"They’re restrained by alien technology that is being plundered by offworlders",
				"The earthquakes are used to generate enormous amounts of energy.",
			},
		},
		Things: roll.List{
			Items: []string{
				"Earthquake generator",
				"Earthquake suppressor",
				"Mineral formed at the core of the world",
				"Earthquake-proof building schematics",
			},
		},
		Places: roll.List{
			Items: []string{
				"Volcanic caldera",
				"Village during an earthquake",
				"Mud slide",
				"Earthquake opening superheated steam fissures",
			},
		},
	},
	{
		Name: "Shackled World",
		Desc: "This world is being systematically contained by an outside power. Some ancient autonomous defense grid, robot law enforcement, alien artifact, or other force is preventing the locals from developing certain technology, or using certain devices, or perhaps from developing interstellar flight. This limit may or may not apply to offworlders; in the former case, the PCs may have to figure out a way to beat the shackles simply to escape the world.",
		Enemies: roll.List{
			Items: []string{
				"Passionless jailer-AI",
				"Paranoid military grid AI",
				"Robot overlord",
				"Enigmatic alien master",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Struggling local researcher",
				"Offworlder trapped here",
				"Scientist with a plan to break the chains",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The shackles come off for certain brief windows of time",
				"The locals think the shackles are imposed by God",
				"An outside power greatly profits from the shackles",
				"The rulers are exempt from the shackles",
			},
		},
		Things: roll.List{
			Items: []string{
				"Keycode to bypass the shackle",
				"Tech shielded from the shackle",
				"Exportable version of the shackle that can affect other worlds",
			},
		},
		Places: roll.List{
			Items: []string{
				"Grim high-tech control center",
				"Factory full of workaround tech",
				"Temple to the power or entity that imposed the shackle",
			},
		},
	},
	{
		Name: "Societal Despair",
		Desc: "The world’s dominant society has lost faith in itself. Whether through some all-consuming war, great catastrophe, overwhelming outside culture, or religious collapse, the natives no longer believe in their old values, and search desperately for something new. Fierce conflict often exists between the last believers in the old dispensation and the nihilistic or searching disciples of the new age.",
		Enemies: roll.List{
			Items: []string{
				"Zealot who blames outsiders for the decay",
				"Nihilistic warlord",
				"Offworlder looking to exploit the local despair",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Struggling messenger of a new way",
				"Valiant paragon of a fading tradition",
				"Local going through the motions of serving a now-irrelevant role",
			},
		},
		Complications: roll.List{
			Items: []string{
				"A massive war discredited all the old values",
				"Outside powers are working to erode societal confidence for their own benefit",
				"A local power is profiting greatly from the despair",
				"The old ways were meant to aid survival on this world and their passing is causing many new woes",
			},
		},
		Things: roll.List{
			Items: []string{
				"Relic that would inspire a renaissance",
				"Art that would inspire new ideas",
				"Priceless artifact of a now-scorned belief",
			},
		},
		Places: roll.List{
			Items: []string{
				"Empty temple",
				"Crowded den of obliviating vice",
				"Smoky hall full of frantic speakers",
			},
		},
	},
	{
		Name: "Sole Supplier",
		Desc: "Some extremely important resource is exported from this world and this world alone. It’s unlikely that the substance is critical for building spike drives unless this world is also the first to begin interstellar flight, but it may be critical to other high-tech processes or devices. The locals make a large amount of money off this trade and control of it is of critical importance to the planet’s rulers, and potentially to outside powers.",
		Enemies: roll.List{
			Items: []string{
				"Resource oligarch",
				"Ruthless smuggler",
				"Resource-controlling warlord",
				"Foreign agent seeking to subvert local government",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Doughty resource miner",
				"Researcher trying to synthesize the stuff",
				"Small-scale resource producer",
				"Harried starport trade overseer",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The substance is slow poison to process",
				"The substance is created by hostile alien natives",
				"The substance is very easy to smuggle in usable amounts",
				"Only the natives have the genes or tech to extract it effectively",
			},
		},
		Things: roll.List{
			Items: []string{
				"Cache of processed resource",
				"Trade permit to buy a load of it",
				"A shipment of nigh-undetectably fake substance",
			},
		},
		Places: roll.List{
			Items: []string{
				"Bustling resource extraction site",
				"Opulent palace built with resource money",
				"Lazy town square where everyone lives on resource payments",
			},
		},
	},
	{
		Name: "Taboo Treasure",
		Desc: "The natives here produce something that is both fabulously valuable and strictly forbidden elsewhere in the sector. It may be a lethally addictive drug, forbidden gengineering tech, vat-grown “perfect slaves”, or a useful substance that can only be made through excruciating human suffering. This treasure is freely traded on the world, but bringing it elsewhere is usually an invitation to a long prison stay or worse.",
		Enemies: roll.List{
			Items: []string{
				"Maker of a vile commodity",
				"Smuggler for a powerful offworlder",
				"Depraved offworlder here for “fun”",
				"Local warlord who controls the treasure",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Reformer seeking to end its use",
				"Innovator trying to repurpose the treasure in innocent ways",
				"Wretched addict unwillingly prey to the treasure",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The treasure is extremely hard to smuggle",
				"Its use visibly marks a user",
				"The natives consider it for their personal use only",
			},
		},
		Things: roll.List{
			Items: []string{
				"Load of the forbidden good",
				"Smuggling tech that could hide the good perfectly",
				"Blackmail data on offworld buyers of the good",
			},
		},
		Places: roll.List{
			Items: []string{
				"Den where the good is used",
				"Market selling the good to locals and a few outsiders",
				"Factory or processing area where the good is created",
			},
		},
	},
	{
		Name: "Terraform Failure",
		Desc: "This world was marginal for human habitation when it was discovered, but the Mandate or the early government put in pretech terraforming engines to correct its more extreme qualities. The terraforming did not entirely work, either failing of its own or suffering the destruction of the engines during the Silence. The natives are only partly adapted to the world’s current state, and struggle with the environment.",
		Enemies: roll.List{
			Items: []string{
				"Brutal ruler who cares only for their people",
				"Offworlder trying to loot the damaged engines",
				"Warlord trying to seize limited habitable land",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Local trying to fix the engines",
				"Offworlder student of the engines",
				"World-wise native survivor",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The engines produced too much of something instead of too little",
				"The engines were hijacked by aliens with different preferences",
				"It was discovered that an Earth-like environment would eventually cause a catastrophic disaster",
			},
		},
		Things: roll.List{
			Items: []string{
				"Parts to repair or restore the engines",
				"Lootable pretech fragments",
				"Valuable local tech devised to cope with the world",
			},
		},
		Places: roll.List{
			Items: []string{
				"Zone of tolerable gravity or temperature",
				"Native settlement built to cope with the environment",
				"Massive ruined terraforming engine",
			},
		},
	},
	{
		Name: "Theocracy",
		Desc: "The planet is ruled by the priesthood of the predominant religion or ideology. The rest of the locals may or may not be terribly pious, but the clergy have the necessary military strength, popular support or control of resources to maintain their rule. Alternative faiths or incompatible ideologies are likely to be both illegal and socially unacceptable.",
		Enemies: roll.List{
			Items: []string{
				"Decadent priest-ruler",
				"Zealous inquisitor",
				"Relentless proselytizer",
				"True Believer",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Heretic",
				"Offworld theologian",
				"Atheistic merchant",
				"Desperate commoner",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The theocracy actually works well",
				"The theocracy is decadent and hated by the common folk",
				"The theocracy is divided into mutually hostile sects",
				"The theocracy is led by aliens",
			},
		},
		Things: roll.List{
			Items: []string{
				"Precious holy text",
				"Martyr’s bones",
				"Secret church records",
				"Ancient church treasures",
			},
		},
		Places: roll.List{
			Items: []string{
				"Glorious temple",
				"Austere monastery",
				"Academy for ideological indoctrination",
				"Decadent pleasure-cathedral",
			},
		},
	},
	{
		Name: "Tomb World",
		Desc: "Tomb worlds are planets that were once inhabited by humans before the Silence. The sudden collapse of the jump gate network and the inability to bring in the massive food supplies required by the planet resulted in starvation, warfare, and death. Most tomb worlds are naturally hostile to human habitation and could not raise sufficient crops to maintain life. The few hydroponic facilities were usually destroyed in the fighting, and all that is left now are ruins, bones, and silence.",
		Enemies: roll.List{
			Items: []string{
				"Demented survivor tribe chieftain",
				"Avaricious scavenger",
				"Automated defense system",
				"Native predator",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Scavenger Fleet captain",
				"Archaeologist",
				"Salvaging historian",
				"Xenophilic native survivor",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The ruins are full of booby-traps left by the final inhabitants",
				"The world’s atmosphere quickly degrades anything in an opened building",
				"A handful of desperate natives survived the Silence",
				"The structures are unstable and collapsing",
			},
		},
		Things: roll.List{
			Items: []string{
				"Lost pretech equipment",
				"Tech caches",
				"Stores of unused munitions",
				"Ancient historical data",
			},
		},
		Places: roll.List{
			Items: []string{
				"Crumbling hive-city",
				"City square carpeted in bones",
				"Ruined hydroponic facility",
				"Cannibal tribe’s lair",
				"Dead orbital jump gate",
			},
		},
	},
	{
		Name: "Trade Hub",
		Desc: "This world is a major crossroads for local interstellar trade. It is well-positioned at the nexus of several short-drill trade routes, and has facilities for easy transfer of valuable cargoes and the fueling and repairing of starships. The natives are accustomed to outsiders, and a polyglot mass of people from every nearby world can be found trading here.",
		Enemies: roll.List{
			Items: []string{
				"Cheating merchant",
				"Thieving dockworker",
				"Commercial spy",
				"Corrupt customs official",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Rich tourist",
				"Hardscrabble free trader",
				"Merchant prince in need of catspaws",
				"Friendly spaceport urchin",
			},
		},
		Complications: roll.List{
			Items: []string{
				"An outworlder faction schemes to seize the trade hub",
				"Saboteurs seek to blow up a rival’s warehouses",
				"Enemies are blockading the trade routes",
				"Pirates lace the hub with spies",
			},
		},
		Things: roll.List{
			Items: []string{
				"Voucher for a warehouse’s contents",
				"Insider trading information",
				"Case of precious offworld pharmaceuticals",
				"Box of legitimate tax stamps indicating customs dues have been paid.",
			},
		},
		Places: roll.List{
			Items: []string{
				"Raucous bazaar",
				"Elegant restaurant",
				"Spaceport teeming with activity",
				"Foggy street lined with warehouses",
			},
		},
	},
	{
		Name: "Tyranny",
		Desc: "The local government is brutal and indifferent to the will of the people. Laws may or may not exist, but the only one that matters is the whim of the rulers on any given day. Their minions swagger through the streets while the common folk live in terror of their appetites. The only people who stay wealthy are friends and servants of the ruling class.",
		Enemies: roll.List{
			Items: []string{
				"Debauched autocrat",
				"Sneering bully-boy",
				"Soulless government official",
				"Occupying army officer",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Conspiring rebel",
				"Oppressed merchant",
				"Desperate peasant",
				"Inspiring religious leader",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The tyrant rules with vastly superior technology",
				"The tyrant is a figurehead for a cabal of powerful men and women",
				"The people are resigned to their suffering",
				"The tyrant is hostile to “meddlesome outworlders”.",
			},
		},
		Things: roll.List{
			Items: []string{
				"Plundered wealth",
				"Beautiful toys of the elite",
				"Regalia of rulership",
			},
		},
		Places: roll.List{
			Items: []string{
				"Impoverished village",
				"Protest rally massacre",
				"Decadent palace",
				"Religious hospital for the indigent",
			},
		},
	},
	{
		Name: "Unbraked AI",
		Desc: "Artificial intelligences are costly and difficult to create, requiring a careful sequence of “growth stages” in order to bring them to sentience before artificial limits on cognition speed and learning development are installed. These “brakes” prevent runaway cognition metastasis. This world has an “unbraked AI” on it, probably with a witting or unwitting corps of servants. Unbraked AIs are quite insane, but they learn and reason with a speed impossible for humans, and can demonstrate a truly distressing subtlety.",
		Enemies: roll.List{
			Items: []string{
				"AI Cultist",
				"Maltech researcher",
				"Government official dependent on the AI",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Perimeter agent",
				"AI researcher",
				"Braked AI",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The AI’s presence is unknown to the locals",
				"The locals depend on the AI for some vital service",
				"The AI appears to be harmless",
				"The AI has fixated on the group’s ship’s computer",
				"The AI wants transport offworld",
			},
		},
		Things: roll.List{
			Items: []string{
				"The room-sized AI core itself",
				"Maltech research files",
				"Perfectly tabulated blackmail on government officials",
				"Pretech computer circuitry",
			},
		},
		Places: roll.List{
			Items: []string{
				"Municipal computing banks",
				"Cult compound",
				"Repair center",
				"Ancient hardcopy library",
			},
		},
	},
	{
		Name: "Urbanized Surface",
		Desc: "The world’s land area is covered with buildings that extend downward for multiple levels. Such worlds either have a population in the trillions, extremely little land area, or are largely-abandoned due to some past catastrophe. Agriculture and resource extraction are part of the urban complex, and there may be an advanced maintenance system that may not be entirely under the control of present natives.",
		Enemies: roll.List{
			Items: []string{
				"Maintenance AI that hates outsiders",
				"Tyrant of a habitation block",
				"Deep-dwelling prophet who considers “the sky” a blasphemy to be quelled",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Local yearning for wild spaces",
				"Grubby urchin of the underlevels",
				"Harried engineer trying to maintain ancient works",
				"Grizzled hab cop",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The urban blocks are needed to survive the environment",
				"The blocks were part of an ancient device of world-spanning size",
				"The blocks require constant maintenance to avoid dangerous types of decay",
			},
		},
		Things: roll.List{
			Items: []string{
				"Massively efficient power source",
				"Map of the secret ways of a zone",
				"Passkey into restricted hab block areas",
			},
		},
		Places: roll.List{
			Items: []string{
				"Giant hab block now devoid of inhabitants",
				"Chemical-reeking underway",
				"Seawater mine full of salt and massive flowing channels",
			},
		},
	},
	{
		Name: "Utopia",
		Desc: "Natural and social conditions on this world have made it a paradise for its inhabitants, a genuine utopia of happiness and fulfillment. This is normally the result of drastic human engineering, including brain-gelding, neurochemical control, personality curbs, or complete “humanity” redefinitions. Even so, the natives are extremely happy with their lot, and may wish to extend that joy to poor, sad outsiders.",
		Enemies: roll.List{
			Items: []string{
				"Compassionate neurotherapist",
				"Proselytizing native missionary to outsiders",
				"Brutal tyrant who rules through inexorable happiness",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Deranged malcontent",
				"Bloody-handed guerrilla leader of a rebellion of madmen",
				"Outsider trying to find a way to reverse the utopian changes",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The natives really are deeply and contentedly happy with their altered lot",
				"The utopia produces something that attracts others",
				"The utopia works on converting outsiders through persuasion and generosity",
				"The utopia involves some sacrifice that’s horrifying to non-members",
			},
		},
		Things: roll.List{
			Items: []string{
				"Portable device that applies the utopian change",
				"Plans for a device that would destroy the utopia",
				"Goods created joyfully by the locals",
			},
		},
		Places: roll.List{
			Items: []string{
				"Plaza full of altered humans",
				"Social ritual site",
				"Secret office where “normal” humans rule",
			},
		},
	},
	{
		Name: "Warlords",
		Desc: "The world is plagued by warlords. Numerous powerful men and women control private armies sufficiently strong to cow whatever local government may exist. On the lands they claim, their word is law. Most spend their time oppressing their own subjects and murderously pillaging those of their neighbors. Most like to wrap themselves in the mantle of ideology, religious fervor, or an ostensibly legitimate right to rule.",
		Enemies: roll.List{
			Items: []string{
				"Warlord",
				"Avaricious lieutenant",
				"Expensive assassin",
				"Aspiring minion",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Vengeful commoner",
				"Government military officer",
				"Humanitarian aid official",
				"Village priest",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The warlords are willing to cooperate to fight mutual threats",
				"The warlords favor specific religions or races over others",
				"The warlords are using substantially more sophisticated tech than others",
				"Some of the warlords are better rulers than the government",
			},
		},
		Things: roll.List{
			Items: []string{
				"Weapons cache",
				"Buried plunder",
				"A warlord’s personal battle harness",
				"Captured merchant shipping",
			},
		},
		Places: roll.List{
			Items: []string{
				"Gory battlefield",
				"Burnt-out village",
				"Barbaric warlord palace",
				"Squalid refugee camp",
			},
		},
	},
	{
		Name: "Xenophiles",
		Desc: "The natives of this world are fast friends with a particular alien race. The aliens may have saved the planet at some point in the past, or awed the locals with superior tech or impressive cultural qualities. The aliens might even be the ruling class on the planet.",
		Enemies: roll.List{
			Items: []string{
				"Offworld xenophobe",
				"Suspicious alien leader",
				"Xenocultural imperialist",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Benevolent alien",
				"Native malcontent",
				"Gone-native offworlder",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The enthusiasm is due to alien psionics or tech",
				"The enthusiasm is based on a lie",
				"The aliens strongly dislike their “groupies”",
				"The aliens feel obliged to rule humanity for its own good",
				"Humans badly misunderstand the aliens",
			},
		},
		Things: roll.List{
			Items: []string{
				"Hybrid alien-human tech",
				"Exotic alien crafts",
				"Sophisticated xenolinguistic and xenocultural research data",
			},
		},
		Places: roll.List{
			Items: []string{
				"Alien district",
				"Alien-influenced human home",
				"Cultural festival celebrating alien artist",
			},
		},
	},
	{
		Name: "Xenophobes",
		Desc: "The natives are intensely averse to dealings with outworlders. Whether through cultural revulsion, fear of tech contamination, or a genuine immunodeficiency, the locals shun foreigners from offworld and refuse to have anything to do with them beyond the bare necessities of contact. Trade may or may not exist on this world, but if it does, it is almost certainly conducted by a caste of untouchables and outcasts.",
		Enemies: roll.List{
			Items: []string{
				"Revulsed local ruler",
				"Native convinced some wrong was done to him",
				"Cynical demagogue",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Curious native",
				"Exiled former ruler",
				"Local desperately seeking outworlder help",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The natives are symptomless carriers of a contagious and dangerous disease",
				"The natives are exceptionally vulnerable to offworld diseases",
				"The natives require elaborate purification rituals after speaking to an offworlder or touching them",
				"The local ruler has forbidden any mercantile dealings with outworlders",
			},
		},
		Things: roll.List{
			Items: []string{
				"Jealously-guarded precious relic",
				"Local product under export ban",
				"Esoteric local technology",
			},
		},
		Places: roll.List{
			Items: []string{
				"Sealed treaty port",
				"Public ritual not open to outsiders",
				"Outcaste slum home",
			},
		},
	},
	{
		Name: "Zombies",
		Desc: "This menace may not take the form of shambling corpses, but some disease, alien artifact, or crazed local practice produces men and women with habits similar to those of murderous cannibal undead. These outbreaks may be regular elements in local society, either provoked by some malevolent creators or the consequence of some local condition.",
		Enemies: roll.List{
			Items: []string{
				"Soulless maltech biotechnology cult",
				"Sinister governmental agent",
				"Crazed zombie cultist",
			},
		},
		Friends: roll.List{
			Items: []string{
				"Survivor of an outbreak",
				"Doctor searching for a cure",
				"Rebel against the secret malefactors",
			},
		},
		Complications: roll.List{
			Items: []string{
				"The zombies retain human intelligence",
				"The zombies can be cured",
				"The process is voluntary among devotees",
				"The condition is infectious",
			},
		},
		Things: roll.List{
			Items: []string{
				"Cure for the condition",
				"Alien artifact that causes it",
				"Details of the cult’s conversion process",
			},
		},
		Places: roll.List{
			Items: []string{
				"House with boarded-up windows",
				"Dead city",
				"Fortified bunker that was overrun from within",
			},
		},
	},
}
