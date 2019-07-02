package utils

import (
	"GoProxyPoll/GoProxyPoll/defs"
	"github.com/parnurzeal/gorequest"
	"log"
)

func CheckIp(ip *defs.Ip) bool {
	var pollUrl string
	var proxy string
	if ip.Type == "HTTP" {
		pollUrl = "http://httpbin.org/get"
		proxy = "http://" + ip.Data
	} else if ip.Type == "HTTPS" {
		pollUrl = "https://httpbin.org/get"
		proxy = "https://" + ip.Data
	} else {
		log.Println("Wrong type")
		return false
	}

	resp, _, err := gorequest.New().Proxy(proxy).Get(pollUrl).End()
	if err != nil {
		log.Printf("Error when request url: %s\n", err)
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		return true
	}
	return false
}