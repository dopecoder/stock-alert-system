package core

import (
	"errors"
	"fmt"

	"github.com/revulcan/stock-alert-system/data_service/grpc/data_service"
)

type MockStream struct {
	Strm   chan *LTP
	Closed bool
}

func NewMockStream() *MockStream {
	return &MockStream{
		Strm:   make(chan *LTP, 5),
		Closed: false,
	}
}

func (m *MockStream) Send(ltp *data_service.LTP) error {
	if !m.Closed {
		m.Strm <- &LTP{
			Ltp:            ltp.Ltp,
			Open:           ltp.Open,
			Close:          ltp.Close,
			Low:            ltp.Low,
			High:           ltp.High,
			ExchangeToken:  ltp.ExchangeToken,
			Scrip:          ltp.Scrip,
			InstrumentType: InstrumentType(ltp.InstrumentType),
		}
		fmt.Println("(MockStream) Send: Wrote to channel")
		return nil
	}
	return errors.New("Stream Closed")
}

func (m *MockStream) Close() {
	if !m.Closed {
		m.Closed = true
	}
}
