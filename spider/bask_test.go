package spider_test

import (
	"testing"

	"github.com/yino/AgentSpider/spider"
)

func TestTcpGather(t *testing.T) {
	spider.TcpGather("60.168.80.97	", "1133")
}
