package beast

import (
	"bytes"
	"fmt"
)

// Beast defines the aggregate descriptors for an animal
type Beast struct {
	Features    string
	BodyPlan    string
	LimbNovelty string
	SkinNovelty string
	MainWeapon  string
	Size        string
}

// New Beast
func New() Beast {
	return Beast{
		Features:    BasicAnimalFeatures.Roll(),
		BodyPlan:    BodyPlan.Roll(),
		LimbNovelty: LimbNovelty.Roll(),
		SkinNovelty: SkinNovelty.Roll(),
		MainWeapon:  MainWeapon.Roll(),
		Size:        Size.Roll(),
	}
}

func (b Beast) String() string {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, "%s\t:\t%s\n", BasicAnimalFeatures.Label(), b.Features)
	fmt.Fprintf(buf, "%s\t:\t%s\n", BodyPlan.Label(), b.BodyPlan)
	fmt.Fprintf(buf, "%s\t:\t%s\n", LimbNovelty.Label(), b.LimbNovelty)
	fmt.Fprintf(buf, "%s\t:\t%s\n", SkinNovelty.Label(), b.SkinNovelty)
	fmt.Fprintf(buf, "%s\t:\t%s\n", MainWeapon.Label(), b.MainWeapon)
	fmt.Fprintf(buf, "%s\t:\t%s\n", Size.Label(), b.Size)

	return buf.String()
}
