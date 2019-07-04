package scheduler

import (
	"GoProxyPoll/GoProxyPoll/dbops"
	"GoProxyPoll/GoProxyPoll/defs"
	"GoProxyPoll/GoProxyPoll/utils"
	"log"
	"sync"
)

func Tester() {
	count, _ := dbops.Count()
	log.Printf("Start testing: %d ips\n", count)
	ips, err := dbops.AllIps()
	if err != nil {
		log.Println(err)
		return
	}
	wg := sync.WaitGroup{}
	for i := 0; i + 2 <= len(ips); i+=2 {
		wg.Add(1)
		go func(ips []*defs.Ip) {
			for _, ip := range ips {
				log.Printf("Testing: %s\n", ip.Data)
				if ok, err := utils.CheckIp(ip); ok {
					log.Printf("Valid ip: %s\n", ip.Data)
				} else {
					log.Println(err)
					log.Printf("Invalid ip: %s\n", ip.Data)
					_ = dbops.Decrease(ip)
				}
			}
			wg.Done()
		}(ips[i:i+2])
	}
	wg.Wait()
}