package alien

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Alien with a Body
type Alien struct {
	Body            string
	Lense           string
	SocialStructure string
}

// New Alien with random characteristics
func New() Alien {
	alien := Alien{
		Body:            Body.Roll(),
		Lense:           Lense.Roll(),
		SocialStructure: SocialStructure.Roll(),
	}
	return alien
}

func (alien Alien) String() string {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, "%s\t:\t%s\n", Body.Name, alien.Body)
	fmt.Fprintf(buf, "%s\t:\t%s\n", Lense.Name, alien.Lense)
	fmt.Fprintf(buf, "%s\t:\t%s\n", SocialStructure.Name, alien.SocialStructure)
	return buf.String()
}
