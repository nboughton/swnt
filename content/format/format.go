// Package format provides formatting functions for text output
package format

import (
	"bytes"
	"fmt"
	"strings"
)

// OutputType defines identifiers for output style categories such as text and markdown
type OutputType string

// OutputType constants
const (
	TEXT     OutputType = "txt"
	MARKDOWN OutputType = "md"
)

// Types of format output currently supported
var Types = []OutputType{TEXT, MARKDOWN}

func (o OutputType) String() string {
	return string(o)
}

// Find attempts retrieve the OutputType value from string
func Find(name string) (OutputType, error) {
	for _, t := range Types {
		if strings.ToLower(name) == t.String() {
			return t, nil
		}
	}

	return "", fmt.Errorf("format not supported, available output types are: %s", Types)
}

// Header formats and returns a header in format t, header sizes are defined in HTML/MARKDOWN terms with 1 being the largest
// and reducing in size as the number increases.
func Header(t OutputType, size int, text string) string {
	out := ""

	switch t {
	case TEXT:
		out = text + "\n"

	case MARKDOWN:
		out = fmt.Sprintf("%s %s\n\n", strings.Repeat("#", size), text)
	}

	return out
}

// Table returns a formatted Table of type t. Headers are optional so that different bits of content
// can be concatenated into a single table through multiple calls to Table. Bear in mind that Markdown
// tables must start with a header though.
func Table(t OutputType, headers []string, rows [][]string) string {
	buf, sep, rowTmpl := new(bytes.Buffer), "", ""

	if len(rows) == 0 {
		return "no table data found"
	}

	switch t {
	case TEXT:
		sep, rowTmpl = "\t:\t", "%s\n"

		if len(headers) > 0 {
			fmt.Fprintf(buf, "%s\n", strings.Join(headers, sep))
		}
	case MARKDOWN:
		sep, rowTmpl = " | ", "| %s |\n"

		if len(headers) > 0 {
			cells := make([]string, len(rows[0]))
			fmt.Fprintf(buf, "| %s |\n| %s --- |\n", strings.Join(headers, " | "), strings.Join(cells, " --- |"))
		}
	}

	for _, row := range rows {
		fmt.Fprintf(buf, rowTmpl, strings.Join(row, sep))
	}

	return buf.String()
}
