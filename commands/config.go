package commands

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
)

type JiraAuth struct {
	SiteUrl string `json:"site_url"`
	Email   string `json:"email"`
	Token   string `json:"token"`
}

func Config() cli.Command {
	return cli.Command{
		Name:  "commands",
		Usage: "init commands",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "site_url",
				Usage:    "commands the site url of your jira site",
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			siteUrl := ctx.String("site_url")
			if siteUrl != "" {
				fmt.Println("adding init configuration......")

				c := JiraAuth{
					SiteUrl: siteUrl,
				}

				j, err := json.MarshalIndent(c, "", "")
				if err != nil {
					fmt.Println(err)
				}

				configFile, err := os.Stat("commands.json")
				if err != nil {
					fmt.Println(err)
				}
				if configFile.Size() > 0 {
					err = os.Remove("commands.json")
					if err != nil {
						fmt.Println(err)
					}
				}

				err = ioutil.WriteFile("commands.json", j, 0644)
				if err != nil {
					fmt.Println(err)
				}

			}
			return nil
		},
	}
}
