package main

import (
    "fmt"
    "github.com/jesusfar/go.mas/agent"
)

func main() {
    agent := agent.New("My agent")

    fmt.Println(agent)
    agent.Run()
}
