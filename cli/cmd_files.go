package cli

import (
	"fmt"

	"github.com/silasb/octoprint-cli/api"
	"github.com/urfave/cli"
)

func CmdFiles() cli.Command {
	return cli.Command{
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
	}
}
