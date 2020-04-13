package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func BatchReplaceInFiles(path, old, new string) error {
	fmt.Println(old, new)
	if err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		matched, err := filepath.Match("*.go", info.Name())
		if err != nil {
			return err
		}
		if matched {
			read, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			newContents := strings.Replace(string(read), old, new, -1)

			err = ioutil.WriteFile(path, []byte(newContents), 0)
			if err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return err
	}
	return nil
}

func EnsureDir(fileName string) error {
	dirName := filepath.Dir(fileName)
	if _, err := os.Stat(dirName); err != nil {
		if err := os.MkdirAll(dirName, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func DirExists(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
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
