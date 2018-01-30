package bitcointrade

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	strUtil "github.com/miguel-fontes/stringutil"
	"github.com/miguel-fontes/timeutil"
	"github.com/pkg/errors"
)

const (
	// ?start_time=2016-10-01T00:00:00-03:00&end_time=2018-10-10T23:59:59-03:00&page_size=1000&current_page=1
	tradesEndpointURL = "https://api.bitcointrade.com.br/v1/public/BTC/trades"
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

// GetPrice returns the unit price of the trade
func (trade Trade) GetPrice() float64 {
	return float64(trade.UnitPrice)
}

// GetDate returns the date of the trade
func (trade Trade) GetDate() time.Time {
	parsedTime, err := time.Parse(time.RFC3339Nano, trade.Date)
	if err != nil {
		errors.Wrap(err, "error while parsing date")
	}

	return parsedTime
}

// GetAmount returns the Amount traded
func (trade Trade) GetAmount() float64 {
	return float64(trade.Amount)
}

// GetType returns the type of the trade: buy or sell
func (trade Trade) GetType() string {
	return trade.Type
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

func buildRequest(diaInicial, diaFinal time.Time, currentPage int) (*http.Request, error) {
	req, err := http.NewRequest("GET", tradesEndpointURL, nil)

	q := req.URL.Query()
	q.Set("start_time", diaInicial.Format(time.RFC3339Nano))
	q.Set("end_time", diaFinal.Format(time.RFC3339Nano))
	q.Set("page_size", "1000")
	q.Set("current_page", strUtil.IntToStr(currentPage))
	req.URL.RawQuery = q.Encode()

	return req, err
}

// GetTrades fetches trades from the given time period (1000 maximum)
func GetTrades(dataInicial, dataFinal time.Time) ([]Trade, error) {
	client := &http.Client{}

	var message TradesMessage
	message.Data.Pagination.Pages = 99
	currentPage := 1
	trades := []Trade{}

	for currentPage <= message.Data.Pagination.Pages && message.Data.Pagination.Pages != 0 {
		req, _ := buildRequest(dataInicial, dataFinal, currentPage)

		resp, getErr := client.Do(req)
		if getErr != nil {
			return nil, errors.Wrap(getErr, "erro efetuando request")
		}

		defer resp.Body.Close()

		body, readErr := ioutil.ReadAll(resp.Body)
		if readErr != nil {
			return nil, errors.Wrap(readErr, "erro ao ler Body de response")
		}

		message = TradesMessage{}

		unmarshalError := json.Unmarshal(body, &message)
		if unmarshalError != nil {
			return nil, errors.Wrap(unmarshalError, "erro durante unmarshalling")
		}

		trades = append(trades, message.Data.Trades...)
		log.Printf("fim da leitura de pagina %d/%d, total de trades [%d]",
			currentPage,
			message.Data.Pagination.Pages,
			len(trades))

		currentPage++
	}

	log.Printf("consumo finalizado, localizados [%d] trades", len(trades))

	return trades, nil

}

func ungroupDates(dataInicial, dataFinal time.Time) []time.Time {
	days := []time.Time{}
	currentDate := dataInicial
	for !timeutil.IsSameDay(currentDate, dataFinal) {
		days = append(days, currentDate)
		currentDate = currentDate.AddDate(0, 0, 1)
	}

	days = append(days, currentDate)

	return days
}
