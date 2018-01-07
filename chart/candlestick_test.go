package chart

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/pkg/errors"
)

const (
	// ISO8601 specifies the layout for ISO8601 formatted dates
	//  Example: "2006-01-02T15:04:05-07:00"
	ISO8601 = "2006-01-02T15:04:05-07:00"
)

// Stub Datatypes

type StubTradeData struct {
	Price  float64
	Date   string
	Amount float64
	Type   string
}

func (trade StubTradeData) GetPrice() float64 {
	return trade.Price
}

func (trade StubTradeData) GetDate() time.Time {
	parsedTime, err := time.Parse(ISO8601, trade.Date)
	if err != nil {
		errors.Wrap(err, "error while parsing date")
	}

	return parsedTime
}

func (trade StubTradeData) GetAmount() float64 {
	return trade.Amount
}

func (trade StubTradeData) GetType() string {
	return trade.Type
}

func TestDates(t *testing.T) {

	parsedTime, err := time.Parse(ISO8601, "1987-07-22T00:00:00-03:00")
	if err != nil {
		errors.Wrap(err, "Error parsing date!")
	}

	assert.Equal(t, 22, parsedTime.Day())
	assert.Equal(t, time.Month(07), parsedTime.Month())
	assert.Equal(t, 1987, parsedTime.Year())
}

func TestCandlesticksFromTradeData(t *testing.T) {
	buyStubTrade1 := StubTradeData{Price: 1654.0,
		Date:   "2006-01-02T15:02:05-07:00",
		Amount: 5.0,
		Type:   "buy"}

	sellStubTrade1 := StubTradeData{Price: 1832.0,
		Date:   "2006-01-02T15:04:05-07:00",
		Amount: 1.0,
		Type:   "sell"}

	buyStubTrade2 := StubTradeData{Price: 1273.0,
		Date:   "2006-01-02T15:04:07-07:00",
		Amount: 5.0,
		Type:   "buy"}

	sellStubTrade2 := StubTradeData{Price: 1715.0,
		Date:   "2006-01-02T15:07:05-07:00",
		Amount: 1.0,
		Type:   "sell"}

	buyStubTrade3 := StubTradeData{Price: 1715.0,
		Date:   "2006-01-03T15:07:05-07:00",
		Amount: 1.0,
		Type:   "sell"}

	trades := []TradeData{buyStubTrade1,
		sellStubTrade1,
		buyStubTrade2,
		sellStubTrade2,
		buyStubTrade3}

	candlesticks, err := CandlesticksFromTradeData(trades)
	if assert.NoErrorf(t, err, "erro ao converter trades em candlesticks") {
		assert.NotNil(t, candlesticks)
		assert.Equal(t, 2, len(candlesticks))

		candlestick := candlesticks[0]

		assert.Equal(t, 1832.0, candlestick.Max)
		assert.Equal(t, 1273.0, candlestick.Min)
		assert.Equal(t, 1654.0, candlestick.Open)
		assert.Equal(t, 1715.0, candlestick.Close)
		assert.Equal(t, "2006/01/02", candlestick.Day)
	}

}
