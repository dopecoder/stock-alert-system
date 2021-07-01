package server

import (
	"context"
	"testing"
	"time"

	"github.com/revulcan/stock-alert-system/core"
	data_core "github.com/revulcan/stock-alert-system/data_service/core"
	"github.com/revulcan/stock-alert-system/grpc/trigger_service"
)

func TestNewMockTriggerSystemServer(t *testing.T) {
	NewMockTriggerSystemServer()
}

func getServer(t *testing.T, start bool) *TriggerServiceServer {

	ts := NewTriggerSystemServer()
	if !start {
		return ts
	}

	s, err := ts.StartService(context.Background(), &trigger_service.StartServiceReq{
		Mock: true,
	})
	if err != nil {
		t.Errorf("StartService failed, with error %v\n", err)
	} else if s.Ok != true {
		t.Errorf("StartService failed, got %t, want %t\n", s.Ok, true)
	}
	return ts
}

func TestStartService(t *testing.T) {
	getServer(t, true)
}

func TestStopService(t *testing.T) {
	ts := getServer(t, true)

	res, e := ts.StopService(context.Background(), &trigger_service.StopServiceReq{Mock: true})
	if e != nil {
		t.Errorf("StartService failed, with error %v\n", e)
	} else if res.Ok != true {
		t.Errorf("StartService failed, got %t, want %t\n", res.Ok, true)
	}
}

func TestCreateTrigger(t *testing.T) {
	ts := getServer(t, true)

	res, e := ts.CreateTrigger(context.Background(), &trigger_service.CreateTriggerReq{Id: "1", TAttrib: trigger_service.TriggerAttrib_LTP, Operator: trigger_service.TriggerOperator_GTE, TPrice: 300.0, Scrip: "SBIN", KiteToken: "", ExchangeToken: "3045", Exchange: trigger_service.Exchange_NSE})
	if e != nil {
		t.Errorf("CreateTrigger failed, with error %v\n", e)
	} else if res.Ok != true {
		t.Errorf("CreateTrigger failed, got %t, want %t\n", res.Ok, true)
	}
}

func TestCreateTriggerWithoutStarting(t *testing.T) {
	ts := getServer(t, false)

	_, e := ts.CreateTrigger(context.Background(), &trigger_service.CreateTriggerReq{Id: "1", TAttrib: trigger_service.TriggerAttrib_LTP, Operator: trigger_service.TriggerOperator_GTE, TPrice: 300.0, Scrip: "SBIN", KiteToken: "", ExchangeToken: "3045", Exchange: trigger_service.Exchange_NSE})
	if e == nil {
		t.Errorf("TestCreateTriggerWithoutStarting failed, error not returned %v\n", e)
	}
	if e.Error() != "service not started" {
		t.Errorf("TestCreateTriggerWithoutStarting failed, got %s, want %s\n", e.Error(), "service not started")
	}
}

func TestCreateTriggerWithInvalidId(t *testing.T) {
	ts := getServer(t, true)

	_, e := ts.CreateTrigger(context.Background(), &trigger_service.CreateTriggerReq{Id: "", TAttrib: trigger_service.TriggerAttrib_LTP, Operator: trigger_service.TriggerOperator_GTE, TPrice: 300.0, Scrip: "SBIN", KiteToken: "", ExchangeToken: "3045", Exchange: trigger_service.Exchange_NSE})
	if e == nil {
		t.Errorf("TestCreateTriggerWithInvalidId failed, error not returned %v\n", e)
	}
	if e.Error() != "invalid trigger id - " {
		t.Errorf("TestCreateTriggerWithInvalidId failed, got %s, want %s\n", e.Error(), "invalid trigger id - ")
	}
}

func TestGetTrigger(t *testing.T) {
	ts := getServer(t, true)
	triggerReq := &trigger_service.CreateTriggerReq{Id: "1", TAttrib: trigger_service.TriggerAttrib_LTP, Operator: trigger_service.TriggerOperator_GTE, TPrice: 300.0, Scrip: "SBIN", KiteToken: "", ExchangeToken: "3045", Exchange: trigger_service.Exchange_NSE}
	getRes, e := ts.CreateTrigger(context.Background(), triggerReq)
	if e != nil {
		t.Errorf("CreateTrigger failed, with error %v\n", e)
	} else if getRes.Ok != true {
		t.Errorf("CreateTrigger failed, got %t, want %t\n", getRes.Ok, true)
	}

	got, e := ts.GetTrigger(context.Background(), &trigger_service.GetTriggerReq{TriggerId: "1"})
	if e != nil {
		t.Errorf("GetTriggerStatus failed, with error %v\n", e)
	}
	expected := triggerReq.String()
	if got.String() != expected {
		t.Errorf("GetTriggerStatus failed, got %s, want %s\n", got.String(), expected)
	}
}

func TestGetTriggerWithoutStarting(t *testing.T) {
	ts := getServer(t, false)

	_, e := ts.GetTrigger(context.Background(), &trigger_service.GetTriggerReq{TriggerId: "1"})
	if e == nil {
		t.Errorf("TestGetTriggerWithoutStarting failed, error not returned %v\n", e)
	}
	if e.Error() != "service not started" {
		t.Errorf("TestGetTriggerWithoutStarting failed, got %s, want %s\n", e.Error(), "service not started")
	}
}

func TestGetTriggerStatus(t *testing.T) {
	ts := getServer(t, true)
	triggerReq := &trigger_service.CreateTriggerReq{Id: "1", TAttrib: trigger_service.TriggerAttrib_LTP, Operator: trigger_service.TriggerOperator_GTE, TPrice: 300.0, Scrip: "SBIN", KiteToken: "", ExchangeToken: "3045", Exchange: trigger_service.Exchange_NSE}
	getRes, e := ts.CreateTrigger(context.Background(), triggerReq)
	if e != nil {
		t.Errorf("CreateTrigger failed, with error %v\n", e)
	} else if getRes.Ok != true {
		t.Errorf("CreateTrigger failed, got %t, want %t\n", getRes.Ok, true)
	}

	res, e := ts.GetTriggerStatus(context.Background(), &trigger_service.TriggerStatusReq{TriggerId: "1"})
	if e != nil {
		t.Errorf("GetTriggerStatus failed, with error %v\n", e)
	}
	exp := &core.Trigger{
		Id:         triggerReq.Id,
		TAttrib:    int(triggerReq.TAttrib),
		Operator:   int(triggerReq.Operator),
		TPrice:     triggerReq.TPrice,
		TNearPrice: triggerReq.TNearPrice,
		Instument: &data_core.Instrument{
			Scrip:          triggerReq.Scrip,
			KiteToken:      triggerReq.KiteToken,
			ExchangeToken:  triggerReq.ExchangeToken,
			Exchange:       triggerReq.Exchange.String(),
			InstrumentType: data_core.Stock,
		},
	}
	expected, err := exp.ToString()
	if err != nil {
		t.Errorf("GetTriggerStatus failed, with error %v\n", e)
	}
	if res.Status != expected {
		t.Errorf("GetTriggerStatus failed, got %s, want %s\n", res.Status, expected)
	}
}

func TestGetTriggerStatusWithoutStarting(t *testing.T) {
	ts := getServer(t, false)

	_, e := ts.GetTriggerStatus(context.Background(), &trigger_service.TriggerStatusReq{TriggerId: "1"})
	if e == nil {
		t.Errorf("TestGetTriggerStatusWithoutStarting failed, error not returned %v\n", e)
	}
	if e.Error() != "service not started" {
		t.Errorf("TestGetTriggerStatusWithoutStarting failed, got %s, want %s\n", e.Error(), "service not started")
	}
}

func TestIsTradingTimeStarting(t *testing.T) {
	location, _ := time.LoadLocation("Asia/Kolkata")
	startTime := time.Date(2021, 6, 29, 9, 0, 1, 0, location)
	ts := IsTradingTime(startTime)

	if ts != true {
		t.Errorf("TestIsTradingTime failed, got %t, wanted true\n", ts)
	}
}

func TestIsTradingTimeInBetween(t *testing.T) {
	location, _ := time.LoadLocation("Asia/Kolkata")
	newTime := time.Date(2021, 6, 29, 13, 0, 0, 0, location)
	ts := IsTradingTime(newTime)

	if ts != true {
		t.Errorf("TestIsTradingTime failed, got %t, wanted true\n", ts)
	}
}

func TestIsTradingTimeInWeekend(t *testing.T) {
	location, _ := time.LoadLocation("Asia/Kolkata")
	newTime := time.Date(2021, 6, 27, 9, 0, 0, 0, location)
	ts := IsTradingTime(newTime)

	if ts != false {
		t.Errorf("TestIsTradingTime failed, got %t, wanted false\n", ts)
	}
}

func TestIsTradingTimeOutside(t *testing.T) {
	location, _ := time.LoadLocation("Asia/Kolkata")
	newTime := time.Date(2021, 6, 29, 8, 0, 0, 0, location)
	ts := IsTradingTime(newTime)

	if ts != false {
		t.Errorf("TestIsTradingTime failed, got %t, wanted false\n", ts)
	}
}

func TestIsTradingTimeEnding(t *testing.T) {
	location, _ := time.LoadLocation("Asia/Kolkata")
	endTime := time.Date(2021, 6, 29, 15, 29, 0, 0, location)
	ts := IsTradingTime(endTime)

	if ts != true {
		t.Errorf("TestIsTradingTime failed, got %t, wanted true\n", ts)
	}
}
