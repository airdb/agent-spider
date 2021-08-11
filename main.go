package main

import (
	"fmt"

	"github.com/yino/AgentSpider/spider"
)

func main() {
	exec := spider.NewGetDataSpider("https://www.89ip.cn/index_2.html", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36")
	// pageSize := spider.GetPageSize()
	pageSize := 185
	fmt.Println(pageSize)

	exec.GetList(int64(pageSize))
}
