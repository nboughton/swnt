package sector

import (
	"bytes"
	"fmt"
	"math/rand"
	"sort"
	"text/tabwriter"
	"time"

	"github.com/nboughton/swnt/gen/culture"
	"github.com/nboughton/swnt/gen/name"
	"github.com/nboughton/swnt/gen/poi"
	"github.com/nboughton/swnt/gen/world"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Star represents a single Star on the Sector map
type Star struct {
	Row, Col int
	Culture  culture.Culture
	Name     string
	Worlds   []world.World
	POIs     []poi.POI
}

// NewStar generates a new Star struct to be added to the map
func NewStar(r, c int, n string, e []string, poiChance, otherWorldChance int) *Star {
	ctr := culture.Random()

	s := &Star{
		Row:     r,
		Col:     c,
		Culture: ctr,
		Name:    n,
		Worlds:  []world.World{world.New(ctr, true, e)},
	}

	// Cascading 10% chance of other worlds
	for rand.Intn(100) < otherWorldChance {
		ctr = culture.Random()
		s.Worlds = append(s.Worlds, world.New(ctr, false, e))
	}

	// 30% chance of a Point of Interest
	if rand.Intn(100) < poiChance {
		s.POIs = append(s.POIs, poi.New())
	}

	return s
}

// String satisfies the stringer interface for a Star
func (s *Star) String() string {
	var (
		buf = new(bytes.Buffer)
		tw  = tabwriter.NewWriter(buf, 1, 1, 2, ' ', 0)
	)

	fmt.Fprintf(tw, "Coords\t:\t%d,%d\n", s.Row, s.Col)
	fmt.Fprintf(tw, "Name\t:\t%s\n", s.Name)
	fmt.Fprintln(tw, "\t")
	fmt.Fprintf(tw, "Primary World\t\n")
	fmt.Fprintf(tw, s.Worlds[0].String())

	for i, w := range s.Worlds[1:] {
		fmt.Fprintf(tw, "World %d\t\n", i+2)
		fmt.Fprintf(tw, w.String())
	}

	for _, p := range s.POIs {
		fmt.Fprintln(tw, "\t")
		fmt.Fprintf(tw, "Point of Interest\t\n")
		fmt.Fprintf(tw, p.String())
	}

	fmt.Fprintln(tw)
	tw.Flush()

	return buf.String()
}

// Stars represents the generated collection of Stars that will be used to populate a hex grid
type Stars struct {
	Rows, Cols int
	Systems    []*Star
}

type byCoord []*Star

func (b byCoord) Len() int { return len(b) }
func (b byCoord) Less(i, j int) bool {
	if b[i].Row == b[j].Row {
		return b[i].Col < b[j].Col
	}

	return b[i].Row < b[j].Row
}
func (b byCoord) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

// ByCoords returns a sorted array of star systems ordered by coordinates.
func (s *Stars) ByCoords() []*Star {
	b := byCoord(s.Systems)
	sort.Sort(b)
	return b
}

// NewSector returns a blank Sector struct and generates tag information according to the guidelines
// in pages 133 - 177 of Stars Without Number (Revised Edition).
func NewSector(excludeTags []string, poiChance, otherWorldChance int) *Stars {
	s := &Stars{
		Rows: 8,
		Cols: 10,
	}

	cells := s.Rows * s.Cols
	stars := (rand.Intn(cells/4) / 2) + (cells / 4)

	for r, c := rand.Intn(s.Rows), rand.Intn(s.Cols); len(s.Systems) <= stars; r, c = rand.Intn(s.Rows), rand.Intn(s.Cols) {
		if !s.IsActive(r, c) {
			s.Systems = append(s.Systems, NewStar(r, c, s.unusedSystemName(), excludeTags, poiChance, otherWorldChance))
		}
	}

	return s
}

// UniqueName ensures rolls on the name.System table until it gets a name that is not currently in use.
func (s *Stars) unusedSystemName() string {
	n := name.System.Roll()
	for {
		if !s.isNameUsed(n) {
			return n
		}

		n = name.System.Roll()
	}
}

func (s *Stars) isNameUsed(n string) bool {
	for _, star := range s.Systems {
		if star.Name == n {
			return true
		}
	}

	return false
}

// IsActive checks to see if there is an active Star at r(ow), c(ol)
func (s *Stars) IsActive(r, c int) bool {
	for _, star := range s.Systems {
		if star.Row == r && star.Col == c {
			return true
		}
	}

	return false
}
