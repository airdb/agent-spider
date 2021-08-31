package request

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/asmcos/requests"
)

const (
	HidemyURL = "https://proxylist.geonode.com/api/proxy-list"
)

func Hidemy() {
	req, err := NewRequest("hidemy_header.txt")

	if err != nil {
		log.Println("Hidemy error:", err)
		return
	}
	resp, err := req.Get("http://hidemy.name/en/proxy-list/#list")

	if err != nil {
		log.Println("Get Total Geonode:", err)
		return
	}
	fmt.Println(resp.Text())
}

func NewRequest(headerFilepath string) (*requests.Request, error) {
	req := requests.Requests()
	req.SetTimeout(60)
	bytes, err := ioutil.ReadFile(headerFilepath)
	if err != nil {
		log.Println("error : %s", err)
		return req, err
	}

	contentLines := strings.Split(string(bytes), "\n")
	if len(contentLines) > 0 {
		for _, line := range contentLines {
			headers := strings.Split(line, ":")
			headerValue := strings.Join(headers[1:], ":")
			req.Header.Add(headers[0], headerValue)
		}
	}
	return req, nil
}
