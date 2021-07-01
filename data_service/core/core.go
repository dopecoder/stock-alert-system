package core

import (
	"github.com/revulcan/stock-alert-system/data_service/grpc/data_service"
)

type InstrumentType int

type Stream interface {
	Send(*data_service.LTP) error
}

const (
	Stock InstrumentType = iota
	Index
)

type LTP struct {
	Ltp            float64
	Open           float64
	Close          float64
	Low            float64
	High           float64
	ExchangeToken  string
	Scrip          string
	InstrumentType InstrumentType
}

type Instrument struct {
	Scrip          string
	KiteToken      string
	ExchangeToken  string
	Exchange       string
	InstrumentType InstrumentType
}

type LtpService interface {
	// GetLTPForInstrument(instrument Instrument) LTP
	// GetLTPforInstruments(instruments []Instrument) []LTP
	WatchLTPforInstr(instrument *Instrument, stream Stream) error
	WatchLTPforInstrs(instruments []*Instrument, stream Stream) error
}

type LtpClient interface {
	Init() error
	Subscribe(instruments []*Instrument, stream chan *LTP) (int, error)
	UnSubscribe(id int) error
}
