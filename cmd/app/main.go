package main

import (
	"github.com/urfave/cli/v2"
	"go-template/internal/pkg/builder"
	"go-template/internal/pkg/config"
	"log"
	"os"
)

func main() {

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "api",
				Aliases: []string{},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "config",
						Aliases: []string{"c"},
						Value:   "internal/pkg/config/default.yaml",
						Usage:   "Configuration file",
					},
				},
				Action: func(c *cli.Context) (err error) {
					configFilePath := c.String("config")
					err = config.Load(configFilePath)
					if err != nil {
						return err
					}

					cmd, err := builder.NewApiBuilder()
					if err != nil {
						return err
					}
					return cmd.Run()
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
