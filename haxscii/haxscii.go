// Package haxscii is a simple libary that overlays cell text over a static template
package haxscii

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

type colourFunc func(string, ...interface{}) string

var (
	coords = regexp.MustCompile(`\d,\d`)
	offset = 5
)

var mapscii = `   __________                __________                __________                __________               
  /0,0       \              /2,0       \              /4,0       \              /6,0       \              
 /            \            /            \            /            \            /            \            
/              \__________/              \__________/              \__________/              \__________
\              /1,0       \              /3,0       \              /5,0       \              /7,0       \
 \            /            \            /            \            /            \            /            \
  \__________/              \__________/              \__________/              \__________/              \
  /0,1       \              /2,1       \              /4,1       \              /6,1       \              /
 /            \            /            \            /            \            /            \            /
/              \__________/              \__________/              \__________/              \__________/
\              /1,1       \              /3,1       \              /5,1       \              /7,1       \
 \            /            \            /            \            /            \            /            \
  \__________/              \__________/              \__________/              \__________/              \
  /0,2       \              /2,2       \              /4,2       \              /6,2       \              /
 /            \            /            \            /            \            /            \            /
/              \__________/              \__________/              \__________/              \__________/
\              /1,2       \              /3,2       \              /5,2       \              /7,2       \
 \            /            \            /            \            /            \            /            \
  \__________/              \__________/              \__________/              \__________/              \
  /0,3       \              /2,3       \              /4,3       \              /6,3       \              /
 /            \            /            \            /            \            /            \            /
/              \__________/              \__________/              \__________/              \__________/
\              /1,3       \              /3,3       \              /5,3       \              /7,3       \
 \            /            \            /            \            /            \            /            \
  \__________/              \__________/              \__________/              \__________/              \
  /0,4       \              /2,4       \              /4,4       \              /6,4       \              /
 /            \            /            \            /            \            /            \            /
/              \__________/              \__________/              \__________/              \__________/
\              /1,4       \              /3,4       \              /5,4       \              /7,4       \
 \            /            \            /            \            /            \            /            \
  \__________/              \__________/              \__________/              \__________/              \
  /0,5       \              /2,5       \              /4,5       \              /6,5       \              /
 /            \            /            \            /            \            /            \            /
/              \__________/              \__________/              \__________/              \__________/
\              /1,5       \              /3,5       \              /5,5       \              /7,5       \
 \            /            \            /            \            /            \            /            \
  \__________/              \__________/              \__________/              \__________/              \
  /0,6       \              /2,6       \              /4,6       \              /6,6       \              /
 /            \            /            \            /            \            /            \            /
/              \__________/              \__________/              \__________/              \__________/
\              /1,6       \              /3,6       \              /5,6       \              /7,6       \
 \            /            \            /            \            /            \            /            \
  \__________/              \__________/              \__________/              \__________/              \
  /0,7       \              /2,7       \              /4,7       \              /6,7       \              /
 /            \            /            \            /            \            /            \            /
/              \__________/              \__________/              \__________/              \__________/
\              /1,7       \              /3,7       \              /5,7       \              /7,7       \
 \            /            \            /            \            /            \            /            \
  \__________/              \__________/              \__________/              \__________/              \
  /0,8       \              /2,8       \              /4,8       \              /6,8       \              /
 /            \            /            \            /            \            /            \            /
/              \__________/              \__________/              \__________/              \__________/
\              /1,8       \              /3,8       \              /5,8       \              /7,8       \
 \            /            \            /            \            /            \            /            \
  \__________/              \__________/              \__________/              \__________/              \
  /0,9       \              /2,9       \              /4,9       \              /6,9       \              /
 /            \            /            \            /            \            /            \            /
/              \__________/              \__________/              \__________/              \__________/
\              /1,9       \              /3,9       \              /5,9       \              /7,9       \
 \            /            \            /            \            /            \            /            \
  \__________/              \__________/              \__________/              \__________/              \
             \              /          \              /          \              /          \              /
              \            /            \            /            \            /            \            /
               \__________/              \__________/              \__________/              \__________/`

// Map represents a 2 dimensional string array of the template
type Map [][]string

// NewMap converts the mapscii raw text into a Map
func NewMap() Map {
	var c Map
	for row, line := range strings.Split(mapscii, "\n") {
		c = append(c, []string{})
		for _, char := range line {
			c[row] = append(c[row], string(char))
		}
	}

	return c
}

// SetTxt sets the text of a given hex
func (m Map) SetTxt(row, col int, text, tag1, tag2, techlevel string) {
	for r, line := range m {
		for c := range line {
			if c+2 < len(line) {
				crd := strings.Join(m[r][c:c+3], "")
				if coords.MatchString(crd) && crd == fmt.Sprintf("%d,%d", row, col) {
					m.print(r+1, c+offset-(len(text)/2), text)
					m.print(r+2, c+offset-(len(tag1)/2), tag1)
					m.print(r+3, c+offset-(len(tag2)/2), tag2)
					m.print(r+4, c+offset-(len(techlevel)/2), techlevel)
				}
			}
		}
	}
}

func (m Map) print(startRow, startCol int, text string) {
	for row, col, i := startRow, startCol, 0; i < len(text); col, i = col+1, i+1 {
		if col < 0 {
			col = 0
		}

		if col < len(m[row]) {
			m[row][col] = string(text[i])
		} else {
			m[row] = append(m[row], string(text[i]))
		}
	}
}

// SetTxtColour sets the text of a given hex
func (m Map) SetTxtColour(row, col int, text, tag1, tag2, techlevel string, colour colourFunc) {
	for r, line := range m {
		for c := range line {
			if c+2 < len(line) {
				crd := strings.Join(m[r][c:c+3], "")
				if coords.MatchString(crd) && crd == fmt.Sprintf("%d,%d", row, col) {
					m.printColour(r+1, c+offset-(len(text)/2), text, colour)
					m.printColour(r+2, c+offset-(len(tag1)/2), tag1, colour)
					m.printColour(r+3, c+offset-(len(tag2)/2), tag2, colour)
					m.printColour(r+4, c+offset-(len(techlevel)/2), techlevel, colour)
				}
			}
		}
	}
}

func (m Map) printColour(startRow, startCol int, text string, colour colourFunc) {
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
