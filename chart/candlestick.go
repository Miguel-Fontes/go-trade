package chart

import (
	"fmt"
	"sort"
	"time"
)

// Candlestick represents a bar on a Candlestick chart, containing
// the max, min, open and close prices of a given day
type Candlestick struct {
	Max   float64
	Min   float64
	Open  float64
	Close float64
	Day   string
}

// TradeData interface defines the contract that any type resembling
// a Trade must comply.
type TradeData interface {
	GetPrice() float64
	GetDate() time.Time
	GetAmount() float64
	GetType() string
}

type tradeDataSlice []TradeData

// Define the Len method for use by the sort package
func (p tradeDataSlice) Len() int {
	return len(p)
}

// Define the Less method for use by the sort package
func (p tradeDataSlice) Less(i, j int) bool {
	return p[i].GetDate().Before(p[j].GetDate())
}

// Define the Swap method for use by the sort package
func (p tradeDataSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// CandlesticksFromTradeData converts an TradeDate set to a set of Candlesticks
func CandlesticksFromTradeData(trades tradeDataSlice) (candlesticks []Candlestick, err error) {
	sort.Sort(trades)

	var closingTrade TradeData
	maximum := trades[0].GetPrice()
	minimum := trades[0].GetPrice()
	openingTrade := trades[0]
	lastDate := trades[0].GetDate()
	for _, trade := range trades {
		if !isFromSameDay(lastDate, trade.GetDate()) {
			candlesticks = append(candlesticks, Candlestick{Day: lastDate.String(),
				Open:  openingTrade.GetPrice(),
				Max:   maximum,
				Min:   minimum,
				Close: closingTrade.GetPrice()})

			maximum = 0
			minimum = 0
			openingTrade = trade
		}

		if trade.GetPrice() > maximum {
			maximum = trade.GetPrice()
		}

		if trade.GetPrice() < minimum {
			minimum = trade.GetPrice()
		}

		lastDate = trade.GetDate()
		closingTrade = trade
	}

	// Append the last trade
	candlesticks = append(candlesticks, Candlestick{Day: lastDate.String(),
		Open:  openingTrade.GetPrice(),
		Max:   maximum,
		Min:   minimum,
		Close: closingTrade.GetPrice()})

	fmt.Printf("Candlesticks: %v\n", candlesticks)

	return candlesticks, nil
}

func isFromSameDay(time1, time2 time.Time) bool {
	return time1.Day() == time2.Day() &&
		time1.Month() == time2.Month() &&
		time1.Year() == time2.Year()
}

func isLastTrade(index int, trades tradeDataSlice) bool {
	return index == trades.Len()-1
}
