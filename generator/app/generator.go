package app

import (
	"github.com/lulucas/hasura-pie-cli/utils"
	"path/filepath"
)

var (
	files = map[string]string{
		"Dockerfile":              dockerfileTpl,
		"main.go":                 mainTpl,
		".rsync":                  rsyncTpl,
		".env":                    envTpl,
		".env.prod":               envProdTpl,
		"docker-compose.yml":      dockerComposeTpl,
		"docker-compose.prod.yml": dockerComposeProdTpl,
		"config.yaml":             configYamlTpl,
	}
)

func GenerateApp(path string) error {
	for filename, content := range files {
		if err := utils.SaveToFile(filepath.Join(path, filename), content); err != nil {
			return err
		}
	}
	return nil
}
