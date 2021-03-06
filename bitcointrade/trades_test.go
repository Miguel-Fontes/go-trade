package bitcointrade

import (
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

const (
	stubTradesResponse = "{\"message\":null,\"data\":{\"pagination\":{\"total_pages\":1,\"current_page\":1,\"page_size\":3,\"registers_count\":116493},\"trades\":[{\"type\":\"sell\",\"amount\":0.00084554,\"unit_price\":48050,\"active_order_code\":\"rkySHruXf\",\"passive_order_code\":\"SkXJrH-7f\",\"date\":\"2018-01-01T23:33:10.960Z\"},{\"type\":\"buy\",\"amount\":0.01,\"unit_price\":48650,\"active_order_code\":\"r1zySSdQG\",\"passive_order_code\":\"Bk7lVBdXf\",\"date\":\"2018-01-01T23:31:37.880Z\"},{\"type\":\"buy\",\"amount\":0.00490924,\"unit_price\":48700,\"active_order_code\":\"r1R7XrOmf\",\"passive_order_code\":\"B1bKWJu7f\",\"date\":\"2018-01-01T23:24:22.067Z\"}]}}"
)

func TestGetTrades(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", tradesEndpointURL,
		httpmock.NewStringResponder(200, stubTradesResponse))

	dataInicial, _ := time.Parse(time.RFC3339Nano, "2017-12-01T00:00:00")
	dataFinal, _ := time.Parse(time.RFC3339Nano, "2017-12-30T23:59:59")

	trades, err := GetTrades(dataInicial, dataFinal)
	if assert.NoError(t, err, "erro durante execução") {
		assert.True(t, len(trades) > 0, "lista de trades deve possuir itens")
	}
}

func TestGetTradesShouldThrowRequestError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", tradesEndpointURL,
		httpmock.NewErrorResponder(errors.New("error")))

	dataInicial, _ := time.Parse(time.RFC3339Nano, "2017-12-01T00:00:00")
	dataFinal, _ := time.Parse(time.RFC3339Nano, "2017-12-30T23:59:59")

	trades, err := GetTrades(dataInicial, dataFinal)
	if assert.Error(t, err, "erro durante execução") {
		assert.True(t, len(trades) == 0, "lista de trades não deve possuir itens")
	}
}

func TestGetTradesShouldThrowUnmarshalError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", tradesEndpointURL,
		httpmock.NewStringResponder(200, "{stubTradesResponse}"))

	dataInicial, _ := time.Parse(time.RFC3339Nano, "2017-12-01T00:00:00")
	dataFinal, _ := time.Parse(time.RFC3339Nano, "2017-12-30T23:59:59")

	trades, err := GetTrades(dataInicial, dataFinal)
	if assert.Error(t, err, "erro durante unmarshalling") {
		assert.True(t, len(trades) == 0, "lista de trades não deve possuir itens")
	}
}
