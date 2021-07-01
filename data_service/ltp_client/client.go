package ltp_client

import (
	"context"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"sync"
	"time"

	SmartApi "github.com/angelbroking-github/smartapigo"
	"github.com/revulcan/stock-alert-system/data_service/angel"
	"github.com/revulcan/stock-alert-system/data_service/core"
	log "github.com/sirupsen/logrus"
)

type InstrumentSubscription map[int]chan *core.LTP

type AngelLTPClient struct {
	mu                    *sync.Mutex
	apiKey                string
	usename               string
	password              string
	socketClient          *angel.SocketClient
	subscriptions         map[int][]*core.Instrument
	subscribedChannels    map[string]InstrumentSubscription
	subscribedInstruments map[string]*core.Instrument
	count                 int
	ctx                   context.Context
	socketCtx             context.Context
	cancel                context.CancelFunc
	socketCancel          context.CancelFunc
	connected             bool
	dataPointCount        map[string]int64
}

func NewAngelLTPClient(apiKey string, usename string, password string) *AngelLTPClient {
	client := &AngelLTPClient{
		apiKey:                apiKey,
		usename:               usename,
		password:              password,
		subscriptions:         make(map[int][]*core.Instrument),
		subscribedChannels:    make(map[string]InstrumentSubscription),
		subscribedInstruments: make(map[string]*core.Instrument),
		count:                 0,
		mu:                    &sync.Mutex{},
		connected:             false,
		dataPointCount:        make(map[string]int64),
	}
	return client
}

func (a *AngelLTPClient) Init() error {
	return nil
}

func (a *AngelLTPClient) subscribe() error {
	if a.socketClient != nil && a.socketClient.Conn != nil {
		// log.Println("Trying to close existing connection")
		err := a.socketClient.Close()
		if err != nil {
			return err
		}
		if a.socketCancel != nil {
			a.socketCancel()
		}
	}

	if a.cancel != nil {
		//log.Println("Trying to close existing context")
		a.cancel()
	}
	ctx, cancel := context.WithCancel(context.Background())
	a.ctx = ctx
	a.cancel = cancel

	angelClient := SmartApi.New(a.usename, a.password, a.apiKey)
	session, err := angelClient.GenerateSession()
	log.Println(session)

	if err != nil {
		log.Println(err.Error())
		return nil
	}
	//Get User Profile
	session.UserProfile, err = angelClient.GetUserProfile()

	if err != nil {
		log.Println(err.Error())
		return err
	}

	instrumentQueryString, err := a.buildInstrumentQueryString()
	log.Printf("Instrument Query String : %v", instrumentQueryString)

	if err != nil {
		return err
	}

	socketCtx, socketCancel := context.WithCancel(a.ctx)
	a.socketCancel = socketCancel
	a.socketCtx = socketCtx

	a.socketClient = angel.New(session.ClientCode, session.FeedToken, instrumentQueryString)

	// Assign callbacks
	a.socketClient.OnError(onErrorAngel(a, socketCtx))
	a.socketClient.OnClose(onCloseAngel(a, socketCtx))
	a.socketClient.OnMessage(onMessageAngel(a, socketCtx))
	a.socketClient.OnConnect(onConnectAngel(a, socketCtx))
	a.socketClient.OnReconnect(onReconnectAngel(a, socketCtx))
	a.socketClient.OnNoReconnect(onNoReconnectAngel(a, socketCtx))

	// Start Consuming Data
	go a.socketClient.Serve(socketCtx)
	return nil
}

// func (a *AngelLTPClient) subscribe() error {
// 	instrumentQueryString, err := a.buildInstrumentQueryString()
// 	//log.Println(instrumentQueryString)

// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (a *AngelLTPClient) subscribe() error {
// 	if a != nil && a.connected && a.socketClient != nil {
// 		instrumentQueryString, err := a.buildInstrumentQueryString()
// 		if err != nil {
// 			return err
// 		}
// 		a.socketClient.Scrips = ""
// 		err = a.socketClient.Subscribe()
// 		if err != nil {
// 			log.Fatalf("(AngelLtpClient) subscribe err : %v\n", err)
// 		}
// 		a.socketClient.Scrips = instrumentQueryString
// 		log.Debugln("Scrips: ", a.socketClient.Scrips)
// 		err = a.socketClient.Subscribe()
// 		if err != nil {
// 			log.Fatalf("(AngelLtpClient) subscribe err : %v\n", err)
// 		}
// 	}
// 	return nil
// }

func (a *AngelLTPClient) Subscribe(instruments []*core.Instrument, stream chan *core.LTP) (int, error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	newCount := a.incrementAndGetCount()
	subscriptionChanged := false
	for _, instrument := range instruments {

		_, pointCountExists := a.dataPointCount[instrument.ExchangeToken]
		if !pointCountExists {
			a.dataPointCount[instrument.ExchangeToken] = 0
		}

		_, exists := a.subscribedChannels[instrument.ExchangeToken]
		if !exists {
			a.subscribedChannels[instrument.ExchangeToken] = make(InstrumentSubscription)
			a.subscribedInstruments[instrument.ExchangeToken] = instrument
			subscriptionChanged = true
		}
		a.subscribedChannels[instrument.ExchangeToken][newCount] = stream
	}
	a.subscriptions[newCount] = instruments

	if subscriptionChanged {
		err := a.subscribe()
		if err != nil {
			return newCount, err
		}
	}
	return newCount, nil
}

func (a *AngelLTPClient) UnSubscribe(id int) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	instrs, exists := a.subscriptions[id]
	log.Infof("UnSubscribe : instrs -> %v", instrs)

	subscriptionChanged := false
	if !exists {
		return fmt.Errorf("SubscriptionHandle with %d does not exist", id)
	}
	for _, instrument := range instrs {
		_, exists := a.subscribedChannels[instrument.ExchangeToken]
		if exists {
			log.Infof("UnSubscribe : with %v\n", a.subscribedChannels[instrument.ExchangeToken])
			delete(a.subscribedChannels[instrument.ExchangeToken], id)
		}

		if len(a.subscribedChannels[instrument.ExchangeToken]) == 0 {
			delete(a.subscribedChannels, instrument.ExchangeToken)
			delete(a.subscribedInstruments, instrument.ExchangeToken)
			subscriptionChanged = true
		}
	}
	delete(a.subscriptions, id)

	if subscriptionChanged {
		err := a.subscribe()
		return err
	}
	return nil
}

func (a *AngelLTPClient) incrementAndGetCount() int {
	a.count = a.count + 1
	return a.count
}

// Triggered when any error is raised
func onErrorAngel(a *AngelLTPClient, c context.Context) func(error) {
	return func(err error) {
		if a == nil {
			return
		}
		select {
		case <-c.Done():
			return
		default:
			log.Println("Error: ", err)
		}
	}

}

// Triggered when websocket connection is closed
func onCloseAngel(a *AngelLTPClient, c context.Context) func(code int, reason string) {
	return func(code int, reason string) {
		if a == nil {
			return
		}
		select {
		case <-c.Done():
			return
		default:
			log.Debugf("Closed Websocket")
			log.Println("Close: ", code, reason)
			a.mu.Lock()
			a.connected = false
			a.mu.Unlock()
		}
	}
}

// Triggered when connection is established and ready to send and accept data
func onConnectAngel(a *AngelLTPClient, c context.Context) func() {
	return func() {
		select {
		case <-c.Done():
			return
		default:
			log.Debugf("Connected to Websocket")
			a.mu.Lock()
			a.connected = true
			a.mu.Unlock()
			err := a.subscribe()
			if err != nil {
				log.Errorln(err)
			}
		}
	}
}

// Triggered when a message is received
func onMessageAngel(a *AngelLTPClient, c context.Context) func(message []map[string]interface{}) {
	return func(message []map[string]interface{}) {
		if a == nil {
			return
		}
		select {
		case <-c.Done():
			return
		default:
			a.broadcastLTP(message)
		}
	}
}

// Triggered when reconnection is attempted which is enabled by default
func onReconnectAngel(a *AngelLTPClient, c context.Context) func(attempt int, delay time.Duration) {
	return func(attempt int, delay time.Duration) {
		if a == nil {
			return
		}
		select {
		case <-c.Done():
			return
		default:
			log.Debugf("Reconnected to Websocket")
			a.mu.Lock()
			a.connected = true
			log.Debugf("Trying to subscribe to instruments")
			err := a.subscribe()
			if err != nil {
				log.Fatalf("(AngelLtpClient) subscribe err : %v\n", err)
			}
			a.mu.Unlock()
			log.Printf("Reconnect attempt %d in %fs\n", attempt, delay.Seconds())
		}
	}
}

// Triggered when maximum number of reconnect attempt is made and the program is terminated
func onNoReconnectAngel(a *AngelLTPClient, c context.Context) func(int) {
	return func(attempt int) {
		if a == nil {
			return
		}
		select {
		case <-c.Done():
			return
		default:
			log.Fatalf("Maximum no of reconnect attempt reached: %d\n", attempt)
		}
	}
}

func (a *AngelLTPClient) buildInstrumentQueryString() (string, error) {
	token := ""
	exchangeMapping := map[string]string{
		"NSE": "nse_cm",
		"BSE": "bse_cm",
		"NFO": "nse_fo",
		"MCX": "mcx_fo",
		"CDS": "cde_fo",
	}

	for exchToken, instr := range a.subscribedInstruments {
		fmtString := "&%s|%s"
		if len(token) == 0 {
			fmtString = "%s|%s"
		}
		token += fmt.Sprintf(fmtString, exchangeMapping[instr.Exchange], exchToken)
	}
	return token, nil
}

// Triggered when a message is received
func (a *AngelLTPClient) broadcastLTP(message []map[string]interface{}) {
	for _, m := range message {
		tk, exchTkExists := m["tk"]
		low := m["lo"]
		high := m["h"]
		open := m["op"]
		close := m["c"]
		ltp := m["ltp"]
		token := fmt.Sprintf("%v", tk)

		a.mu.Lock()
		a.dataPointCount[token] += 1
		log.Debugf("Received Ticker %d for  %s -> %v\n", a.dataPointCount[token], token, m)
		if exchTkExists {
			exchToken := fmt.Sprintf("%v", token)

			subs, exists := a.subscribedChannels[exchToken]

			if exists {
				for _, channel := range subs {
					a.mu.Unlock()
					log.Debugf("Sending Ticker %d for  %s -> %v\n", a.dataPointCount[token], token, m)
					channel <- &core.LTP{
						Ltp:            getFloat(ltp),
						Open:           getFloat(open),
						Close:          getFloat(close),
						Low:            getFloat(low),
						High:           getFloat(high),
						ExchangeToken:  exchToken,
						InstrumentType: core.Stock,
					}
					log.Debugf("Sent Ticker %d for  %s -> %v\n", a.dataPointCount[token], token, m)
					a.mu.Lock()
				}
			}
		}
		a.mu.Unlock()
		// else {
		// 	log.Printf("Received an invalid ticker for %s -> %v\n", m["tk"], m)
		// }
	}
}

var floatType = reflect.TypeOf(float64(0.0))
var stringType = reflect.TypeOf("")

func getFloat(unk interface{}) float64 {
	switch i := unk.(type) {
	case float64:
		return i
	case float32:
		return float64(i)
	case int64:
		return float64(i)
	case int32:
		return float64(i)
	case int:
		return float64(i)
	case uint64:
		return float64(i)
	case uint32:
		return float64(i)
	case uint:
		return float64(i)
	case string:
		num, e := strconv.ParseFloat(i, 64)
		if e != nil {
			return 0.0
		}
		return num
	default:
		return math.NaN()
		// v := reflect.ValueOf(unk)
		// v = reflect.Indirect(v)
		// if v.Type().ConvertibleTo(floatType) {
		// 	fv := v.Convert(floatType)
		// 	return fv.Float()
		// } else if v.Type().ConvertibleTo(stringType) {
		// 	sv := v.Convert(stringType)
		// 	s := sv.String()
		// 	num, e := strconv.ParseFloat(s, 64)
		// 	if e != nil {
		// 		return 0.0
		// 	}
		// 	return num
		// } else {
		// 	return math.NaN()
		// }
	}
}
