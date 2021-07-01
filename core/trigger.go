package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"github.com/revulcan/stock-alert-system/data_service/core"
	log "github.com/sirupsen/logrus"
)

const (
	LTP = iota
)

const (
	NSE = iota
	BSE
)

const (
	LT = iota
	LTE
	GT
	GTE
)

type Trigger struct {
	Id         string
	TAttrib    int
	Operator   int
	TPrice     float64
	TNearPrice float64
	Instument  *core.Instrument
	Status     int
	mu         sync.Mutex
	DataPoint  int64
}

func (t *Trigger) ToString() (string, error) {
	b, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(b), nil
}

func (t *Trigger) GetStatusString() (string, error) {
	if t.Status == NOT_HIT {
		return "NOT_HIT", nil
	} else if t.Status == NEAR_HIT {
		return "NEAR_HIT", nil
	} else if t.Status == HIT {
		return "HIT", nil
	}
	return "", errors.New("status is invalid")
}

const (
	ALL = iota
	NOT_HIT
	NEAR_HIT
	HIT
)

func (t *Trigger) Process(ltp *core.LTP, c chan *Trigger, index int) {
	t.mu.Lock()
	t.DataPoint++
	point := t.DataPoint
	log.Debugf("Received-d Ticker %d for %s -> %v\n", point, ltp.ExchangeToken, ltp)
	t.mu.Unlock()
	if t.Operator == LT {
		if ltp.Ltp < t.TPrice {
			hitTriggered(t, c)
		} else if ltp.Ltp < t.TNearPrice {
			nearHitTriggered(t, c)
		}
	} else if t.Operator == LTE {
		if ltp.Ltp <= t.TPrice {
			hitTriggered(t, c)
		} else if ltp.Ltp <= t.TNearPrice {
			nearHitTriggered(t, c)
		}
	} else if t.Operator == GT {
		if ltp.Ltp > t.TPrice {
			hitTriggered(t, c)
		} else if ltp.Ltp > t.TNearPrice {
			nearHitTriggered(t, c)
		}
	} else if t.Operator == GTE {
		if ltp.Ltp >= t.TPrice {
			hitTriggered(t, c)
		} else if ltp.Ltp >= t.TNearPrice {
			nearHitTriggered(t, c)
		}
	} else {
		// error
		fmt.Println("Reached a invalid Atrrib")
	}
	log.Debugf("Processed-d Ticker %d for %s -> %v\n", point, ltp.ExchangeToken, ltp)
}

func nearHitTriggered(t *Trigger, c chan *Trigger) {
	t.mu.Lock()
	defer t.mu.Unlock()
	hit := t.Status == HIT
	nearHit := t.Status == NEAR_HIT
	if !nearHit && !hit {
		log.Debugf("NEAR HIT-a Trigged Ticker %d for %s\n", t.DataPoint, t.Instument.ExchangeToken)
		t.Status = NEAR_HIT
		c <- t
	}
}

func hitTriggered(t *Trigger, c chan *Trigger) {
	t.mu.Lock()
	defer t.mu.Unlock()
	hit := t.Status == HIT
	if !hit {
		log.Debugf("HIT-a Trigged Ticker %d for %s\n", t.DataPoint, t.Instument.ExchangeToken)
		t.Status = HIT
		c <- t
	}
}
