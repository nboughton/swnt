package sector

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"

	"github.com/nboughton/swnt/content"
	"github.com/nboughton/swnt/content/culture"
	"github.com/nboughton/swnt/content/format"
	"github.com/nboughton/swnt/content/name"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Star represents a single Star on the Sector map
type Star struct {
	Row, Col int
	Culture  culture.Culture
	Name     string
	Worlds   []content.World
	POIs     []content.POI
}

// NewStar generates a new Star struct to be added to the map
func NewStar(row, col int, name string, exclude []string, poiChance, otherWorldChance int) *Star {
	ctr := culture.Random()

	s := &Star{
		Row:     row,
		Col:     col,
		Culture: ctr,
		Name:    name,
		Worlds:  []content.World{content.NewWorld(ctr, true, exclude)},
	}

	// Cascading 10% chance of other worlds
	for rand.Intn(100) < otherWorldChance {
		ctr = culture.Random()
		s.Worlds = append(s.Worlds, content.NewWorld(ctr, false, exclude))
	}

	// 30% chance of a Point of Interest
	if rand.Intn(100) < poiChance {
		s.POIs = append(s.POIs, content.NewPOI())
	}

	return s
}

// Format returns the details of a Star formatted as type t
func (s *Star) Format(t format.OutputType) string {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, format.Header(t, 2, fmt.Sprintf("Hex:  %d,%d", s.Row, s.Col)))
	fmt.Fprintf(buf, format.Header(t, 3, "Primary World"))
	fmt.Fprintf(buf, s.Worlds[0].Format(t))

	if len(s.Worlds) > 1 {
		fmt.Fprintf(buf, format.Header(t, 3, "Other Worlds"))
		for _, w := range s.Worlds[1:] {
			fmt.Fprintln(buf, w.Format(t))
		}
	}

	if len(s.POIs) > 0 {
		fmt.Fprintf(buf, format.Header(t, 3, "Points of Interest"))
		for _, p := range s.POIs {
			fmt.Fprintln(buf, p.Format(t))
		}
	}

	return buf.String()
}

// Stars represents the generated collection of Stars that will be used to populate a hex grid
type Stars struct {
	Rows, Cols int
	Systems    []*Star
}

// NewSector returns a blank Sector struct and generates tag information according to the guidelines
// in pages 133 - 177 of Stars Without Number (Revised Edition).
func NewSector(rows, cols int, excludeTags []string, poiChance, otherWorldChance int) *Stars {
	s := &Stars{
		Rows: rows,
		Cols: cols,
	}

	cells := s.Rows * s.Cols
	stars := (rand.Intn(cells/4) / 2) + (cells / 4)

	for row, col := rand.Intn(s.Rows), rand.Intn(s.Cols); len(s.Systems) <= stars; row, col = rand.Intn(s.Rows), rand.Intn(s.Cols) {
		if !s.active(row, col) {
			s.Systems = append(s.Systems, NewStar(row, col, s.systemName(), excludeTags, poiChance, otherWorldChance))
		}
	}

	return s
}

// UniqueName ensures rolls on the name.System table until it gets a name that is not currently in use.
func (s *Stars) systemName() string {
	//n := name.System.Roll() // Try system first
	n := name.Generate(rand.Intn(4) + 3)
	for {
		if !s.nameUsed(n) {
			return n
		}

		n = name.Generate(rand.Intn(4) + 3)
	}
}

func (s *Stars) nameUsed(n string) bool {
	for _, star := range s.Systems {
		if star.Name == n {
			return true
		}
	}

	return false
}

// active checks to see if there is an active Star at r(ow), c(ol)
func (s *Stars) active(row, col int) bool {
	for _, star := range s.Systems {
		if star.Row == row && star.Col == col {
			return true
		}
	}

	return false
}
