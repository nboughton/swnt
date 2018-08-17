package heresy

import (
	"bytes"
	"fmt"

	"github.com/nboughton/swnt/content/format"
)

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

// Format return Heresy h as type t
func (h Heresy) Format(t format.OutputType) string {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, format.Table(t, true, "Heresy", [][]string{
		{Founder.Name, h.Founder},
		{MajorHeresy.Name, h.MajorHeresy},
		{Attitude.Name, h.Attitude},
		{Quirk.Name, h.Quirk},
	}))

	return buf.String()
}
