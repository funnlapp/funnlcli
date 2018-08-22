package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func CompleteTask(c *cli.Context) error {
	fmt.Println("completed task: ", c.Args().First())
	return nil
}
