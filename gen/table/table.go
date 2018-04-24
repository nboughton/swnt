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
	D4  NamedList
	D6  NamedList
	D8  NamedList
	D10 NamedList
	D12 NamedList
	D20 NamedList
}

// Roll performs all rolls for a OneRoll and returns the results
func (o OneRoll) Roll() string {
	var (
		buf = new(bytes.Buffer)
	)

	fmt.Fprintf(buf, "%s\t:\t%s\n", o.D4.Name, o.D4.Table.Roll())
	fmt.Fprintf(buf, "%s\t:\t%s\n", o.D6.Name, o.D6.Table.Roll())
	fmt.Fprintf(buf, "%s\t:\t%s\n", o.D8.Name, o.D8.Table.Roll())
	fmt.Fprintf(buf, "%s\t:\t%s\n", o.D10.Name, o.D10.Table.Roll())
	fmt.Fprintf(buf, "%s\t:\t%s\n", o.D12.Name, o.D12.Table.Roll())
	fmt.Fprintf(buf, "%s\t:\t%s\n", o.D20.Name, o.D20.Table.Roll())

	return buf.String()
}

func (o OneRoll) String() string {
	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n", o.D4, o.D6, o.D8, o.D10, o.D12, o.D20)
}

// NamedList represents a named List. It's kind of obvious
type NamedList struct {
	Name  string
	Table rollt.List
}

func (n NamedList) String() string {
	var (
		buf = new(bytes.Buffer)
	)

	fmt.Fprintf(buf, "D%d\t|\t%s\n", len(n.Table), n.Name)
	for i, text := range n.Table {
		fmt.Fprintf(buf, "%d\t|\t%s\n", i+1, text)
	}

	return buf.String()
}

// Roll on the list table
func (n NamedList) Roll() string {
	return n.Table.Roll()
}
