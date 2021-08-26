package po_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/airdb/AgentSpider/po"
)

type User struct {
	Id           int64
	UserNickanme string `json:user_nickname`
	Mobile       string `json:mobile`
}

func TestPool(t *testing.T) {
	po.NewConnPool()
	db := po.GetDB()

	for i := 0; i < 10000; i++ {
		go func(i int) {
			var data User
			db.Table("table_user").First(&data)
			fmt.Println(i, data)
		}(i)
	}
	fmt.Println(db)
	for {
		time.Sleep(time.Hour)
	}
}
