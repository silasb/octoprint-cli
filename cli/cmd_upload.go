package cli

import (
	"errors"
	"fmt"
	"log"

	"github.com/silasb/octoprint-cli/api"
	"github.com/urfave/cli"
)

func CmdUpload() cli.Command {
	return cli.Command{
		Name:      "upload",
		Aliases:   []string{"u"},
		Usage:     "upload files",
		ArgsUsage: "[files]",
		Action: func(c *cli.Context) error {
			api := c.App.Metadata["api"].(*api.Client)

			if c.NArg() > 0 {
				for _, file := range c.Args() {
					fmt.Print("Uploading file: ", file)
					status, err := api.UploadFile(file)
					fmt.Println(" =>", status)

					if err != nil {
						log.Fatal(err)
					}
				}
				return nil
			}

			return errors.New("missing file")
		},
	}
}
