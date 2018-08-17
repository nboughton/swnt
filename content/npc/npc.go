package npc

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"

	"github.com/nboughton/swnt/content/culture"
	"github.com/nboughton/swnt/content/format"
	"github.com/nboughton/swnt/content/name"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Gender is a shorthand type for IDing general labels
type Gender byte

// Gender ids
const (
	Male        Gender = 'm'
	Female      Gender = 'f'
	Androgynous Gender = 'a'
	Any         Gender = 'y'
)

var gender = map[Gender]string{
	Male:        "Male",
	Female:      "Female",
	Androgynous: "Andro",
	Any:         "Any",
}

// NPC represents an NPC
type NPC struct {
	Name       string
	Gender     string
	Culture    culture.Culture
	Manner     string
	Outcome    string
	Motivation string
	Want       string
	Power      string
	Hook       string
	NPC        [][]string
	Reaction   string
	IsPatron   bool
	Patron     [][]string
}

// New roll a new NPC
func New(ctr culture.Culture, g Gender, isPatron bool) NPC {
	n := NPC{
		IsPatron:   isPatron,
		Manner:     Manner.Roll(),
		Outcome:    Outcome.Roll(),
		Motivation: Motivation.Roll(),
		Want:       Want.Roll(),
		Power:      Power.Roll(),
		Hook:       Hook.Roll(),
		Reaction:   Reaction.Roll(),
		NPC:        NPCTable.Roll(),
	}

	if n.IsPatron {
		n.Patron = PatronTable.Roll()
	}

	nm := name.Names.ByCulture(ctr)
	switch g {
	case Male:
		n.Name = fmt.Sprintf("%s %s", nm.Male.Roll(), nm.Surname.Roll())
	case Female:
		n.Name = fmt.Sprintf("%s %s", nm.Female.Roll(), nm.Surname.Roll())
	case Androgynous, Any:
		switch rand.Intn(2) {
		case 0:
			n.Name = fmt.Sprintf("%s %s", nm.Male.Roll(), nm.Surname.Roll())
		case 1:
			n.Name = fmt.Sprintf("%s %s", nm.Female.Roll(), nm.Surname.Roll())
		}
	}

	if g, ok := gender[g]; ok {
		n.Gender = g
	}

	n.Culture = ctr

	return n
}

// Format returns a string output in the specified format t
func (n NPC) Format(t format.OutputType) string {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, format.Table(t, true, n.Name, [][]string{{"Culture", string(n.Culture)}}))
	fmt.Fprintf(buf, format.Table(t, false, "", n.NPC))
	fmt.Fprintf(buf, format.Table(t, false, "", [][]string{
		{"", ""},
		{Manner.Name, n.Manner},
		{Outcome.Name, n.Outcome},
		{Motivation.Name, n.Motivation},
		{Want.Name, n.Want},
		{Power.Name, n.Power},
		{Hook.Name, n.Hook},
		{Reaction.Name, n.Reaction},
	}))

	if n.IsPatron {
		fmt.Fprintln(buf)
		fmt.Fprintf(buf, format.Table(t, true, "Patron", n.Patron))
	}

	return buf.String()
}
