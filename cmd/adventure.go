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

	"github.com/nboughton/swnt/gen/adventure"
	"github.com/nboughton/swnt/gen/world"
	"github.com/spf13/cobra"
)

const (
	flTag  = "tag"
	flTags = "tags"
)

// adventureCmd represents the adventure command
var adventureCmd = &cobra.Command{
	Use:   "adventure",
	Short: "Generate an Adventure",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		tag, _ := cmd.Flags().GetString(flTag)
		list, _ := cmd.Flags().GetBool(flTags)

		if list {
			for _, t := range world.Tags {
				fmt.Println(t.Name)
			}
			return
		}

		fmt.Println(adventure.Generate(tag))
	},
}

func init() {
	newCmd.AddCommand(adventureCmd)
	adventureCmd.Flags().StringP(flTag, "t", "Zombies", "Specify world tag to base roll on.")
	adventureCmd.Flags().BoolP(flTags, "l", false, "List available tags.")
}
