package chart

import (
	"sort"
	"time"

	"com.miguelmf/stringutil"
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

func (candlestick Candlestick) String() string {
	return "Candlestick{" +
		"Max: " + stringutil.Float64ToStr(candlestick.Max, 2) + ", " +
		"Min: " + stringutil.Float64ToStr(candlestick.Min, 2) + ", " +
		"Open: " + stringutil.Float64ToStr(candlestick.Open, 2) + ", " +
		"Close: " + stringutil.Float64ToStr(candlestick.Close, 2) + ", " +
		"Day: " + candlestick.Day +
		"}"
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

type tradeIterationInfo struct {
	maximum      float64
	minimum      float64
	openingTrade TradeData
	closingTrade TradeData
	lastDate     time.Time
}

// getDateAsyyyyMMdd is a helper method to extract only the yyyyMMdd portion of the date
func (info tradeIterationInfo) getDateAsyyyyMMdd() string {
	return info.lastDate.Format("2006/01/02")
}

// CandlesticksFromTradeData converts a TradeData set to a set of Candlesticks
func CandlesticksFromTradeData(trades tradeDataSlice) (candlesticks []Candlestick, err error) {
	// sorts the data by day
	sort.Sort(trades)

	info := tradeIterationInfo{}
	info.maximum = trades[0].GetPrice()
	info.minimum = trades[0].GetPrice()
	info.openingTrade = trades[0]
	info.lastDate = trades[0].GetDate()

	for _, trade := range trades {
		if !isFromSameDay(info.lastDate, trade.GetDate()) {
			candlesticks = append(candlesticks, newCandlestick(info))

			info.maximum = 0
			info.minimum = 0
			info.openingTrade = trade
		}

		info.maximum = maximumOf(trade.GetPrice(), info.maximum)
		info.minimum = minimumOf(trade.GetPrice(), info.minimum)
		info.lastDate = trade.GetDate()
		info.closingTrade = trade
	}

	// Append the last trade that won't be added inside the for above, since
	//  we append only when the current trade date is different from the last one
	candlesticks = append(candlesticks, newCandlestick(info))

	return candlesticks, nil
}

func maximumOf(value1, value2 float64) float64 {
	if value1 > value2 {
		return value1
	}

	return value2
}

func minimumOf(value1, value2 float64) float64 {
	if value1 < value2 {
		return value1
	}

	return value2
}

func newCandlestick(info tradeIterationInfo) Candlestick {
	return Candlestick{Day: info.getDateAsyyyyMMdd(),
		Open:  info.openingTrade.GetPrice(),
		Max:   info.maximum,
		Min:   info.minimum,
		Close: info.closingTrade.GetPrice()}
}

func isFromSameDay(time1, time2 time.Time) bool {
	return time1.Day() == time2.Day() &&
		time1.Month() == time2.Month() &&
		time1.Year() == time2.Year()
}
