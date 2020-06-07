package commands

import (
	"encoding/json"
	"fmt"
	"github.com/marti1125/jiracli/jira"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
)

func Config() cli.Command {
	return cli.Command{
		Name:  "config",
		Usage: "init config",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "site_url",
				Usage:    "set site url of your jira site",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "email",
				Usage:    "set your email",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "token",
				Usage:    "set your token",
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {

			siteUrl := ctx.String("site_url")
			email := ctx.String("email")
			token := ctx.String("token")

			if siteUrl != "" || email != "" || token != "" {
				fmt.Println("adding init configuration......")

				c := jira.Auth{
					SiteUrl: siteUrl,
					Email:   email,
					Token:   token,
				}

				j, err := json.MarshalIndent(c, "", "")
				if err != nil {
					fmt.Println(err)
				}

				configFile, _ := os.Stat("config.json")

				if configFile != nil && configFile.Size() > 0 {
					err = os.Remove("config.json")
					if err != nil {
						fmt.Println(err)
					}
				}

				err = ioutil.WriteFile("config.json", j, 0644)
				if err != nil {
					fmt.Println(err)
				}

			}
			return nil
		},
	}
}
