package cli

import (
	"fmt"

	"github.com/silasb/octoprint-cli/api"
	"github.com/urfave/cli"
)

func CmdInfo() cli.Command {
	return cli.Command{
		Name:    "info",
		Aliases: []string{"i"},
		Usage:   "info",
		Action: func(c *cli.Context) error {
			printer, err := api.Info()

			if err != nil {
				return err
			}

			fmt.Printf("%+v", printer)

			return nil
		},
	}
}
