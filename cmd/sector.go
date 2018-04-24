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
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/nboughton/swnt/gen/name"
	"github.com/nboughton/swnt/gen/sector"
	"github.com/nboughton/swnt/haxscii"
	"github.com/spf13/cobra"
)

const (
	flRows   = "rows"
	flCols   = "cols"
	flColour = "colour"
)

var ansRegex = regexp.MustCompile("(?i)^(y|n|r)$")

// sectorCmd represents the sector command
var sectorCmd = &cobra.Command{
	Use:   "sector",
	Short: "Create the skeleton of a Sector",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// get flags
		var (
			colour, _ = cmd.Flags().GetBool(flColour)
			secData   = sector.NewSector().ByCoords()
			secName   = genSectorName()
		)

		fmt.Println(secName)
		if colour {
			fmt.Println(hexmapColor(secData))
		} else {
			fmt.Println(hexmap(secData))
		}

		ans := "r"
		for {
			fmt.Printf("Write Sector? [y]es, [n]o, [r]eroll: [%s] ", ans)
			fmt.Scanf("%s", &ans)
			if ansRegex.MatchString(ans) {
				switch ans {
				case "y":
					os.Mkdir(secName, 0777)

					for _, system := range secData {
						dir := fmt.Sprintf("%d,%d-%s", system.Row, system.Col, system.Name)
						os.Mkdir(filepath.Join(secName, dir), 0777)
						ioutil.WriteFile(filepath.Join(secName, dir, fmt.Sprintf("%s.%s", system.Name, "txt")), []byte(system.String()), 0755)
					}

					ioutil.WriteFile(filepath.Join(secName, "map.txt"), []byte(hexmap(secData)), 0755)
					fmt.Printf("%s written\n", secName)
					return

				case "n":
					return

				case "r":
					secData = sector.NewSector().ByCoords()
					secName = genSectorName()
					fmt.Println(secName)
					if colour {
						fmt.Println(hexmapColor(secData))
					} else {
						fmt.Println(hexmap(secData))
					}
				}
			}
		}
	},
}

func genSectorName() string {
	secName := fmt.Sprintf("%s Sector", name.System.Roll())
	_, err := os.Stat(secName)
	for os.IsExist(err) {
		secName = fmt.Sprintf("%s Sector", name.System.Roll())
		_, err = os.Stat(secName)
	}

	return secName
}

// I should probably fold these into a single function at some point
func hexmap(data []*sector.Star) string {
	h := haxscii.NewMap()
	for _, s := range data {
		h.SetTxt(s.Row, s.Col, s.Name, s.Worlds[0].Tags[0].Name, s.Worlds[0].Tags[1].Name, strings.Split(s.Worlds[0].TechLevel, ",")[0])
	}

	return h.String()
}

func hexmapColor(data []*sector.Star) string {
	h := haxscii.NewMap()
	for _, s := range data {
		name, tag1, tag2, tl := s.Name, s.Worlds[0].Tags[0].Name, s.Worlds[0].Tags[1].Name, strings.Split(s.Worlds[0].TechLevel, ",")[0]
		c := color.WhiteString
		switch tl {
		case "TL0":
			c = color.WhiteString
		case "TL1":
			c = color.RedString
		case "TL2":
			c = color.YellowString
		case "TL3":
			c = color.MagentaString
		case "TL4", "TL4+":
			c = color.GreenString
		case "TL5":
			c = color.CyanString
		}

		h.SetTxtColour(s.Row, s.Col, name, tag1, tag2, tl, c)
	}

	return h.String()
}

func init() {
	newCmd.AddCommand(sectorCmd)
	sectorCmd.Flags().IntP(flRows, "r", 8, "Set rows")
	sectorCmd.Flags().IntP(flCols, "c", 10, "Set columns")
	sectorCmd.Flags().BoolP(flColour, "l", false, "Toggle colour output")
}
