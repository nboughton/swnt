// Package haxscii is a simple libary that overlays cell text over a static template
package haxscii

import (
	"bytes"
	"fmt"
	"regexp"
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

func strToMatrix(s string) [][]string {
	var m [][]string

	for row, line := range strings.Split(s, "\n") {
		m = append(m, []string{})
		for _, char := range line {
			m[row] = append(m[row], string(char))
		}
	}

	return m
}

type cell struct {
	raw         string
	widthTop    int
	widthMiddle int
	height      int
}

var tmpl = cell{
	raw: `  \__________/  
  /rr,cc     \  
 /            \ 
/              \
\              /
 \            / 
  \__________/  `,
	height:      7,
	widthTop:    10,
	widthMiddle: 16,
}

func newCell(row, col int) cell {
	c := cell{
		raw: `  \__________/  
  /r         \  
 /            \ 
/              \
\              /
 \            / 
  \__________/  `,
		height:      7,
		widthTop:    10,
		widthMiddle: 16,
	}

	return c.setCrds(row, col)
}

func (c cell) setCrds(row, col int) cell {
	text := genCrdText(row, col)
	strA := make([]string, len(c.raw))

	// Write string Array
	for i, char := range c.raw {
		strA[i] = string(char)
	}

	// Rewrite values
	for i, char := range strA {
		if string(char) == "r" {
			for j, sub := range text {
				//fmt.Println(i+j, string(sub), strA[i+j])
				strA[i+j] = string(sub)
			}
		}
	}

	c.raw = strings.Join(strA, "")
	return c
}

func genCrdText(row, col int) string {
	rStr, cStr := fmt.Sprintf("%d", row), fmt.Sprintf("%d", col)
	if row < 10 {
		rStr = fmt.Sprintf("0%d", row)
	}
	if col < 10 {
		cStr = fmt.Sprintf("0%d", col)
	}

	return fmt.Sprintf("%s,%s", rStr, cStr)
}

// Return string array of column characters
func (c cell) col(n int) []string {
	if n > c.widthMiddle {
		return []string{}
	}

	chars := []string{}
	for _, row := range strToMatrix(c.raw) {
		chars = append(chars, row[n])
	}

	return chars
}

// Return string array of row characters
func (c cell) row(n int) []string {
	if n > c.height {
		return []string{}
	}

	return strToMatrix(c.raw)[n]
}

// Map represents a 2 dimensional string array of the template
type Map [][]string

// NewMap generates a mapscii template so that text can be superimposed on it
func NewMap(height, width int) Map {
	h := (height * (tmpl.height - 1)) + tmpl.height/2 + 1 // Shared borders reduce total height
	w := width * (tmpl.widthMiddle - 2)
	m := make(Map, h)

	// Create blank template space
	for r := 0; r < h; r++ {
		m[r] = make([]string, w)
		for c := range m[r] {
			m[r][c] = " "
		}
	}

	for r := 0; r < height; r++ {
		row := r * (tmpl.height - 1)

		for c := 0; c < width; c++ {
			col := c * (tmpl.widthMiddle - 3)
			if c%2 != 0 {
				row = r*(tmpl.height-1) + tmpl.height/2
			}

			m.blankCell(row, col, r, c)
			if c%2 != 0 {
				row = row - tmpl.height/2
			}
		}
	}

	return m
}

// blankCell sets writes a blank cell over the map coordinate space listed
func (m Map) blankCell(row, col, rLabel, cLabel int) Map {
	r, c := row, col
	cl := newCell(rLabel, cLabel)

	for cellRow := 0; cellRow < cl.height; cellRow++ {
		for _, char := range cl.row(cellRow) {
			if r >= len(m) || c >= len(m[r]) { // Bounds check, motherfucker.
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

	for r, line := range m {
		for c := range line {
			if c+4 < len(line) {
				crd := strings.Join(m[r][c:c+5], "")
				if coords.MatchString(crd) && crd == genCrdText(row, col) {
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
