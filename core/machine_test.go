package core

import (
	"context"
	"testing"
	"time"

	"github.com/revulcan/stock-alert-system/data_service/core"
	"github.com/revulcan/stock-alert-system/data_service/ltp_client"
	"github.com/revulcan/stock-alert-system/data_service/service"
)

func getNewService(c *ltp_client.MockLTPClient) *service.Service {
	s := service.New(c)
	s.Init()
	return s
}

func TestStartStopMachine(t *testing.T) {
	c := ltp_client.NewMockLTPClientWithControl(298.0, 301.0, 1.0, time.Millisecond*1000)
	i := &core.Instrument{
		Scrip: "SBIN", KiteToken: "", ExchangeToken: "3045", Exchange: "NSE", InstrumentType: core.Stock,
	}
	tm := NewTriggerMachine(context.Background(), i, getNewService(c), nil)
	tm.Start()
	if tm.started != true {
		t.Errorf("TestStartStopMachine failed, wanted = %t; got = %t", true, tm.started)
	}

	tm.Stop()
	time.Sleep(time.Millisecond * 500)
	if tm.finished != true {
		t.Errorf("TestStartStopMachine failed, wanted = %t; got = %t", true, tm.finished)
	}
}

func TestAddTrigger(t *testing.T) {
	c := ltp_client.NewMockLTPClientWithControl(298.0, 302.0, 1.0, time.Millisecond*300)
	i := &core.Instrument{
		Scrip: "SBIN", KiteToken: "", ExchangeToken: "3045", Exchange: "NSE", InstrumentType: core.Stock,
	}
	tm := NewTriggerMachine(context.Background(), i, getNewService(c), nil)
	tm.Start()
	if tm.started != true {
		t.Errorf("TestAddTrigger failed, wanted = %t; got = %t", true, tm.started)
	}

	trigger := &Trigger{Id: "1", TAttrib: LTP, Operator: GTE, TNearPrice: 299.0, TPrice: 300.0, Instument: i}
	err := tm.AddTrigger(trigger)
	if err != nil {
		t.Errorf("TestAddTrigger failed, got error = %v ", err)
	}
	time.Sleep(time.Second * 3)
	if trigger.Status != HIT {
		t.Errorf("TestAddTrigger failed, wanted = %d; got = %d", HIT, trigger.Status)
	}

	tm.Stop()
	time.Sleep(time.Millisecond * 500)
	if tm.finished != true {
		t.Errorf("TestAddTrigger failed, wanted = %t; got = %t", true, tm.finished)
	}
}

func TestNearHitGT(t *testing.T) {
	c := ltp_client.NewMockLTPClientWithControl(297.0, 300.0, 0.5, time.Millisecond*300)
	i := &core.Instrument{
		Scrip: "SBIN", KiteToken: "", ExchangeToken: "3045", Exchange: "NSE", InstrumentType: core.Stock,
	}
	tm := NewTriggerMachine(context.Background(), i, getNewService(c), nil)
	tm.Start()
	if tm.started != true {
		t.Errorf("TestNearHitGT failed, wanted = %t; got = %t", true, tm.started)
	}

	trigger := &Trigger{Id: "1", TAttrib: LTP, Operator: GT, TNearPrice: 299.0, TPrice: 300.0, Instument: i}
	err := tm.AddTrigger(trigger)
	if err != nil {
		t.Errorf("TestNearHitGT failed, got error = %v ", err)
	}
	time.Sleep(time.Second * 3)
	if trigger.Status != NEAR_HIT {
		t.Errorf("TestNearHitGT failed, wanted = %d; got = %d", NEAR_HIT, trigger.Status)
	}

	tm.Stop()
	time.Sleep(time.Millisecond * 500)
	if tm.finished != true {
		t.Errorf("TestNearHitGT failed, wanted = %t; got = %t", true, tm.finished)
	}
}

func TestNearHitGTE(t *testing.T) {
	c := ltp_client.NewMockLTPClientWithControl(297.0, 299.0, 0.5, time.Millisecond*300)
	i := &core.Instrument{
		Scrip: "SBIN", KiteToken: "", ExchangeToken: "3045", Exchange: "NSE", InstrumentType: core.Stock,
	}
	tm := NewTriggerMachine(context.Background(), i, getNewService(c), nil)
	tm.Start()
	if tm.started != true {
		t.Errorf("TestNearHitGTE failed, wanted = %t; got = %t", true, tm.started)
	}

	trigger := &Trigger{Id: "1", TAttrib: LTP, Operator: GTE, TNearPrice: 299.0, TPrice: 300.0, Instument: i}
	err := tm.AddTrigger(trigger)
	if err != nil {
		t.Errorf("TestNearHitGTE failed, got error = %v ", err)
	}
	time.Sleep(time.Second * 3)
	if trigger.Status != NEAR_HIT {
		t.Errorf("TestNearHitGTE failed, wanted = %d; got = %d", NEAR_HIT, trigger.Status)
	}

	tm.Stop()
	time.Sleep(time.Millisecond * 500)
	if tm.finished != true {
		t.Errorf("TestNearHitGTE failed, wanted = %t; got = %t", true, tm.finished)
	}
}

func TestNearHitLT(t *testing.T) {
	c := ltp_client.NewMockLTPClientWithControl(302.0, 300.5, -0.5, time.Millisecond*300)
	i := &core.Instrument{
		Scrip: "SBIN", KiteToken: "", ExchangeToken: "3045", Exchange: "NSE", InstrumentType: core.Stock,
	}
	tm := NewTriggerMachine(context.Background(), i, getNewService(c), nil)
	tm.Start()
	if tm.started != true {
		t.Errorf("TestNearHitLT failed, wanted = %t; got = %t", true, tm.started)
	}

	trigger := &Trigger{Id: "1", TAttrib: LTP, Operator: LT, TNearPrice: 301.0, TPrice: 300.0, Instument: i}
	err := tm.AddTrigger(trigger)
	if err != nil {
		t.Errorf("TestNearHitLT failed, got error = %v ", err)
	}
	time.Sleep(time.Second * 3)
	if trigger.Status != NEAR_HIT {
		t.Errorf("TestNearHitLT failed, wanted = %d; got = %d", NEAR_HIT, trigger.Status)
	}

	tm.Stop()
	time.Sleep(time.Millisecond * 500)
	if tm.finished != true {
		t.Errorf("TestNearHitLT failed, wanted = %t; got = %t", true, tm.finished)
	}
}

func TestNearHitLTE(t *testing.T) {
	c := ltp_client.NewMockLTPClientWithControl(302.0, 301, -0.5, time.Millisecond*300)
	i := &core.Instrument{
		Scrip: "SBIN", KiteToken: "", ExchangeToken: "3045", Exchange: "NSE", InstrumentType: core.Stock,
	}
	tm := NewTriggerMachine(context.Background(), i, getNewService(c), nil)
	tm.Start()
	if tm.started != true {
		t.Errorf("TestNearHitLTE failed, wanted = %t; got = %t", true, tm.started)
	}

	trigger := &Trigger{Id: "1", TAttrib: LTP, Operator: LTE, TNearPrice: 301.0, TPrice: 300.0, Instument: i}
	err := tm.AddTrigger(trigger)
	if err != nil {
		t.Errorf("TestNearHitLTE failed, got error = %v ", err)
	}
	time.Sleep(time.Second * 3)
	if trigger.Status != NEAR_HIT {
		t.Errorf("TestNearHitLTE failed, wanted = %d; got = %d", NEAR_HIT, trigger.Status)
	}

	tm.Stop()
	time.Sleep(time.Millisecond * 500)
	if tm.finished != true {
		t.Errorf("TestNearHitLTE failed, wanted = %t; got = %t", true, tm.finished)
	}
}

func TestHitGT(t *testing.T) {
	c := ltp_client.NewMockLTPClientWithControl(298.0, 300.5, 0.5, time.Millisecond*300)
	i := &core.Instrument{
		Scrip: "SBIN", KiteToken: "", ExchangeToken: "3045", Exchange: "NSE", InstrumentType: core.Stock,
	}
	tm := NewTriggerMachine(context.Background(), i, getNewService(c), nil)
	tm.Start()
	if tm.started != true {
		t.Errorf("TestHitGT failed, wanted = %t; got = %t", true, tm.started)
	}

	trigger := &Trigger{Id: "1", TAttrib: LTP, Operator: GT, TNearPrice: 299.0, TPrice: 300.0, Instument: i}
	err := tm.AddTrigger(trigger)
	if err != nil {
		t.Errorf("TestHitGT failed, got error = %v ", err)
	}
	time.Sleep(time.Second * 3)
	if trigger.Status != HIT {
		t.Errorf("TestHitGT failed, wanted = %d; got = %d", HIT, trigger.Status)
	}

	tm.Stop()
	time.Sleep(time.Millisecond * 500)
	if tm.finished != true {
		t.Errorf("TestHitGT failed, wanted = %t; got = %t", true, tm.finished)
	}
}

func TestHitGTE(t *testing.T) {
	c := ltp_client.NewMockLTPClientWithControl(298.0, 300.0, 0.5, time.Millisecond*300)
	i := &core.Instrument{
		Scrip: "SBIN", KiteToken: "", ExchangeToken: "3045", Exchange: "NSE", InstrumentType: core.Stock,
	}
	tm := NewTriggerMachine(context.Background(), i, getNewService(c), nil)
	tm.Start()
	if tm.started != true {
		t.Errorf("TestHitGTE failed, wanted = %t; got = %t", true, tm.started)
	}

	trigger := &Trigger{Id: "1", TAttrib: LTP, Operator: GTE, TNearPrice: 299.0, TPrice: 300.0, Instument: i}
	err := tm.AddTrigger(trigger)
	if err != nil {
		t.Errorf("TestHitGTE failed, got error = %v ", err)
	}
	time.Sleep(time.Second * 3)
	if trigger.Status != HIT {
		t.Errorf("TestHitGTE failed, wanted = %d; got = %d", HIT, trigger.Status)
	}

	tm.Stop()
	time.Sleep(time.Millisecond * 500)
	if tm.finished != true {
		t.Errorf("TestHitGTE failed, wanted = %t; got = %t", true, tm.finished)
	}
}

func TestHitLT(t *testing.T) {
	c := ltp_client.NewMockLTPClientWithControl(300.5, 299.5, -0.5, time.Millisecond*300)
	i := &core.Instrument{
		Scrip: "SBIN", KiteToken: "", ExchangeToken: "3045", Exchange: "NSE", InstrumentType: core.Stock,
	}
	tm := NewTriggerMachine(context.Background(), i, getNewService(c), nil)
	tm.Start()
	if tm.started != true {
		t.Errorf("TestHitLT failed, wanted = %t; got = %t", true, tm.started)
	}

	trigger := &Trigger{Id: "1", TAttrib: LTP, Operator: LT, TNearPrice: 301.0, TPrice: 300.0, Instument: i}
	err := tm.AddTrigger(trigger)
	if err != nil {
		t.Errorf("TestHitLT failed, got error = %v ", err)
	}
	time.Sleep(time.Second * 3)
	if trigger.Status != HIT {
		t.Errorf("TestHitLT failed, wanted = %d; got = %d", HIT, trigger.Status)
	}

	tm.Stop()
	time.Sleep(time.Millisecond * 500)
	if tm.finished != true {
		t.Errorf("TestHitLT failed, wanted = %t; got = %t", true, tm.finished)
	}
}

func TestHitLTE(t *testing.T) {
	c := ltp_client.NewMockLTPClientWithControl(301, 300, -0.5, time.Millisecond*300)
	i := &core.Instrument{
		Scrip: "SBIN", KiteToken: "", ExchangeToken: "3045", Exchange: "NSE", InstrumentType: core.Stock,
	}
	tm := NewTriggerMachine(context.Background(), i, getNewService(c), nil)
	tm.Start()
	if tm.started != true {
		t.Errorf("TestHitLTE failed, wanted = %t; got = %t", true, tm.started)
	}

	trigger := &Trigger{Id: "1", TAttrib: LTP, Operator: LTE, TNearPrice: 301.0, TPrice: 300.0, Instument: i}
	err := tm.AddTrigger(trigger)
	if err != nil {
		t.Errorf("TestHitLTE failed, got error = %v ", err)
	}
	time.Sleep(time.Second * 3)
	if trigger.Status != HIT {
		t.Errorf("TestHitLTE failed, wanted = %d; got = %d", HIT, trigger.Status)
	}

	tm.Stop()
	time.Sleep(time.Millisecond * 500)
	if tm.finished != true {
		t.Errorf("TestHitLTE failed, wanted = %t; got = %t", true, tm.finished)
	}
}

// todo, multiple triggers
