package export

import (
	"fmt"
	"os"
	"strings"

	"github.com/nboughton/swnt/content/sector"
	"github.com/nboughton/swnt/haxscii"
)

var (
	dirPerm  = os.FileMode(0755)
	filePerm = os.FileMode(0644)
)

// Exporter represents any type that can Setup an export directory and output data to it.
type Exporter interface {
	Write() error
}

// New returns a new Exporter. Export types currently supported are: hugo and text
func New(exportType, name string, data *sector.Stars) (Exporter, error) {
	switch exportType {
	case "hugo":
		return &Hugo{
			Name:  name,
			Stars: data,
		}, nil

	case "txt":
		return &Text{
			Name:  name,
			Stars: data,
		}, nil
	case "json":
		return &JSON{
			Name:  name,
			Stars: data,
		}, nil
	}

	return nil, fmt.Errorf("no Exporter found for [%s], available options are [%s]", exportType, []string{"hugo", "txt", "json"})
}

// Hexmap returns the ASCII representation of a Sector map
func Hexmap(data *sector.Stars, useColour bool, playerMap bool) string {
	haxscii.Colour(useColour)
	h := haxscii.NewMap(data.Rows, data.Cols)
	for _, s := range data.Systems {
		name, tag1, tag2, tl := s.Name, s.Worlds[0].Tags[0].Name, s.Worlds[0].Tags[1].Name, strings.Split(s.Worlds[0].TechLevel, ",")[0]
		c := haxscii.White // I default to black/dark terminals, this might be problematic for weirdos that use light terms

		switch tl {
		case "TL0":
			c = haxscii.White
		case "TL1":
			c = haxscii.Red
		case "TL2":
			c = haxscii.Yellow
		case "TL3":
			c = haxscii.Magenta
		case "TL4", "TL4+":
			c = haxscii.Green
		case "TL5":
			c = haxscii.Cyan
		}

		if playerMap {
			h.SetTxt(s.Row, s.Col, [4]string{name, "", "", ""}, c)
		} else {
			h.SetTxt(s.Row, s.Col, [4]string{name, tag1, tag2, tl}, c)
		}
	}

	return h.String()
}
