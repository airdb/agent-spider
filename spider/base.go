package spider

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

type Agent struct {
	Ip         string
	Port       string
	Address    string
	Operator   string
	Time       string
	Url        string
	Index      int
	ConnStatus string
}

func TcpGather(ip string, port string) {
	// 检查 emqx 1883, 8083, 8080, 18083 端口
	c := colly.NewCollector(
		colly.AllowedDomains("icanhazip.com"),
		colly.Debugger(&debug.LogDebugger{}),
	)
	c.OnRequest(func(r *colly.Request) {
		r.ProxyURL = fmt.Sprintf("%s:%s", ip, port)
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println(string(r.Body))
	})

	c.Visit("http://icanhazip.com/")

	c.Wait()
}
