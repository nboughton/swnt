package poi

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"

	"github.com/nboughton/swnt/content/format"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// POI Point of Interest
type POI struct {
	Point     string
	Occupied  string
	Situation string
}

// New roll a new point of interest
func New() POI {
	t := Table.Tables[rand.Intn(len(Table.Tables))]

	return POI{
		Point:     t.Name,
		Occupied:  t.SubTable1.Roll(),
		Situation: t.SubTable2.Roll(),
	}
}

// Format returns the POI formatted as type t
func (p POI) Format(t format.OutputType) string {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, format.Table(t, true, p.Point, [][]string{
		{Table.Headers[1], p.Occupied},
		{Table.Headers[2], p.Situation},
	}))

	return buf.String()
}
