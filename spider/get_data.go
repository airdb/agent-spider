package spider

import (
	"fmt"
	"log"
	"math"
	"net/http"
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
	Host      string
}

func NewGetDataSpider(userAgent, host string) *Spider {
	c := colly.NewCollector(
		colly.Async(true),
		colly.Debugger(&debug.LogDebugger{}),
	)
	extensions.RandomUserAgent(c)
	// extensions.Referer(c)
	return &Spider{
		UserAgent: userAgent,
		Collector: c,
		Host:      host,
	}
}

// GetPage  获取分页 返回总页数
func (s *Spider) GetPageSize() int64 {
	pageSize := float64(total) / float64(limit)
	return int64(math.Ceil(pageSize))
}

// GetList 获取ip列表
func (s *Spider) GetList(pageSize int64) {
	fmt.Println("get list")

	s.Collector.SetCookies(s.Host, setCookieRaw("__utmc=104525399; __utmz=104525399.1629272868.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); __gads=ID=bdadd9748805971a-225235e0d2ca00ce:T=1629272868:RT=1629272868:S=ALNI_MZ0m2cM8Xt19r5rFEcZiAskQme0nw; fp=95c1ad43903e66095f2e6b2582d0b954; __utma=104525399.1732060880.1629272868.1629283102.1629286898.3; __utmt=1; __utmb=104525399.5.10.1629286898"))
	s.Collector.Limit(&colly.LimitRule{
		DomainGlob:  s.DomainGlob(),
		Parallelism: 2,
		Delay:       5 * time.Second,
	})

	for _, url := range s.RunUrlList(pageSize) {
		s.Collector.Visit(url)
	}

	s.Collector.OnRequest(func(r *colly.Request) {
		fmt.Println("req url", r.URL.Host+r.URL.Path)
	})

	//收到响应后
	s.Collector.OnResponse(func(r *colly.Response) {
		fmt.Println("on response")
		reqUrl := r.Request.URL.Host + r.Request.URL.Path
		fmt.Println("HOST:", r.Request.URL.Host)
		doc, err := htmlquery.Parse(strings.NewReader(string(r.Body)))
		if err != nil {
			log.Fatal(err)
		}
		switch r.Request.URL.Host {
		case Host89IPCN:
			Spider89IPCNHTMLParse(doc, reqUrl)

		case HostFreeProxyCN:
			SpiderFreeProxyCnParse(doc, reqUrl)
		}
	})

	s.Collector.OnError(func(r *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
		fmt.Println("url", r.Request.URL)
		fmt.Println("body", string(r.Body))
		fmt.Println("code", r.StatusCode)

	})

	s.Collector.Wait()
}

func (s *Spider) RunUrlList(pageSize int64) []string {
	var urlArr []string
	switch s.Host {
	case Host89IPCN:
		fmt.Println("host", Host89IPCN)
		for i := 1; i <= int(pageSize); i++ {
			urlArr = append(urlArr, fmt.Sprintf("https://www.89ip.cn/index_%d", i))
		}

	case HostFreeProxyCN:
		fmt.Println("host", HostFreeProxyCN)
		for i := 1; i <= int(pageSize); i++ {
			if i == 1 {
				urlArr = append(urlArr, "http://free-proxy.cz/en/proxylist/country/VN/all/ping/all")
			} else {
				urlArr = append(urlArr, fmt.Sprintf("http://free-proxy.cz/en/proxylist/country/VN/all/ping/all/%d", i))
			}

		}
	}
	return urlArr
}

func (s *Spider) DomainGlob() string {
	var domainGlob string
	switch s.Host {
	case Host89IPCN:
		fmt.Println("host", Host89IPCN)
		domainGlob = "*89.cn*"

	case HostFreeProxyCN:
		fmt.Println("host", HostFreeProxyCN)
		domainGlob = "*free-proxy*"
	}
	fmt.Println("domainGlob", domainGlob)
	return domainGlob
}

// set cookies raw
func setCookieRaw(cookieRaw string) []*http.Cookie {
	// 可以添加多个cookie
	var cookies []*http.Cookie
	cookieList := strings.Split(cookieRaw, "; ")
	for _, item := range cookieList {
		keyValue := strings.Split(item, "=")
		// fmt.Println(keyValue)
		name := keyValue[0]
		valueList := keyValue[1:]
		cookieItem := http.Cookie{
			Name:  name,
			Value: strings.Join(valueList, "="),
		}
		cookies = append(cookies, &cookieItem)
	}
	return cookies
}
