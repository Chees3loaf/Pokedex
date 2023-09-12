package Potion

import (
    "os"
    
)

// Reads a file and returns its content as a string
func Load(filename string) (string, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return "", err
    }
    return string(data), nil
}

// Writes data to a file
func Save(filename, data string) error {
    err := os.WriteFile(filename, []byte(data), 0644)
    if err != nil {
        return err // Return the error, not an empty string
    }
    return nil
}

// Reads the named directory, returning all its directory entries sorted by filename
func Map(dirname string) ([]string, error) {
    files, err := os.ReadDir(dirname)
    if err != nil {
        return nil, err // Return nil slice and error
    }

    var fileNames []string
    for _, file := range files {
        fileNames = append(fileNames, file.Name())
    }

    return fileNames, nil
}
