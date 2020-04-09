package module

import (
	"github.com/lulucas/hasura-pie-cli/errors"
	"github.com/lulucas/hasura-pie-cli/utils"
	"os"
	"path/filepath"
	"text/template"
)

func GenerateModule(path string) error {
	list := filepath.SplitList(path)
	if len(list) < 2 {
		path = filepath.Join(list[0], path)
	}
	filename := path + ".go"

	if utils.FileExists(filename) {
		return errors.ErrAlreadyExists
	}

	if err := utils.EnsureDir(filename); err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	t, err := template.New("module").Parse(Module)
	if err != nil {
		return err
	}

	if err := t.Execute(f, map[string]string{
		"Module": filepath.Base(path),
	}); err != nil {
		return err
	}

	return nil
}
