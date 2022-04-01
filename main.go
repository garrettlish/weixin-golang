package main

import (
	"log"
	"net/http"
	"weixin-golang/service"
)

func main() {
	http.HandleFunc("/", service.IndexHandler)
	http.HandleFunc("/api/stock", service.StockHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}
