package content

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"

	"github.com/nboughton/rollt"
	"github.com/nboughton/swnt/content/format"
	"github.com/nboughton/swnt/content/table"
)

func init() {
	table.Registry.Add(beastFeaturesTable.basicFeatures)
	table.Registry.Add(beastFeaturesTable.bodyPlan)
}

// Beast defines the aggregate descriptors for an animal
type Beast struct {
	Type        string
	Behaviour   string
	Features    string
	BodyPlan    string
	LimbNovelty string
	SkinNovelty string
	MainWeapon  string
	Size        string
}

// NewBeast for terrorising players
func NewBeast() Beast {
	b := Beast{
		Features:    beastFeaturesTable.basicFeatures.Roll(),
		BodyPlan:    beastFeaturesTable.bodyPlan.Roll(),
		LimbNovelty: beastFeaturesTable.limbNovelty.Roll(),
		SkinNovelty: beastFeaturesTable.skinNovelty.Roll(),
		MainWeapon:  beastFeaturesTable.mainWeapon.Roll(),
		Size:        beastFeaturesTable.size.Roll(),
	}

	switch rand.Intn(3) {
	case 0:
		b.Type = beastBehaviourTable.predator.Label()
		b.Behaviour = beastBehaviourTable.predator.Roll()

	case 1:
		b.Type = beastBehaviourTable.prey.Label()
		b.Behaviour = beastBehaviourTable.prey.Roll()

	case 2:
		b.Type = beastBehaviourTable.scavenger.Label()
		b.Behaviour = beastBehaviourTable.scavenger.Roll()
	}

	return b
}

// Format output as format type t
func (b Beast) Format(t format.OutputType) string {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, format.Table(t, []string{"Beast", ""}, [][]string{
		{b.Type, b.Behaviour},
		{beastFeaturesTable.basicFeatures.Label(), b.Features},
		{beastFeaturesTable.bodyPlan.Label(), b.BodyPlan},
		{beastFeaturesTable.limbNovelty.Label(), b.LimbNovelty},
		{beastFeaturesTable.skinNovelty.Label(), b.SkinNovelty},
		{beastFeaturesTable.mainWeapon.Label(), b.MainWeapon},
		{beastFeaturesTable.size.Label(), b.Size},
	}))

	return buf.String()
}

func (b Beast) String() string {
	return b.Format(format.TEXT)
}

var beastFeaturesTable = struct {
	basicFeatures rollt.Table
	bodyPlan      rollt.Table
	limbNovelty   rollt.Table
	skinNovelty   rollt.Table
	mainWeapon    rollt.Table
	size          rollt.Table
}{
	// BasicAnimalFeatures is pretty fucking self-explanatory
	rollt.Table{
		Name: "Basic Features",
		ID:   "beast.BasicFeatures",
		Dice: "1d10",
		Items: []rollt.Item{
			{Match: []int{1}, Text: "Amphibian, froggish or newtlike"},
			{Match: []int{2}, Text: "Bird, winged and feathered"},
			{Match: []int{3}, Text: "Fish, scaled and torpedo-bodied"},
			{Match: []int{4}, Text: "Insect, beetle-like or fly-winged"},
			{Match: []int{5}, Text: "Mammal, hairy and fanged"},
			{Match: []int{6}, Text: "Reptile, lizardlike and long-bodied"},
			{Match: []int{7}, Text: "Spider, many-legged and fat"},
			{Match: []int{8}, Text: "Exotic, made of wholly alien elements"},
			{Match: []int{9, 10}, Text: "Mixed", Action: func() string {
				tbl, _ := table.Registry.Get("beast.BasicFeatures")
				tbl.Dice = "1d8"

				res := make(map[string]bool)
				for len(res) < 2 {
					res[tbl.Roll()] = true
				}

				text := []string{}
				for k := range res {
					text = append(text, k)
				}

				return strings.Join(text, " and ")
			}},
		},
	},

	// BodyPlan p201 SWN:RE Free edition
	rollt.Table{
		Name: "Body Plan",
		ID:   "beast.BodyPlan",
		Dice: "1d6",
		Items: []rollt.Item{
			{Match: []int{1}, Text: "Humanoid"},
			{Match: []int{2}, Text: "Quadruped"},
			{Match: []int{3}, Text: "Many-legged"},
			{Match: []int{4}, Text: "Bulbous"},
			{Match: []int{5}, Text: "Amorphous"},
			{Match: []int{6}, Text: "", Action: func() string {
				tbl, _ := table.Registry.Get("beast.BodyPlan")
				tbl.Dice = "1d5"

				types, res := 2, make(map[string]bool)
				for len(res) < types {
					res[tbl.Roll()] = true
				}

				text := []string{}
				for k := range res {
					text = append(text, k)
				}

				return strings.Join(text, " and ")
			}},
		},
	},

	// LimbNovelty p201 SWN:RE Free edition
	rollt.Table{
		Name: "Limb Novelty",
		ID:   "beast.LimbNovelty",
		Dice: "1d6",
		Items: []rollt.Item{
			{Match: []int{1}, Text: "Wings"},
			{Match: []int{2}, Text: "Many joints"},
			{Match: []int{3}, Text: "Tentacles"},
			{Match: []int{4}, Text: "Opposable thumbs"},
			{Match: []int{5}, Text: "Retractable"},
			{Match: []int{6}, Text: "Varying sizes"},
		},
	},

	// SkinNovelty p201 SWN:RE Free edition
	rollt.Table{
		Name: "Skin Novelty",
		ID:   "beast.SkinNovelty",
		Dice: "1d6",
		Items: []rollt.Item{
			{Match: []int{1}, Text: "Hard shell"},
			{Match: []int{2}, Text: "Exoskeleton"},
			{Match: []int{3}, Text: "Odd texture"},
			{Match: []int{4}, Text: "Molts regularly"},
			{Match: []int{5}, Text: "Harmful to touch"},
			{Match: []int{6}, Text: "Wet or slimy"},
		},
	},

	// MainWeapon p201 SWN:RE Free edition
	rollt.Table{
		Name: "Main Weapon",
		ID:   "beast.MainWeapon",
		Dice: "1d6",
		Items: []rollt.Item{
			{Match: []int{1}, Text: "Teeth or mandibles"},
			{Match: []int{2}, Text: "Claws"},
			{Match: []int{3}, Text: "Poison", Action: poisonAction},
			{Match: []int{4}, Text: "Harmful discharge", Action: func() string {
				return harmfulDischargesTable.Roll()
			}},
			{Match: []int{5}, Text: "Pincers"},
			{Match: []int{6}, Text: "Horns"},
		},
	},

	// Size p201 SWN:RE Free edition
	rollt.Table{
		Name: "Size",
		ID:   "beast.Size",
		Dice: "1d6",
		Items: []rollt.Item{
			{Match: []int{1}, Text: "Cat-sized"},
			{Match: []int{2}, Text: "Wolf-sized"},
			{Match: []int{3}, Text: "Calf-sized"},
			{Match: []int{4}, Text: "Bull-sized"},
			{Match: []int{5}, Text: "Hippo-sized"},
			{Match: []int{6}, Text: "Elephant-sized"},
		},
	},
}

// Behavioural traits

// Predator p201 SWN:RE Free edition
var beastBehaviourTable = struct {
	predator  rollt.Table
	prey      rollt.Table
	scavenger rollt.Table
}{
	rollt.Table{
		Name: "Predator",
		ID:   "beast.Predator",
		Dice: "1d8",
		Items: []rollt.Item{
			{Match: []int{1}, Text: "Hunts in kin-group packs"},
			{Match: []int{2}, Text: "Favors ambush attacks"},
			{Match: []int{3}, Text: "Cripples prey and waits for death"},
			{Match: []int{4}, Text: "Pack supports alpha-beast attack"},
			{Match: []int{5}, Text: "Lures or drives prey into danger"},
			{Match: []int{6}, Text: "Hunts as a lone, powerful hunter"},
			{Match: []int{7}, Text: "Only is predator at certain times"},
			{Match: []int{8}, Text: "Mindlessly attacks humans"},
		},
	},

	// Prey p201 SWN:RE Free edition
	rollt.Table{
		Name: "Prey",
		ID:   "beast.Prey",
		Dice: "1d8",
		Items: []rollt.Item{
			{Match: []int{1}, Text: "Moves in vigilant herds"},
			{Match: []int{2}, Text: "Exists in small family groups"},
			{Match: []int{3}, Text: "They all team up on a single foe"},
			{Match: []int{4}, Text: "They go berserk when near death"},
			{Match: []int{5}, Text: "They’re violent in certain seasons"},
			{Match: []int{6}, Text: "They’re vicious if threatened"},
			{Match: []int{7}, Text: "Symbiotic creature protects them"},
			{Match: []int{8}, Text: "Breeds at tremendous rates"},
		},
	},

	// Scavenger p201 SWN:RE Free edition
	rollt.Table{
		Name: "Scavenger",
		ID:   "beast.Scavenger",
		Dice: "1d8",
		Items: []rollt.Item{
			{Match: []int{1}, Text: "Never attacks unwounded prey"},
			{Match: []int{2}, Text: "Uses other beasts as harriers"},
			{Match: []int{3}, Text: "Always flees if significantly hurt"},
			{Match: []int{4}, Text: "Poisons prey, waits for it to die"},
			{Match: []int{5}, Text: "Disguises itself as its prey"},
			{Match: []int{6}, Text: "Remarkably stealthy"},
			{Match: []int{7}, Text: "Summons predators to weak prey"},
			{Match: []int{8}, Text: "Steals prey from weaker predator"},
		},
	},
}

// HarmfulDischarges p201 SWN:RE Free edition
var harmfulDischargesTable = rollt.Table{
	Name: "HarmfulDischarges",
	ID:   "beast.HarmfulDischarges",
	Dice: "1d8",
	Items: []rollt.Item{
		{Match: []int{1}, Text: "Acidic spew doing its damage on a hit"},
		{Match: []int{2}, Text: "Toxic spittle or cloud; ", Action: poisonAction},
		{Match: []int{3}, Text: "Super-heated or super-chilled spew"},
		{Match: []int{4}, Text: "Sonic drill or other disabling noise"},
		{Match: []int{5}, Text: "Natural laser or plasma discharge"},
		{Match: []int{6}, Text: "Nauseating stench or disabling chemical"},
		{Match: []int{7}, Text: "Equipment-melting corrosive"},
		{Match: []int{8}, Text: "Explosive pellets or chemical catalysts"},
	},
}

// Poisons

// Poison p201 SWN:RE Free edition
var poisonTable = struct {
	effect   rollt.Table
	onset    rollt.Table
	duration rollt.Table
}{
	rollt.Table{
		Name: "Poison",
		ID:   "beast.Poison",
		Dice: "1d6",
		Items: []rollt.Item{
			{Match: []int{1}, Text: "Death"},
			{Match: []int{2}, Text: "Paralysis"},
			{Match: []int{3}, Text: "1d4 dmg per onset interval"},
			{Match: []int{4}, Text: "Convulsions"},
			{Match: []int{5}, Text: "Blindness"},
			{Match: []int{6}, Text: "Hallucinations"},
		},
	},
	rollt.Table{
		Name: "Onset",
		ID:   "beast.Onset",
		Dice: "1d6",
		Items: []rollt.Item{
			{Match: []int{1}, Text: "Instant"},
			{Match: []int{2}, Text: "1 round"},
			{Match: []int{3}, Text: "1d6 rounds"},
			{Match: []int{4}, Text: "1 minute"},
			{Match: []int{5}, Text: "1d6 minutes"},
			{Match: []int{6}, Text: "1 hour"},
		},
	},
	rollt.Table{
		Name: "Duration",
		ID:   "beast.Duration",
		Dice: "1d6",
		Items: []rollt.Item{
			{Match: []int{1}, Text: "1d6 rounds"},
			{Match: []int{2}, Text: "1 minute"},
			{Match: []int{3}, Text: "10 minutes"},
			{Match: []int{4}, Text: "1 hour"},
			{Match: []int{5}, Text: "1d6 hours"},
			{Match: []int{6}, Text: "1d6 days"},
		},
	},
}

func poisonAction() string {
	return fmt.Sprintf("in %s the target suffers from %s over %s.",
		poisonTable.onset.Roll(),
		poisonTable.effect.Roll(),
		poisonTable.duration.Roll())
}
