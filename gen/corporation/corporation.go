package corporation

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Corporation with a Body
type Corporation struct {
	Name               string
	Organization       string
	Business           string
	ReputationAndRumor string
}

// New Corporation with random characteristics
func New() Corporation {
	corporation := Corporation{
		Name:               Name.Roll(),
		Organization:       Organization.Roll(),
		Business:           Business.Roll(),
		ReputationAndRumor: ReputationAndRumor.Roll(),
	}
	return corporation
}

func (corporation Corporation) String() string {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, "%s\t:\t%s %s\n", Name.Name, corporation.Name, corporation.Organization)
	fmt.Fprintf(buf, "%s\t:\t%s\n", Business.Name, corporation.Business)
	fmt.Fprintf(buf, "%s\t:\t%s", ReputationAndRumor.Name, corporation.ReputationAndRumor)
	return buf.String()
}
