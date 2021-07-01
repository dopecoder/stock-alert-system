package ltp_client

import (
	"testing"

	"github.com/revulcan/stock-alert-system/data_service/core"
)

func createNewClient() core.LtpClient {
	return NewMockLTPClient()
	// return NewAngelLTPClient("pVKn0AoU", "Y41983", "Yashas711811")
}
func TestInitClient(t *testing.T) {
	createNewClient()
}

func TestSubscribeClient(t *testing.T) {
	ch := make(chan *core.LTP)
	c := createNewClient()
	handle, err := c.Subscribe([]*core.Instrument{
		{ExchangeToken: "3045", Exchange: "NSE"},
	}, ch)
	if err != nil {
		t.Error(err)
	}

	if handle != 1 {
		t.Errorf("Client TestSubscribeClient failed, got = %d; want = %d", handle, 1)
	}
	got := <-ch
	if got.ExchangeToken != "3045" {
		t.Errorf("Client TestSubscribeClient failed, got = %v;", got)
	}
}

func TestUnSubscribeClient(t *testing.T) {
	ch := make(chan *core.LTP)
	c := createNewClient()
	handle, err := c.Subscribe([]*core.Instrument{
		{ExchangeToken: "3045", Exchange: "NSE"},
	}, ch)
	if err != nil {
		t.Error(err)
	}

	if handle != 1 {
		t.Errorf("Client TestUnSubscribeClient failed, got = %d; want = %d", handle, 1)
	}
	got := <-ch
	if got.ExchangeToken != "3045" {
		t.Errorf("Client TestUnSubscribeClient failed, got = %v;", got)
	}

	unSubErr := c.UnSubscribe(handle)
	if err != nil {
		t.Errorf("Client TestUnSubscribeClient failed, err = %v;", unSubErr)
	}
}

// simple call -> check the state
// concurrent multiple calls to watch -> check the state
