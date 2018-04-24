package culture

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// ID represents the id int of a statically defined culture
type ID int

// Culture constants
const (
	Arabic ID = iota
	Chinese
	English
	Greek
	Indian
	Japanese
	Latin
	Nigerian
	Russian
	Spanish
	Any
)

// Cultures list
var Cultures = []string{"Arabic", "Chinese", "English", "Greek", "Indian", "Japanese", "Latin", "Nigerian", "Russian", "Spanish"}

// Random returns a cultures' identifier and string name at random
func Random() (ID, string) {
	n := rand.Intn(len(Cultures))
	return ID(n), Cultures[n]
}

// IDByName returns the appropriate ID for string s
func IDByName(s string) (ID, error) {
	for id, name := range Cultures {
		if name == s {
			return ID(id), nil
		}
	}

	if s == "Any" {
		id, _ := Random()
		return id, nil
	}

	return ID(0), fmt.Errorf("No culture found for %s", s)
}

// NameByID returns the name of a Culture by its ID
func NameByID(i ID) (string, error) {
	if int(i) < len(Cultures) {
		return Cultures[i], nil
	}

	if i == Any {
		_, str := Random()
		return str, nil
	}

	return "", fmt.Errorf("No culture found for %d", i)
}
