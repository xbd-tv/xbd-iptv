package spider

import "github.com/gocolly/colly/v2"

// http://male66.com/ 爬虫
const domain = "male66.com"

type MaleSpider struct {
	*SpiderBase
}

func (c *MaleSpider) Execute() {
	c.listTv()
}

func (c *MaleSpider) listTv() {
	colly := colly.NewCollector(
		colly.AllowedDomains(domain),
	)
	c.detailTv()
	colly.Visit(domain)
}

func (c *MaleSpider) detailTv() {
	c.detailVideo()
}
func (c *MaleSpider) detailVideo() {
}
