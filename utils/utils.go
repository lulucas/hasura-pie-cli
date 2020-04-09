package utils

import (
	"os"
	"path/filepath"
	"strings"
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

func Snake2Camel(word string) string {
	words := strings.Split(word, "_")
	return strings.Replace(strings.Title(strings.Join(words, " ")), " ", "", -1)
}

func SaveToFile(name string, s string) error {
	if err := EnsureDir(name); err != nil {
		return err
	}
	f, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	if _, err := f.Write([]byte(s)); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}
	return nil
}
