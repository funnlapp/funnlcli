package main

import (
	"github.com/urfave/cli"
)

// LoadCommands returns an array of available commands (and subcommands) to the user. These are linked to corresponding actions
func LoadCommands() []cli.Command {
	return []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a task to the list",
			Action:  TestAction,
		},
		{
			Name:    "complete",
			Aliases: []string{"c"},
			Usage:   "complete a task on the list",
			Action:  TestAction,
		},
		{
			Name:    "template",
			Aliases: []string{"t"},
			Usage:   "options for task templates",
			Subcommands: []cli.Command{
				{
					Name:   "add",
					Usage:  "add a new template",
					Action: TestAction,
				},
				{
					Name:   "remove",
					Usage:  "remove an existing template",
					Action: TestAction,
				},
			},
		},
	}
}
