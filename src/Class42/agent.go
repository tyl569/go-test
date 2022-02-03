package Class42

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
)

type Event struct {
	Source  string
	content string
}

type EventReceiver interface {
	OnEvent(evt Event)
}

type Collector interface {
	Init(evtReceive EventReceiver) error
	Start(agt context.Context) error
	Stop() error
	Destory() error
}

type Agent struct {
	collectors map[string]Collector
	evtBuf     chan Event
	cancel     context.CancelFunc
	ctx        context.Context
	state      int
}

func (agent *Agent) OnEvent(event Event) {
	agent.evtBuf <- event
}

const (
	Waiting = iota
	Running
)

var WrongStateError = errors.New("state code error")

type CollectorsError struct {
	CollectorErrors []error
}

func (ce *CollectorsError) Error() string {
	var str []string
	for _, err := range ce.CollectorErrors {
		str = append(str, err.Error())
	}
	return strings.Join(str, ";")
}

func (agent *Agent) Start() error {
	if agent.state != Waiting {
		return WrongStateError
	}
	agent.state = Running
	agent.ctx, agent.cancel = context.WithCancel(context.Background())
	go agent.EventProcessGroutine()
	return agent.StartCollectors()
}

func (agent *Agent) EventProcessGroutine() {
	var evtSeg [10]Event
	for {
		for i := 0; i < 10; i++ {
			select {
			case evtSeg[i] = <-agent.evtBuf:
			case <-agent.ctx.Done():
				return
			}
		}
		fmt.Println(evtSeg)
	}
}

func (agent *Agent) Stop() error {
	if agent.state != Waiting {
		return WrongStateError
	}
	agent.state = Waiting
	agent.cancel()
	return agent.StopCollector()
}

func (agent *Agent) Destory() error {
	if agent.state != Waiting {
		return WrongStateError
	}
	return agent.DestoryCollector()
}

func (agent *Agent) RegisterCollectors(name string, collector Collector) error {
	if agent.state != Waiting {
		return WrongStateError
	}
	agent.collectors[name] = collector
	return collector.Init(agent)
}
func (agent *Agent) StartCollectors() error {
	var (
		err   error
		mutex sync.Mutex
		errs  CollectorsError
	)
	for name, collector := range agent.collectors {
		go func(name string, collector2 Collector, ctx context.Context) {
			defer func() {
				mutex.Unlock()
			}()
			err = collector2.Start(ctx)
			mutex.Lock()
			if err != nil {
				errs.CollectorErrors = append(errs.CollectorErrors, errors.New(name+":"+err.Error()))
			}
		}(name, collector, agent.ctx)
	}
	if len(errs.CollectorErrors) == 0 {
		return nil
	}
	return errs
}

func (agent *Agent) StopCollector() error {
	var (
		err  error
		errs CollectorsError
	)
	for name, collector := range agent.collectors {
		if err = collector.Stop(); err != nil {
			errs.CollectorErrors = append(errs.CollectorErrors, errors.New(name+":"+err.Error()))
		}
	}
	if len(errs.CollectorErrors) == 0 {
		return nil
	}
	return errs
}

func (agent *Agent) DestoryCollector() error {
	var (
		err  error
		errs CollectorsError
	)
	for name, collector := range agent.collectors {
		if err = collector.Destory(); err != nil {
			errs.CollectorErrors = append(errs.CollectorErrors, errors.New(name+":"+err.Error()))
		}
	}
	if len(errs.CollectorErrors) == 0 {
		return nil
	}
	return errs
}

func NewAgent(sizeBuf int) *Agent {
	agent := Agent{
		collectors: map[string]Collector{},
		evtBuf:     make(chan Event, sizeBuf),
		state:      Waiting,
	}
	return &agent
}
