package request

import (
	"fmt"
	"log"

	"github.com/asmcos/requests"
)

const (
	HidemyURL = "https://proxylist.geonode.com/api/proxy-list"
)

func Hidemy() {
	req := requests.Requests()
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("accept-encoding", "gzip, deflate, br")
	req.Header.Set("accept-language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	req.Header.Set("cache-control", "max-age=0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:91.0) Gecko/20100101 Firefox/91.0")
	req.Header.Set("cookie", "t=225335792; PAPVisitorId=68729dcb527799284d6aqp5SfFB3dtWf; PAPVisitorId=68729dcb527799284d6aqp5SfFB3dtWf; _ym_uid=1629367856532942687; _ym_d=1629367856; _ga=GA1.2.1970326936.1629367856; _fbp=fb.1.1629367858196.877914826; _gid=GA1.2.955463117.1630416824; _ym_isad=2; _dc_gtm_UA-90263203-1=1; _gat_UA-90263203-1=1")
	req.Header.Set("referer", "https://hidemy.name/en/proxy-list/?start=64")
	req.Header.Set("sec-ch-ua", "\"Chromium\";v=\"92\", \" Not A;Brand\";v=\"99\", \"Google Chrome\";v=\"92\"")
	req.Header.Set("accept-language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua", "sec-fetch-dest")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.SetTimeout(50)
	resp, err := req.Get("http://hidemy.name/en/proxy-list/#list")

	if err != nil {
		log.Println("Get Total Geonode:", err)
		return
	}
	fmt.Println(resp.Text())
}
