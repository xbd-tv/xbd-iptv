package cmd

import (
	"os"

	"github.com/urfave/cli/v2"
	"iptv.xbd/internal/platform/m3u"
)

var CmdParseM3u = cli.Command{
	Name:   "m3u",
	Usage:  "parse .m3u file",
	Action: runParse,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "in",
			Required: true,
			Usage:    "file.m3u",
		},
		&cli.StringFlag{
			Name:    "csv",
			Aliases: []string{"out-csv"},
			Usage:   "",
		},
	},
}

func runParse(c *cli.Context) error {
	fileIn := c.String("in")
	csv := c.String("csv")
	pl, err := parseFile(fileIn)

	if err != nil {
		panic(err)
	}
	m3u.ToCsv(&pl, csv)
	return nil
}

func parseFile(filePath string) (m3u.PlayList, error) {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	return m3u.Parse(f)
}
