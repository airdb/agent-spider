package spider

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

//  总代理量
var total = 4296
var limit = 20

type Spider struct {
	Referer   string
	UserAgent string
	Collector *colly.Collector
}

func New89Spider(referer, userAgent string) *Spider {
	c := colly.NewCollector(
		colly.AllowedDomains("www.89ip.cn"),
		colly.Async(true),
		colly.Debugger(&debug.LogDebugger{}),
	)
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
	f, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// url := "https://www.89ip.cn/index_2.html"
	fmt.Println("get list ")
	s.Collector.Limit(&colly.LimitRule{
		DomainGlob:  "*89ip*",
		Parallelism: 2,
		Delay:       5 * time.Second,
	})

	s.Collector.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", s.UserAgent)
		r.Headers.Set("Referer", s.Referer)
		fmt.Println("req url", r.URL.Host+r.URL.Path)
	})

	//收到响应后
	s.Collector.OnResponse(func(r *colly.Response) {
		reqUrl := r.Request.URL.Host + r.Request.URL.Path
		doc, err := htmlquery.Parse(strings.NewReader(string(r.Body)))
		if err != nil {
			log.Fatal(err)
		}
		nodes := htmlquery.Find(doc, `//table[@class="layui-table"]/tbody/tr`)
		for index, node := range nodes {
			tdArr := htmlquery.Find(node, "./td")
			ip := htmlquery.InnerText(tdArr[0])
			port := htmlquery.InnerText(tdArr[1])
			address := htmlquery.InnerText(tdArr[2])
			operator := htmlquery.InnerText(tdArr[3])
			time := htmlquery.InnerText(tdArr[4])
			// 去除空格、换行、tab
			data := Agent{
				Ip: strings.Replace(
					strings.Replace(
						strings.Replace(ip, " ", "", -1),
						"\n", "", -1),
					"	", "", -1),
				Port: strings.Replace(
					strings.Replace(
						strings.Replace(port, " ", "", -1),
						"\n", "", -1),
					"	", "", -1),
				Address: strings.Replace(
					strings.Replace(
						strings.Replace(address, " ", "", -1),
						"\n", "", -1),
					"	", "", -1),
				Operator: strings.Replace(
					strings.Replace(
						strings.Replace(operator, " ", "", -1),
						"\n", "", -1),
					"	", "", -1),
				Time: strings.Replace(
					strings.Replace(
						strings.Replace(time, " ", "", -1),
						"\n", "", -1),
					"	", "", -1),
				Url:   reqUrl,
				Index: index,
			}

			byteArr, err := json.Marshal(data)
			if err != nil {
				log.Println("error:", err.Error())
			}
			f.WriteString(string(byteArr) + "\n")
		}
	})

	for i := 1; i <= int(pageSize); i++ {
		s.Collector.Visit(fmt.Sprintf("https://www.89ip.cn/index_%d", i))
	}
	// s.Collector.Visit(url)

	s.Collector.Wait()
}
