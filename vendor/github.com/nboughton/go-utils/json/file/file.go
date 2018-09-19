// Package file provides utility methods for reading/parsing and writing json data
// to and from files.
package file

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
)

// Scan decodes a JSON file into a struct pointer
func Scan(file string, o interface{}) error {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	return json.NewDecoder(bytes.NewReader(b)).Decode(&o)
}

// Write marshalls json data in order to output it to a file. If the file exists it
// is truncated to 0 bytes so that all data is overwritten.
func Write(path string, v interface{}) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	_, err = f.Write(b)
	return err
}
