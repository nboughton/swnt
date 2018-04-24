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

	"github.com/nboughton/swnt/gen/culture"
	"github.com/nboughton/swnt/gen/npc"
	"github.com/spf13/cobra"
)

const (
	flCulture = "culture"
	flGender  = "gender"
	flPatron  = "is-patron"
)

// npcCmd represents the person command
var npcCmd = &cobra.Command{
	Use:   "npc",
	Short: "Generate a NPC",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			clt, _      = cmd.Flags().GetString(flCulture)
			gender, _   = cmd.Flags().GetString(flGender)
			isPatron, _ = cmd.Flags().GetBool(flPatron)
			err         error
			cID         culture.ID
			gID         npc.GenderID
		)

		if clt == "" {
			cID, clt = culture.Random()
		} else {
			cID, err = culture.IDByName(clt)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		if gender == "" {
			gID = npc.Any
		} else {
			gID = npc.GenderID(byte(gender[0]))
		}

		fmt.Fprintf(tw, npc.New(cID, gID, isPatron).String())
		tw.Flush()
	},
}

func init() {
	rand.Seed(time.Now().UnixNano())

	newCmd.AddCommand(npcCmd)
	npcCmd.Flags().StringP(flCulture, "c", "", "Select Culture, choices are: "+strings.Join(culture.Cultures, ", "))
	npcCmd.Flags().StringP(flGender, "g", "", "Select Gender, choices are m, f, a, y (male, female, androgynous, any")
	npcCmd.Flags().BoolP(flPatron, "p", false, "NPC is a Patron")
}
