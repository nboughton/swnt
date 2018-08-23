package beast

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"

	"github.com/nboughton/swnt/content/format"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Beast defines the aggregate descriptors for an animal
type Beast struct {
	Type        string
	Behaviour   string
	Features    string
	BodyPlan    string
	LimbNovelty string
	SkinNovelty string
	MainWeapon  string
	Size        string
}

// New Beast
func New() Beast {
	b := Beast{
		Features:    BasicAnimalFeatures.Roll(),
		BodyPlan:    BodyPlan.Roll(),
		LimbNovelty: LimbNovelty.Roll(),
		SkinNovelty: SkinNovelty.Roll(),
		MainWeapon:  MainWeapon.Roll(),
		Size:        Size.Roll(),
	}

	switch rand.Intn(3) {
	case 0:
		b.Type = Predator.Label()
		b.Behaviour = Predator.Roll()

	case 1:
		b.Type = Prey.Label()
		b.Behaviour = Prey.Roll()

	case 2:
		b.Type = Scavenger.Label()
		b.Behaviour = Scavenger.Roll()
	}

	return b
}

// Format output as format type t
func (b Beast) Format(t format.OutputType) string {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, format.Table(t, true, "Beast", [][]string{
		{b.Type, b.Behaviour},
		{BasicAnimalFeatures.Label(), b.Features},
		{BodyPlan.Label(), b.BodyPlan},
		{LimbNovelty.Label(), b.LimbNovelty},
		{SkinNovelty.Label(), b.SkinNovelty},
		{MainWeapon.Label(), b.MainWeapon},
		{Size.Label(), b.Size},
	}))

	return buf.String()
}

func (b Beast) String() string {
	return b.Format(format.TEXT)
}
