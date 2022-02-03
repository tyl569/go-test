package Class42

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

type DemoCollector struct {
	evtReceiver EventReceiver
	agentCtx    context.Context
	stopChan    chan struct{}
	name        string
	content     string
}

func NewCollector(name string, content string) *DemoCollector {
	return &DemoCollector{
		stopChan: make(chan struct{}),
		name:     name,
		content:  content,
	}
}

func (dc *DemoCollector) Init(receiver EventReceiver) error {
	fmt.Println("init collector ", dc.name)
	dc.evtReceiver = receiver
	return nil
}

func (dc *DemoCollector) Start(ctx context.Context) {
	fmt.Println("start collector", dc.name)
	for {
		select {
		case <-ctx.Done():
			dc.stopChan <- struct{}{}
		default:
			time.Sleep(time.Millisecond * 50)
			dc.evtReceiver.OnEvent(Event{dc.name, dc.content})
		}
	}
}

func (dc *DemoCollector) Stop(ctx context.Context) error {
	fmt.Println("stop collector", dc.name)
	select {
	case <-dc.stopChan:
		return nil
	case <-time.After(time.Second * 1):
		return errors.New("failed to stop for timeout")
	}
}

func (dc *DemoCollector) Destory() error {
	fmt.Println("destory collector", dc.name)
	return nil
}

func TestAgent(t *testing.T) {
	agent := NewAgent(100)
	c1 := NewCollector("c1", "1")
	c2 := NewCollector("c2", "2")
	agent.RegisterCollectors("c1", c1)
	agent.RegisterCollectors("c2", c2)
	agent.Start()
	fmt.Println(agent.Start())
	time.Sleep(time.Second * 1)
	agent.Stop()
}
