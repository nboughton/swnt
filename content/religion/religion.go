package religion

import (
	"bytes"
	"fmt"

	"github.com/nboughton/swnt/content/format"
)

// Religion is pretty self explanatory
type Religion struct {
	Evolution       string
	Leadership      string
	OriginTradition string
}

// New Religion with random characteristics
func New() Religion {
	r := Religion{
		Evolution:       Evolution.Roll(),
		Leadership:      Leadership.Roll(),
		OriginTradition: OriginTradition.Roll(),
	}
	return r
}

// Format returns the religion formatted as type t
func (r Religion) Format(t format.OutputType) string {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, format.Table(t, true, "Religion", [][]string{
		{OriginTradition.Name, r.OriginTradition},
		{Evolution.Name, r.Evolution},
		{Leadership.Name, r.Leadership},
	}))

	return buf.String()
}
