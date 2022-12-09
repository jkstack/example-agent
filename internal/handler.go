package internal

import (
	"context"

	"github.com/jkstack/anet"
)

// OnConnect on connect callback
func (agent *Agent) OnConnect() {
}

// OnDisconnect on disconnect callback
func (agent *Agent) OnDisconnect() {
}

// OnReportMonitor on report monitor callback
func (agent *Agent) OnReportMonitor() {
}

// OnMessage on receive message callback
func (agent *Agent) OnMessage(msg *anet.Msg) error {
	switch msg.Type {
	case anet.TypeFoo:
		id := msg.TaskID
		var msg anet.Msg
		msg.Type = anet.TypeBar
		msg.TaskID = id
		agent.chWrite <- &msg
	}
	return nil
}

// LoopWrite loop write
func (agent *Agent) LoopWrite(ctx context.Context, ch chan *anet.Msg) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case msg := <-agent.chWrite:
			ch <- msg
		}
	}
}
