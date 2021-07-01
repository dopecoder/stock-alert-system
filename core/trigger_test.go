package core

import (
	"testing"

	"github.com/revulcan/stock-alert-system/data_service/core"
)

func TestTriggerNearHitForGTOperator(t *testing.T) {
	c := make(chan *Trigger)
	instr := &core.Instrument{
		Scrip: "SBIN", KiteToken: "", ExchangeToken: "3045", Exchange: "NSE", InstrumentType: core.Stock,
	}
	trigger := &Trigger{Id: "1", TAttrib: LTP, Operator: GTE, TNearPrice: 290, TPrice: 300.0, Instument: instr}
	go trigger.Process(&core.LTP{
		Ltp: 290.05,
	}, c, 1)
	tUpdate := <-c
	if tUpdate.Status != NEAR_HIT {
		t.Errorf("TestTriggerNearHitForGTOperator failed, wanted = %d; got = %d", NEAR_HIT, tUpdate.Status)
	}
}

func TestTriggerNearHitForGTEOperator(t *testing.T) {
	c := make(chan *Trigger)
	instr := &core.Instrument{
		Scrip: "SBIN", KiteToken: "", ExchangeToken: "3045", Exchange: "NSE", InstrumentType: core.Stock,
	}
	trigger := &Trigger{Id: "1", TAttrib: LTP, Operator: GTE, TNearPrice: 290, TPrice: 300.0, Instument: instr}
	go trigger.Process(&core.LTP{
		Ltp: 290,
	}, c, 1)
	tUpdate := <-c
	if tUpdate.Status != NEAR_HIT {
		t.Errorf("TestTriggerNearHitForGTEOperator failed, wanted = %d; got = %d", NEAR_HIT, tUpdate.Status)
	}
}

func TestTriggerNearHitForLTOperator(t *testing.T) {
	c := make(chan *Trigger)
	instr := &core.Instrument{
		Scrip: "SBIN", KiteToken: "", ExchangeToken: "3045", Exchange: "NSE", InstrumentType: core.Stock,
	}
	trigger := &Trigger{Id: "1", TAttrib: LTP, Operator: LT, TNearPrice: 310, TPrice: 300.0, Instument: instr}
	go trigger.Process(&core.LTP{
		Ltp: 309.95,
	}, c, 1)
	tUpdate := <-c
	if tUpdate.Status != NEAR_HIT {
		t.Errorf("TestTriggerNearHitForLTOperator failed, wanted = %d; got = %d", NEAR_HIT, tUpdate.Status)
	}
}

func TestTriggerNearHitForLTEOperator(t *testing.T) {
	c := make(chan *Trigger)
	instr := &core.Instrument{
		Scrip: "SBIN", KiteToken: "", ExchangeToken: "3045", Exchange: "NSE", InstrumentType: core.Stock,
	}
	trigger := &Trigger{Id: "1", TAttrib: LTP, Operator: LTE, TNearPrice: 310, TPrice: 300.0, Instument: instr}
	go trigger.Process(&core.LTP{
		Ltp: 310,
	}, c, 1)
	tUpdate := <-c
	if tUpdate.Status != NEAR_HIT {
		t.Errorf("TestTriggerNearHitForLTEOperator failed, wanted = %d; got = %d", NEAR_HIT, tUpdate.Status)
	}
}

func TestTriggerHitForGTOperator(t *testing.T) {
	c := make(chan *Trigger)
	instr := &core.Instrument{
		Scrip: "SBIN", KiteToken: "", ExchangeToken: "3045", Exchange: "NSE", InstrumentType: core.Stock,
	}
	trigger := &Trigger{Id: "1", TAttrib: LTP, Operator: GTE, TNearPrice: 290, TPrice: 300.0, Instument: instr}
	go trigger.Process(&core.LTP{
		Ltp: 300.05,
	}, c, 1)
	tUpdate := <-c
	if tUpdate.Status != HIT {
		t.Errorf("TestTriggerHitForGTOperator failed, wanted = %d; got = %d", HIT, tUpdate.Status)
	}
}

func TestTriggerHitForGTEOperator(t *testing.T) {
	c := make(chan *Trigger)
	instr := &core.Instrument{
		Scrip: "SBIN", KiteToken: "", ExchangeToken: "3045", Exchange: "NSE", InstrumentType: core.Stock,
	}
	trigger := &Trigger{Id: "1", TAttrib: LTP, Operator: GTE, TNearPrice: 290, TPrice: 300.0, Instument: instr}
	go trigger.Process(&core.LTP{
		Ltp: 300,
	}, c, 1)
	tUpdate := <-c
	if tUpdate.Status != HIT {
		t.Errorf("TestTriggerHitForGTEOperator failed, wanted = %d; got = %d", HIT, tUpdate.Status)
	}
}

func TestTriggerHitForLTOperator(t *testing.T) {
	c := make(chan *Trigger)
	instr := &core.Instrument{
		Scrip: "SBIN", KiteToken: "", ExchangeToken: "3045", Exchange: "NSE", InstrumentType: core.Stock,
	}
	trigger := &Trigger{Id: "1", TAttrib: LTP, Operator: LT, TNearPrice: 310, TPrice: 300.0, Instument: instr}
	go trigger.Process(&core.LTP{
		Ltp: 299.95,
	}, c, 1)
	tUpdate := <-c
	if tUpdate.Status != HIT {
		t.Errorf("TestTriggerHitForLTOperator failed, wanted = %d; got = %d", HIT, tUpdate.Status)
	}
}

func TestTriggerHitForLTEOperator(t *testing.T) {
	c := make(chan *Trigger)
	instr := &core.Instrument{
		Scrip: "SBIN", KiteToken: "", ExchangeToken: "3045", Exchange: "NSE", InstrumentType: core.Stock,
	}
	trigger := &Trigger{Id: "1", TAttrib: LTP, Operator: LTE, TNearPrice: 310, TPrice: 300.0, Instument: instr}
	go trigger.Process(&core.LTP{
		Ltp: 300,
	}, c, 1)
	tUpdate := <-c
	if tUpdate.Status != HIT {
		t.Errorf("TestTriggerHitForLTEOperator failed, wanted = %d; got = %d", HIT, tUpdate.Status)
	}
}
