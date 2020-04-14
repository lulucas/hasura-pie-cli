package project

import (
	"github.com/lulucas/hasura-pie-cli/generator/app"
	"github.com/lulucas/hasura-pie-cli/generator/ci"
	"github.com/lulucas/hasura-pie-cli/utils"
	"os/exec"
	"path/filepath"
)

var (
	files = map[string]string{
		"go.mod":                    modTpl,
		".gitignore":                ignoreTpl,
		"/app/business/config.yaml": configYamlTpl,
	}
)

func GenerateProject(path string) error {
	if err := exec.Command("hasura", "init", path+"/app/business").Run(); err != nil {
		return err
	}
	if err := app.GenerateApp(path + "/app/business"); err != nil {
		return err
	}
	if err := ci.GenerateGithubAction(path + "/.github/workflows/ci.yml"); err != nil {
		return err
	}
	for filename, content := range files {
		if err := utils.SaveToFile(filepath.Join(path, filename), content); err != nil {
			return err
		}
	}

	return nil
}
