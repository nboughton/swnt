package export

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/nboughton/swnt/content/format"
	"github.com/nboughton/swnt/content/sector"
)

// Hugo represents the Exporter for Hugo projects
type Hugo struct {
	Name  string
	Stars *sector.Stars
}

// Write satisfies the Setup requirement of the Exporter interface
func (h *Hugo) Write() error {
	fmt.Println("Exporting as hugo site...")
	wdir, _ := os.Getwd()

	hugoDir := "hugo"
	if err := os.Mkdir(hugoDir, dirPerm); err != nil {
		return err
	}

	if err := os.Chdir(hugoDir); err != nil {
		return err
	}

	fmt.Println("Creating new hugo site...")
	o, err := exec.Command("hugo", "new", "site", ".").CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Print(string(o))

	o, err = exec.Command("git", "init").CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Print(string(o))

	o, err = exec.Command("git", "submodule", "add", "https://github.com/nboughton/hugo-theme-docdock.git", "themes/docdock").CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Print(string(o))

	o, err = exec.Command("git", "submodule", "init").CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Print(string(o))

	o, err = exec.Command("git", "submodule", "update").CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Print(string(o))

	fmt.Println("Copying config...")
	_, err = exec.Command("cp", "themes/docdock/exampleSite/config.toml", ".").CombinedOutput()
	if err != nil {
		return err
	}

	fmt.Println("Setting Title...")
	_, err = exec.Command("sed", "-i", fmt.Sprintf("s/TITLE/%s/", h.Name), "config.toml").CombinedOutput()
	if err != nil {
		return err
	}

	fmt.Println("Copying in default archetype...")
	_, err = exec.Command("cp", "themes/docdock/archetypes/default.md", "archetypes/").CombinedOutput()
	if err != nil {
		return err
	}

	fmt.Println("Creating Stars dir...")
	starsDir := "content/Stars"
	if err := os.Mkdir(starsDir, dirPerm); err != nil {
		return err
	}

	fmt.Println("Populating Stars dir...")
	for _, star := range h.Stars.Systems {
		// Create article stub
		o, err := exec.Command("hugo", "new", fmt.Sprintf("Stars/%s.md", star.Name)).CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Print(string(o))

		// Append Star data using Markdown
		f, err := os.OpenFile(fmt.Sprintf("%s/%s.md", starsDir, star.Name), os.O_APPEND|os.O_WRONLY, filePerm)
		if err != nil {
			return err
		}

		if _, err := f.Write([]byte(star.Format(format.MARKDOWN))); err != nil {
			return err
		}

		f.Close()
	}

	// Print hexmap to index.md
	o, err = exec.Command("hugo", "new", "_index.md").CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(string(o))

	f, err := os.OpenFile("content/_index.md", os.O_APPEND|os.O_WRONLY, filePerm)
	if err != nil {
		return err
	}

	if _, err := f.Write([]byte("# " + h.Name + "\n\n```\n" + Hexmap(h.Stars, false, false) + "\n```")); err != nil {
		return err
	}

	return os.Chdir(wdir)
}
