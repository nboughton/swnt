package religion

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Religion
type Religion struct {
	Evolution       string
	Leadership      string
	OriginTradition string
}

// New Religion with random characteristics
func New() Religion {
	religion := Religion{
		Evolution:       Evolution.Roll(),
		Leadership:      Leadership.Roll(),
		OriginTradition: OriginTradition.Roll(),
	}
	return religion
}

func (religion Religion) String() string {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, "%s\t:\t%s\n", Evolution.Name, religion.Evolution)
	fmt.Fprintf(buf, "%s\t:\t%s\n", Leadership.Name, religion.Leadership)
	fmt.Fprintf(buf, "%s\t:\t%s", OriginTradition.Name, religion.OriginTradition)
	return buf.String()
}
