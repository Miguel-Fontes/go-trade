package chart

import (
	"encoding/json"
	"log"
	"net/http"
)

var Candlesticks []Candlestick

func Serve(candlesticks []Candlestick) {
	Candlesticks = candlesticks

	fs := http.FileServer(http.Dir("./viewer/"))
	http.Handle("/", fs)

	http.HandleFunc("/candlesticks", getCandlesticks)

	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}

func getCandlesticks(rw http.ResponseWriter, req *http.Request) {
	jsonData, _ := json.Marshal(Candlesticks)
	rw.Write(jsonData)
}
