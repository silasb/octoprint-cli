package cli

import (
	"errors"
	"fmt"

	"github.com/codegangsta/cli"
	api "github.com/silasb/octoprint-cli/api"
)

var Commands []cli.Command = []cli.Command{
	{
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
	},
	{
		Name:    "files",
		Aliases: []string{"f"},
		Usage:   "list files",
		Action: func(c *cli.Context) error {
			job := api.GetJob()
			files := api.ListFiles()
			for _, file := range files {
				fmt.Print(file.Name)
				if job.Job.File.Name == file.Name {
					fmt.Println(" *")
				} else {
					println()
				}
			}
			return nil
		},
	},
}
