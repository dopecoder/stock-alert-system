package core

import (
	"context"
	"errors"
	"sync"

	core "github.com/revulcan/stock-alert-system/data_service/core"
	"github.com/revulcan/stock-alert-system/data_service/service"
	log "github.com/sirupsen/logrus"
)

type TriggerMachine struct {
	Instrument  *core.Instrument
	PricePoints map[float64]*Trigger
	Triggers    []*Trigger
	ds          *service.Service
	ctx         context.Context
	cancel      context.CancelFunc
	started     bool
	finished    bool
	mu          *sync.Mutex
	tUpdates    chan *Trigger
	onTrigger   chan *Trigger
	DataPoint   int
}

func NewTriggerMachine(ctx context.Context, instrument *core.Instrument, ds *service.Service, onTrigger chan *Trigger) *TriggerMachine {
	thisCtx, cancelFunc := context.WithCancel(ctx)
	return &TriggerMachine{
		Instrument:  instrument,
		PricePoints: make(map[float64]*Trigger),
		Triggers:    make([]*Trigger, 0),
		ds:          ds,
		ctx:         thisCtx,
		cancel:      cancelFunc,
		mu:          &sync.Mutex{},
		tUpdates:    make(chan *Trigger, 5),
		onTrigger:   onTrigger,
		DataPoint:   0,
	}
}
func (tm *TriggerMachine) Start() {
	tm.mu.Lock()
	if !tm.started {
		go tm.runner()
		go tm.processTriggerUpdates()
		tm.started = true
	}
	tm.mu.Unlock()
}

func (tm *TriggerMachine) Stop() {
	// todo: shold we use lock here?
	//log.Infoln("Stopping machine")
	tm.cancel()
}

func (tm *TriggerMachine) AddTrigger(t *Trigger) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	if !tm.finished {
		tm.Triggers = append(tm.Triggers, t)
		tm.PricePoints[t.TPrice] = t
		log.Infoln("New Trigger Added - ", t)
		return nil
	}
	return errors.New("machine has finished")
}

func (tm *TriggerMachine) DeleteTrigger(t *Trigger) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	if !tm.finished {
		delete(tm.PricePoints, t.TPrice)
		tm.removeTriggerFromList(t)
		log.Infoln("Deleted Trigger - ", t)
		return nil
	}
	return errors.New("machine has finished")
}

func (tm *TriggerMachine) DestroyIfEmpty() bool {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	if len(tm.Triggers) == 0 {
		tm.cancel()
		return true
	}
	return false
}

func (tm *TriggerMachine) runner() {
	//log.Infoln("Started runner")
	s := core.NewMockStream()
	go tm.ds.WatchLTPforInstrs([]*core.Instrument{
		tm.Instrument,
	}, s)
	for {
		//log.Infoln("Entered Runner Loop")
		select {
		case <-tm.ctx.Done():
			//log.Infoln("Runner Loop -> Context Closed")
			tm.mu.Lock()
			s.Close()
			tm.finished = true
			tm.mu.Unlock()
			//log.Infoln("Runner Loop -> Context Closed, Returning")
			return
		case data := <-s.Strm:
			tm.mu.Lock()
			tm.DataPoint += 1
			log.Debugf("Received-c Ticker %d for %s -> %v\n", tm.DataPoint, data.ExchangeToken, data)
			triggers := tm.Triggers
			for _, trigger := range triggers {
				go trigger.Process(data, tm.tUpdates, tm.DataPoint)
			}
			log.Debugf("Executed-c trigger.Process %d for  %s -> %v\n", tm.DataPoint, data.ExchangeToken, data)
			tm.mu.Unlock()
		}
	}
}

func (tm *TriggerMachine) processTriggerUpdates() {
	for {
		select {
		case <-tm.ctx.Done():
			//log.Infoln("(TriggerMachine) Exiting processTriggerUpdates")
			return
		case t := <-tm.tUpdates:
			if t.Status == NEAR_HIT {
				log.Debugf("NEAR HIT-b Trigged Ticker for %s\n", t.Instument.ExchangeToken)
			} else {
				log.Debugf("HIT-b Trigged Ticker for %s\n", t.Instument.ExchangeToken)
			}
			if tm.onTrigger != nil {
				tm.onTrigger <- t
			}
			if t.Status == HIT {
				tm.mu.Lock()
				delete(tm.PricePoints, t.TPrice)
				tm.removeTriggerFromList(t)
				tm.mu.Unlock()
			}
		}
	}
}

func (tm *TriggerMachine) removeTriggerFromList(trigger *Trigger) {
	for idx, t := range tm.Triggers {
		if t.Id == trigger.Id {
			tm.Triggers = remove(tm.Triggers, idx)
			return
		}
	}
}

func remove(s []*Trigger, i int) []*Trigger {
	s[i] = s[len(s)-1]
	// We do not need to put s[i] at the end, as it will be discarded anyway
	return s[:len(s)-1]
}
