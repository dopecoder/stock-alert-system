package ltp_client

import (
	"fmt"
	"sync"
	"time"

	"github.com/revulcan/stock-alert-system/data_service/core"
)

type MockLTPClient struct {
	mu                    *sync.RWMutex
	subscriptions         map[int][]*core.Instrument
	subscribedChannels    map[string]InstrumentSubscription
	subscribedInstruments map[string]*core.Instrument
	count                 int
}

func NewMockLTPClient() *MockLTPClient {

	c := &MockLTPClient{
		subscriptions:         make(map[int][]*core.Instrument),
		subscribedChannels:    make(map[string]InstrumentSubscription),
		subscribedInstruments: make(map[string]*core.Instrument),
		count:                 0,
		mu:                    &sync.RWMutex{},
	}
	go c.startBroadcast()
	return c
}

func NewMockLTPClientWithControl(start float64, end float64, step float64, sleepDuration time.Duration) *MockLTPClient {

	c := &MockLTPClient{
		subscriptions:         make(map[int][]*core.Instrument),
		subscribedChannels:    make(map[string]InstrumentSubscription),
		subscribedInstruments: make(map[string]*core.Instrument),
		count:                 0,
		mu:                    &sync.RWMutex{},
	}
	go c.startBroadcastWithControl(start, end, step, sleepDuration)
	return c
}

func (a *MockLTPClient) Init() error {
	return nil
}

func (a *MockLTPClient) Subscribe(instruments []*core.Instrument, stream chan *core.LTP) (int, error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	fmt.Println(instruments)
	newCount := a.incrementAndGetCount()
	// subscriptionChanged := false
	for _, instrument := range instruments {

		_, exists := a.subscribedChannels[instrument.ExchangeToken]
		if !exists {
			a.subscribedChannels[instrument.ExchangeToken] = make(InstrumentSubscription)
			a.subscribedInstruments[instrument.ExchangeToken] = instrument
			// subscriptionChanged = true
		}
		a.subscribedChannels[instrument.ExchangeToken][newCount] = stream
	}
	a.subscriptions[newCount] = instruments
	return newCount, nil
}

func (a *MockLTPClient) UnSubscribe(id int) error {
	a.mu.Lock()
	instrs, exists := a.subscriptions[id]
	fmt.Printf("(MockLTPClient) UnSubscribe : instrs -> %v", instrs)
	// subscriptionChanged := false
	if !exists {
		return fmt.Errorf("SubscriptionHandle with %d does not exist", id)
	}
	for _, instrument := range instrs {
		_, exists := a.subscribedChannels[instrument.ExchangeToken]
		if exists {
			fmt.Printf("(MockLTPClient) UnSubscribe : with %v\n", a.subscribedChannels[instrument.ExchangeToken])
			delete(a.subscribedChannels[instrument.ExchangeToken], id)
		}

		if len(a.subscribedChannels[instrument.ExchangeToken]) == 0 {
			delete(a.subscribedChannels, instrument.ExchangeToken)
			delete(a.subscribedInstruments, instrument.ExchangeToken)
			fmt.Printf("(MockLTPClient) UnSubscribe : deleting subscribed instruments %v\n", a.subscribedInstruments)
			// subscriptionChanged = true
		}
	}
	delete(a.subscriptions, id)
	a.mu.Unlock()
	return nil
}

func (a *MockLTPClient) startBroadcast() {

	for {
		arr := make([]map[string]interface{}, 0)
		a.mu.RLock()
		instrMap := a.subscribedInstruments
		for _, instrument := range instrMap {
			msg := make(map[string]interface{})
			msg["tk"] = instrument.ExchangeToken
			arr = append(arr, msg)
		}
		a.mu.RUnlock()

		if len(arr) > 0 {
			a.broadcastLTP(arr, 100.0)
		}
		time.Sleep(time.Second * 2)
	}
}

func (a *MockLTPClient) startBroadcastWithControl(start float64, end float64, step float64, sleepDuration time.Duration) {

	//fmt.Println("startBroadcastWithControl : started")
	for {
		for price := start; (step > 0 && price <= end) || (step < 0 && price >= end); price = price + step {
			arr := make([]map[string]interface{}, 0)

			a.mu.RLock()
			instrMap := a.subscribedInstruments
			for _, instrument := range instrMap {
				msg := make(map[string]interface{})
				msg["tk"] = instrument.ExchangeToken
				arr = append(arr, msg)
			}
			a.mu.RUnlock()

			if len(arr) > 0 {
				//fmt.Println("startBroadcastWithControl : Broadcasting LTP")
				a.broadcastLTP(arr, price)
				time.Sleep(sleepDuration)
				price = price + step
				if (step > 0 && price >= end) || (step < 0 && price <= end) {
					break
				}
			} else {
				//fmt.Println("startBroadcastWithControl : no Instruments")
				time.Sleep(sleepDuration)
			}

		}

	}
}

func (a *MockLTPClient) incrementAndGetCount() int {
	a.count = a.count + 1
	return a.count
}

// Triggered when a message is received
func (a *MockLTPClient) broadcastLTP(message []map[string]interface{}, ltp float64) {
	for _, m := range message {
		fmt.Printf("Received ticker for %s -> %v\n", m["tk"], m)
		tk, exchTkExists := m["tk"]
		if exchTkExists {
			exchToken := fmt.Sprintf("%v", tk)

			a.mu.Lock()
			subs, exists := a.subscribedChannels[exchToken]
			a.mu.Unlock()

			if exists {
				for _, channel := range subs {
					fmt.Printf("Trying to send ticker over channel %s -> %v\n", m["tk"], m)
					channel <- &core.LTP{
						Ltp:            ltp,
						Open:           100,
						Close:          100,
						Low:            100,
						High:           100,
						ExchangeToken:  exchToken,
						Scrip:          "Scrip",
						InstrumentType: core.Stock,
					}
					fmt.Printf("Sent ticker over channel %s -> %v\n", m["tk"], ltp)
				}
			}
		}
		// else {
		// 	fmt.Printf("Received an invalid ticker for %s -> %v\n", m["tk"], m)
		// }
	}
}
