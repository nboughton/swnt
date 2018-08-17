package table

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nboughton/rollt"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// ThreePart represents a relatively common structure for multi-layer tables in
// SWN (RE)
type ThreePart struct {
	Headers [3]string
	Tables  []ThreePartSubTable
}

// Roll performs all rolls on a ThreePart table
func (t ThreePart) Roll() [][]string {
	i := rand.Intn(len(t.Tables))

	return [][]string{
		{t.Headers[0], t.Tables[i].Name},
		{t.Headers[1], t.Tables[i].SubTable1.Roll()},
		{t.Headers[2], t.Tables[i].SubTable2.Roll()},
	}
}

func (t ThreePart) String() string {
	s := ""
	for _, sub := range t.Tables {
		s += sub.String()
	}

	return s
}

// ThreePartSubTable represent the subtables of a ThreePart
type ThreePartSubTable struct {
	Name      string
	SubTable1 rollt.Able
	SubTable2 rollt.Able
}

// String satisfies the Stringer interface for threePartSubTables
func (t ThreePartSubTable) String() string {
	return fmt.Sprintf("%s\n%s\n%s", t.Name, t.SubTable1, t.SubTable2)
}

// OneRoll represents the oft used one-roll systems spread throughout SWN
type OneRoll struct {
	D4  rollt.Able
	D6  rollt.Able
	D8  rollt.Able
	D10 rollt.Able
	D12 rollt.Able
	D20 rollt.Able
}

// Roll performs all rolls for a OneRoll and returns the results
func (o OneRoll) Roll() [][]string {
	return [][]string{
		{o.D4.Label(), o.D4.Roll()},
		{o.D6.Label(), o.D6.Roll()},
		{o.D8.Label(), o.D8.Roll()},
		{o.D10.Label(), o.D10.Roll()},
		{o.D12.Label(), o.D12.Roll()},
		{o.D20.Label(), o.D20.Roll()},
	}
}

func (o OneRoll) String() string {
	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n", o.D4, o.D6, o.D8, o.D10, o.D12, o.D20)
}
