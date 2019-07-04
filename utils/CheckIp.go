package utils

import (
	"GoProxyPoll/GoProxyPoll/defs"
	"errors"
	"github.com/parnurzeal/gorequest"
	"time"
)

func CheckIp(ip *defs.Ip) (bool, error) {
	var pollUrl string
	var proxy string
	if ip.Type == "HTTP" {
		pollUrl = "http://httpbin.org/get"
		proxy = "http://" + ip.Data
	} else if ip.Type == "HTTPS" {
		pollUrl = "https://httpbin.org/get"
		proxy = "https://" + ip.Data
	} else {
		err := errors.New("Wrong protocol type")
		return false, err
	}
	resp, _, err := gorequest.New().
						Proxy(proxy).
						Get(pollUrl).
						Timeout(time.Second*5).
						End()
	if err != nil {
		err2 := errors.New("Error when request url")
		return false, err2
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		return true, nil
	}
	return false, errors.New("Wrong status code.")
}