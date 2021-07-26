package utils

import (
	"os"
	"path/filepath"
	"strings"
)

// Create the given file path. The path to the directory will be created
// as well as the file itself.
func Create(path string) error {
	if strings.Contains(path, string(os.PathSeparator)) {
		folderPath := filepath.Dir(path)
		err := os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			return err
		}
	}
    f, err := os.Create(path)
    if err != nil {
        return err
    }
    f.Close()
    return nil
}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
