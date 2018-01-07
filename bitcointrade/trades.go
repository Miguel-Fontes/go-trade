package bitcointrade

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	strUtil "com.miguelmf/stringutil"
	"github.com/pkg/errors"
)

const (
	tradesEndpointURL = "https://api.bitcointrade.com.br/v1/public/BTC/trades?start_time=2016-10-01T00:00:00-03:00&end_time=2018-10-10T23:59:59-03:00&page_size=100&current_page=1"
)

// TradesMessage represents the envolope of a message received from Bitcointrade
type TradesMessage struct {
	Message string
	Data    TradesPage
}

// TradesPage contains a page of Trades data and pagination metadata
type TradesPage struct {
	Pagination Pagination `json:"pagination"`
	Trades     []Trade    `json:"trades"`
}

// Pagination contains the pagination Metadata, making it possible to work
// with large Trades Datasets in chunks
type Pagination struct {
	Pages          int `json:"total_pages"`
	Current        int `json:"current_page"`
	Size           int `json:"page_size"`
	RegistersCount int `json:"registers_count"`
}

// Trade represents a single Trade data
type Trade struct {
	Type        string  `json:"type"`
	Amount      float32 `json:"amount"`
	UnitPrice   float32 `json:"unit_price"`
	ActiveCode  string  `json:"active_order_code"`
	PassiveCode string  `json:"passive_order_code"`
	Date        string  `json:"date"`
}

// String returns a String representation of the type TradesMessage
func (message TradesMessage) String() string {
	return "TradesMessage{" +
		"Message: " + message.Message + ", " +
		"Data: " + message.Data.String() +
		"}"
}

// String returns a String representation of the type TradesPage
func (tradesPage TradesPage) String() string {
	return "TradesPage{" +
		"Page: " + tradesPage.Pagination.String() + ", " +
		"Trades: [" + tradeArrayToString(tradesPage.Trades) + "]" +
		"}"
}

// String returns a String representation of the type Pagination
func (pagination Pagination) String() string {
	return "Pagination{" +
		"Pages: " + strUtil.IntToStr(pagination.Pages) + ", " +
		"Current: " + strUtil.IntToStr(pagination.Current) + ", " +
		"Size: " + strUtil.IntToStr(pagination.Size) + ", " +
		"Count: " + strUtil.IntToStr(pagination.RegistersCount) +
		"}"
}

func tradeArrayToString(trades []Trade) (result string) {
	for index, trade := range trades {
		result += trade.String()
		if len(trades)-1 != index {
			result += ", "
		}
	}

	return result
}

// String returns a String representation of the type Trade
func (trade Trade) String() string {
	return "Trade{" +
		"Type: " + trade.Type + ", " +
		"Amount: " + strUtil.FloatToStr(trade.Amount, 6) + ", " +
		"Unit_price: " + strUtil.FloatToStr(trade.UnitPrice, 2) + ", " +
		"Active_order_code: " + trade.ActiveCode + ", " +
		"Passive_order_code: " + trade.PassiveCode + ", " +
		"Date: " + trade.Date +
		"}"
}

// GetTrades fetches trades from the given time period (1000 maximum)
func GetTrades(diaInicial, diaFinal string) ([]Trade, error) {
	resp, getErr := http.Get(tradesEndpointURL)
	if getErr != nil {
		return nil, errors.Wrap(getErr, "erro efetuando request")
	}

	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return nil, errors.Wrap(readErr, "erro ao ler Body de response")
	}

	var message TradesMessage

	unmarshalError := json.Unmarshal(body, &message)
	if unmarshalError != nil {
		return nil, errors.Wrap(unmarshalError, "erro durante Unmarshalling")
	}

	return message.Data.Trades, nil
}
