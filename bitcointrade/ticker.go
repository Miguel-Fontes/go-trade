package bitcointrade

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	strUtil "com.miguelmf/stringutil"
)

const (
	tickerEndpointURL = "https://api.bitcointrade.com.br/v1/public/BTC/ticker"
)

// GetTicker gets the current Ticker from BitcoinTrade's Api.
// This function makes a HTTP request to the endpoint, retrieves and
// Unmarshal the data, returning it as a Ticker type.
func GetTicker() (*Ticker, error) {
	resp, err := http.Get(tickerEndpointURL)
	if err != nil {
		log.Fatal("Erro ao efetuar request! [" + err.Error() + "]")
		return nil, errors.New("Erro efetuando request [" + err.Error() + "]")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("Erro ao ler Body de Response [" + err.Error() + "]")
	}

	var message Message

	unmarshalError := json.Unmarshal(body, &message)
	if unmarshalError != nil {
		log.Fatal(unmarshalError)
		return nil, errors.New("Erro durante Unmarshalling [" + err.Error() + "]")
	}

	return &message.Data, nil
}

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
