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
	"strings"

	"github.com/nboughton/swnt/content"
	"github.com/nboughton/swnt/content/format"
	"github.com/spf13/cobra"
)

// encounterCmd represents the encounter command
var encounterCmd = &cobra.Command{
	Use:   "encounter",
	Short: "Generate a quick encounter",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		wild, _ := cmd.Flags().GetBool(flWilderness)
		fmc, _ := cmd.Flags().GetString(flFormat)

		e := content.NewEncounter(wild)
		for _, f := range strings.Split(fmc, ",") {
			fID, err := format.Find(f)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Fprintf(tw, e.Format(fID))
			fmt.Fprintln(tw)
			tw.Flush()
		}
	},
}

func init() {
	newCmd.AddCommand(encounterCmd)
	encounterCmd.Flags().BoolP(flWilderness, "w", false, "Set encounter type to Wilderness (default is Urban)")
	encounterCmd.Flags().StringP(flFormat, "f", "text", "Set output format. (--format text,markdown)")
}
