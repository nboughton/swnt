package export

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"text/tabwriter"

	"github.com/nboughton/swnt/content/format"
	"github.com/nboughton/swnt/content/sector"
)

// Text represents the Exporter for text based output
type Text struct {
	Name  string
	Stars *sector.Stars
}

func (t *Text) Write() error {
	fmt.Println("Exporting as plain text...")
	wdir, _ := os.Getwd()

	textDir := "text"
	if err := os.Mkdir(textDir, dirPerm); err != nil {
		return err
	}

	if err := os.Chdir(textDir); err != nil {
		return err
	}

	fmt.Println("Creating Stars dir...")
	starsDir := "Stars"
	if err := os.Mkdir(starsDir, dirPerm); err != nil {
		return err
	}

	for _, system := range t.Stars.Systems {
		buf := new(bytes.Buffer)
		tab := tabwriter.NewWriter(buf, 1, 2, 1, ' ', 0)

		fmt.Fprint(tab, system.Format(format.TEXT))
		tab.Flush()

		ioutil.WriteFile(starsDir+"/"+system.Name+".txt", buf.Bytes(), filePerm)
	}

	mapDir := "Maps"
	if err := os.Mkdir(mapDir, dirPerm); err != nil {
		return err
	}

	ioutil.WriteFile(mapDir+"/gm-map.txt", []byte(Hexmap(t.Stars, false, false)), filePerm)
	ioutil.WriteFile(mapDir+"/pc-map.txt", []byte(Hexmap(t.Stars, false, true)), filePerm)
	ioutil.WriteFile(mapDir+"/gm-map-ansi.txt", []byte(Hexmap(t.Stars, true, false)), filePerm)
	ioutil.WriteFile(mapDir+"/pc-map-ansi.txt", []byte(Hexmap(t.Stars, true, true)), filePerm)

	return os.Chdir(wdir)
}
