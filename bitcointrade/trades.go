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

// TradesPage is the Top Level data Structure that contains a page of Trades info and
// pagination metadata
type TradesPage struct {
	Pagination Pagination `json:"pagination"`
	Trades     []Trade    `json:"trades"`
}

// Trades contains the list of trades received from Bitcontrade
type Trades struct {
	Data []Trade
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

func (message TradesMessage) String() string {
	return "TradesMessage{" +
		"Message: " + message.Message +
		"Data: " + message.Data.String() +
		"}"
}

func (tradesPage TradesPage) String() string {
	return "TradesPage{" +
		"Message: " + tradesPage.Pagination.String() +
		// "List: [" + tradesPage.Trades.String() + "]" +
		"}"
}

func (pagination Pagination) String() string {
	return "Pagination{" +
		"Pages: " + strUtil.IntToStr(pagination.Pages) + ", " +
		"Current: " + strUtil.IntToStr(pagination.Current) + ", " +
		"Size: " + strUtil.IntToStr(pagination.Size) + ", " +
		"Count: " + strUtil.IntToStr(pagination.RegistersCount) +
		"}"
}

func (trades Trades) String() (result string) {
	for _, trade := range trades.Data {
		result += trade.String()
	}

	return result
}

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

// GetTrades busca os últimos 1000 trades de acordo com os critérios
// especificados por parâmetro.
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

	println(string(body))

	var message TradesMessage

	unmarshalError := json.Unmarshal(body, &message)
	if unmarshalError != nil {
		return nil, errors.Wrap(unmarshalError, "erro durante Unmarshalling")
	}

	return message.Data.Trades, nil
}

/*
{
	"message": null,
	"data": {
	  "pagination": {
		"total_pages": 1,
		"current_page": 1,
		"page_size": 100,
		"registers_count": 3
	  },
	  "trades": [
		{
		  "type": "buy",
		  "amount": 0.005,
		  "unit_price": 9999.99,
		  "active_order_code": "By4qWV-p_",
		  "passive_order_code": "HJ-ddb-6_",
		  "date": "2017-10-07T01:25:42.307Z"
		},
		{
		  "type": "sell",
		  "amount": 0.00861033,
		  "unit_price": 9999.99,
		  "active_order_code": "Ay4qWV-p_",
		  "passive_order_code": "AJ-ddb-6_",
		  "date": "2017-10-06T14:22:24.020Z"
		},
		{
		  "type": "sell",
		  "amount": 0.002,
		  "unit_price": 9999.99,
		  "active_order_code": "Cy4qWV-p_",
		  "passive_order_code": "CJ-ddb-6_",
		  "date": "2017-10-06T14:22:24.017Z"
		}
	  ]
	}
  }
*/
