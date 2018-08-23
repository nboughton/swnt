package name

import (
	"math/rand"
	"strings"

	"github.com/nboughton/rollt"
)

// Name generator. This is a primitive first effort that will be refined over time
var vl = rollt.Table{
	Name: "vowels",
	Dice: "3d12",
	Items: []rollt.Item{
		{Match: []int{3}, Text: "y"},
		{Match: []int{4}, Text: "ya"},
		{Match: []int{5}, Text: "ue"},
		{Match: []int{6}, Text: "ey"},
		{Match: []int{7}, Text: "ei"},
		{Match: []int{8}, Text: "ao"},
		{Match: []int{9}, Text: "ae"},
		{Match: []int{10}, Text: "ai"},
		{Match: []int{11}, Text: "eo"},
		{Match: []int{12}, Text: "iu"},
		{Match: []int{14}, Text: "oo"},
		{Match: []int{15}, Text: "ie"},
		{Match: []int{16}, Text: "ea"},
		{Match: []int{17}, Text: "a"},
		{Match: []int{18}, Text: "e"},
		{Match: []int{19}, Text: "i"},
		{Match: []int{20}, Text: "o"},
		{Match: []int{21}, Text: "u"},
		{Match: []int{13}, Text: "oa"},
		{Match: []int{22}, Text: "ou"},
		{Match: []int{23}, Text: "io"},
		{Match: []int{24}, Text: "ia"},
		{Match: []int{25}, Text: "oi"},
		{Match: []int{26}, Text: "ay"},
		{Match: []int{27}, Text: "oe"},
		{Match: []int{28}, Text: "oy"},
		{Match: []int{29}, Text: "au"},
		{Match: []int{30}, Text: "ui"},
		{Match: []int{31}, Text: "uo"},
		{Match: []int{32}, Text: "ua"},
		{Match: []int{33}, Text: "eu"},
		{Match: []int{34}, Text: "yo"},
		{Match: []int{35}, Text: "yi"},
		{Match: []int{36}, Text: "yu"},
	},
}

var con = rollt.List{
	Name: "consonants",
	Items: []string{
		"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "z",
		"bh", "bl", "br",
		"ch", "cl", "cr",
		"dh", "dr",
		"fl", "fr",
		"gh", "gl", "gr", "gn",
		"kh", "kl", "kr", "kn",
		"nk",
		"ph", "pl", "pr",
		"qu",
		"sh", "sl", "st",
		"th"},
}

var badSuffix = []string{
	"bl", "br", "cl", "cr", "dh", "dr", "fl", "fr", "gl", "gr", "gn", "iw", "kl", "kr", "kn",
	"pl", "pr", "qu", "sl",
}

// Generate creates a random name by combining alternating vowels and consonants
func Generate(ln int) string {
	name := ""
	for i := rand.Intn(2); len(name) <= ln; i++ {
		if i%2 != 0 {
			name += con.Roll()
		} else {
			name += vl.Roll()
		}
	}

	for _, suf := range badSuffix {
		if string(name[len(name)-2:]) == suf {
			name = string(name[:len(name)-2])
		}
	}

	return strings.ToUpper(string(name[0])) + string(name[1:])
}
