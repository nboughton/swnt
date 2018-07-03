// Package haxscii is a simple libary that overlays cell text over a static template
package haxscii

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type colourFunc func(string, ...interface{}) string

var (
	coords = regexp.MustCompile(`\d{2},\d{2}`)
	offset = 5
)

// Colour funcs
var (
	White   = color.WhiteString
	Red     = color.RedString
	Yellow  = color.YellowString
	Magenta = color.MagentaString
	Green   = color.GreenString
	Blue    = color.BlueString
	Cyan    = color.CyanString
)

func genCrdText(row, col int) string {
	rStr, cStr := strconv.Itoa(row), strconv.Itoa(col)
	if row < 10 {
		rStr = "0" + rStr
	}
	if col < 10 {
		cStr = "0" + cStr
	}

	return fmt.Sprintf("%s,%s", rStr, cStr)
}

type cell struct {
	text        [][]string
	widthTop    int
	widthMiddle int
	height      int
}

/* A cell:
`  \__________/
  /r         \
 /            \
/              \
\              /
 \            /
	\__________/  `,
*/

func newCell(row, col int) *cell {
	c := &cell{
		text: [][]string{
			{" ", " ", `\`, "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "/", " ", " "},
			{" ", " ", "/", "r", " ", " ", " ", " ", " ", " ", " ", " ", " ", `\`, " ", " "},
			{" ", "/", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", `\`, " "},
			{"/", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", `\`},
			{`\`, " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "/"},
			{" ", `\`, " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "/", " "},
			{" ", " ", `\`, "_", "_", "_", "_", "_", "_", "_", "_", "_", "_", "/", " ", " "},
		},
		height:      7,
		widthTop:    10,
		widthMiddle: 16,
	}

	c.setCrds(row, col)

	return c
}

func (c *cell) setCrds(row, col int) *cell {
	text := genCrdText(row, col)

	for row := range c.text {
		for col, char := range c.text[row] {
			if char == "r" {
				for j, sub := range text {
					c.text[row][col+j] = string(sub)
				}
				return c
			}
		}
	}

	return c
}

// Return string array of column characters
func (c *cell) col(n int) []string {
	if n < 0 || n >= c.widthMiddle {
		return []string{}
	}

	chars := []string{}
	for _, row := range c.text {
		chars = append(chars, row[n])
	}

	return chars
}

// Return string array of row characters
func (c *cell) row(n int) []string {
	if n < 0 || n >= c.height {
		return []string{}
	}

	return c.text[n]
}

// Map represents a 2 dimensional string matrix of the template
type Map [][]string

// NewMap generates a mapscii template so that text can be superimposed on it
func NewMap(height, width int) Map {
	var (
		cl    = newCell(0, 0) // Use a cell as a reference
		wDiff = (cl.widthMiddle - cl.widthTop) / 2
		w     = (width * (cl.widthMiddle - wDiff)) + wDiff
		h     = (height * (cl.height - 1)) + cl.height/2 + 1 // Shared borders reduce total height
		m     = make(Map, h)
	)

	// Create blank template space
	for r := 0; r < h; r++ {
		m[r] = make([]string, w)
		for c := range m[r] {
			m[r][c] = " "
		}
	}

	for r := 0; r < height; r++ {
		row := r * (cl.height - 1)

		for c := 0; c < width; c++ {
			col := c * (cl.widthMiddle - 3)
			if c%2 != 0 {
				row = r*(cl.height-1) + cl.height/2
			}

			m.blankCell(row, col, r, c)
			if c%2 != 0 {
				row = row - cl.height/2
			}
		}
	}

	return m
}

// blankCell writes a blank cell to the Map matrix
func (m Map) blankCell(row, col, rLabel, cLabel int) Map {
	r, c := row, col
	cl := newCell(rLabel, cLabel)

	for cellRow := 0; cellRow < cl.height; cellRow++ {
		for _, char := range cl.row(cellRow) {
			if r >= len(m) || c >= len(m[r]) { // Bounds check matrices references
				return m
			}

			m[r][c] = char
			c++
		}
		r++
		c = col
	}

	return m
}

// SetTxt sets the text of a given hex
func (m Map) SetTxt(row, col int, lines [4]string, color colourFunc) {
	crds := genCrdText(row, col)
	cLen := len(crds)

	for r, line := range m {
		for c := range line {
			if c+len(lines) < len(line) {
				crd := strings.Join(m[r][c:c+cLen], "")
				if coords.MatchString(crd) && crd == crds {
					for i, line := range lines {
						m.print(r+i+1, c+offset-(len(line)/2), line, color)
					}
				}
			}
		}
	}
}

func (m Map) print(startRow, startCol int, text string, colour colourFunc) {
	for row, col, i := startRow, startCol, 0; i < len(text); col, i = col+1, i+1 {
		if col < 0 {
			col = 0
		}

		if col < len(m[row]) {
			m[row][col] = colour(string(text[i]))
		} else {
			m[row] = append(m[row], colour(string(text[i])))
		}
	}
}

func (m Map) String() string {
	b := new(bytes.Buffer)

	for _, line := range m {
		fmt.Fprintf(b, "%s\n", strings.Join(line, ""))
	}

	return b.String()
}

// Colour toggles the colour output on/off (true/false)
func Colour(b bool) {
	color.NoColor = !b
}
