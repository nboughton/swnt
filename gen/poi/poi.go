package poi

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
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

func (p POI) String() string {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, "%s\t:\t%s\n", Table.Headers[0], p.Point)
	fmt.Fprintf(buf, "%s\t:\t%s\n", Table.Headers[1], p.Occupied)
	fmt.Fprintf(buf, "%s\t:\t%s\n", Table.Headers[2], p.Situation)

	return buf.String()
}
