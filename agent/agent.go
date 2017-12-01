package agent

import (
	"github.com/satori/go.uuid"
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
	id      uuid.UUID
	name    string
	beliefs []belief
}

func New(name string) *agent {
	agent := &agent{}
	agent.id = uuid.NewV4()
	agent.name = name

	return agent
}

func (a *agent) GetId() uuid.UUID {
	return a.id
}
