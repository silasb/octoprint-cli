package cli

import (
	"errors"
	"os"
	"time"

	"github.com/silasb/octoprint-cli/api"
	"github.com/urfave/cli"
)

var defaultFlags = []cli.Flag{
	cli.StringFlag{
		Name:   "host",
		Usage:  "Octoprint host *REQUIRED*",
		EnvVar: "HOST",
	},
	cli.StringFlag{
		Name:   "key",
		Usage:  "Octoprint API key *REQUIRED*",
		EnvVar: "KEY",
	},
}

var beforeFunc = func(c *cli.Context) error {
	host := c.String("host")
	key := c.String("key")

	if host == "" || key == "" {
		return errors.New("missing required flags")
	}

	host = host + "/api"

	c.App.Metadata = map[string]interface{}{
		"Host": host,
		"Key":  key,
	}

	cfg := api.Config{
		Key:      key,
		Endpoint: host,
	}
	api, err := api.New(cfg)

	if err != nil {
		panic(err)
	}
	c.App.Metadata["api"] = api

	return nil
}

func Main() {
	app := cli.NewApp()
	app.Name = "octoprint-cli"
	app.Version = "0.0.1"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Silas Baronda",
			Email: "silas.baronda@gmail.com",
		},
	}
	app.Usage = ""
	app.Flags = defaultFlags
	app.Before = beforeFunc
	app.Commands = Commands

	app.Run(os.Args)
}
