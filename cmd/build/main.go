package main

import (
	"os"
)

type BuildEnvironment struct {
	environment string
	language    build.Language
	compiler    build.Compiler
}

func main() {
	cmd := cli.New(&cli.CLI{
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "git, g",
				Usage: "Remote git repository",
			},
			cli.BoolFlag{
				Name:  "production, prod",
				Usage: "Include to enable production environment",
			},
		},
		Commands: []cli.Command{
			{
				Name:    "init",
				Aliases: []string{"n"},
				Usage:   "Initialize project 'build' file in working directory",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "build",
				Aliases: []string{"compile", "b"},
				Usage:   "Build ",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "clean",
				Aliases: []string{"c"},
				Usage:   "Build ",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "issues",
				Aliases: []string{"todo", "tasks", "t"},
				Usage:   "Publish the compiled project",
				Action: func(c *cli.Context) error {
					// TODO: Scan all files for TODO's and use them to build a task list,
					// in addition to any issues posted to Github or other repository
					// locations.
					return nil
				},
			},
			{
				Name:    "publish",
				Aliases: []string{"p"},
				Usage:   "Publish the compiled project",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "install",
				Aliases: []string{"i"},
				Usage:   "Build ",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
		},
	})

	cmd.Run(os.Args)
}
