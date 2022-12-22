package spider

type ISpider interface {
	Run()
}
type SpiderBase struct {
	Domain string
}

func (s *SpiderBase) Run() {

}
