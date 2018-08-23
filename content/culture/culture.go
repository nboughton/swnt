package culture

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Culture represents the name of supported culture
type Culture string

// Culture constants
const (
	Arabic   Culture = "Arabic"
	Chinese  Culture = "Chinese"
	English  Culture = "English"
	Greek    Culture = "Greek"
	Indian   Culture = "Indian"
	Japanese Culture = "Japanese"
	Latin    Culture = "Latin"
	Nigerian Culture = "Nigerian"
	Russian  Culture = "Russian"
	Spanish  Culture = "Spanish"
	Any      Culture = "Any"
)

// Cultures list
var Cultures = []Culture{Arabic, Chinese, English, Greek, Indian, Japanese, Latin, Nigerian, Russian, Spanish}

// Random returns a cultures' identifier and string name at random
func Random() Culture {
	n := rand.Intn(len(Cultures))
	return Cultures[n]
}

// Find returns the correct constant or an error if it does not exist
func Find(name string) (Culture, error) {
	if strings.ToLower(name) == strings.ToLower(Any.String()) || name == "" {
		return Random(), nil
	}

	for _, c := range Cultures {
		if strings.ToLower(c.String()) == strings.ToLower(name) {
			return c, nil
		}
	}

	return Culture(""), fmt.Errorf("no culture found for \"%s\", options available are %s", name, Cultures)
}

func (c Culture) String() string {
	return string(c)
}
