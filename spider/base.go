package spider

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Agent struct {
	Ip         string
	Port       string
	Address    string
	Operator   string
	Time       string
	Url        string
	Index      int
	ConnStatus string
}

type AuthAgent struct {
	Agent
	Http  error
	Https error
}

// 检测 TCP
func TcpGather(httpType, ip, port string) error {
	proxyAddress := fmt.Sprintf("%s://%s:%s", httpType, ip, port)
	//访问查看ip的一个网址
	httpUrl := "http://icanhazip.com"
	proxy, err := url.Parse(proxyAddress)
	if err != nil {
		return err
	}
	netTransport := &http.Transport{
		Proxy:                 http.ProxyURL(proxy),
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * time.Duration(15),
	}
	httpClient := &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}
	res, err := httpClient.Get(httpUrl)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return err
	}
	c, _ := ioutil.ReadAll(res.Body)

	if string(c) != ip {
		return errors.New("fail ip")
	}
	return nil
}

// 加载  IP文件
func LoadIp(path string) (agentArr []Agent, err error) {
	_, err = os.Stat(path)
	if err != nil {
		return
	}
	if os.IsNotExist(err) {
		err = errors.New("file or dir not exist")
		return
	}
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	chunks := make([]byte, 1024, 1024)
	buf := make([]byte, 1024)
	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
		// fmt.Println(string(buf[:n]))
	}
	contentLineArr := strings.Split(string(chunks), "\n")
	for _, val := range contentLineArr {
		if len(val) > 0 {
			var agentData Agent
			err := json.Unmarshal([]byte(val), &agentData)
			if err != nil {
				log.Println(err, fmt.Sprintf("unmarshal fail %s", val))
				continue
			}
			agentArr = append(agentArr, agentData)
		}

	}
	return
}

// 批量检测 TCP
func BatchTcpGather(path string) {
	agentArr, err := LoadIp(path)
	if err != nil {
		panic(err)
	}
	ch := make(chan AuthAgent)
	// agentlen := len(agentArr)
	for index, val := range agentArr {
		go func(val Agent) {
			err := TcpGather("http", val.Ip, val.Port)
			AuthAgentData := AuthAgent{
				Agent: val,
				Http:  err,
				Https: nil,
			}

			if err != nil {
				err = TcpGather("https", val.Ip, val.Port)
				AuthAgentData.Https = err
			}
			ch <- AuthAgentData
		}(val)
	}
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("已经读取完所有数据了")
			break
		}

		if v.Http != nil || v.Https != nil {
			fmt.Println("无效IP地址", v.Ip, v.Port)
		} else {
			fmt.Println("有效IP地址", v.Ip, v.Port)
		}
	}
}
