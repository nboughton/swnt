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
	"math/rand"
	"strings"
	"time"

	"github.com/nboughton/swnt/content"
	"github.com/nboughton/swnt/content/culture"
	"github.com/nboughton/swnt/content/format"
	"github.com/nboughton/swnt/content/gender"
	"github.com/spf13/cobra"
)

// npcCmd represents the person command
var npcCmd = &cobra.Command{
	Use:   "npc",
	Short: "Generate a NPC",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			clt, _      = cmd.Flags().GetString(flCulture)
			gdr, _      = cmd.Flags().GetString(flGender)
			isPatron, _ = cmd.Flags().GetBool(flPatron)
			fmc, _      = cmd.Flags().GetString(flFormat)
			err         error
		)

		cID, err := culture.Find(clt)
		if err != nil {
			fmt.Println(err)
			return
		}

		gID, err := gender.Find(gdr)
		if err != nil {
			fmt.Println(err)
			return
		}

		n := content.NewNPC(cID, gID, isPatron)
		for _, f := range strings.Split(fmc, ",") {
			fID, err := format.Find(f)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Fprintf(tw, n.Format(fID))
			fmt.Fprintln(tw)
			tw.Flush()
		}
	},
}

func init() {
	rand.Seed(time.Now().UnixNano())

	newCmd.AddCommand(npcCmd)
	npcCmd.Flags().StringP(flCulture, "c", "any", fmt.Sprintf("Select Culture, choices are: %v", culture.Cultures))
	npcCmd.Flags().StringP(flGender, "g", "", fmt.Sprintf("Select Gender, choices are: %v", gender.Genders))
	npcCmd.Flags().BoolP(flPatron, "p", false, "NPC is a Patron")
}
