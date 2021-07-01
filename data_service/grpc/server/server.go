package server

import (
	"github.com/revulcan/stock-alert-system/data_service/core"
	"github.com/revulcan/stock-alert-system/data_service/grpc/data_service"
	"github.com/revulcan/stock-alert-system/data_service/ltp_client"
	"github.com/revulcan/stock-alert-system/data_service/service"
)

type LtpStreamingServer struct {
	data_service.UnimplementedLTPServiceServer
	ltpService *service.Service
}

func NewLTPStreamingServer(apiKey string, usename string, password string) *LtpStreamingServer {
	service := service.New(
		// ltp_client.NewMockLTPClient(),
		ltp_client.NewAngelLTPClient(apiKey, usename, password),
	)
	service.Init()
	return &LtpStreamingServer{
		ltpService: service,
	}
}

func (s *LtpStreamingServer) WatchLTPforInstruments(instruments *data_service.Instruments, strm data_service.LTPService_WatchLTPforInstrumentsServer) error {

	//fmt.Println("(Server) WatchLTPforInstruments: Started")
	var instrs []*core.Instrument = make([]*core.Instrument, 0)
	for _, i := range instruments.Items {
		instrs = append(instrs, &core.Instrument{
			Scrip:          i.Scrip,
			ExchangeToken:  i.ExchangeToken,
			Exchange:       i.Exchange,
			InstrumentType: core.InstrumentType(i.InstrumentType),
			KiteToken:      i.KiteToken,
		})
	}

	// ch := core.NewMockStream()
	s.ltpService.WatchLTPforInstrs(instrs, strm)
	// if err != nil {
	// 	return err
	// }

	// for {
	// 	val, more := <-ch.Strm

	// 	if !more {
	// 		break
	// 	}

	// 	err := strm.Send(&data_service.LTP{
	// 		Ltp:            val.Ltp,
	// 		Open:           val.Open,
	// 		Close:          val.Close,
	// 		Low:            val.Low,
	// 		High:           val.High,
	// 		ExchangeToken:  val.ExchangeToken,
	// 		Scrip:          val.Scrip,
	// 		InstrumentType: data_service.InstrumentType(val.InstrumentType),
	// 	})

	// 	if err != nil {
	// 		fmt.Printf("(Server) WatchLTPforInstruments: Finished, %v\n", err)
	// 		return err
	// 	}
	// }
	//fmt.Println("(Server) WatchLTPforInstruments: Finished")
	return nil
}
