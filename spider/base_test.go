package spider_test

import (
	"fmt"
	"testing"

	"github.com/airdb/AgentSpider/spider"
)

func TestTcpGather(t *testing.T) {
	err := spider.TcpGather("http", "49.89.86.140", "8888")
	fmt.Println(err)
}

func TestLoadIp(t *testing.T) {
	agent, err := spider.LoadIp("../data.txt.bak")
	fmt.Println(agent, err)
}

func TestBatchTcpGather(t *testing.T) {
	spider.BatchTcpGather("../data.txt.bak")
}
