package spider

import (
	"github.com/yino/AgentSpider/po"
)

func TimerSyncSpider(){
	po.InitDB()

	exec := NewGetDataSpider("https://www.89ip.cn/index_2.html", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36")
	pageSize := 185
	exec.GetList(int64(pageSize))
}
