package internal

import "github.com/jkstack/anet"

// AgentName agent name
var AgentName string

// Agent agent object
type Agent struct {
	cfgDir  string
	cfg     *Configure
	version string
	chWrite chan *anet.Msg
}

// New create agent object
func New(dir, version string) *Agent {
	return &Agent{
		cfgDir:  dir,
		cfg:     load(dir),
		version: version,
		chWrite: make(chan *anet.Msg),
	}
}

// AgentName get agent name
func (agent *Agent) AgentName() string {
	return "example-agent"
}

// Version get agent version
func (agent *Agent) Version() string {
	return agent.version
}
