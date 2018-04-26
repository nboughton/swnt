package adventure

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/nboughton/swnt/gen/world"
)

// Adventure represents the elements of an Adventure outline
type Adventure struct {
	Seed string
	Tag  tag
}

type tag struct {
	Name         string
	Enemy        string
	Friend       string
	Thing        string
	Place        string
	Complication string
}

// New throws together a random adventure seed using the table available from Stars Without Number
func New(worldTag string) Adventure {
	a := Adventure{
		Seed: Seed.Roll(),
		Tag:  tag{Name: worldTag},
	}

	for _, t := range world.Tags {
		if t.Name == worldTag {
			a.Tag.Enemy = color.RedString(t.Enemies.Roll())
			a.Tag.Friend = color.GreenString(t.Friends.Roll())
			a.Tag.Thing = color.MagentaString(t.Things.Roll())
			a.Tag.Place = color.CyanString(t.Places.Roll())
			a.Tag.Complication = color.YellowString(t.Complications.Roll())
		}
	}

	return a
}

func (a Adventure) String() string {
	buf, str := new(bytes.Buffer), a.Seed

	fmt.Fprintf(buf, "Tag: %s\n", a.Tag.Name)
	str = strings.Replace(str, "Enemy", fmt.Sprintf("Enemy (%s)", a.Tag.Enemy), -1)
	str = strings.Replace(str, "Friend", fmt.Sprintf("Friend (%s)", a.Tag.Friend), -1)
	str = strings.Replace(str, "Thing", fmt.Sprintf("Thing (%s)", a.Tag.Thing), -1)
	str = strings.Replace(str, "Place", fmt.Sprintf("Place (%s)", a.Tag.Place), -1)
	str = strings.Replace(str, "Complication", fmt.Sprintf("Complication (%s)", a.Tag.Complication), -1)
	fmt.Fprintf(buf, "%s", str)

	return buf.String()
}
