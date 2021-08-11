package spider
import (
	"encoding/json"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"log"
	"os"
	"strings"
)
const (
	Host89IPCN = "www.89ip.cn"
)

func Spider89IPCNHTMLParse(doc *html.Node, reqUrl string)  {
	f, err := os.OpenFile("89ip.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	nodes := htmlquery.Find(doc, `//table[@class="layui-table"]/tbody/tr`)
	for index, node := range nodes {
		tdArr := htmlquery.Find(node, "./td")
		ip := htmlquery.InnerText(tdArr[0])
		port := htmlquery.InnerText(tdArr[1])
		address := htmlquery.InnerText(tdArr[2])
		operator := htmlquery.InnerText(tdArr[3])
		time := htmlquery.InnerText(tdArr[4])
		// 去除空格、换行、tab
		data := Agent{
			Ip: strings.Replace(
				strings.Replace(
					strings.Replace(ip, " ", "", -1),
					"\n", "", -1),
				"	", "", -1),
			Port: strings.Replace(
				strings.Replace(
					strings.Replace(port, " ", "", -1),
					"\n", "", -1),
				"	", "", -1),
			Address: strings.Replace(
				strings.Replace(
					strings.Replace(address, " ", "", -1),
					"\n", "", -1),
				"	", "", -1),
			Operator: strings.Replace(
				strings.Replace(
					strings.Replace(operator, " ", "", -1),
					"\n", "", -1),
				"	", "", -1),
			Time: strings.Replace(
				strings.Replace(
					strings.Replace(time, " ", "", -1),
					"\n", "", -1),
				"	", "", -1),
			Url:   reqUrl,
			Index: index,
		}

		byteArr, err := json.Marshal(data)
		if err != nil {
			log.Println("error:", err.Error())
		}
		f.WriteString(string(byteArr) + "\n")
	}
}
