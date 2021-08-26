package po_test

import (
	"testing"

	"github.com/airdb/AgentSpider/po"
)

func TestBatchFindIp(t *testing.T) {
	po.InitDB()

	po.BatchFindIp([]string{"127.0.0.1", "128.0.0.1"})
}
