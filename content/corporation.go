package content

import (
	"bytes"
	"fmt"

	"github.com/nboughton/go-roll"
	"github.com/nboughton/swnt/content/format"
)

// Corporation with a Body
type Corporation struct {
	Name               string
	Organization       string
	Business           string
	ReputationAndRumor string
}

// NewCorporation with random characteristics
func NewCorporation() Corporation {
	c := Corporation{
		Name:               corpTable.name.Roll(),
		Organization:       corpTable.organization.Roll(),
		Business:           corpTable.business.Roll(),
		ReputationAndRumor: corpTable.reputation.Roll(),
	}
	return c
}

// Format returns Corporation c formatted as type t
func (c Corporation) Format(t format.OutputType) string {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, format.Table(t, []string{"Corporation", ""}, [][]string{
		{corpTable.name.Name, fmt.Sprintf("%s %s", c.Name, c.Organization)},
		{corpTable.business.Name, c.Business},
		{corpTable.reputation.Name, c.ReputationAndRumor},
	}))

	return buf.String()
}

func (c Corporation) String() string {
	return c.Format(format.TEXT)
}

var corpTable = struct {
	name         roll.List
	organization roll.List
	business     roll.List
	reputation   roll.List
}{
	// Name SWN Revised Free Edition p192
	roll.List{
		Name: "Name",
		Items: []string{
			"Ad Astra",
			"Colonial",
			"Compass",
			"Daybreak",
			"Frontier",
			"Guo Yin",
			"Highbeam",
			"Imani",
			"Magnus",
			"Meteor",
			"Neogen",
			"New Dawn",
			"Omnitech",
			"Outertech",
			"Overwatch",
			"Panstellar",
			"Shogun",
			"Silverlight",
			"Spiker",
			"Stella",
			"Striker",
			"Sunbeam",
			"Terra Prime",
			"Wayfarer",
			"West Wind",
		},
	},

	// Organization SWN Revised Free Edition p192
	roll.List{
		Name: "Organization",
		Items: []string{
			"Alliance",
			"Association",
			"Band",
			"Circle",
			"Clan",
			"Combine",
			"Company",
			"Cooperative",
			"Corporation",
			"Enterprises",
			"Faction",
			"Group",
			"Megacorp",
			"Multistellar",
			"Organization",
			"Outfit",
			"Pact",
			"Partnership",
			"Ring",
			"Society",
			"Sodality",
			"Syndicate",
			"Union",
			"Unity",
			"Zaibatsu",
		},
	},

	// Business SWN Revised Free Edition p192
	roll.List{
		Name: "Business",
		Items: []string{
			"Aeronautics",
			"Agriculture",
			"Art",
			"Assassination",
			"Asteroid Mining",
			"Astrotech",
			"Biotech",
			"Bootlegging",
			"Computer Hardware",
			"Construction",
			"Cybernetics",
			"Electronics",
			"Energy Weapons",
			"Entertainment",
			"Espionage",
			"Exploration",
			"Fishing",
			"Fuel Refining",
			"Gambling",
			"Gemstones",
			"Gengineering",
			"Grav Vehicles",
			"Heavy Weapons",
			"Ideology",
			"Illicit Drugs",
			"Journalism",
			"Law Enforcement",
			"Liquor",
			"Livestock",
			"Maltech",
			"Mercenary Work",
			"Metallurgy",
			"Pharmaceuticals",
			"Piracy",
			"Planetary Mining",
			"Plastics",
			"Pretech",
			"Prisons",
			"Programming",
			"Projectile Guns",
			"Prostitution",
			"Psionics",
			"Psitech",
			"Robotics",
			"Security",
			"Shipyards",
			"Snacks",
			"Telcoms",
			"Transport",
			"Xenotech",
		},
	},

	// ReputationAndRumor SWN Revised Free Edition p192
	roll.List{
		Name: "Reputation and Rumors",
		Items: []string{
			"Reckless with the lives of their employees",
			"Have a dark secret about their board of directors",
			"Notoriously xenophobic towards aliens",
			"Lost much money to an embezzler who evaded arrest",
			"Reliable and trustworthy goods",
			"Stole a lot of R&D from a rival corporation",
			"They have high-level political connections",
			"Rumored cover-up of a massive industrial accident",
			"Stodgy and very conservative in their business plans",
			"Stodgy and very conservative in their business plans",
			"The company’s owner is dangerously insane",
			"Rumored ties to a eugenics cult",
			"Said to have a cache of pretech equipment",
			"Possibly teetering on the edge of bankruptcy",
			"Front for a planetary government’s espionage arm",
			"Secretly run by a psychic cabal",
			"Secretly run by hostile aliens",
			"Secretly run by an unbraked AI",
			"They’ve turned over a new leaf with the new CEO",
			"Deeply entangled with the planetary underworld",
		},
	},
}
