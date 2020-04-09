package utils

import (
	"os"
	"path/filepath"
)

func EnsureDir(fileName string) error {
	dirName := filepath.Dir(fileName)
	if _, err := os.Stat(dirName); err != nil {
		if err := os.MkdirAll(dirName, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
