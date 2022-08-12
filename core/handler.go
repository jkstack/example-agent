package core

import (
	"context"

	"github.com/jkstack/anet"
)

func (agent *Agent) OnConnect() {
}

func (agent *Agent) OnDisconnect() {
}

func (agent *Agent) OnReportMonitor() {
}

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
