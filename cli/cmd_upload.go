package cli

import (
	"errors"
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/silasb/octoprint-cli/api"
)

func CmdUpload() cli.Command {
	return cli.Command{
		Name:      "upload",
		Aliases:   []string{"u"},
		Usage:     "upload files",
		ArgsUsage: "[files]",
		Action: func(c *cli.Context) error {
			if c.NArg() > 0 {
				for _, file := range c.Args() {
					fmt.Print("Uploading file: ", file)
					status := api.UploadFile(file)
					fmt.Println(" =>", status)
				}
				return nil
			}

			return errors.New("missing file")
		},
	}
}
