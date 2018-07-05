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
	heresy := Heresy{
		Founder:     Founder.Roll(),
		MajorHeresy: MajorHeresy.Roll(),
		Attitude:    Attitude.Roll(),
		Quirk:       Quirk.Roll(),
	}
	return heresy
}

func (heresy Heresy) String() string {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, "%s\t:\t%s\n", Founder.Name, heresy.Founder)
	fmt.Fprintf(buf, "%s\t:\t%s\n", MajorHeresy.Name, heresy.MajorHeresy)
	fmt.Fprintf(buf, "%s\t:\t%s\n", Attitude.Name, heresy.Attitude)
	fmt.Fprintf(buf, "%s\t:\t%s", Quirk.Name, heresy.Quirk)
	return buf.String()
}
