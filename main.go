package main

import (
	"context"
	"fmt"
	"time"

	"github.com/airdb/AgentSpider/spider"
	"github.com/airdb/sailor/deployutil"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
	"github.com/tencentyun/scf-go-lib/events"
)

func main() {
	spider.TimerSyncFreeProxy()
	if deployutil.IsStageDev() {
		spider.TimerSyncSpider()
	}

	cloudfunction.Start(Run)
}

// Refer: https://xuthus.cc/go/scf-go-runtime.html
func Run(ctx context.Context, event events.TimerEvent) {
	fmt.Println("hello", time.Now())

	spider.TimerSyncSpider()
}
