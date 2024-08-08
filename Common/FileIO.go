package Common

import (
	"fmt"
	"io/ioutil"
	"os"
)

// WriteToFile writes the provided content to the specified file
func WriteToFile(filename, content string) error {
	// Write the content to the file
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("failed to write to file %s: %w", filename, err)
	}
	return nil
}

// ReadFromFile reads the content from the specified file
func ReadFromFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to read from file %s: %w", filename, err)
	}
	return string(content), nil
}

// CreateFileIfNotExist creates the file if it does not already exist
func CreateFileIfNotExist(filename string) error {
	if !fileExists(filename) {
		err := createFile(filename)
		if err != nil {
			return err
		}
	}
	return nil
}

// fileExists checks if a file exists and is not a directory
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// createFile creates a new file
func createFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filename, err)
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			err = cerr
		}
	}()
	return nil
}
