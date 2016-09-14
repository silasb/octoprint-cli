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

			fmt.Print("Bed:")
			fmt.Printf("\tActual: %0.2f", printer.Temperature.Bed.Actual)
			fmt.Printf("\tTarget: %0.2f\n", printer.Temperature.Bed.Target)

			fmt.Print("Tool0:")
			fmt.Printf("\tActual: %0.2f", printer.Temperature.Tool0.Actual)
			fmt.Printf("\tTarget: %0.2f\n", printer.Temperature.Tool0.Target)

			return nil
		},
	}
}
