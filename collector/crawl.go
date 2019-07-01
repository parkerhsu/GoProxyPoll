package collector

import (
	"GoProxyPoll/GoProxyPoll/defs"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func CrawlData5u() (res []*defs.Ip) {
	pollUrl := "http://www.data5u.com/"
	resp, err := http.Get(pollUrl)
	if err != nil {
		log.Println("Error when request data5u.")
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Println("Wrong status code of data5u.")
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	doc.Find()
}