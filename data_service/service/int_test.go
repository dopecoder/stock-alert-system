package service

import (
	"testing"
	"time"

	"github.com/revulcan/stock-alert-system/data_service/ltp_client"
)

func createNewClient() *ltp_client.AngelLTPClient {
	return ltp_client.NewAngelLTPClient("pVKn0AoU", "Y41983", "Yashas711811")
}

func TestIntStartService(t *testing.T) {
	c := createNewClient()
	service := New(c)
	service.Init()
	time.Sleep(time.Second * 1)
	got := service.listening
	if !got {
		t.Errorf("Service start failed, listening = %t; want true", got)

	}
}

func TestIntStopService(t *testing.T) {
	var c *ltp_client.AngelLTPClient = createNewClient()
	service := New(c)
	service.Init()
	time.Sleep(time.Second * 1)
	service.Close()
	time.Sleep(time.Second * 1)
	got := service.finished
	if !got {
		t.Errorf("Service close failed, finised = %t; want true", got)
	}
}

// func TestIntWatchLTPforInstrument(t *testing.T) {
// 	var c *ltp_client.AngelLTPClient = createNewClient()
// 	service := New(c)
// 	service.Init()
// 	time.Sleep(time.Second * 1)
// 	strm := core.NewMockStream()
// 	go service.WatchLTPforInstrs(
// 		[]core.Instrument{{ExchangeToken: "3045", Exchange: "NSE"}}, strm)

// 	got := <-strm.Strm
// 	ltp := core.LTP{
// 		Ltp:            100,
// 		Open:           100,
// 		Close:          100,
// 		Low:            100,
// 		High:           100,
// 		ExchangeToken:  "3045",
// 		Scrip:          "Scrip",
// 		InstrumentType: 0,
// 	}
// 	if got.ExchangeToken != ltp.ExchangeToken {
// 		t.Errorf("Service WatchLTPforInstruments failed, got = %v; want = %v", got, ltp)
// 	}
// 	service.Close()
// }

// func TestIntWatchLTPforInstruments(t *testing.T) {
// 	var c *ltp_client.AngelLTPClient = createNewClient()
// 	service := New(c)
// 	service.Init()
// 	time.Sleep(time.Second * 1)
// 	strm := core.NewMockStream()
// 	go service.WatchLTPforInstrs([]core.Instrument{
// 		{ExchangeToken: "3045", Exchange: "NSE"},
// 	}, strm)
// 	got := <-strm.Strm
// 	ltp := core.LTP{
// 		Ltp:            100,
// 		Open:           100,
// 		Close:          100,
// 		Low:            100,
// 		High:           100,
// 		ExchangeToken:  "3045",
// 		Scrip:          "Scrip",
// 		InstrumentType: 0,
// 	}
// 	if got.ExchangeToken != ltp.ExchangeToken {
// 		t.Errorf("Service WatchLTPforInstruments failed, got = %v; want = %v", got, ltp)
// 	}
// 	service.Close()
// }

// func TestIntWatchLTPforInstrumentsMutiple(t *testing.T) {
// 	var c *ltp_client.AngelLTPClient = createNewClient()
// 	service := New(c)
// 	service.Init()
// 	time.Sleep(time.Second * 1)
// 	strm := core.NewMockStream()
// 	expected := map[string]bool{
// 		"3041": true,
// 		"3042": true,
// 		"3043": true,
// 		"3044": true,
// 		"3045": true,
// 	}
// 	go service.WatchLTPforInstrs([]core.Instrument{
// 		{ExchangeToken: "3041", Exchange: "NSE"},
// 		{ExchangeToken: "3042", Exchange: "NSE"},
// 		{ExchangeToken: "3043", Exchange: "NSE"},
// 		{ExchangeToken: "3044", Exchange: "NSE"},
// 		{ExchangeToken: "3045", Exchange: "NSE"},
// 	}, strm)

// 	got := make([]core.LTP, 5)
// 	for i := 0; i < 5; i++ {
// 		got[i] = <-strm.Strm
// 		_, exists := expected[got[i].ExchangeToken]
// 		if !exists {
// 			t.Errorf("Service WatchLTPforInstruments failed, got = %v; want = %v", got, expected)
// 		}
// 	}
// 	service.Close()
// }
