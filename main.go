package main

import (
	"fmt"
	"github.com/marti1125/jiracli/commands"
	"github.com/urfave/cli"
	"os"
)

var (
	cl *cli.App
)

func init()  {
	cl = cli.NewApp()
	cl.Name = "JiraAuth Tool"
	cl.Usage = ""
	cl.Version = ""
	cl.Description = ""
}

func main() {
	cl.Commands = []cli.Command{
		commands.Config(),
	}

	err := cl.Run(os.Args)

	if err != nil {
		fmt.Println("run has error:", err.Error())
	}
}
