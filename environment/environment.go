package environment

import (
	"github.com/jesusfar/go.mas/agent"
	"fmt"
	"log"
)

type Environment struct {
	name string
	agents []*agent.Agent
}

func NewEnvironment(name string) *Environment {
	env := Environment{
		name: name,
	}
	return &env
}

func (e *Environment) RegisterAgent(agent *agent.Agent)  {
	e.agents = append(e.agents, agent)
}

func (e *Environment) ShowAgents()  {
	for _, agent := range e.agents {
		fmt.Println(agent.GetName())
	}
}

func (e *Environment) Run()  {
	log.Println("[Environment] Run environment: " + e.name)
	for _, agent := range e.agents  {
		agent.Run()
	}
}
