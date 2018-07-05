package heresy

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Heresy with a Body
type Heresy struct {
	Founder     string
	MajorHeresy string
	Attitude    string
	Quirk       string
}

// New Heresy with random characteristics
func New() Heresy {
	h := Heresy{
		Founder:     Founder.Roll(),
		MajorHeresy: MajorHeresy.Roll(),
		Attitude:    Attitude.Roll(),
		Quirk:       Quirk.Roll(),
	}
	return h
}

func (h Heresy) String() string {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, "%s\t:\t%s\n", Founder.Name, h.Founder)
	fmt.Fprintf(buf, "%s\t:\t%s\n", MajorHeresy.Name, h.MajorHeresy)
	fmt.Fprintf(buf, "%s\t:\t%s\n", Attitude.Name, h.Attitude)
	fmt.Fprintf(buf, "%s\t:\t%s", Quirk.Name, h.Quirk)
	return buf.String()
}
