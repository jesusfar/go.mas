package agent

import (
	"fmt"
	"github.com/satori/go.uuid"
	"time"
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
}

// New creates a valid agent
func New(name string) *agent {
	agent := &agent{}
	agent.id = uuid.NewV4()
	agent.name = name
	agent.status = CREATED

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
	go process()
	a.ChangeStatus(RUNNING)
}

// Stop the agent
func (a *agent) Stop() {
	a.ChangeStatus(STOPPED)
}

func process() {
	for {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("Running ...")
	}
}
