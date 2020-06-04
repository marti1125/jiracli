package config

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
)

type Config struct {
	SiteUrl string `json:"site_url"`
}

func InitCommand() cli.Command {
	return cli.Command{
		Name: "config",
		Usage: "init config",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "site_url",
				Usage: "config the site url of your jira site",
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			siteUrl := ctx.String("site_url")
			if siteUrl!= "" {
				fmt.Println("adding init configuration......")

				c := Config{
					SiteUrl: siteUrl,
				}

				j, err := json.MarshalIndent(c, "", "")
				if err != nil {
					fmt.Println(err)
				}

				configFile, err := os.Stat("config.json")
				if err != nil {
					fmt.Println(err)
				}
				if configFile.Size() > 0 {
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
