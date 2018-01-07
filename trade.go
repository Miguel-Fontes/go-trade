package main

import (
	"fmt"
	"os"

	"com.miguelmf/trade/bitcointrade"
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

	println(ticker.String())
	fmt.Printf("Trades: %v", trades)

}
