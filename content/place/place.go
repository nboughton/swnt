package place

import (
	"bytes"
	"fmt"

	"github.com/nboughton/swnt/content/format"
)

// Place represents the aggregate details of a generated place
type Place struct {
	Reward   string
	Ongoings string
	Hazard   [][]string
}

// New roll a new place, default to urban
func New(wilderness bool) Place {
	og := ""
	if wilderness {
		og = Ongoings.Wilderness.Roll()
	} else {
		og = Ongoings.Civilised.Roll()
	}

	return Place{
		Reward:   Reward.Roll(),
		Ongoings: og,
		Hazard:   Hazard.Roll(),
	}
}

// Format returns Place formatted as type t
func (p Place) Format(t format.OutputType) string {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, format.Table(format.TEXT, true, "Place", p.Hazard))
	fmt.Fprintf(buf, format.Table(format.TEXT, false, "", [][]string{
		{"Ongoings", p.Ongoings},
		{"Reward", p.Reward},
	}))

	return buf.String()
}
