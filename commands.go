package main

import (
	"errors"
	"fmt"

	"github.com/codegangsta/cli"
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
					status := uploadFile(file)
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
			job := getJob()
			files := listFiles()
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
