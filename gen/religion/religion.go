package religion

import (
	"bytes"
	"fmt"
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

func (r Religion) String() string {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, "%s\t:\t%s\n", OriginTradition.Name, r.OriginTradition)
	fmt.Fprintf(buf, "%s\t:\t%s\n", Evolution.Name, r.Evolution)
	fmt.Fprintf(buf, "%s\t:\t%s\n", Leadership.Name, r.Leadership)

	return buf.String()
}
