package request

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/airdb/AgentSpider/po"
	"github.com/airdb/AgentSpider/spider"
	"github.com/asmcos/requests"
)

const (
	GeonodeURL = "https://proxylist.geonode.com/api/proxy-list"
)

func Geonode() {
	req := requests.Requests()
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:91.0) Gecko/20100101 Firefox/91.0")
	resp, err := req.Get("https://proxylist.geonode.com/api/proxy-list")

	if err != nil {
		log.Println("Get Total Geonode:", err)
		return
	}

	m := make(map[string]interface{})
	err = json.Unmarshal([]byte(resp.Text()), &m)

	if err != nil {
		fmt.Println("Get Total Geonode Umarshal failed:", err)
		return
	}
	// total := m["total"].(float64)

	url := fmt.Sprintf("https://proxylist.geonode.com/api/proxy-list?limit=%d&page=1&sort_by=lastChecked&sort_type=desc", (int(10)))
	// url := fmt.Sprintf("https://proxylist.geonode.com/api/proxy-list?limit=%d&page=1&sort_by=lastChecked&sort_type=desc", (int(total)))
	fmt.Println(url)
	resp, err = req.Get(url)

	if err != nil {
		log.Println("Geonode Data:", err)
		return
	}

	data := make(map[string]interface{})
	err = json.Unmarshal([]byte(resp.Text()), &data)

	if err != nil {
		fmt.Println("Geonode Umarshal failed:", err)
		return
	}
	poChan := make(chan po.AgentIp)

	dataLen := len(data["data"].([]interface{}))
	// 解析数据
	for _, val := range data["data"].([]interface{}) {
		item := val.(map[string]interface{})
		agentPo := po.AgentIp{
			IP:            item["ip"].(string),
			Port:          item["port"].(string),
			Anonymity:     item["anonymityLevel"].(string),
			Country:       item["country"].(string),
			City:          item["city"].(string),
			LastCheckedAt: uint(time.Now().Unix()),
		}

		switch item["speed"].(type) {
		case float64:
			agentPo.Speed = strconv.FormatFloat(item["speed"].(float64), 'E', -1, 64)
		}

		// 开启 goroutine 验证 ip的有效性
		go func(agentPo po.AgentIp, poChan chan po.AgentIp) {
			fmt.Println(agentPo)
			err := spider.TcpGather("http", agentPo.IP, agentPo.Port)
			if err == nil {
				agentPo.ProxyType = "http"
				poChan <- agentPo
				fmt.Println("http 校验成功")
				return
			}

			err = spider.TcpGather("https", agentPo.IP, agentPo.Port)
			if err == nil {
				agentPo.ProxyType = "https"
				poChan <- agentPo
				fmt.Println("https 校验成功")
				return
			}

			poChan <- agentPo
			fmt.Println("https、http 校验失败")
		}(agentPo, poChan)
	}

	var batchData, insertBatchData []po.AgentIp
	var ipPortArr []string

	// 通过channel 获取数据
	count := 0
	for {
		if count == dataLen {
			break
		}
		poItem := <-poChan
		if poItem.ProxyType == "http" || poItem.ProxyType == "https" {
			batchData = append(batchData, poItem)
			ipPortArr = append(ipPortArr, poItem.IP+":"+poItem.Port)
		}
		count += 1
	}
	ipMap := make(map[string]int)
	// 查询是否有存在的 ip 有的话 update操作
	if len(ipPortArr) > 0 {
		iPArr := po.BatchFindIp(ipPortArr)
		for _, val := range iPArr {
			ipMap[val.IpPort] = 1
			po.UpdateAgent(val)
		}
	}

	// 去除已存在数据库的数据
	if len(batchData) > 0 {
		for _, val := range batchData {
			if _, ok := ipMap[val.IP+":"+val.Port]; !ok {
				insertBatchData = append(insertBatchData, val)
			}
		}
	}

	// 批量插入数据
	if len(insertBatchData) > 0 {
		po.BatchAgentInsert(insertBatchData)
	}
}
