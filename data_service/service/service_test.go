package service

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
	"time"

	"github.com/revulcan/stock-alert-system/data_service/core"
	"github.com/revulcan/stock-alert-system/data_service/ltp_client"
)

func TestStartStopService(t *testing.T) {
	var c *ltp_client.MockLTPClient = ltp_client.NewMockLTPClient()
	service := New(c)
	service.Init()
	time.Sleep(time.Second * 1)
	got := service.listening
	if !got {
		t.Errorf("Service start failed, listening = %t; want true", got)
	}

	service.Close()
	time.Sleep(time.Second * 1)
	got = service.finished
	if !got {
		t.Errorf("Service close failed, finised = %t; want true", got)
	}
}

func TestWatchLTPforInstrument(t *testing.T) {
	var c *ltp_client.MockLTPClient = ltp_client.NewMockLTPClient()
	service := New(c)
	service.Init()
	time.Sleep(time.Second * 1)
	strm := core.NewMockStream()
	go service.WatchLTPforInstrs(
		[]*core.Instrument{{ExchangeToken: "3045"}}, strm)

	got := <-strm.Strm
	ltp := core.LTP{
		Ltp:            100,
		Open:           100,
		Close:          100,
		Low:            100,
		High:           100,
		ExchangeToken:  "3045",
		Scrip:          "Scrip",
		InstrumentType: 0,
	}
	if *got != ltp {
		t.Errorf("Service WatchLTPforInstruments failed, got = %v; want = %v", got, ltp)
	}
	service.Close()
}

func TestWatchLTPforInstruments(t *testing.T) {
	var c *ltp_client.MockLTPClient = ltp_client.NewMockLTPClient()
	service := New(c)
	service.Init()
	time.Sleep(time.Second * 1)
	strm := core.NewMockStream()
	go service.WatchLTPforInstrs([]*core.Instrument{
		{ExchangeToken: "3045"},
	}, strm)
	got := <-strm.Strm
	ltp := core.LTP{
		Ltp:            100,
		Open:           100,
		Close:          100,
		Low:            100,
		High:           100,
		ExchangeToken:  "3045",
		Scrip:          "Scrip",
		InstrumentType: 0,
	}
	if *got != ltp {
		t.Errorf("Service WatchLTPforInstruments failed, got = %v; want = %v", got, ltp)
	}
	service.Close()
}

func TestWatchLTPforInstrumentsMutiple(t *testing.T) {
	var c *ltp_client.MockLTPClient = ltp_client.NewMockLTPClient()
	service := New(c)
	service.Init()
	time.Sleep(time.Second * 1)
	strm := core.NewMockStream()
	go service.WatchLTPforInstrs([]*core.Instrument{
		{ExchangeToken: "3041"},
		{ExchangeToken: "3042"},
		{ExchangeToken: "3043"},
		{ExchangeToken: "3044"},
		{ExchangeToken: "3045"},
	}, strm)

	expected := make([]core.LTP, 5)
	for i := 0; i < 5; i++ {
		expected[i] = core.LTP{
			Ltp:            100,
			Open:           100,
			Close:          100,
			Low:            100,
			High:           100,
			ExchangeToken:  fmt.Sprintf("304%d", i+1),
			Scrip:          "Scrip",
			InstrumentType: 0,
		}
	}

	got := make([]core.LTP, 5)
	for i := 0; i < 5; i++ {
		got[i] = *<-strm.Strm
		fmt.Printf("Got - %v\n", got[i])
	}

	if !array_sorted_equal(expected, got) {
		t.Errorf("Service WatchLTPforInstruments failed, got = %v; want = %v", got, expected)
	}
	service.Close()
}

type byLTP []core.LTP

func (s byLTP) Len() int {
	return len(s)
}
func (s byLTP) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLTP) Less(i, j int) bool {
	return s[i].ExchangeToken < s[j].ExchangeToken
}

func array_sorted_equal(a, b []core.LTP) bool {
	if len(a) != len(b) {
		return false
	}

	a_copy := make([]core.LTP, len(a))
	b_copy := make([]core.LTP, len(b))

	copy(a_copy, a)
	copy(b_copy, b)

	sort.Sort(byLTP(a_copy))
	sort.Sort(byLTP(b_copy))

	return reflect.DeepEqual(a_copy, b_copy)
}

// func array_sorted_equal_token(a, b []core.LTP) bool {
// 	if len(a) != len(b) {
// 		return false
// 	}

// 	a_copy := make([]core.LTP, len(a))
// 	b_copy := make([]core.LTP, len(b))

// 	copy(a_copy, a)
// 	copy(b_copy, b)

// 	sort.Sort(byLTP(a_copy))
// 	sort.Sort(byLTP(b_copy))

// 	if len(a_copy) != len(b_copy) {
// 		return false
// 	}

// 	for i := 0; i < len(a_copy); i++ {
// 		if a_copy[i].ExchangeToken != b_copy[i].ExchangeToken {
// 			return false
// 		}
// 	}

// 	return true
// }

// simple call -> check the state
// concurrent multiple calls to watch -> check the state
