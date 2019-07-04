package scheduler

import (
	"log"
	"time"
)

func runGetter() {
	for {
		log.Println("Starting getter")
		Getter()
		time.Sleep(time.Second*20)
	}
}

func runTester() {
	for {
		log.Println("Starting tester")
		Tester()
		time.Sleep(time.Second*3)
	}
}

func Run() {
	var wait chan int
	go runGetter()
	go runTester()
	<-wait
}