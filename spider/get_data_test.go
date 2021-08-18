package spider_test

import (
	"testing"

	"github.com/yino/AgentSpider/spider"
)

func TestGet89ipCnList(t *testing.T) {
	exec := spider.NewGetDataSpider("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36", spider.Host89IPCN)
	pageSize := 1
	exec.GetList(int64(pageSize))
}
func TestGetFreeProxyCnList(t *testing.T) {
	exec := spider.NewGetDataSpider("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36", spider.HostFreeProxyCN)
	pageSize := 2
	exec.GetList(int64(pageSize))
}
