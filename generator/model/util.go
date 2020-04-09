package model

import (
	"os"
	"strings"
)

const id = "Id"

func snake2Camel(word string) string {
	words := strings.Split(word, "_")
	return strings.Replace(strings.Title(strings.Join(words, " ")), " ", "", -1)
}

func saveToFile(name string, b []byte) error {
	f, err := os.OpenFile(name+".go", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	if _, err := f.Write(b); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}
	return nil
}
