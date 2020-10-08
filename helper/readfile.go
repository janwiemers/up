package helper

import (
	"fmt"
	"io/ioutil"
)

// ReadFile opens and reads a given file
func ReadFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil
	}
	return data
}
