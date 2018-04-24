package adventure

import (
	"fmt"
	"strings"

	"github.com/nboughton/swnt/gen/world"

	"github.com/fatih/color"
)

// New throws together a random adventure seed using the table available from Stars Without Number
func New(tag string) string {
	for _, t := range world.Tags {
		if t.Name == tag {
			str := Seed.Roll()
			str = strings.Replace(str, "Thing", color.MagentaString(t.Things.Roll()), -1)
			str = strings.Replace(str, "Friend", color.GreenString(t.Friends.Roll()), -1)
			str = strings.Replace(str, "Enemy", color.RedString(t.Enemies.Roll()), -1)
			str = strings.Replace(str, "Complication", color.YellowString(t.Complications.Roll()), -1)
			str = strings.Replace(str, "Place", color.CyanString(t.Places.Roll()), -1)

			return str
		}
	}

	return fmt.Sprintf("No tag data for %s", tag)
}
