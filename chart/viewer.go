package chart

import (
	"encoding/json"
	"log"
	"net/http"
)

var Candlesticks []Candlestick

// Serve put up a webserver at port 8080 that serves the chart visualization
// at http://localhost:8080/viewer
func Serve(candlesticks []Candlestick) {
	Candlesticks = candlesticks

	fs := http.FileServer(http.Dir("/home/miguel/tools/viewer/"))
	http.Handle("/", fs)

	http.HandleFunc("/candlesticks", getCandlesticks)

	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}

func getCandlesticks(rw http.ResponseWriter, req *http.Request) {
	jsonData, _ := json.Marshal(Candlesticks)
	rw.Write(jsonData)
}
