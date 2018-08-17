package corporation

import (
	"bytes"
	"fmt"

	"github.com/nboughton/swnt/content/format"
)

// Corporation with a Body
type Corporation struct {
	Name               string
	Organization       string
	Business           string
	ReputationAndRumor string
}

// New Corporation with random characteristics
func New() Corporation {
	c := Corporation{
		Name:               Name.Roll(),
		Organization:       Organization.Roll(),
		Business:           Business.Roll(),
		ReputationAndRumor: ReputationAndRumor.Roll(),
	}
	return c
}

// Format returns Corporation c formatted as type t
func (c Corporation) Format(t format.OutputType) string {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, format.Table(t, true, "Corporation", [][]string{
		{Name.Name, fmt.Sprintf("%s %s", c.Name, c.Organization)},
		{Business.Name, c.Business},
		{ReputationAndRumor.Name, c.ReputationAndRumor},
	}))

	return buf.String()
}
