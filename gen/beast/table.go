package beast

import (
	"strings"

	"github.com/nboughton/rollt"
)

var reg = rollt.NewRegistry()

func init() {
	reg.Add(BasicAnimalFeatures)
}

// BasicAnimalFeatures is pretty fucking self-explanatory
var BasicAnimalFeatures = rollt.Table{
	Name: "Basic Animal Features",
	ID:   "beast.BasicAnimalFeatures",
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
			tbl, _ := reg.Get("beast.BasicAnimalFeatures")
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
}
