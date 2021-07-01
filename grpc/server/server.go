package server

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/revulcan/stock-alert-system/core"
	data_core "github.com/revulcan/stock-alert-system/data_service/core"
	"github.com/revulcan/stock-alert-system/data_service/ltp_client"
	"github.com/revulcan/stock-alert-system/grpc/trigger_service"
	"github.com/revulcan/stock-alert-system/service"
	log "github.com/sirupsen/logrus"
)

var NotStartedErr error = fmt.Errorf("service not started")

type TriggerServiceServer struct {
	trigger_service.UnimplementedTriggerServiceServer
	ts      service.TriggerService
	started bool
	stopped bool
	mu      *sync.Mutex
}

func NewTriggerSystemServer() *TriggerServiceServer {
	log.Info("NewTriggerSystem Created, returning...")
	return &TriggerServiceServer{
		started: false,
		stopped: true,
		mu:      &sync.Mutex{},
	}
}

func NewMockTriggerSystemServer() *TriggerServiceServer {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := &core.Config{
		UpdatePostEndpoint: os.Getenv("UPDATES_POST_ENDPOINT"),
		KiteToken:          "",
		KiteApiKey:         "",
		AngelApiKey:        "",
		AngelClientId:      "",
		AngelPassword:      "",
	}

	ts := service.NewTriggerSystem(context.Background(), ltp_client.NewMockLTPClient(), &core.MockRepo{}, config)
	ts.SetupService()
	return &TriggerServiceServer{
		ts:      ts,
		started: false,
		stopped: true,
	}
}

func (s *TriggerServiceServer) CreateTrigger(c context.Context, req *trigger_service.CreateTriggerReq) (*trigger_service.CreateTriggerRes, error) {
	log.Info("Entered CreateTrigger service")
	err := s.checkIfStarted()
	if err != nil {
		return nil, err
	}

	if len(strings.Trim(req.Id, " ")) == 0 || s.idAlreadyExists(req.Id) {
		return nil, fmt.Errorf("invalid trigger id - %s", req.Id)
	}

	log.Info("Calling CreateTrigger fn")
	err = s.ts.CreateTrigger(
		&core.Trigger{
			Id:         req.Id,
			TAttrib:    int(req.TAttrib),
			Operator:   int(req.Operator),
			TPrice:     req.TPrice,
			TNearPrice: req.TNearPrice,
			DataPoint:  0,
			Instument: &data_core.Instrument{
				Scrip:          req.Scrip,
				KiteToken:      req.KiteToken,
				ExchangeToken:  req.ExchangeToken,
				Exchange:       req.Exchange.String(),
				InstrumentType: data_core.Stock,
			},
		})
	if err != nil {
		log.Error("CreateTrigger : ", err)
		return nil, err
	}
	log.Info("Completed CreateTrigger fn, returning...")
	return &trigger_service.CreateTriggerRes{Ok: true}, nil
}

func (s *TriggerServiceServer) DeleteTrigger(c context.Context, req *trigger_service.DeleteTriggerReq) (*trigger_service.DeleteTriggerRes, error) {
	log.Info("Entered DeleteTrigger service")
	err := s.checkIfStarted()
	if err != nil {
		return nil, err
	}

	if len(strings.Trim(req.TriggerId, " ")) == 0 || !s.idAlreadyExists(req.TriggerId) {
		return nil, fmt.Errorf("invalid trigger id - %s", req.TriggerId)
	}

	log.Info("Calling DeleteTrigger fn")
	err = s.ts.DeleteTrigger(req.TriggerId)
	if err != nil {
		log.Error("DeleteTrigger : ", err)
		return nil, err
	}
	log.Info("Completed DeleteTrigger fn, returning...")
	return &trigger_service.DeleteTriggerRes{Ok: true}, nil
}

func (s *TriggerServiceServer) StartService(c context.Context, req *trigger_service.StartServiceReq) (*trigger_service.StartServiceRes, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	log.Info("Entered StartService")
	if !s.started && (req.Override || (IsTradingTime(time.Now()) || req.Mock)) {

		log.Info("Calling StartService fn")

		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		config := &core.Config{
			UpdatePostEndpoint: os.Getenv("UPDATES_POST_ENDPOINT"),
			KiteToken:          os.Getenv("KITE_APIKEY"),
			KiteApiKey:         os.Getenv("KITE_TOKEN"),
			AngelApiKey:        os.Getenv("ANGEL_APIKEY"),
			AngelClientId:      os.Getenv("ANGEL_CLIENTID"),
			AngelPassword:      os.Getenv("ANGEL_PASSWORD"),
		}
		var client data_core.LtpClient
		if req.Mock {
			client = ltp_client.NewMockLTPClient()
		} else {
			client = ltp_client.NewAngelLTPClient(config.AngelApiKey, config.AngelClientId, config.AngelPassword)

		}
		clientInitErr := client.Init()
		if clientInitErr != nil {
			log.Fatalf("clientInitErr: %v", clientInitErr)
			return &trigger_service.StartServiceRes{
				Ok: false,
			}, clientInitErr
		}

		log.Info("Completed StartService fn, returning...")
		s.ts = service.NewTriggerSystem(context.Background(), client, &core.MockRepo{}, config)
		s.started = true
		s.stopped = false
		return &trigger_service.StartServiceRes{
			Ok: true,
		}, nil
	} else if s.started {
		return &trigger_service.StartServiceRes{
			Ok: false,
		}, errors.New("service already started")
	}
	return &trigger_service.StartServiceRes{
		Ok: false,
	}, errors.New("override to start service")
}

func (s *TriggerServiceServer) StopService(c context.Context, req *trigger_service.StopServiceReq) (*trigger_service.StopServiceRes, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	log.Info("Entered StopService")
	if s.started && (req.Override || (!IsTradingTime(time.Now()) || req.Mock)) {

		log.Info("Calling StopService fn")
		s.ts.Stop()
		log.Info("Completed StopService fn, returning...")
		s.started = false
		s.stopped = true
		return &trigger_service.StopServiceRes{
			Ok: true,
		}, nil
	} else if s.stopped {
		return &trigger_service.StopServiceRes{
			Ok: false,
		}, errors.New("service already stopped")
	}
	return &trigger_service.StopServiceRes{
		Ok: false,
	}, errors.New("override to stop service")
}

func (s *TriggerServiceServer) GetTrigger(c context.Context, req *trigger_service.GetTriggerReq) (*trigger_service.GetTriggerRes, error) {
	log.Info("Entered GetTrigger service")
	err := s.checkIfStarted()
	if err != nil {
		return nil, err
	}

	log.Info("Calling GetTrigger fn")
	t, err := s.ts.GetTrigger(req.TriggerId)
	if err != nil {
		log.Error("GetTrigger : ", err)
		return nil, err
	}
	log.Info("Completed GetTrigger fn, returning...")
	exchng := 0
	if t.Instument.Exchange == "BSE" {
		exchng = 1
	}

	return &trigger_service.GetTriggerRes{
		Id:            t.Id,
		TAttrib:       trigger_service.TriggerAttrib(t.TAttrib),
		Operator:      trigger_service.TriggerOperator(t.Operator),
		TPrice:        t.TPrice,
		TNearPrice:    t.TNearPrice,
		Scrip:         t.Instument.Scrip,
		KiteToken:     t.Instument.KiteToken,
		ExchangeToken: t.Instument.ExchangeToken,
		Exchange:      trigger_service.Exchange(exchng),
	}, nil
}

func (s *TriggerServiceServer) GetTriggerStatus(c context.Context, req *trigger_service.TriggerStatusReq) (*trigger_service.TriggerStatusRes, error) {
	log.Info("Entered GetTriggerStatus service")
	err := s.checkIfStarted()
	if err != nil {
		return nil, err
	}

	log.Info("Calling GetTriggerStatus fn")
	st, err := s.ts.GetTriggerStatus(req.TriggerId)
	if err != nil {
		log.Error("GetTriggerStatus : ", err)
		return &trigger_service.TriggerStatusRes{Status: ""}, err
	}
	log.Info("Completed GetTriggerStatus fn, returning...")
	return &trigger_service.TriggerStatusRes{Status: st}, nil
}

func (s *TriggerServiceServer) idAlreadyExists(id string) bool {
	trig, _ := s.ts.GetTrigger(id)
	log.Infof("ID %s already exists? - %t, returning...", id, trig != nil)
	return trig != nil
}

func IsTradingTime(t time.Time) bool {
	location, err := time.LoadLocation("Asia/Kolkata")
	startTime := time.Date(t.Year(), t.Month(), t.Day(), 9, 0, 0, 0, location)
	endTime := time.Date(t.Year(), t.Month(), t.Day(), 15, 30, 0, 0, location)

	if err != nil {
		fmt.Println(err)
	}

	istTime := t.In(location)
	fmt.Println("ZONE : ", location, " Time : ", istTime) // IST
	isTradingTime := istTime.After(startTime) && istTime.Before(endTime)
	notSaturday := istTime.Weekday() != time.Saturday
	notSunday := istTime.Weekday() != time.Sunday
	if isTradingTime && notSaturday && notSunday {
		return true
	}
	return false
}

func (s *TriggerServiceServer) checkIfStarted() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.started {
		return NotStartedErr
	}
	return nil
}
