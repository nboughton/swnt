package beast

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
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

func (b Beast) String() string {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, "%s\t:\t%s\n", b.Type, b.Behaviour)
	fmt.Fprintf(buf, "%s\t:\t%s\n", BasicAnimalFeatures.Label(), b.Features)
	fmt.Fprintf(buf, "%s\t:\t%s\n", BodyPlan.Label(), b.BodyPlan)
	fmt.Fprintf(buf, "%s\t:\t%s\n", LimbNovelty.Label(), b.LimbNovelty)
	fmt.Fprintf(buf, "%s\t:\t%s\n", SkinNovelty.Label(), b.SkinNovelty)
	fmt.Fprintf(buf, "%s\t:\t%s\n", MainWeapon.Label(), b.MainWeapon)
	fmt.Fprintf(buf, "%s\t:\t%s\n", Size.Label(), b.Size)

	return buf.String()
}
