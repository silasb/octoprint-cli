package cli

import (
	"fmt"
	"log"
	"strings"

	"github.com/silasb/octoprint-cli/api"
	"github.com/urfave/cli"
)

func CmdRun() cli.Command {
	return cli.Command{
		Name:    "run",
		Aliases: []string{"r"},
		Usage:   "run gcode commands",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name: "verbose, v",
				//Usage:  "Octoprint host *REQUIRED*",
			},
		},
		//ArgsUsage: "",
		Action: func(c *cli.Context) error {
			if c.NArg() > 0 {
				// handle both stdin and argument based input

				gcode := c.Args().First()
				commands := strings.Split(gcode, ";")
				for i, command := range commands {
					commands[i] = strings.TrimSpace(command)
				}

				if c.Bool("verbose") {
					for _, command := range commands {
						fmt.Printf("%+v\n", command)
					}
					return nil
				}

				err := api.Run(commands)
				if err != nil {
					log.Panic(err)
				}
				return nil
			} else {
				// TODO: interactive mode
				// -i --interactive

				fmt.Println("TODO: Interactive mode")
				return nil
			}

		},
	}
}
