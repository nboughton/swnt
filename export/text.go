package export

import (
	"fmt"
	"io/ioutil"
	"os"

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
		ioutil.WriteFile(starsDir+"/"+system.Name+".txt", []byte(system.Format(format.TEXT)), filePerm)
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
