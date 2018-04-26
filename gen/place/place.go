package place

import (
	"bytes"
	"fmt"
)

type Place struct {
	Reward   string
	Ongoings string
	Hazard   string
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

func (p Place) String() string {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, p.Hazard)
	fmt.Fprintf(buf, "Ongoings\t:\t%s\n", p.Ongoings)
	fmt.Fprintf(buf, "Reward\t:\t%s\n", p.Reward)

	return buf.String()
}
