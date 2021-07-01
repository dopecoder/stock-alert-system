package service

import (
	"context"
	"testing"

	"github.com/revulcan/stock-alert-system/core"
	ds_core "github.com/revulcan/stock-alert-system/data_service/core"
	"github.com/revulcan/stock-alert-system/data_service/ltp_client"
)

// func (ts *TriggerSystem) setupTrigger(t *core.Trigger, tm *core.TriggerMachine) error
// func (ts *TriggerSystem) setupService() []error

// func (ts *TriggerSystem) GetTrigger(id string) (*core.Trigger, error)
// func (ts *TriggerSystem) CreateTrigger(t *core.Trigger) error
// func (ts *TriggerSystem) GetTriggerStatus(id string) (string, error)

func TestCreateTrigger(t *testing.T) {
	c := ltp_client.NewMockLTPClient()
	repo := &core.MockRepo{}
	ts := NewTriggerSystem(context.Background(), c, repo, &core.Config{
		UpdatePostEndpoint: "https://events.resignal.app/onTriggerSuccess",
	})
	instr := &ds_core.Instrument{
		Scrip: "SBIN", KiteToken: "", ExchangeToken: "3045", Exchange: "NSE", InstrumentType: ds_core.Stock,
	}
	trigger := &core.Trigger{Id: "1", TAttrib: core.LTP, Operator: core.GTE, TNearPrice: 290, TPrice: 300.0, Instument: instr}
	err := ts.CreateTrigger(trigger)
	if err != nil {
		t.Errorf("TestCreateTrigger failed, error returned %v\n", err)
	}
}

func TestDeleteTrigger(t *testing.T) {
	c := ltp_client.NewMockLTPClient()
	repo := &core.MockRepo{}
	ts := NewTriggerSystem(context.Background(), c, repo, &core.Config{
		UpdatePostEndpoint: "https://events.resignal.app/onTriggerSuccess",
	})
	instr := &ds_core.Instrument{
		Scrip: "SBIN", KiteToken: "", ExchangeToken: "3045", Exchange: "NSE", InstrumentType: ds_core.Stock,
	}
	trigger := &core.Trigger{Id: "1", TAttrib: core.LTP, Operator: core.GTE, TNearPrice: 290, TPrice: 300.0, Instument: instr}
	err := ts.CreateTrigger(trigger)
	if err != nil {
		t.Errorf("TestCreateTrigger failed, error returned %v\n", err)
	}

	deleteErr := ts.DeleteTrigger("1")
	if deleteErr != nil {
		t.Errorf("TestDeleteTrigger failed, error returned %v\n", err)
	}
}

func TestDeleteTriggerWithoutAdding(t *testing.T) {
	c := ltp_client.NewMockLTPClient()
	repo := &core.MockRepo{}
	ts := NewTriggerSystem(context.Background(), c, repo, &core.Config{
		UpdatePostEndpoint: "https://events.resignal.app/onTriggerSuccess",
	})
	err := ts.DeleteTrigger("1")
	if err == nil {
		t.Errorf("TestDeleteTriggerWithoutAdding failed, expected error\n")
	} else if err.Error() != "trigger not found - 1" {
		t.Errorf("TestDeleteTriggerWithoutAdding failed, expected %v, got %v\n", "trigger not found - 1", err.Error())
	}
}

func TestGetTrigger(t *testing.T) {
	c := ltp_client.NewMockLTPClient()
	repo := &core.MockRepo{}
	ts := NewTriggerSystem(context.Background(), c, repo, &core.Config{
		UpdatePostEndpoint: "https://events.resignal.app/onTriggerSuccess",
	})
	instr := &ds_core.Instrument{
		Scrip: "SBIN", KiteToken: "", ExchangeToken: "3045", Exchange: "NSE", InstrumentType: ds_core.Stock,
	}
	trigger := &core.Trigger{Id: "1", TAttrib: core.LTP, Operator: core.GTE, TNearPrice: 290, TPrice: 300.0, Instument: instr}
	err := ts.CreateTrigger(trigger)
	if err != nil {
		t.Errorf("TestGetTrigger failed, error returned %v\n", err)
	}

	tr, err := ts.GetTrigger("1")
	if err != nil {
		t.Errorf("TestGetTrigger failed, error returned %v\n", err)
	}

	got, err := tr.ToString()
	if err != nil {
		t.Errorf("TestGetTrigger failed, error returned %v\n", err)
	}

	expected, err := trigger.ToString()
	if err != nil {
		t.Errorf("TestGetTrigger failed, error returned %v\n", err)
	}

	if got != expected {
		t.Errorf("TestGetTrigger failed, wanted = %s; got = %s", got, expected)
	}
}

func TestGetTriggerStatus(t *testing.T) {
	c := ltp_client.NewMockLTPClient()
	repo := &core.MockRepo{}
	ts := NewTriggerSystem(context.Background(), c, repo, &core.Config{
		UpdatePostEndpoint: "https://events.resignal.app/onTriggerSuccess",
	})
	instr := &ds_core.Instrument{
		Scrip: "SBIN", KiteToken: "", ExchangeToken: "3045", Exchange: "NSE", InstrumentType: ds_core.Stock,
	}
	trigger := &core.Trigger{Id: "1", TAttrib: core.LTP, Operator: core.GTE, TNearPrice: 290, TPrice: 300.0, Instument: instr}
	err := ts.CreateTrigger(trigger)
	if err != nil {
		t.Errorf("TestGetTriggerStatus failed, error returned %v\n", err)
	}

	st, err := ts.GetTriggerStatus("1")
	if err != nil {
		t.Errorf("TestGetTriggerStatus failed, error returned %v\n", err)
	}

	repr, err := trigger.ToString()
	if err != nil {
		t.Errorf("TestGetTriggerStatus failed, error returned %v\n", err)
	}

	if st != repr {
		t.Errorf("TestGetTriggerStatus failed, wanted = %s; got = %s", "", st)
	}
}
