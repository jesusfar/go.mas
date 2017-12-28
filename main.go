package main

import (
	"github.com/jesusfar/go.mas/environment"
	"fmt"
	"github.com/jesusfar/go.mas/agent"
	"github.com/jesusfar/go.mas/messaging"
)

func main() {

	env := environment.NewEnvironment("mas-go")

	agent1 := agent.New("agent1", messaging.NewRedisMessaging())
	agent2 := agent.New("agent2", messaging.NewRedisMessaging())

	env.RegisterAgent(agent1)
	env.RegisterAgent(agent2)
	env.ShowAgents()
	env.Run()

	var input string
	fmt.Scanln(&input)
}
