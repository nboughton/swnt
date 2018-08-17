package world

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/nboughton/rollt"
	"github.com/nboughton/swnt/content/culture"
	"github.com/nboughton/swnt/content/format"
	"github.com/nboughton/swnt/content/name"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// TagsTable represents the collection of Tags
type TagsTable []Tag

// Roll selects a random Tag
func (t TagsTable) Roll() string {
	return fmt.Sprint(t[rand.Intn(len(t))])
}

// Random selects a random tag (used in Adventure seed generation)
func (t TagsTable) Random() string {
	return Tags[rand.Intn(len(Tags))].Name
}

// Find returns the tag specified. The search is case insensitive for convenience
func (t TagsTable) Find(name string) (Tag, error) {
	for _, tag := range t {
		if strings.ToLower(tag.Name) == strings.ToLower(name) {
			return tag, nil
		}
	}

	return Tag{}, fmt.Errorf("no tag with name \"%s\"", name)
}

func selectTags(exclude []string) (Tag, Tag) {
	var t TagsTable
	for _, tag := range Tags {
		if !tag.match(exclude) {
			t = append(t, tag)
		}
	}

	t1Idx, t2Idx := rand.Intn(len(t)), rand.Intn(len(t))
	for t1Idx == t2Idx { // Ensure the same tag isn't selected twice
		t2Idx = rand.Intn(len(t))
	}

	return t[t1Idx], t[t2Idx]
}

func (t Tag) match(s []string) bool {
	for _, x := range s {
		if strings.ToLower(t.Name) == strings.ToLower(x) {
			return true
		}
	}

	return false
}

// Tag represents a complete World Tag structure as extracted from the Stars Without Number core book.
type Tag struct {
	Name          string
	Desc          string
	Enemies       rollt.List
	Friends       rollt.List
	Complications rollt.List
	Things        rollt.List
	Places        rollt.List
}

func (t Tag) String() string {
	return fmt.Sprintf(
		"Name\t:\t%s\nDesc\t:\t%s\nEnemies\t:\t%s\nFriends\t:\t%s\nComplications\t:\t%s\nThings\t:\t%s\nPlaces\t:\t%s\n",
		t.Name, t.Desc, t.Enemies, t.Friends, t.Complications, t.Things, t.Places,
	)
}

// World represents a generated world
type World struct {
	Primary      bool
	Name         string
	Culture      culture.Culture
	Tags         [2]Tag
	Atmosphere   string
	Temperature  string
	Population   string
	Biosphere    string
	TechLevel    string
	Origin       string
	Relationship string
	Contact      string
}

// New World, set culture to culture.Any for a random culture and primary to false
// to include relationship information
func New(c culture.Culture, primary bool, exclude []string) World {
	t1, t2 := selectTags(exclude)

	w := World{
		Primary:     primary,
		Name:        name.Names.ByCulture(c).Place.Roll(),
		Culture:     c,
		Tags:        [2]Tag{t1, t2},
		Atmosphere:  Atmosphere.Roll(),
		Temperature: Temperature.Roll(),
		Population:  Population.Roll(),
		Biosphere:   Biosphere.Roll(),
		TechLevel:   TechLevel.Roll(),
	}

	if !w.Primary {
		w.Origin = Other.Origin.Roll()
		w.Relationship = Other.Relationship.Roll()
		w.Contact = Other.Contact.Roll()
	}

	return w
}

// Format returns the content of World w in format t
func (w World) Format(t format.OutputType) string {
	var buf = new(bytes.Buffer)

	fmt.Fprintf(buf, format.Table(t, true, w.Name, [][]string{
		{"Atmosphere", w.Atmosphere},
		{"Temperature", w.Temperature},
		{"Biosphere", w.Biosphere},
		{"Population", w.Population},
		{"Culture", string(w.Culture)},
		{"Tech Level", w.TechLevel},
		{"Tags", ""},
		{w.Tags[0].Name, w.Tags[0].Desc},
		{w.Tags[1].Name, w.Tags[1].Desc},
	}))

	if !w.Primary {
		fmt.Fprintf(buf, format.Table(t, false, "", [][]string{
			{"Origins", ""},
			{Other.Origin.Name, w.Origin},
			{Other.Relationship.Name, w.Relationship},
			{Other.Contact.Name, w.Contact},
		}))
	}

	return buf.String()
}
