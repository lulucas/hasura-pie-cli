package project

import (
	"github.com/lulucas/hasura-pie-cli/generator/app"
	"github.com/lulucas/hasura-pie-cli/generator/ci"
	"github.com/lulucas/hasura-pie-cli/utils"
)

func GenerateProject(path string) error {
	if err := app.GenerateApp(path + "/app/business"); err != nil {
		return err
	}
	if err := ci.GenerateGithubAction(path + "/.github/workflows/ci.yml"); err != nil {
		return err
	}
	if err := utils.SaveToFile(path+"/go.mod", modTpl); err != nil {
		return err
	}
	if err := utils.SaveToFile(path+"/.gitignore", ignoreTpl); err != nil {
		return err
	}
	return nil
}
