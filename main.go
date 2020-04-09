//go:generate go build -o example/pie.exe
package main

import (
	"github.com/jinzhu/configor"
	"github.com/lulucas/hasura-pie-cli/v1/errors"
	"github.com/lulucas/hasura-pie-cli/v1/generator/model"
	"github.com/lulucas/hasura-pie-cli/v1/generator/module"
	"github.com/lulucas/hasura-pie-cli/v1/utils"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

type Config struct {
	Postgres model.Options
}

var (
	config Config
)

const (
	defaultConfigFilename = "config.yml"
)

func main() {
	if utils.FileExists(defaultConfigFilename) {
		if err := configor.Load(&config, defaultConfigFilename); err != nil {
			log.Fatal(err)
		}
	}
	app := &cli.App{
		Name:    "pie",
		Usage:   "hasura-pie cli",
		Version: "0.0.1",
		Commands: []*cli.Command{
			{
				Name:    "generate",
				Aliases: []string{"g"},
				Usage:   "generate code",
				Subcommands: []*cli.Command{
					{
						Name:      "module",
						Aliases:   []string{"m"},
						Usage:     "generate module",
						ArgsUsage: "path to generate",
						Action: func(c *cli.Context) error {
							if c.NArg() < 1 {
								return errors.ErrMissingPath
							}
							return module.GenerateModule(c.Args().First())
						},
					},
				},
			},
			{
				Name:    "sync",
				Aliases: []string{"s"},
				Usage:   "sync code from hasura or postgres",
				Subcommands: []*cli.Command{
					{
						Name:      "model",
						Aliases:   []string{"m"},
						Usage:     "sync postgres table to model struct",
						ArgsUsage: "tables to sync, sync all tables if empty",
						Action: func(c *cli.Context) error {
							if c.NArg() == 0 {
								model.GenerateModel(config.Postgres)
							} else {
								model.GenerateModel(config.Postgres, c.Args().Slice()...)
							}
							return nil
						},
					},
				},
			},
		},
	}

	app.EnableBashCompletion = true

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
