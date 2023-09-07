package sleeptalk

import (
	"fmt"
	"os"
	"log"
)

//Reads a file and returns its content as a string
func Readfile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

//Wites data to a file
func WriteFile(filename, data string) error {
	err := os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return "", err
	}
	return nil
}

//Reads the named directory, returning all its directory entries sorted by filename
func ReadDir(dirname string) ([]string, error) {
	files, err := os.ReadDir(dirname)
	if err != nil {
		return "", err
	}

	return filenames, nil
}