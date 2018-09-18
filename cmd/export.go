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
	"strings"

	"github.com/nboughton/go-utils/json/file"
	"github.com/nboughton/swnt/content/sector"
	"github.com/nboughton/swnt/export"
	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export a json dump to hugo or text",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		jsonFile, _ := cmd.Flags().GetString(flFile)
		exportTypes, _ := cmd.Flags().GetString(flExport)

		secName := strings.Replace(jsonFile, ".json", "", -1)

		secData := new(sector.Stars)
		if err := file.Scan(jsonFile, &secData); err != nil {
			fmt.Println("Error reading file. You may need to reformat the JSON data to make it more easily readable.", err)
			return
		}

		for _, t := range strings.Split(exportTypes, ",") {
			if exporter, err := export.New(t, secName, secData); exporter != nil {
				if err != nil {
					log.Fatal(err)
				}

				if err = exporter.Write(); err != nil {
					log.Fatal(err)
				}
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(exportCmd)
	exportCmd.Flags().StringP(flFile, "i", "", "Path to json file")
	exportCmd.Flags().StringP(flExport, "x", "hugo,txt", "Set export format")
}
