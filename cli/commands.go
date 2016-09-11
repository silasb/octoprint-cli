package cli

import "github.com/urfave/cli"

var Commands []cli.Command = []cli.Command{
	CmdUpload(),
	CmdFiles(),
	CmdRun(),
}
