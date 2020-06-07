package commands

import (
	"fmt"
	"github.com/marti1125/jiracli/jira"
	"github.com/urfave/cli"
)

func Info() cli.Command {
	return cli.Command{
		Name: "info",
		Usage: "get server info",
		Action: func(ctx *cli.Context) error {
			fmt.Println(jira.GetServerInfo())
			return nil
		},
	}
}