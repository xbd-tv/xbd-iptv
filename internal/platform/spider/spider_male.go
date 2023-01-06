package spider

import "github.com/gocolly/colly/v2"

// http://male66.com/ 爬虫
const domain = "male66.com"

type SpiderMale struct {
	*SpiderBase
}

func (s *SpiderMale) Execute() {
	colly := colly.NewCollector(
		colly.AllowedDomains(s.Domain),
	)
	s.listTv(colly)
}

func (s *SpiderMale) listTv(colly *colly.Collector) {
	s.detailTv(colly.Clone())
	colly.Visit(domain)
}

func (s *SpiderMale) detailTv(colly *colly.Collector) {
	s.detailVideo(colly.Clone())
}

func (s *SpiderMale) detailVideo(colly *colly.Collector) {
}
