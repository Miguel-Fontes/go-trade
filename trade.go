package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"com.miguelmf/trade/bitcointrade"
	"com.miguelmf/trade/chart"
	"github.com/pkg/errors"
)

func main() {
	args := os.Args

	// ticker, err := bitcointrade.GetTicker()
	// if err != nil {
	// 	fmt.Printf("erro: %v", err)
	// 	os.Exit(1)
	// }

	dataInicial, errDataInicial := time.Parse(time.RFC3339Nano, args[1]+"T00:00:00-03:00")
	if errDataInicial != nil {
		errors.Wrap(errDataInicial, "error while parsing date")
	}

	dataFinal, errDataFinal := time.Parse(time.RFC3339Nano, args[2]+"T23:59:59-03:00")
	if errDataInicial != nil {
		errors.Wrap(errDataFinal, "error while parsing date")
	}

	trades, errTrades := bitcointrade.GetTrades(dataInicial, dataFinal)
	if errTrades != nil {
		fmt.Printf("erro: %v", errTrades)
		os.Exit(1)
	}

	tradesData := tradesToTradesData(trades)
	candlesticks, _ := chart.CandlesticksFromTradeData(tradesData)

	log.Printf("Candlesticks: %v", candlesticks)

	chart.Serve(candlesticks)
}

func tradesToTradesData(trades []bitcointrade.Trade) []chart.TradeData {
	tradesData := make([]chart.TradeData, len(trades))
	for index, trade := range trades {
		tradesData[index] = trade
	}

	return tradesData
}
