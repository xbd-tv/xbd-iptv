package spider

import "fmt"

// bdys.me 爬虫

type SpiderBdys struct {
	*SpiderBase
}

func (s *SpiderBdys) Execute() {
	fmt.Println("bdys execute")
}
