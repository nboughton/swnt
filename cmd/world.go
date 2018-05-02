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

	"github.com/nboughton/swnt/gen/culture"
	"github.com/nboughton/swnt/gen/world"
	"github.com/spf13/cobra"
)

const (
	flExclude = "exclude"
)

// worldCmd represents the world command
var worldCmd = &cobra.Command{
	Use:   "world",
	Short: "Generate a secondary World for a Sector cell",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			ctr, _ = cmd.Flags().GetString(flCulture)
			exc, _ = cmd.Flags().GetStringArray(flExclude)
			cID    culture.Culture
			err    error
		)

		if ctr == "" {
			cID = culture.Random()
		} else {
			cID, err = culture.Find(ctr)
			if err != nil {
				fmt.Printf("No Culture found for \"%s\", options are %v\n", ctr, culture.Cultures)
				return
			}
		}

		fmt.Fprintf(tw, world.New(cID, false, exc).String())
		tw.Flush()
	},
}

func init() {
	newCmd.AddCommand(worldCmd)
	worldCmd.Flags().StringP(flCulture, "c", "", "Set Culture of world")
	worldCmd.Flags().StringArrayP(flExclude, "x", []string{}, "Exclude tags (-x zombies -x \"regional hegemon\" etc)")
}
