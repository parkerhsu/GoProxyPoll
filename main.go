package main

import (
	"GoProxyPoll/GoProxyPoll/api"
	"GoProxyPoll/GoProxyPoll/scheduler"
)

func main() {
	go scheduler.Run()
	api.Run()
}