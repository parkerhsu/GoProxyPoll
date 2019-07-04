package api

import (
	"GoProxyPoll/GoProxyPoll/dbops"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func getIp(w http.ResponseWriter, r *http.Request) {
	ip, err := dbops.Random()
	if err != nil {
		io.WriteString(w, "Internal error happened")
		return
	}
	resp, err := json.Marshal(ip)
	if err != nil {
		io.WriteString(w, "Internal error happened")
		return
	}
	io.WriteString(w, string(resp))
}

func count(w http.ResponseWriter, r *http.Request) {
	cnt, err := dbops.Count()
	if err != nil {
		io.WriteString(w, "Internal error happened")
		return
	}
	io.WriteString(w, strconv.Itoa(cnt))
}