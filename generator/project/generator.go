package project

import (
	"github.com/lulucas/hasura-pie-cli/generator/app"
	"github.com/lulucas/hasura-pie-cli/generator/ci"
	"github.com/lulucas/hasura-pie-cli/utils"
)

func GenerateProject() error {
	if err := app.GenerateApp("app/business"); err != nil {
		return err
	}
	if err := ci.GenerateGithubAction(); err != nil {
		return err
	}
	if err := utils.SaveToFile("go.mod", modTpl); err != nil {
		return err
	}
	if err := utils.SaveToFile(".gitignore", ignoreTpl); err != nil {
		return err
	}
	return nil
}
