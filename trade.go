package main

import (
	"fmt"
	"log"
	"os"

	"com.miguelmf/trade/bitcointrade"
	"com.miguelmf/trade/chart"
)

func main() {
	args := os.Args

	// ticker, err := bitcointrade.GetTicker()
	// if err != nil {
	// 	fmt.Printf("erro: %v", err)
	// 	os.Exit(1)
	// }

	dataInicial := args[1] + "T00:00:00-03:00"
	dataFinal := args[2] + "T23:59:59-03:00"

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
