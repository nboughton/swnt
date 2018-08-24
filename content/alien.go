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
	table.Registry.Add(alienTable.body)
	table.Registry.Add(alienTable.socialStructure)
}

// Alien with a Body
type Alien struct {
	Body            string
	Lense           string
	SocialStructure string
}

// NewAlien with random characteristics
func NewAlien() Alien {
	a := Alien{
		Body:            alienTable.body.Roll(),
		Lense:           alienTable.lense.Roll(),
		SocialStructure: alienTable.socialStructure.Roll(),
	}
	return a
}

// Format returns Alien a formatted as type t
func (a Alien) Format(t format.OutputType) string {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, format.Table(t, true, "Alien", [][]string{
		{alienTable.body.Name, a.Body},
		{alienTable.lense.Name, a.Lense},
		{alienTable.socialStructure.Name, a.SocialStructure},
	}))

	return buf.String()
}

func (a Alien) String() string {
	return a.Format(format.TEXT)
}

var alienTable = struct {
	body            rollt.Table
	lense           rollt.Table
	socialStructure rollt.Table
}{
	// Body SWN Revised Free Edition p203
	rollt.Table{
		Name: "Body",
		ID:   "alien.Body",
		Dice: "1d6",
		Items: []rollt.Item{
			{Match: []int{1}, Text: "Avian, bat-like, pterodactylian"},
			{Match: []int{2}, Text: "Reptilian, amphibian, draconic"},
			{Match: []int{3}, Text: "Insectile, beetle-like, spiderish, wasp-like"},
			{Match: []int{4}, Text: "Mammalian, furred or bare-skinned"},
			{Match: []int{5}, Text: "Exotic, composed of some novel substance"},
			{Match: []int{6}, Text: "Hybrid of two or more types", Action: func() string {
				tbl, _ := table.Registry.Get("alien.Body")
				tbl.Dice = "1d5"

				types, res := 2+rand.Intn(4), make(map[string]bool)
				for len(res) < types {
					res[tbl.Roll()] = true
				}

				text := []string{}
				for k := range res {
					text = append(text, k)
				}

				return "\n\t\t" + strings.Join(text, "\n\t\t")
			}},
		},
	},

	// Lense from SWN Revised Free Edition p205
	rollt.Table{
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
	},

	// SocialStructure from SWN Revised Free Edition p207
	rollt.Table{
		Name: "Social Structure",
		ID:   "alien.SocialStructure",
		Dice: "1d8",
		Items: []rollt.Item{
			{Match: []int{1}, Text: "Democratic"},
			{Match: []int{2}, Text: "Monarchic"},
			{Match: []int{3}, Text: "Tribal"},
			{Match: []int{4}, Text: "Oligarchic"},
			{Match: []int{5, 6}, Text: "Multipolar Competitive", Action: actionSocialStructure},
			{Match: []int{7, 8}, Text: "Multipolar Cooperative", Action: actionSocialStructure},
		},
	},
}

func actionSocialStructure() string {
	tbl, _ := table.Registry.Get("alien.SocialStructure")
	tbl.Dice = "1d4"

	types, res := 2+rand.Intn(3), make(map[string]bool)
	for len(res) < types {
		res[tbl.Roll()] = true
	}

	text := []string{}
	for k := range res {
		text = append(text, k)
	}

	return strings.Join(text, ", ")
}
