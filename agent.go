package main

import (
    "github.com/satori/go.uuid"
)

// Agent struct
type Agent struct {
    Id   uuid.UUID
    Name string
}

func New(name string) *Agent {
    agent := &Agent{}
    agent.Id = uuid.NewV4()
    agent.Name = name

    return agent
}
