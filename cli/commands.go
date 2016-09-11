package cli

import "github.com/codegangsta/cli"

var Commands []cli.Command = []cli.Command{
	CmdUpload(),
	CmdFiles(),
	CmdRun(),
}
