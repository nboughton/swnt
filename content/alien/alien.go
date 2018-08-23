package alien

import (
	"bytes"
	"fmt"

	"github.com/nboughton/swnt/content/format"
)

// Alien with a Body
type Alien struct {
	Body            string
	Lense           string
	SocialStructure string
}

// New Alien with random characteristics
func New() Alien {
	a := Alien{
		Body:            Body.Roll(),
		Lense:           Lense.Roll(),
		SocialStructure: SocialStructure.Roll(),
	}
	return a
}

// Format returns Alien a formatted as type t
func (a Alien) Format(t format.OutputType) string {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, format.Table(t, true, "Alien", [][]string{
		{Body.Name, a.Body},
		{Lense.Name, a.Lense},
		{SocialStructure.Name, a.SocialStructure},
	}))

	return buf.String()
}

func (a Alien) String() string {
	return a.Format(format.TEXT)
}
