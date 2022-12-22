package spider

import "github.com/gocolly/colly/v2"

// http://male66.com/ 爬虫
const domain = "male66.com"

type SpiderMale struct {
	*SpiderBase
}

func (c *SpiderMale) Execute() {
	c.listTv()
}

func (c *SpiderMale) listTv() {
	colly := colly.NewCollector(
		colly.AllowedDomains(domain),
	)
	c.detailTv()
	colly.Visit(domain)
}

func (c *SpiderMale) detailTv() {
	c.detailVideo()
}
func (c *SpiderMale) detailVideo() {
}
