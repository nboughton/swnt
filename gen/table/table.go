package table

import (
	"bytes"
	//"encoding/json"
	"fmt"
	//"io/ioutil"
	"math/rand"
	"time"

	"github.com/nboughton/rollt"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Table interface is used to define all tables that can be rolled or printed.
// Currently unused but might be useful at a later date.
type Table interface {
	Label() string
	Roll() string
	String() string
}

// ThreePart represents a relatively common structure for multi-layer tables in
// SWN (RE)
type ThreePart struct {
	Headers [3]string
	Tables  []ThreePartSubTable
}

// Roll performs all rolls on a ThreePart table
func (t ThreePart) Roll() string {
	var (
		buf = new(bytes.Buffer)
		i   = rand.Intn(len(t.Tables))
	)

	fmt.Fprintf(buf, "%s\t:\t%s\n", t.Headers[0], t.Tables[i].Name)
	fmt.Fprintf(buf, "%s\t:\t%s\n", t.Headers[1], t.Tables[i].SubTable1.Roll())
	fmt.Fprintf(buf, "%s\t:\t%s\n", t.Headers[2], t.Tables[i].SubTable2.Roll())

	return buf.String()
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
	SubTable1 rollt.List
	SubTable2 rollt.List
}

// String satisfies the Stringer interface for threePartSubTables
func (t ThreePartSubTable) String() string {
	return fmt.Sprintf("%s\n%s\n%s", t.Name, t.SubTable1, t.SubTable2)
}

// OneRoll represents the oft used one-roll systems spread throughout SWN
type OneRoll struct {
	D4  Table
	D6  Table
	D8  Table
	D10 Table
	D12 Table
	D20 Table
}

// Roll performs all rolls for a OneRoll and returns the results
func (o OneRoll) Roll() string {
	var (
		buf = new(bytes.Buffer)
	)

	fmt.Fprintf(buf, "%s\t:\t%s\n", o.D4.Label(), o.D4.Roll())
	fmt.Fprintf(buf, "%s\t:\t%s\n", o.D6.Label(), o.D6.Roll())
	fmt.Fprintf(buf, "%s\t:\t%s\n", o.D8.Label(), o.D8.Roll())
	fmt.Fprintf(buf, "%s\t:\t%s\n", o.D10.Label(), o.D10.Roll())
	fmt.Fprintf(buf, "%s\t:\t%s\n", o.D12.Label(), o.D12.Roll())
	fmt.Fprintf(buf, "%s\t:\t%s\n", o.D20.Label(), o.D20.Roll())

	return buf.String()
}

func (o OneRoll) String() string {
	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n", o.D4, o.D6, o.D8, o.D10, o.D12, o.D20)
}
