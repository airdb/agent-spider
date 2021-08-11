package spider

import (
	"fmt"
	"log"
	"math"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"github.com/gocolly/colly/extensions"
)

//  总代理量
var total = 4296
var limit = 20

type Spider struct {
	Referer   string
	UserAgent string
	Collector *colly.Collector
}

func NewGetDataSpider(referer, userAgent string) *Spider {
	c := colly.NewCollector(
		colly.AllowedDomains("www.89ip.cn"),
		colly.Async(true),
		colly.Debugger(&debug.LogDebugger{}),
	)
	extensions.RandomUserAgent(c)
	extensions.Referer(c)
	return &Spider{
		Referer:   referer,
		UserAgent: userAgent,
		Collector: c,
	}
}

// GetPage  获取分页 返回总页数
func (s *Spider) GetPageSize() int64 {
	pageSize := float64(total) / float64(limit)
	return int64(math.Ceil(pageSize))
}

// GetList 获取ip列表
func (s *Spider) GetList(pageSize int64) {
	// url := "https://www.89ip.cn/index_2.html"
	fmt.Println("get list ")
	s.Collector.Limit(&colly.LimitRule{
		DomainGlob:  "*89ip*",
		Parallelism: 2,
		Delay:       5 * time.Second,
	})

	s.Collector.OnRequest(func(r *colly.Request) {
		fmt.Println("req url", r.URL.Host+r.URL.Path)
	})

	//收到响应后
	s.Collector.OnResponse(func(r *colly.Response){
		reqUrl := r.Request.URL.Host + r.Request.URL.Path
		doc, err := htmlquery.Parse(strings.NewReader(string(r.Body)))
		if err != nil {
			log.Fatal(err)
		}
		switch r.Request.URL.Host{
			case Host89IPCN:
				Spider89IPCNHTMLParse(doc,reqUrl)
		}
	})

	for i := 1; i <= int(pageSize); i++ {
		s.Collector.Visit(fmt.Sprintf("https://www.89ip.cn/index_%d", i))
	}

	s.Collector.Wait()
}
