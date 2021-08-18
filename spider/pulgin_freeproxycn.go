package spider

import (
	"fmt"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

const (
	HostFreeProxyCN = "free-proxy.cz"
)

func SpiderFreeProxyCnParse(doc *html.Node, reqUrl string) {
	nodes := htmlquery.Find(doc, `//table[@id="proxy_list"]/tbody/tr`)

	for _, val := range nodes {
		fmt.Println(val)
	}
}
