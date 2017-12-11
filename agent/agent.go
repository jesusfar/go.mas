package agent

import (
	"fmt"
	"github.com/satori/go.uuid"
	"time"
	"github.com/garyburd/redigo/redis"
)

type Status int

const (
	CREATED Status = 1 + iota
	RUNNING
	STOPPED
)

type belief struct {
	something string
}

type desire struct {
	something string
}

type goal struct {
	desc string
}

type plan struct {
	desc string
}

type intention struct {
	desc string
}

type event struct {
	name string
}

type agent struct {
	id     uuid.UUID
	name   string
	status Status

	pubSubConn *redis.PubSubConn
}

// New creates a valid agent
func New(name string, pubSubConn *redis.PubSubConn) *agent {
	agent := &agent{}
	agent.id = uuid.NewV4()
	agent.name = name
	agent.status = CREATED

	agent.pubSubConn = pubSubConn
	return agent
}

// GetId returns uuid agent
func (a *agent) GetId() uuid.UUID {
	return a.id
}

// GetName returns agent's name
func (a *agent) GetName() string {
	return a.name
}

// GetStatus returns agent's name
func (a *agent) GetStatus() Status {
	return a.status
}

func (a *agent) ChangeStatus(status Status) {
	a.status = status
}

// Run the agent
func (a *agent) Run() {
	a.ChangeStatus(RUNNING)
	a.subscribe()

	fmt.Printf("Agent: %s is %s", a.name, a.status)

	for {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("Fetch from channel")
	}
}

// Stop the agent
func (a *agent) Stop() {
	a.ChangeStatus(STOPPED)
}

func (a *agent) subscribe()  {
	// TODO handle subscribe
	a.pubSubConn.Subscribe("mas-channel")
}