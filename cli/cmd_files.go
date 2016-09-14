package cli

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/silasb/octoprint-cli/api"
	"github.com/urfave/cli"
)

func CmdFiles() cli.Command {
	return cli.Command{
		Name:    "files",
		Aliases: []string{"f"},
		Subcommands: []cli.Command{
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "list files",
				Action: func(c *cli.Context) error {
					job := api.GetJob()
					files := api.ListFiles()

					for i, file := range files {
						fmt.Printf("%d %s", i+1, file.Name)
						if job.Job.File.Name == file.Name {
							fmt.Println(" *")
						} else {
							println()
						}
					}

					return nil
				},
			},
			{
				Name:    "select",
				Aliases: []string{"s"},
				Usage:   "select file for printing",
				Action: func(c *cli.Context) error {
					if c.NArg() > 0 {
						idx := c.Args().First()
						i, err := strconv.Atoi(idx)
						if err != nil {
							return err
						}

						files := api.ListFiles()
						file := files[i-1]

						err = api.SelectFile(file)
					} else {
						fmt.Println("Missing idx")
						return errors.New("Missing argument")
					}

					return nil
				},
			},
		},
	}
}
