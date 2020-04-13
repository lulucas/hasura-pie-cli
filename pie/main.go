package main

import (
	"github.com/jinzhu/configor"
	"github.com/lulucas/hasura-pie-cli/errors"
	"github.com/lulucas/hasura-pie-cli/generator/app"
	"github.com/lulucas/hasura-pie-cli/generator/ci"
	"github.com/lulucas/hasura-pie-cli/generator/module"
	"github.com/lulucas/hasura-pie-cli/generator/project"
	"github.com/lulucas/hasura-pie-cli/utils"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

type Config struct {
	Sync SyncConfig
}

type SyncConfig struct {
	Module []module.SyncModuleConfig
}

var (
	config Config
)

const (
	configFile = "pie.yml"
)

func main() {
	if utils.FileExists(configFile) {
		if err := configor.Load(&config, configFile); err != nil {
			log.Fatal(err)
		}
	}
	a := &cli.App{
		Name:    "pie",
		Usage:   "hasura-pie cli",
		Version: "0.1.5",
		Commands: []*cli.Command{
			{
				Name:      "init",
				Usage:     "initialize a project",
				ArgsUsage: "todo",
				Action: func(c *cli.Context) error {
					if c.NArg() < 1 {
						return errors.ErrMissingPath
					}
					return project.GenerateProject(c.Args().First())
				},
			},
			{
				Name:    "generate",
				Aliases: []string{"g"},
				Usage:   "generate code",
				Subcommands: []*cli.Command{
					{
						Subcommands: []*cli.Command{
							{
								Name: "github",
								Action: func(c *cli.Context) error {
									if c.NArg() < 1 {
										return errors.ErrMissingPath
									}
									return ci.GenerateGithubAction(c.Args().First())
								},
							},
						},
					},
					{
						Name:      "app",
						Aliases:   []string{"a"},
						Usage:     "generate app",
						ArgsUsage: "path to generate",
						Action: func(c *cli.Context) error {
							if c.NArg() < 1 {
								return errors.ErrMissingPath
							}
							return app.GenerateApp(c.Args().First())
						},
					},
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
						Name:      "module",
						Aliases:   []string{"m"},
						Usage:     "sync module from git",
						ArgsUsage: "[git module path] [local path]",
						Action: func(c *cli.Context) error {
							if c.NArg() < 1 {
								// sync from config
								if err := module.SyncGit(); err != nil {
									return err
								}
								return module.SyncModuleByConfig(config.Sync.Module)
							}
							// specific path
							localPath := c.Args().First()
							if c.NArg() > 1 {
								localPath = c.Args().Get(1)
							}
							if err := module.SyncGit(); err != nil {
								return err
							}
							return module.SyncModule(c.Args().First(), localPath)
						},
					},
				},
			},
		},
	}

	a.EnableBashCompletion = true

	err := a.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
