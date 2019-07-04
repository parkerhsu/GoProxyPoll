package scheduler

import (
	"GoProxyPoll/GoProxyPoll/collector"
	"GoProxyPoll/GoProxyPoll/dbops"
	"log"
)

func isOverThreshold() bool {
	count, _ := dbops.Count()
	if count > THRESHOLD {
		return true
	}
	return false
}

func Getter() {
	if !isOverThreshold() {
		ips := collector.CrawlAll()
		for _, ip := range ips {
			if ok, err := dbops.NotExist(ip); ok && err == nil {
				err := dbops.AddIp(ip)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}