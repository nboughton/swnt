package npc

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"

	"github.com/nboughton/swnt/gen/culture"
	"github.com/nboughton/swnt/gen/name"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

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
	NPC        string
	IsPatron   bool
	Patron     string
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

func (n NPC) String() string {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, "Name\t:\t%s\n", n.Name)
	fmt.Fprintf(buf, "Culture\t:\t%s\n", n.Culture)
	fmt.Fprint(buf, n.NPC)
	fmt.Fprintln(buf, "\t")
	fmt.Fprintf(buf, "%s\t:\t%s\n", Manner.Name, n.Manner)
	fmt.Fprintf(buf, "%s\t:\t%s\n", Outcome.Name, n.Outcome)
	fmt.Fprintf(buf, "%s\t:\t%s\n", Motivation.Name, n.Motivation)
	fmt.Fprintf(buf, "%s\t:\t%s\n", Want.Name, n.Want)
	fmt.Fprintf(buf, "%s\t:\t%s\n", Power.Name, n.Power)
	fmt.Fprintf(buf, "%s\t:\t%s\n", Hook.Name, n.Hook)

	fmt.Fprintln(buf, "\t")

	if n.IsPatron {
		fmt.Fprintf(buf, n.Patron)
		fmt.Fprintln(buf)
	}

	return buf.String()
}
