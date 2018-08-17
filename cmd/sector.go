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
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/nboughton/swnt/content/name"
	"github.com/nboughton/swnt/content/sector"
	"github.com/nboughton/swnt/export"
	"github.com/spf13/cobra"
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
			excludeTags, _      = cmd.Flags().GetStringArray(flExclude)
			poiChance, _        = cmd.Flags().GetInt(flPoi)
			otherWorldChance, _ = cmd.Flags().GetInt(flOW)
			secHeight, _        = cmd.Flags().GetInt(flSecHeight)
			secWidth, _         = cmd.Flags().GetInt(flSecWidth)
			exportTypes, _      = cmd.Flags().GetString(flExport)
		)

		if secHeight < 2 || secHeight > 99 || secWidth < 2 || secWidth > 99 {
			fmt.Println("Sectors larger than 99, or smaller than 2, hexes in either direction are not supported")
			return
		}

		var (
			secData = sector.NewSector(secHeight, secWidth, excludeTags, poiChance, otherWorldChance)
			secName = genSectorName()
		)

		fmt.Println(secName)
		fmt.Println(export.Hexmap(secData, true, false))

		ans := "r"
		for {
			fmt.Printf("Write Sector? [y]es, [n]o, [r]eroll: [%s] ", ans)
			fmt.Scanf("%s", &ans)
			if ansRegex.MatchString(ans) {
				switch ans {
				case "y":
					if err := os.Mkdir(secName, dirPerm); err != nil {
						log.Fatal(err)
					}

					if err := os.Chdir(secName); err != nil {
						log.Fatal(err)
					}

					for _, t := range strings.Split(exportTypes, ",") {
						if ex := export.New(t, secName, secData); ex != nil {
							if err := ex.Write(); err != nil {
								log.Fatal(err)
							}
						}
					}

					return

				case "n":
					return

				case "r":
					secData = sector.NewSector(secHeight, secWidth, excludeTags, poiChance, otherWorldChance)
					secName = genSectorName()
					fmt.Println(secName)
					fmt.Println(export.Hexmap(secData, true, false))
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

func init() {
	newCmd.AddCommand(sectorCmd)
	sectorCmd.Flags().StringArrayP(flExclude, "x", []string{}, "Exclude tags (-x zombies -x \"regional hegemon\" etc)")
	sectorCmd.Flags().IntP(flPoi, "p", 30, "Set % chance of a POI being generated for any given star in the sector")
	sectorCmd.Flags().IntP(flOW, "o", 10, "Set % chance for a secondary world to be generated for any given star in the sector")
	sectorCmd.Flags().IntP(flSecHeight, "e", 10, "Set height of sector in hexes")
	sectorCmd.Flags().IntP(flSecWidth, "w", 8, "Set width of sector in hexes")
	sectorCmd.Flags().String(flExport, "text", "Set export formats. (--export text,hugo,json) format types must be comma separated without spaces.")
}
