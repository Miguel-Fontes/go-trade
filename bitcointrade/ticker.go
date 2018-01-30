package bitcointrade

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"

	strUtil "github.com/miguel-fontes/stringutil"
)

const (
	tickerEndpointURL = "https://api.bitcointrade.com.br/v1/public/BTC/ticker"
)

// Message represents the envolope of a message received from Bitcointrade
type Message struct {
	Message string
	Data    Ticker
}

// Ticker data, received from Bitcointrade's Ticker endpoint
type Ticker struct {
	High   float32
	Low    float32
	Volume float32
	Trades int `json:"trades_quantity"`
	Last   float32
	Sell   float32
	Buy    float32
	Date   string
}

// String converts an Message to a string
func (message Message) String() string {
	return "Message{" +
		"Message: \"" + message.Message + "\", " +
		"Data: " + message.Data.String() +
		"}"
}

// String converts an Ticker to a tring
func (ticker Ticker) String() string {
	return "Ticker{" +
		"High: " + strUtil.FloatToStr(ticker.High, 2) + ", " +
		"Low: " + strUtil.FloatToStr(ticker.Low, 2) + ", " +
		"Volume: " + strUtil.FloatToStr(ticker.Volume, 6) + ", " +
		"Trades: " + strUtil.IntToStr(ticker.Trades) + ", " +
		"Last: " + strUtil.FloatToStr(ticker.Last, 2) + ", " +
		"Sell: " + strUtil.FloatToStr(ticker.Sell, 2) + ", " +
		"Buy: " + strUtil.FloatToStr(ticker.Buy, 2) + ", " +
		"Date: \"" + ticker.Date + "\"" +
		"}"
}

// GetTicker gets the current Ticker from BitcoinTrade's Api.
// This function makes a HTTP request to the endpoint, retrieves and
// Unmarshal the data, returning it as a Ticker type.
func GetTicker() (*Ticker, error) {
	resp, getErr := http.Get(tickerEndpointURL)
	if getErr != nil {
		return nil, errors.Wrap(getErr, "erro efetuando request")
	}

	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return nil, errors.Wrap(readErr, "erro ao ler Body de response")
	}

	var message Message

	unmarshalError := json.Unmarshal(body, &message)
	if unmarshalError != nil {
		return nil, errors.Wrap(unmarshalError, "erro durante Unmarshalling")
	}
	return &message.Data, nil
}
