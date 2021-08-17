package spider

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/yino/AgentSpider/po"
	"golang.org/x/net/html"
)

const (
	Host89IPCN = "www.89ip.cn"
)

func Spider89IPCNHTMLParse(doc *html.Node, reqUrl string) {
	nodes := htmlquery.Find(doc, `//table[@class="layui-table"]/tbody/tr`)
	var ipArr []string
	var AgentIpArrPo []po.AgentIp
	for _, node := range nodes {
		tdArr := htmlquery.Find(node, "./td")
		ip := htmlquery.InnerText(tdArr[0])
		port := htmlquery.InnerText(tdArr[1])
		address := htmlquery.InnerText(tdArr[2])
		operator := htmlquery.InnerText(tdArr[3])
		//time := htmlquery.InnerText(tdArr[4])
		Actived := true
		var httpType string
		data := po.AgentIp{
			// 去除空格、换行、tab
			IP: strings.Replace(
				strings.Replace(
					strings.Replace(ip, " ", "", -1),
					"\n", "", -1),
				"	", "", -1),
			Port: strings.Replace(
				strings.Replace(
					strings.Replace(port, " ", "", -1),
					"\n", "", -1),
				"	", "", -1),
			ProxyType: "http",
			Country:   "China",
			City: strings.Replace(
				strings.Replace(
					strings.Replace(address, " ", "", -1),
					"\n", "", -1),
				"	", "", -1),
			Operator: strings.Replace(
				strings.Replace(
					strings.Replace(operator, " ", "", -1),
					"\n", "", -1),
				"	", "", -1),
			Origin:        Host89IPCN,
			Actived:       &Actived,
			LastCheckedAt: uint(time.Now().Unix()),
		}

		httpErr := TcpGather("http", data.IP, data.Port)
		httpsErr := TcpGather("https", data.IP, data.Port)
		ipArr = append(ipArr, data.IP)

		if httpErr == nil {
			httpType = "http"
			Actived = true
		}
		if httpsErr == nil {
			httpType = "https"
			Actived = true
		}
		data.ProxyType = httpType
		data.Actived = &Actived

		AgentIpArrPo = append(AgentIpArrPo, data)
		fmt.Println(data)
	}

	ipData := po.BatchFindIp(ipArr)
	ipDbDataMap := make(map[string]*po.AgentIp)
	for _, val := range ipData {
		ipDbDataMap[val.IP+":"+val.Port] = val
	}

	// 拼接 insert data、update data
	var updatePo, insertPo []po.AgentIp
	for _, agentIpPo := range AgentIpArrPo {
		// update
		if val, ok := ipDbDataMap[agentIpPo.IP+":"+agentIpPo.Port]; ok {
			agentIpPo.ID = val.ID
			agentIpPo.CreatedAt = val.CreatedAt
			agentIpPo.UpdatedAt = val.UpdatedAt
			agentIpPo.LastCheckedAt = val.LastCheckedAt
			updatePo = append(updatePo, agentIpPo)
			continue
		}
		// insert
		insertPo = append(insertPo, agentIpPo)
	}
	// batch insert
	var insertErr, updateErr error
	if len(insertPo) > 0 {
		insertErr = po.BatchAgentInsert(insertPo)
	}

	if len(updatePo) > 0 {
		for _, val := range updatePo {
			updateErr = po.UpdateAgent(&val)
			if updateErr != nil {
				log.Println("【update error】", updateErr, val)
			}
		}
	}

	if insertErr != nil {
		log.Println("【insert error】", insertErr, insertPo)
	}
}
