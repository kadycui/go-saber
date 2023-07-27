package main

import (
	"gostart/web/http/utils"
	"net/http"
)

// 127.0.0.1:8080/server?server_id=1

func main() {
	http.HandleFunc("/server", utils.ServerHandler)
	http.ListenAndServe(":8080", nil)
}
