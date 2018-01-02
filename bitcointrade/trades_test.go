package bitcointrade

import (
	"testing"

	"github.com/stretchr/testify/assert"

	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

const (
	stubTradesResponse = "{\"message\":null,\"data\":{\"pagination\":{\"total_pages\":38831,\"current_page\":1,\"page_size\":3,\"registers_count\":116493},\"trades\":[{\"type\":\"sell\",\"amount\":0.00084554,\"unit_price\":48050,\"active_order_code\":\"rkySHruXf\",\"passive_order_code\":\"SkXJrH-7f\",\"date\":\"2018-01-01T23:33:10.960Z\"},{\"type\":\"buy\",\"amount\":0.01,\"unit_price\":48650,\"active_order_code\":\"r1zySSdQG\",\"passive_order_code\":\"Bk7lVBdXf\",\"date\":\"2018-01-01T23:31:37.880Z\"},{\"type\":\"buy\",\"amount\":0.00490924,\"unit_price\":48700,\"active_order_code\":\"r1R7XrOmf\",\"passive_order_code\":\"B1bKWJu7f\",\"date\":\"2018-01-01T23:24:22.067Z\"}]}}"
)

func TestGetTrades(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", tradesEndpointURL,
		httpmock.NewStringResponder(200, stubTradesResponse))

	trades, err := GetTrades("2017-12-01 00:00:00", "2017-12-30 23:59:59")
	if assert.NoError(t, err, "erro durante execução") {
		assert.True(t, len(trades) > 0, "lista de trades deve possuir itens")
	}
}
