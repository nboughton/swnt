// Copyright © 2018 Nick Boughton <nicholasboughton@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/nboughton/swnt/gen/name"
	"github.com/nboughton/swnt/gen/sector"
	"github.com/nboughton/swnt/haxscii"
	"github.com/spf13/cobra"
)

const (
	flColour = "colour"
	flPoi    = "poi-chance"
	flOW     = "other-worlds-chance"
)

var (
	ansRegex = regexp.MustCompile("(?i)^(y|n|r)$")
	dirPerm  = os.FileMode(0755)
	filePerm = os.FileMode(0644)
)

// sectorCmd represents the sector command
var sectorCmd = &cobra.Command{
	Use:   "sector",
	Short: "Create the skeleton of a Sector",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// get flags
		var (
			colour, _           = cmd.Flags().GetBool(flColour)
			excludeTags, _      = cmd.Flags().GetStringArray(flExclude)
			poiChance, _        = cmd.Flags().GetInt(flPoi)
			otherWorldChance, _ = cmd.Flags().GetInt(flOW)
			secData             = sector.NewSector(excludeTags, poiChance, otherWorldChance).ByCoords()
			secName             = genSectorName()
			mapDir              = "Maps"
		)

		fmt.Println(secName)
		fmt.Println(hexmap(secData, colour, false))

		ans := "r"
		for {
			fmt.Printf("Write Sector? [y]es, [n]o, [r]eroll: [%s] ", ans)
			fmt.Scanf("%s", &ans)
			if ansRegex.MatchString(ans) {
				switch ans {
				case "y":
					ensure(os.Mkdir(secName, dirPerm))

					for _, system := range secData {
						dir := fmt.Sprintf("%d,%d-%s", system.Row, system.Col, system.Name)
						ensure(os.Mkdir(filepath.Join(secName, dir), dirPerm))
						ensure(ioutil.WriteFile(filepath.Join(secName, dir, fmt.Sprintf("%s.%s", system.Name, "txt")), []byte(system.String()), filePerm))
					}

					ensure(os.Mkdir(filepath.Join(secName, mapDir), dirPerm))
					ensure(ioutil.WriteFile(filepath.Join(secName, mapDir, "gm-map.txt"), []byte(hexmap(secData, false, false)), filePerm))
					ensure(ioutil.WriteFile(filepath.Join(secName, mapDir, "pc-map.txt"), []byte(hexmap(secData, false, true)), filePerm))
					if colour {
						ensure(ioutil.WriteFile(filepath.Join(secName, mapDir, "gm-map-ansi.txt"), []byte(hexmap(secData, true, false)), filePerm))
						ensure(ioutil.WriteFile(filepath.Join(secName, mapDir, "pc-map-ansi.txt"), []byte(hexmap(secData, true, true)), filePerm))
					}
					fmt.Printf("%s written\n", secName)
					return

				case "n":
					return

				case "r":
					secData = sector.NewSector(excludeTags, poiChance, otherWorldChance).ByCoords()
					secName = genSectorName()
					fmt.Println(secName)
					fmt.Println(hexmap(secData, colour, false))
				}
			}
		}
	},
}

func genSectorName() string {
	secName := fmt.Sprintf("%s Sector", name.System.Roll())
	_, err := os.Stat(secName) // Ensure that there isn't already a sector of this name in the working directory
	for os.IsExist(err) {
		secName = fmt.Sprintf("%s Sector", name.System.Roll())
		_, err = os.Stat(secName)
	}

	return secName
}

func hexmap(data []*sector.Star, useColour bool, playerMap bool) string {
	haxscii.Colour(useColour)
	h := haxscii.NewMap()
	for _, s := range data {
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

func ensure(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	newCmd.AddCommand(sectorCmd)
	sectorCmd.Flags().BoolP(flColour, "l", false, "Toggle colour output")
	sectorCmd.Flags().StringArrayP(flExclude, "x", []string{}, "Exclude tags (-x zombies -x \"regional hegemon\" etc)")
	sectorCmd.Flags().IntP(flPoi, "p", 30, "Set % chance of a POI being generated for any given star in the sector")
	sectorCmd.Flags().IntP(flOW, "o", 10, "Set % chance for a secondary world to be generated for any given star in the sector")
}
