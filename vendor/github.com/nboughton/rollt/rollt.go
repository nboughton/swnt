// Package rollt implements a small library for creating Roll tables (commonly used in TTRPGs)
package rollt

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/nboughton/go-dice"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Able interface is used to define all rollt.Ables that can be rolled or printed.
type Able interface {
	Label() string
	Roll() string
	String() string
}

// Registry allows multiple tables to be registered and called as
// item actions from other tables.
/* For Example:

    var r = rollt.NewRegistry()

    var t = rollt.Table{
	    Name: "Test",
	    ID:   "Parent",
	    Dice: "1d6",
	    Reroll: rollt.Reroll{
		    Match: []int{5, 6},
		    Dice:  "1d4",
	    },
	    Items: []rollt.Item{
		    {Match: []int{1}, Text: "Item 1", Action: func() string {
			    tbl, _ := r.Get(t1.ID)
			    return tbl.Roll()
		    }},
		    {Match: []int{2}, Text: "Item 2"},
		    {Match: []int{3}, Text: "Item 3"},
		    {Match: []int{4}, Text: "Item 4"},
		    {Match: []int{5}, Text: "Item 5"},
		    {Match: []int{6}, Text: "Item 6"},
	    },
    }

    var t1 = rollt.Table{
	    Name: "Subtable 1",
	    ID:   "Child 1",
	    Dice: "1d6",
	    Items: []rollt.Item{
		    {Match: []int{1}, Text: "Item 1.1"},
		    {Match: []int{2}, Text: "Item 2.1"},
		    {Match: []int{3}, Text: "Item 3.1"},
		    {Match: []int{4}, Text: "Item 4.1"},
		    {Match: []int{5}, Text: "Item 5.1"},
		    {Match: []int{6}, Text: "Item 6.1"},
	    },
    }

    func main() {
	    r.Add(t)
	    r.Add(t1)

	    //	fmt.Printf("%+v\n", t.Items)
	    for i := 0; i < 10; i++ {
		    fmt.Println(t.Roll())
    	}
		}

*/
type Registry map[string]Table

// NewRegistry returns a new registry
func NewRegistry() Registry {
	return make(Registry)
}

// Add a table to the Registry
func (r Registry) Add(t Table) error {
	if _, ok := r[t.ID]; !ok {
		r[t.ID] = t
		return nil
	}

	return fmt.Errorf("table %s already registered", t.ID)
}

// Remove a table from the Registry
func (r Registry) Remove(id string) error {
	if _, ok := r[id]; ok {
		delete(r, id)
		return nil
	}

	return fmt.Errorf("table %s is not registered", id)
}

// Get a table from the registry
func (r Registry) Get(id string) (Table, error) {
	t, ok := r[id]
	if !ok {
		return Table{}, fmt.Errorf("no table registered with id [%s]", id)
	}

	return t, nil
}

// Table represents a table of text options that can be rolled on. Name is
// optional. Tables are preferable to Lists when using multiple dice to achieve
// a result (i.e 2d6) because their results fall on a bell curve whereas single-die
// rolls have an even probability.
type Table struct {
	ID     string // Shorthand ID for finding subtables
	Name   string
	Dice   string
	Reroll Reroll
	Items  []Item
}

// Reroll describes conditions under which the table should be rolled on again, using a different dice value
type Reroll struct {
	Match matchSet
	Dice  string
}

// Item represents the text and matching numbers from the table
type Item struct {
	Match  matchSet
	Text   string
	Action func() string
}

type matchSet []int

func (m matchSet) contains(n int) bool {
	for _, i := range m {
		if i == n {
			return true
		}
	}

	return false
}

func (m matchSet) String() string {
	var s []string

	for _, n := range m {
		s = append(s, strconv.Itoa(n))
	}

	return strings.Join(s, ", ")
}

// Roll on the table and return the option drawn.
func (t Table) Roll() string {
	out := ""

	d, err := dice.NewBag(t.Dice)
	if err != nil {
		return "Error: " + err.Error()
	}

	n, _ := d.Roll()
	// Record initial roll result
	for _, i := range t.Items {
		if i.Match.contains(n) {
			out = i.Text
		}
	}

	// Check for a reroll
	if t.Reroll.Match.contains(n) {
		d, err = dice.NewBag(t.Reroll.Dice)
		if err != nil {
			return "Error: " + err.Error()
		}

		n, _ = d.Roll()
		for _, i := range t.Items {
			if i.Match.contains(n) {
				if out != "" {
					out += "; "
				}
				out += i.Text
			}
		}
	}

	// Append text for final roll result
	for _, i := range t.Items {
		if i.Match.contains(n) {
			if i.Action != nil {
				if out != "" {
					out += "; "
				}

				out += i.Action()
			}
			return out
		}
	}

	return ""
}

func (t Table) String() string {
	var (
		buf = new(bytes.Buffer)
		tw  = tabwriter.NewWriter(buf, 2, 2, 1, ' ', 0)
	)

	fmt.Fprintln(tw, "Dice\t|\tText")
	for _, i := range t.Items {
		fmt.Fprintf(tw, "%s\t|\t%s\n", i.Match, i.Text)
	}
	tw.Flush()

	return buf.String()
}

// Label returns the table Name
func (t Table) Label() string {
	return t.Name
}

// List represents a List of strings from which something can be selected at random
type List struct {
	Name  string
	Items []string
}

// Roll returns a random string from List
func (l List) Roll() string {
	if len(l.Items) > 0 {
		return l.Items[rand.Intn(len(l.Items))]
	}

	return ""
}

func (l List) String() string {
	return strings.Join(l.Items, ", ")
}

// Label returns the list Name
func (l List) Label() string {
	return l.Name
}
