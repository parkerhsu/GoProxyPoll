package api

import (
	"log"
	"net/http"
)

func Run() {
	http.HandleFunc("/", getIp)
	http.HandleFunc("/get", getIp)
	http.HandleFunc("/count", count)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}