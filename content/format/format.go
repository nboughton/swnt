package format

import (
	"bytes"
	"fmt"
	"strings"
)

// OutputType defines identifiers for output style categories such as text and markdown
type OutputType int

// OutputType constants
const (
	TEXT OutputType = iota
	MARKDOWN
)

var types = map[string]OutputType{
	"text":     TEXT,
	"markdown": MARKDOWN,
}

// Find attempts retrieve the OutputType value from string
func Find(t string) (OutputType, error) {
	if o, ok := types[strings.ToLower(t)]; ok {
		return o, nil
	}

	return 0, fmt.Errorf("no such output type")
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
// tables must start with a header though. If headers are omitted "name" can also be omitted.
func Table(t OutputType, header bool, name string, rows [][]string) string {
	buf, sep, rowTmpl := new(bytes.Buffer), "", ""

	switch t {
	case TEXT:
		sep, rowTmpl = "\t:\t", "%s\n"

		if header {
			fmt.Fprintf(buf, "%s\n", name)
		}
	case MARKDOWN:
		sep, rowTmpl = " | ", "| %s |\n"

		if header {
			cells := make([]string, len(rows[0]))
			fmt.Fprintf(buf, "| %s %s |\n| %s --- |\n", name, strings.Join(cells, sep), strings.Join(cells, " --- |"))
		}
	}

	for _, row := range rows {
		fmt.Fprintf(buf, rowTmpl, strings.Join(row, sep))
	}

	return buf.String()
}
