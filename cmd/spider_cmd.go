package cmd

import (
	"github.com/urfave/cli/v2"
	"iptv.xbd/internal/platform/spider"
)

var CmdSpider = cli.Command{
	Name:   "spider",
	Usage:  "spider male66.com or bdys.me",
	Action: runSpider,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "male",
			Usage: "spider male66.com tv",
		},
		&cli.BoolFlag{
			Name:  "bdys",
			Usage: "spider bdys.me tv",
		},
	},
}

func runSpider(c *cli.Context) error {
	male := c.Bool("male")
	bdys := c.Bool("bdys")
	if male {
		s := spider.SpiderMale{
			Domain: "male16.com",
		}
		s.Execute()
	}

	if bdys {
		s := spider.SpiderBdys{}
		s.Execute()
	}

	return nil
}
