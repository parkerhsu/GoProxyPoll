package collector

import (
	"GoProxyPoll/GoProxyPoll/defs"
	"fmt"
	"github.com/pkg/errors"
	"regexp"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestCrawl(t *testing.T) {
	//t.Run("CrawlData5u", testCrawlData5u)
	t.Run("CrawlFeiyiProxy", testCrawlFeiyiProxy)
}

func checkIp(ips []*defs.Ip) (bool, error) {
	if len(ips) == 0 {
		err := errors.New("No result")
		return false, err
	}
	for _, ip := range ips {
		if ip.Type != "HTTP" && ip.Type != "HTTPS" {
			err := errors.New("Wrong http type")
			return false, err
		}
		if match, err := regexp.MatchString(`\d+\.\d+\.\d+\.\d+:\d+`, ip.Data); err != nil || !match {
			err := errors.New("Match ip failed")
			return false, err
		}
		fmt.Println(ip)
	}
	return true, nil
}

func testCrawlData5u(t *testing.T) {
	res := CrawlData5u()
	if ok, err := checkIp(res); !ok {
		t.Error(err)
	}
}

func testCrawlFeiyiProxy(t *testing.T) {
	res := CrawlFeiyiProxy()
	if ok, err := checkIp(res); !ok {
		t.Error(err)
	}
}