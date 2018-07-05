package alien

import (
	"bytes"
	"fmt"
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

func (a Alien) String() string {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, "%s\t:\t%s\n", Body.Name, a.Body)
	fmt.Fprintf(buf, "%s\t:\t%s\n", Lense.Name, a.Lense)
	fmt.Fprintf(buf, "%s\t:\t%s\n", SocialStructure.Name, a.SocialStructure)
	return buf.String()
}
