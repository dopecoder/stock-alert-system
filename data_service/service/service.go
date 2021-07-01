package service

import (
	"fmt"
	"sync"

	"github.com/revulcan/stock-alert-system/data_service/core"
	"github.com/revulcan/stock-alert-system/data_service/grpc/data_service"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	mu                 *sync.Mutex
	client             core.LtpClient
	instrumentsWatched map[string]map[int]chan *core.LTP
	ltpRecvChannel     chan *core.LTP
	prevTickers        map[string]*core.LTP
	done               chan bool
	listening          bool
	finished           bool
	recvPointCount     map[string]int64
	sendPointCount     map[string]int64
}

func New(c core.LtpClient) *Service {
	return &Service{
		client:             c,
		instrumentsWatched: make(map[string]map[int]chan *core.LTP, 5),
		ltpRecvChannel:     make(chan *core.LTP, 5),
		done:               make(chan bool, 5),
		listening:          false,
		finished:           false,
		mu:                 &sync.Mutex{},
		prevTickers:        make(map[string]*core.LTP),
		recvPointCount:     make(map[string]int64),
		sendPointCount:     make(map[string]int64),
	}
}

func (s *Service) Init() {
	// start listner
	go s.startListner()
}

// func (s *Service) WatchLTPforInstruments(instruments *data_service.Instruments, stream data_service.LTPService_WatchLTPforInstrumentsServer) error {

// 	watcherChannel := make(chan core.LTP)

// 	s.subscribe(instruments, watcherChannel)

// 	for {
// 		ticker, more := <-watcherChannel
// 		if more {
// 			//log.Infoln("received ticker", ticker)
// 			if err := stream.Send(ticker); err != nil {
// 				s.unSubscribe()
// 				return err
// 			}
// 		} else {
// 			//log.Infoln("received all jobs")
// 			break
// 		}
// 	}

// 	s.unSubscribe()

// 	return nil
// }

// func (s *Service) WatchLTPforInstr(instrument core.Instrument, stream core.Stream) error {
// 	return s.WatchLTPforInstrs([]core.Instrument{instrument}, stream)
// }

func (s *Service) WatchLTPforInstrs(instruments []*core.Instrument, stream core.Stream) error {
	watcherChannel := make(chan *core.LTP, 5)

	handle, err := s.subscribe(instruments, watcherChannel)
	if err != nil {
		close(watcherChannel)
		return err
	}

	for {
		ticker, more := <-watcherChannel
		if more {
			//log.Infoln("received ticker", ticker)
			// prevTicker, tickExists := s.prevTickers[ticker.ExchangeToken]
			// if tickExists && prevTicker.Ltp == ticker.Ltp {
			// 	continue
			// } else {
			// 	s.prevTickers[ticker.ExchangeToken] = ticker
			// }

			s.mu.Lock()
			s.sendPointCount[ticker.ExchangeToken] += 1
			cnt := s.sendPointCount[ticker.ExchangeToken]
			s.mu.Unlock()
			log.Debugf("Received-c Ticker %d for  %s -> %v\n", cnt, ticker.ExchangeToken, ticker)

			if err := stream.Send(&data_service.LTP{
				Ltp:            ticker.Ltp,
				Open:           ticker.Open,
				Close:          ticker.Close,
				Low:            ticker.Low,
				High:           ticker.High,
				ExchangeToken:  ticker.ExchangeToken,
				Scrip:          ticker.Scrip,
				InstrumentType: data_service.InstrumentType(ticker.InstrumentType),
			}); err != nil {
				s.unSubscribe(handle)
				close(watcherChannel)
				return err
			}
			log.Debugf("Sent-c Ticker %d for  %s -> %v\n", cnt, ticker.ExchangeToken, ticker)

		} else {
			//log.Infoln("received all jobs")
			break
		}
	}

	s.unSubscribe(handle)
	close(watcherChannel)
	return nil
}

func (s *Service) subscribe(instruments []*core.Instrument, c chan *core.LTP) (int, error) {

	sd, err := s.client.Subscribe(instruments, s.ltpRecvChannel)
	if err != nil {
		return -1, err
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	for _, i := range instruments {
		_, exists := s.instrumentsWatched[i.ExchangeToken]
		if !exists {
			s.instrumentsWatched[i.ExchangeToken] = make(map[int]chan *core.LTP)
		}
		s.instrumentsWatched[i.ExchangeToken][sd] = c
	}
	fmt.Printf("(Service) subscribe: subscribed to %d\n", sd)
	return sd, nil
}

func (s *Service) unSubscribe(handle int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	fmt.Printf("(Service) unSubscribe: unsubscribe from %d\n", handle)
	if handle != -1 {
		err := s.client.UnSubscribe(handle)
		for k := range s.instrumentsWatched {
			_, exists := s.instrumentsWatched[k][handle]
			if exists {
				fmt.Printf("(Service) unSubscribe: deleting handle in %v\n", s.instrumentsWatched[k])
				delete(s.instrumentsWatched[k], handle)
			}
		}
		for k := range s.instrumentsWatched {
			if len(s.instrumentsWatched[k]) == 0 {
				delete(s.instrumentsWatched, k)
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) startListner() {
	//log.Infoln("(Service) startListner: Started listner")

	s.mu.Lock()
	s.listening = true
	s.mu.Unlock()

	for {
		if s.finished {
			break
		}

		select {
		case <-s.done:
			//log.Infoln("(Service) startListner: started <-s.done")

			s.mu.Lock()
			s.finished = true
			s.mu.Unlock()

			//log.Infoln("(Service) startListner: finished <-s.done")

		case ltp, more := <-s.ltpRecvChannel:
			if !more {
				//log.Infoln("(Service) startListner: closing ltp := <-s.ltpRecvChannel")
				// todo: need to restart the client if this happens, or it means someone closed the channel
				s.mu.Lock()
				s.finished = true
				s.mu.Unlock()
			} else {
				s.mu.Lock()
				s.recvPointCount[ltp.ExchangeToken] += 1
				log.Debugf("Received-b Ticker %d for  %s -> %v\n", s.recvPointCount[ltp.ExchangeToken], ltp.ExchangeToken, ltp)

				instrChannelMap := s.instrumentsWatched[ltp.ExchangeToken]
				for _, c := range instrChannelMap {
					c <- ltp
				}
				log.Debugf("Sent-b Ticker %d for  %s -> %v\n", s.recvPointCount[ltp.ExchangeToken], ltp.ExchangeToken, ltp)
				s.mu.Unlock()
			}
		}
	}

	s.mu.Lock()
	s.listening = false
	s.mu.Unlock()
}

func (s *Service) Close() {
	//log.Infoln("(Service) Close: Closing listner")
	s.mu.Lock()
	s.done <- true
	s.mu.Unlock()
	// todo : fix the sender to exit if channel is closed and then close this
	// close(s.ltpRecvChannel)
	//log.Infoln("(Service) Close: Closed listner")
}
