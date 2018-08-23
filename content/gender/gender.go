// Package gender is a convenience package for handling gender identifiers.
package gender

/* Why only 3 options?

Gender as a spectrum is an ever growing list of self identifying labels that no developer
could ever hope to keep up with. That said the number of people that identify as non-binary
is a very, very small subset of the population. I'm not trying to marginalise anyone that
doesn't like the term "other" for non-binary genders. I just don't have the time or spoons
to attempt to cater to them.
*/

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Gender is a shorthand type for IDing general labels
type Gender string

// Gender ids
const (
	Male   Gender = "Male"
	Female Gender = "Female"
	Other  Gender = "Other"
	Any    Gender = "Any"
)

// Genders supported
var Genders = []Gender{Male, Female, Other}

// Random returns a random Gender
func Random() Gender {
	n := rand.Intn(len(Genders))
	return Genders[n]
}

// Find returns the id constant or an error if it doesn't exist
func Find(name string) (Gender, error) {
	if strings.ToLower(name) == strings.ToLower(Any.String()) || name == "" {
		return Random(), nil
	}

	for _, g := range Genders {
		if strings.ToLower(g.String()) == strings.ToLower(name) {
			return g, nil
		}
	}

	return Gender(""), fmt.Errorf("no gender found for \"%s\", options available are %s", name, Genders)
}

func (g Gender) String() string {
	return string(g)
}
