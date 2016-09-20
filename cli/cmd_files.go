package cli

import (
	"errors"
	"fmt"
	"log"
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
					job, err := api.GetJob()
					if err != nil {
						log.Println(err)
						return err
					}

					files, err := api.ListFiles()
					if err != nil {
						log.Println(err)
						return err
					}

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
							log.Println(err)
							return err
						}

						files, err := api.ListFiles()
						if err != nil {
							log.Println(err)
							return err
						}
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
