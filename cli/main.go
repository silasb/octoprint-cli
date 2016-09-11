package cli

import (
	"errors"
	"os"
	"time"

	"github.com/codegangsta/cli"
	"github.com/silasb/octoprint-cli/api"
)

//var host string
//var api_key string

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
	api.Host = c.String("host")
	api.Api_key = c.String("key")

	if api.Host == "" || api.Api_key == "" {
		return errors.New("missing required flags")
	}

	api.Host = api.Host + "/api/"

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
