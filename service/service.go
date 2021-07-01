package service

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/revulcan/stock-alert-system/core"
	ds_core "github.com/revulcan/stock-alert-system/data_service/core"
	"github.com/revulcan/stock-alert-system/data_service/service"
	log "github.com/sirupsen/logrus"
)

type TriggerService interface {
	GetTrigger(id string) (*core.Trigger, error)
	CreateTrigger(*core.Trigger) error
	GetTriggerStatus(id string) (string, error)
	DeleteTrigger(id string) error
	Stop() error
}

type TriggerSystem struct {
	machineCount int
	triggerCount int
	machines     map[string]*core.TriggerMachine
	triggers     map[string]*core.Trigger
	repo         core.TriggerSystemRepo
	ds           *service.Service
	onTrigger    chan *core.Trigger
	ctx          context.Context
	cancel       context.CancelFunc
	mu           *sync.Mutex
	config       *core.Config
}

func NewTriggerSystem(ctx context.Context, c ds_core.LtpClient, repo core.TriggerSystemRepo, config *core.Config) *TriggerSystem {
	thisCtx, cancel := context.WithCancel(ctx)
	ts := &TriggerSystem{
		machineCount: 0,
		triggerCount: 0,
		machines:     make(map[string]*core.TriggerMachine),
		triggers:     make(map[string]*core.Trigger),
		ds:           service.New(c),
		repo:         repo,
		onTrigger:    make(chan *core.Trigger, 5),
		ctx:          thisCtx,
		cancel:       cancel,
		mu:           &sync.Mutex{},
		config:       config,
	}
	ts.ds.Init()
	go ts.onTriggerHandler()
	return ts
}

func (ts *TriggerSystem) SetupService() []error {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	errors := make([]error, 0)
	triggers, err := ts.repo.GetTriggers(core.GetTriggerOptions{
		Filter: core.ALL,
	})
	if err != nil {
		errors = append(errors, err)
		return errors
	}
	for _, trigger := range triggers {
		if trigger.Status != core.HIT {
			_, tmExists := ts.machines[trigger.Instument.Scrip]
			if !tmExists {
				tm := core.NewTriggerMachine(context.Background(), trigger.Instument, ts.ds, ts.onTrigger)
				tm.Start()
				ts.machines[trigger.Instument.Scrip] = tm
			}
			err := ts.setupTrigger(&trigger)
			if err != nil {
				errors = append(errors, err)
			}
		}
	}
	return errors
}

func (ts *TriggerSystem) GetTrigger(id string) (*core.Trigger, error) {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	trigger, exists := ts.triggers[id]
	if !exists {
		return nil, fmt.Errorf("invalid trigger id - %s", id)
	}
	return trigger, nil
}

func (ts *TriggerSystem) CreateTrigger(t *core.Trigger) error {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	machine, machineExists := ts.machines[t.Instument.Scrip]
	if machineExists {
		// check if trigger with exact same PricePoint exist
		existingT, triggerExists := machine.PricePoints[t.TPrice]
		if triggerExists && (existingT.Status == core.NOT_HIT || existingT.Status == core.NEAR_HIT) {
			return fmt.Errorf("trigger already exists - %s", existingT.Id)
		} else {
			err := ts.setupTrigger(t)
			return err
		}
	} else {
		tm := core.NewTriggerMachine(context.Background(), t.Instument, ts.ds, ts.onTrigger)
		tm.Start()
		ts.machines[t.Instument.Scrip] = tm
		err := ts.setupTrigger(t)
		return err
	}
}

func (ts *TriggerSystem) DeleteTrigger(id string) error {

	t, tExists := ts.triggers[id]
	if !tExists {
		return fmt.Errorf("trigger not found - %s", id)
	}
	m, tmExists := ts.machines[t.Instument.Scrip]
	if !tmExists {
		return fmt.Errorf("trigger machine not found - %s", id)
	}

	err := ts.repo.DeleteTrigger(t)
	if err != nil {
		return err
	}

	err = m.DeleteTrigger(t)
	if err != nil {
		return err
	}
	delete(ts.triggers, t.Id)
	return nil
}

func (ts *TriggerSystem) GetTriggerStatus(id string) (string, error) {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	t, tExists := ts.triggers[id]
	if !tExists {
		return "", fmt.Errorf("trigger not found - %s", id)
	}
	repr, err := t.ToString()
	if err != nil {
		return "", err
	}
	return repr, nil
}

func (ts *TriggerSystem) setupTrigger(t *core.Trigger) error {
	err := ts.repo.CreateTrigger(t)
	if err != nil {
		return err
	}

	m, tmExists := ts.machines[t.Instument.Scrip]
	if !tmExists {
		return errors.New("trigger machine doesn't exist")
	}
	err = m.AddTrigger(t)
	if err != nil {
		return err
	}
	ts.triggers[t.Id] = t

	return nil
}

func (ts *TriggerSystem) onTriggerHandler() {
	for {
		select {
		case <-ts.ctx.Done():
			return
		case triggerUpdate := <-ts.onTrigger:
			go func() {
				// todo : delete it or leave it to get status for later
				// if triggerUpdate.Status == core.HIT {
				// 	delete(ts.triggers, triggerUpdate.Id)
				// }
				if triggerUpdate.Status == core.NEAR_HIT {
					log.Debugf("NEAR HIT-b Trigged Ticker for %s\n", triggerUpdate.Instument.ExchangeToken)
				} else {
					log.Debugf("HIT-b Trigged Ticker for %s\n", triggerUpdate.Instument.ExchangeToken)
				}

				ts.mu.Lock()
				machine := ts.machines[triggerUpdate.Instument.Scrip]
				if machine != nil {
					destroyed := machine.DestroyIfEmpty()
					if destroyed {
						delete(ts.machines, triggerUpdate.Instument.Scrip)
					}
				}
				ts.mu.Unlock()

				err := ts.repo.OnTrigger(triggerUpdate)
				if err != nil {
					status, _ := triggerUpdate.GetStatusString()
					fmt.Printf("Error on invoking OnTrigger %s - %s\n", triggerUpdate.Id, status)
				} else {
					fmt.Printf("Trigger with ID %s got triggered with status - %d\n", triggerUpdate.Id, triggerUpdate.Status)
				}
			}()
		}
	}
}

func (ts *TriggerSystem) Stop() error {
	ts.ds.Close()
	ts.cancel()
	return nil
}
