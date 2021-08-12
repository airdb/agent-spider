package main

import (
	"fmt"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
	"time"
	"context"

	"github.com/tencentyun/scf-go-lib/events"
	"github.com/yino/AgentSpider/spider"
)

func main() {
	cloudfunction.Start(Run)
}

// Refer: https://xuthus.cc/go/scf-go-runtime.html
func Run(ctx context.Context, event events.TimerEvent) {
	fmt.Println("hello", time.Now())

	spider.TimerSyncSpider()
}
