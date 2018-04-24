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

type GenderID byte

// GenderID ids
const (
	Male        GenderID = 'm'
	Female      GenderID = 'f'
	Androgynous GenderID = 'a'
	Any         GenderID = 'y'
)

var gender = map[GenderID]string{
	Male:        "Male",
	Female:      "Female",
	Androgynous: "Andro",
	Any:         "Any",
}

// NPC represents an NPC
type NPC struct {
	Name       string
	Gender     string
	Culture    string
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
func New(c culture.ID, g GenderID, isPatron bool) NPC {
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

	nm := name.Names.ByCulture(c)
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

	n.Culture, _ = culture.NameByID(c)

	return n
}

func (n NPC) String() string {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, "Their Name\t:\t%s\n", n.Name)
	fmt.Fprintf(buf, "Their Culture\t:\t%s\n", n.Culture)
	fmt.Fprint(buf, n.NPC)
	fmt.Fprintln(buf, "\t")
	fmt.Fprintf(buf, "Initial Manner\t:\t%s\n", n.Manner)
	fmt.Fprintf(buf, "Default Deal Outcome\t:\t%s\n", n.Outcome)
	fmt.Fprintf(buf, "Their Motivation\t:\t%s\n", n.Motivation)
	fmt.Fprintf(buf, "Their Want\t:\t%s\n", n.Want)
	fmt.Fprintf(buf, "Their Power\t:\t%s\n", n.Power)
	fmt.Fprintf(buf, "Their Hook\t:\t%s\n", n.Hook)

	fmt.Fprintln(buf, "\t")

	if n.IsPatron {
		fmt.Fprintf(buf, n.Patron)
		fmt.Fprintln(buf)
	}

	return buf.String()
}
