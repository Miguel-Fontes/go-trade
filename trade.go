package main

import (
	"fmt"
	"os"

	"com.miguelmf/trade/bitcointrade"
	"com.miguelmf/trade/chart"
)

func main() {
	ticker, err := bitcointrade.GetTicker()
	if err != nil {
		fmt.Printf("erro: %v", err)
		os.Exit(1)
	}

	trades, errTrades := bitcointrade.GetTrades("2017-12-01", "2018-01-07")
	if errTrades != nil {
		fmt.Printf("erro: %v", errTrades)
		os.Exit(1)
	}

	tradesData := tradesToTradesData(trades)
	candlesticks, _ := chart.CandlesticksFromTradeData(tradesData)

	println(ticker.String())
	// fmt.Printf("Trades: %v\n\n", trades)
	println(len(trades))
	fmt.Printf("Candlesticks: %v\n\n", candlesticks)

	chart.Serve(candlesticks)

}

func tradesToTradesData(trades []bitcointrade.Trade) []chart.TradeData {
	tradesData := make([]chart.TradeData, len(trades))
	for index, trade := range trades {
		tradesData[index] = trade
	}

	return tradesData
}
