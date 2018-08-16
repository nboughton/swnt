package export

import (
	"fmt"

	"github.com/nboughton/go-utils/json/file"
	"github.com/nboughton/swnt/gen/sector"
)

// JSON represents the Exporter for JSON data
type JSON struct {
	Name  string
	Stars *sector.Stars
}

func (j *JSON) Write() error {
	fmt.Println("Exporting as json...")

	return file.Write(j.Name+".json", j.Stars)
}
