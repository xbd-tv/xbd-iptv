package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"iptv.xbd/cmd"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			&cmd.CmdWeb,
			&cmd.CmdParseM3u,
		},
	}
	app.Flags = append(app.Flags, cmd.CmdWeb.Flags...)
	app.Action = cmd.CmdWeb.Action

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}