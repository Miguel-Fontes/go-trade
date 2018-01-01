// Package bitcointrade contains functions that talks to Bitcointrade's Public API
// and retrieve Trade Data
package bitcointrade

import (
	"errors"
	"testing"

	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

const (
	stubTickerHigh     = 49273 // This variable represents the High Value of the below stub
	stubTickerResponse = "{\"message\":null,\"data\":{\"high\":49273,\"low\":47504,\"volume\":16.41967878,\"trades_quantity\":503,\"last\":48500,\"sell\":48500,\"buy\":48400,\"date\":\"2018-01-01T19:48:46.486Z\"}}"
)

var testTicker = Ticker{
	High:   49275.000000,
	Low:    47504.000000,
	Volume: 20.808434,
	Trades: 534,
	Last:   48599.000000,
	Sell:   48599.000000,
	Buy:    48150.000000,
	Date:   "2018-01-01T18:37:39.358Z"}

func TestTickerToString(t *testing.T) {
	ticker := testTicker

	expected := "Ticker{" +
		"High: 49275.00, " +
		"Low: 47504.00, " +
		"Volume: 20.808434, " +
		"Trades: 534, " +
		"Last: 48599.00, " +
		"Sell: 48599.00, " +
		"Buy: 48150.00, " +
		"Date: \"2018-01-01T18:37:39.358Z\"" +
		"}"

	if ticker.String() != expected {
		t.Errorf("Ticker em formato de String [%s] é diferente do esperado [%s]!", ticker.String(), expected)
	}
}

func TestSuccessMessageToString(t *testing.T) {
	message := Message{
		Message: "",
		Data:    testTicker}

	expected := "Message{" +
		"Message: \"\", " +
		"Data: " + testTicker.String() +
		"}"

	if message.String() != expected {
		t.Errorf("Message em formato String [%s] é diferente do esperado [%s]!", message.String(), expected)
	}
}

func TestFailMessageToString(t *testing.T) {
	message := Message{
		Message: "Failure!",
		Data:    testTicker}

	expected := "Message{" +
		"Message: \"Failure!\", " +
		"Data: " + testTicker.String() +
		"}"

	if message.String() != expected {
		t.Errorf("Message em formato String [%s] é diferente do esperado [%s]!", message.String(), expected)
	}

}

func TestSuccessfulGetTicker(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", tickerEndpointURL,
		httpmock.NewStringResponder(200, stubTickerResponse))

	ticker, err := GetTicker()
	if err != nil {
		t.Errorf("Erro na execução [%s]", err.Error())
	}

	if ticker.High != stubTickerHigh {
		t.Errorf("Valor de Ticker.High não é o esperado! Recebido [%f], esperado [%d].", ticker.High, stubTickerHigh)
	}
}

func TestGetTickerRequestError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", tickerEndpointURL,
		httpmock.NewErrorResponder(errors.New("error executing request")))

	_, err := GetTicker()
	if err == nil {
		t.Errorf("erro em unmarshal de conteúdo inválido não lançado!")
	}
}
