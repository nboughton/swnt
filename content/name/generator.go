package name

import (
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/nboughton/go-roll"
)

var badPrefix = regexp.MustCompile(`[cflmnr][^aeiouyh]`)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Name generator. This is a primitive first effort that will be refined over time

// Generate creates a random name by combining alternating vowels and consonants
func Generate(ln int) string {
	name := ""
	for i := rand.Intn(2); len(name) <= ln; i++ {
		if i%2 != 0 {
			c := con.Roll()

			if name == "" {
				for badPrefix.MatchString(c) { // I don't like starting a name with these
					c = con.Roll()
				}
			}

			name += c
		} else {
			name += vl.Roll()
		}
	}

	return strings.ToUpper(string(name[0])) + string(name[1:])
}

var vl = roll.Table{
	Name: "vowels",
	Dice: "10d5",
	Items: []roll.TableItem{
		{Match: []int{10}, Text: "ii"},
		{Match: []int{11}, Text: "yu"},
		{Match: []int{12}, Text: "uy"},
		{Match: []int{13}, Text: "oy"},
		{Match: []int{14}, Text: "ao"},
		{Match: []int{15}, Text: "ye"},
		{Match: []int{16}, Text: "ae"},
		{Match: []int{17}, Text: "oe"},
		{Match: []int{18}, Text: "eo"},
		{Match: []int{19}, Text: "oi"},
		{Match: []int{20}, Text: "ua"},
		{Match: []int{21}, Text: "au"},
		{Match: []int{22}, Text: "ia"},
		{Match: []int{23}, Text: "ey"},
		{Match: []int{24}, Text: "oo"},
		{Match: []int{25}, Text: "io"},
		{Match: []int{26}, Text: "ea"},
		{Match: []int{27}, Text: "y"},
		{Match: []int{28}, Text: "o"},
		{Match: []int{29}, Text: "a"},
		{Match: []int{30}, Text: "e"},
		{Match: []int{31}, Text: "i"},
		{Match: []int{32}, Text: "u"},
		{Match: []int{33}, Text: "ou"},
		{Match: []int{34}, Text: "ee"},
		{Match: []int{35}, Text: "ai"},
		{Match: []int{36}, Text: "ie"},
		{Match: []int{37}, Text: "ei"},
		{Match: []int{38}, Text: "ue"},
		{Match: []int{39}, Text: "ay"},
		{Match: []int{40}, Text: "ui"},
		{Match: []int{41}, Text: "oa"},
		{Match: []int{42}, Text: "yi"},
		{Match: []int{43}, Text: "ya"},
		{Match: []int{44}, Text: "eu"},
		{Match: []int{45}, Text: "iu"},
		{Match: []int{46}, Text: "yo"},
		{Match: []int{47}, Text: "aa"},
		{Match: []int{48}, Text: "uo"},
		{Match: []int{49}, Text: "uu"},
		{Match: []int{50}, Text: "'"},
	},
}

var con = roll.Table{
	Name: "consonants",
	Dice: "10d5",
	Items: []roll.TableItem{
		{Match: []int{10}, Text: "tt"},
		{Match: []int{11}, Text: "rr"},
		{Match: []int{12}, Text: "ct"},
		{Match: []int{13}, Text: "pr"},
		{Match: []int{14}, Text: "ns"},
		{Match: []int{15}, Text: "bl"},
		{Match: []int{16}, Text: "sh"},
		{Match: []int{17}, Text: "ld"},
		{Match: []int{18}, Text: "k"},
		{Match: []int{19}, Text: "nd"},
		{Match: []int{20}, Text: "ll"},
		{Match: []int{21}, Text: "nt"},
		{Match: []int{22}, Text: "st"},
		{Match: []int{23}, Text: "f"},
		{Match: []int{24}, Text: "ng"},
		{Match: []int{25}, Text: "w"},
		{Match: []int{26}, Text: "th"},
		{Match: []int{27}, Text: "m"},
		{Match: []int{28}, Text: "n"},
		{Match: []int{29}, Text: "d"},
		{Match: []int{30}, Text: "r"},
		{Match: []int{31}, Text: "s"},
		{Match: []int{32}, Text: "t"},
		{Match: []int{33}, Text: "l"},
		{Match: []int{34}, Text: "c"},
		{Match: []int{35}, Text: "v"},
		{Match: []int{36}, Text: "b"},
		{Match: []int{37}, Text: "p"},
		{Match: []int{38}, Text: "h"},
		{Match: []int{39}, Text: "g"},
		{Match: []int{40}, Text: "wh"},
		{Match: []int{41}, Text: "ch"},
		{Match: []int{42}, Text: "ss"},
		{Match: []int{43}, Text: "rs"},
		{Match: []int{44}, Text: "nc"},
		{Match: []int{45}, Text: "fr"},
		{Match: []int{46}, Text: "rt"},
		{Match: []int{47}, Text: "gr"},
		{Match: []int{48}, Text: "rd"},
		{Match: []int{49}, Text: "sp"},
		{Match: []int{50}, Text: "ck"},
	},
}
