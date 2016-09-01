package main

import (
	"errors"
	"os"
	"time"

	"github.com/codegangsta/cli"
)

var host string
var api_key string

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
	host = c.String("host")
	api_key = c.String("key")

	if host == "" || api_key == "" {
		return errors.New("missing required flags")
	}

	host = host + "/api/"

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "octoprint"
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
