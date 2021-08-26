package spider

import (
	"github.com/airdb/AgentSpider/po"
)

func TimerSyncSpider() {
	po.InitDB()

	exec := NewGetDataSpider("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36", Host89IPCN)
	pageSize := 185
	exec.GetList(int64(pageSize))

}

func TimerSyncFreeProxy() {
	exec := NewGetDataSpider("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36", HostFreeProxyCN)
	pageSize := 2
	exec.GetList(int64(pageSize))
}
