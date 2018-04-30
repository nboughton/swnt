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

	t, err := world.Tags.Find(worldTag)
	if err == nil {
		a.Tag.Enemy = t.Enemies.Roll()
		a.Tag.Friend = t.Friends.Roll()
		a.Tag.Thing = t.Things.Roll()
		a.Tag.Place = t.Places.Roll()
		a.Tag.Complication = t.Complications.Roll()
	}

	return a
}

func (a Adventure) String() string {
	buf, str := new(bytes.Buffer), a.Seed

	fmt.Fprintf(buf, "Tag: %s\n", a.Tag.Name)
	str = strings.Replace(str, "Enemy", color.RedString("Enemy (%s)", a.Tag.Enemy), -1)
	str = strings.Replace(str, "Friend", color.GreenString("Friend (%s)", a.Tag.Friend), -1)
	str = strings.Replace(str, "Thing", color.MagentaString("Thing (%s)", a.Tag.Thing), -1)
	str = strings.Replace(str, "Place", color.CyanString("Place (%s)", a.Tag.Place), -1)
	str = strings.Replace(str, "Complication", color.YellowString("Complication (%s)", a.Tag.Complication), -1)
	fmt.Fprintf(buf, "%s", str)

	return buf.String()
}
