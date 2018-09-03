package content

import (
	"fmt"
	"strings"

	"github.com/nboughton/swnt/content/format"
)

// statBlock data for a beast/bot/npc
type statBlock struct {
	Name   string
	HD     int
	AC     ac
	Atk    atk
	Dmg    string
	Move   string
	ML     int
	Skills int
	Saves  int
	Cost   int
}

// Format statBlock s as OutputType t
func (s statBlock) Format(t format.OutputType) string {
	return format.Table(t,
		[]string{"Name", "HD", "AC", "Atk", "Dmg", "Move", "ML", "Skills", "Saves", "Cost (robot/VI only)"},
		[][]string{{s.Name, fmt.Sprintf("%d", s.HD), s.AC.String(), s.Atk.String(), s.Dmg, s.Move, fmt.Sprintf("%d", s.ML), fmt.Sprintf("%d", s.Skills), fmt.Sprintf("%d", s.Saves), fmt.Sprintf("%d", s.Cost)}},
	)
}

func (s statBlock) String() string {
	return s.Format(format.TEXT)
}

type ac struct {
	Val   int
	Notes string
}

func (a ac) String() string {
	n := ""
	if len(a.Notes) > 0 {
		n = fmt.Sprintf(" (%s)", a.Notes)
	}

	return fmt.Sprintf("%d%s", a.Val, n)
}

type atk struct {
	Val int
	X   int
}

func (a atk) String() string {
	x := ""
	if a.X > 1 {
		x = fmt.Sprintf(" x %d", a.X)
	}

	return fmt.Sprintf("+%d%s", a.Val, x)
}

type statBlockTable []statBlock

func (s statBlockTable) Filter(terms ...string) statBlockTable {
	if len(terms) == 0 {
		return s
	}

	out := statBlockTable{}

	for _, row := range s {
		for _, term := range terms {
			if strings.Contains(strings.ToLower(row.Name), strings.ToLower(term)) {
				out = append(out, row)
			}
		}
	}

	return out
}

func (s statBlockTable) Format(t format.OutputType) string {
	rows := [][]string{}

	for _, row := range s {
		rows = append(rows, []string{row.Name, fmt.Sprintf("%d", row.HD), row.AC.String(), row.Atk.String(), row.Dmg, row.Move, fmt.Sprintf("%d", row.ML), fmt.Sprintf("%d", row.Skills), fmt.Sprintf("%d", row.Saves), fmt.Sprintf("%d", row.Cost)})
	}

	return format.Table(t, []string{"Name", "HD", "AC", "Atk", "Dmg", "Move", "ML", "Skills", "Saves", "Cost (robot/VI only)"}, rows)
}

// StatBlocks for typical NPC combat encounters as per SWN pg195
var StatBlocks = statBlockTable{
	// Humans
	{"Peaceful Human", 1, ac{10, ""}, atk{0, 1}, "Unarmed", "10m", 6, 1, 15, 0},
	{"Martial Human", 1, ac{10, ""}, atk{1, 1}, "By weapon", "10m", 8, 1, 15, 0},
	{"Veteran Fighter", 2, ac{14, ""}, atk{2, 1}, "By weapon +1", "10m", 9, 1, 14, 0},
	{"Elite Fighter", 3, ac{16, "combat"}, atk{4, 1}, "By weapon +1", "10m", 10, 2, 14, 0},
	{"Heroic Fighter", 6, ac{16, "combat"}, atk{8, 1}, "By weapon +3", "10m", 11, 3, 12, 0},
	{"Barbarian Hero", 6, ac{16, "primitive"}, atk{8, 1}, "By weapon +3", "10m", 11, 3, 12, 0},
	{"Barbarian Tribal", 1, ac{12, "primitive"}, atk{2, 1}, "By weapon", "10m", 8, 1, 15, 0},
	{"Gang Boss", 3, ac{14, ""}, atk{4, 1}, "By weapon +1", "10m", 9, 2, 14, 0},
	{"Gang Member", 1, ac{12, ""}, atk{1, 1}, "By weapon", "10m", 7, 1, 15, 0},
	{"Gengineered Killer", 4, ac{16, ""}, atk{5, 1}, "By weapon +1", "15m", 10, 2, 13, 0},
	{"Legendary Fighter", 10, ac{20, "powered"}, atk{12, 2}, "By weapon +4", "10m", 12, 5, 10, 0},
	{"Military Elite", 3, ac{16, "combat"}, atk{4, 1}, "By weapon +1", "10m", 10, 2, 14, 0},
	{"Military Soldier", 1, ac{16, "combat"}, atk{1, 1}, "By weapon", "10m", 9, 1, 15, 0},
	{"Normal Human", 1, ac{10, ""}, atk{0, 1}, "Unarmed", "10m", 6, 1, 15, 0},
	{"Pirate King", 7, ac{18, "powered"}, atk{9, 1}, "By weapon +2", "10m", 11, 3, 12, 0},
	{"Police Officer", 1, ac{14, ""}, atk{1, 1}, "By weapon", "10m", 8, 1, 15, 0},
	{"Serial Killer", 6, ac{12, ""}, atk{8, 1}, "By weapon +3", "10m", 12, 3, 12, 0},
	{"Skilled Professional", 1, ac{10, ""}, atk{0, 1}, "By weapon", "10m", 6, 2, 15, 0},
	{"Warrior Tyrant", 8, ac{20, "powered"}, atk{10, 1}, "By weapon +3", "10m", 11, 3, 11, 0},
	// Bots
	{"Janitor Bot", 1, ac{14, ""}, atk{0, 0}, "N/A", "5m", 8, 1, 15, 1000},
	{"Civilian Security Bot", 1, ac{15, ""}, atk{1, 1}, "1d8 stun", "10m", 12, 1, 15, 5000},
	{"Repair Bot", 1, ac{14, ""}, atk{0, 1}, "1d6 tool", "10m", 8, 1, 15, 5000},
	{"Industrial Work Bot", 2, ac{15, ""}, atk{0, 1}, "1d10 crush", "5m", 8, 1, 14, 2000},
	{"Companion Bot", 1, ac{12, ""}, atk{0, 1}, "1d2 unarmed", "10m", 6, 1, 15, 2500},
	{"Soldier Bot", 2, ac{16, ""}, atk{1, 1}, "By weapon", "10m", 10, 1, 14, 10000},
	{"Heavy Warbot", 6, ac{18, ""}, atk{8, 2}, "2d8 plasma", "15m", 10, 2, 12, 50000},
	// Beasts
	{"Small Vicious Beast", 0, ac{14, ""}, atk{1, 1}, "1d2", "10m", 7, 1, 15, 0},
	{"Small Pack Hunter", 1, ac{13, ""}, atk{1, 1}, "1d4", "15m", 8, 1, 15, 0},
	{"Large Pack Hunter", 2, ac{14, ""}, atk{2, 1}, "1d6", "15m", 9, 1, 14, 0},
	{"Large Aggressive Prey Animal", 5, ac{13, ""}, atk{4, 1}, "1d10", "15m", 8, 1, 12, 0},
	{"Lesser Lone Predator", 3, ac{14, ""}, atk{4, 2}, "1d8 each", "15m", 8, 2, 14, 0},
	{"Greater Lone Predator", 5, ac{15, ""}, atk{6, 2}, "1d10 each", "10m", 9, 2, 12, 0},
	{"Terrifying Apex Predator", 8, ac{16, ""}, atk{8, 2}, "1d10 each", "20m", 9, 2, 11, 0},
	{"Gengineered Murder Beast", 10, ac{18, ""}, atk{10, 4}, "1d10 each", "20m", 11, 3, 10, 0},
}
