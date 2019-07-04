package collector

import (
	"GoProxyPoll/GoProxyPoll/defs"
	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func CrawlAll() []*defs.Ip {
	return CrawlFeiyiProxy()
}

func getDocument(url string) (doc *goquery.Document, err error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error when request data5u.")
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Printf("Wrong status code of data5u: %d\n", resp.StatusCode)
		err = errors.New("Wrong status code.")
		return
	}

	doc, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func CrawlData5u() (res []*defs.Ip) {
	pollUrl := "http://www.data5u.com"
	doc, err := getDocument(pollUrl)
	if err != nil {
		return
	}
	doc.Find("div.wlist > ul > li:nth-child(2) > ul").Each(func(i int, s *goquery.Selection) {
		node := strconv.Itoa(i+1)
		ip := s.Find("ul:nth-child("+ node +") > span:nth-child(1)").Text()
		port := s.Find("ul:nth-child("+ node +") > span:nth-child(2)").Text()
		tp := s.Find("ul:nth-child("+ node +") > span:nth-child(4)").Text()
		data := strings.Replace(ip+":"+port, " ", "", -1)
		res = append(res, &defs.Ip{Data:data, Type:tp})
	})
	log.Println("Crawl Data5u Done.")
	return
}

func CrawlFeiyiProxy() (res []*defs.Ip) {
	pollUrl := "http://www.feiyiproxy.com/?page_id=1457"
	doc, err := getDocument(pollUrl)
	if err != nil {
		return
	}

	doc.Find(".entry-content table tr").Each(func(i int, s *goquery.Selection) {
		if i > 0 {
			ip := s.Find("td:nth-child(1)").Text()
			port := s.Find("td:nth-child(2)").Text()
			type2 := s.Find("td:nth-child(4)").Text()
			ip = strings.Replace(ip, " ", "", -1)
			port = strings.Replace(port, " ", "", -1)
			log.Printf("Get ip: %s\n", ip)
			res = append(res, &defs.Ip{Data:ip+":"+port, Type:type2})
		}
	})
	log.Println("CrawlFeiyiProxy Done.")
	return
}