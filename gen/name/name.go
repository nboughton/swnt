package name

import (
	"github.com/nboughton/rollt"
	"github.com/nboughton/swnt/gen/culture"
)

// Table represents the collection of roll lists keyed by cultural background
type Table struct {
	Culture culture.ID
	Male    rollt.List
	Female  rollt.List
	Surname rollt.List
	Place   rollt.List
}

// Tables represents a set of Table structs
type Tables []Table

// ByCulture returns a name table that matches the given culture
func (t Tables) ByCulture(c culture.ID) Table {
	for _, tbl := range t {
		if tbl.Culture == c {
			return tbl
		}
	}

	return Table{}
}
