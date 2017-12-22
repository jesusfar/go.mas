package agent

import (
	"github.com/satori/go.uuid"
	"time"
	"github.com/garyburd/redigo/redis"
	"log"
	"github.com/jesusfar/go.mas/aclmessage"
	"encoding/json"
	"github.com/jesusfar/go.mas/environment"
	"fmt"
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

	friends []string

	env *environment.Environment
}

// New creates a valid agent
func New(name string, env *environment.Environment) *agent {
	agent := &agent{}
	agent.id = uuid.NewV4()
	agent.name = name
	agent.status = CREATED
	agent.env = env

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

	a.logger(fmt.Sprintf("Status: %s", a.status))

	a.process()
}

// Subscribe agent to channel
func (a *agent) Subscribe(channel string)  {
	err := a.env.PubSubConn.Subscribe(channel)

	// TODO handle error subscribe
	if err != nil {
		panic(err)
	}
}

func (a *agent) IsFriend(agentName string) bool {
	for _,friend := range a.friends {
		if friend == agentName {
			return true
		}
	}

	return false
}

func (a *agent) AddFriend(agentName string)  {
	a.logger(fmt.Sprintf("Add friend to %s", agentName))
	a.friends = append(a.friends, agentName)
}

func (a *agent) processMessage(messageData []byte)  {
	a.logger(fmt.Sprintf("Message received: %s", string(messageData)))
	var message aclmessage.Message

	err := json.Unmarshal(messageData, &message)

	if err != nil {
		a.logger("Error parsing message")
		return
	}

	// Only process message if agent name is equal to receiver.
	if message.Receiver == a.name {
		a.logger(fmt.Sprintf("Message received from %s", message.Sender))
		switch message.Performative {
		case aclmessage.REQUEST:
			a.logger("Try to process REQUEST")
			go a.processRequest(message)
		case aclmessage.AGREE:
			a.logger("Try to process AGREE")
			go a.processAgree(message)
		case aclmessage.REFUSE:
			a.logger("Try to process REFUSE")
		case aclmessage.INFORM:
			a.logger("Try to process INFORM")
		}
	}
}

func (a *agent) processRequest(message aclmessage.Message)  {
	messageResponse := aclmessage.Message{
		Sender: a.name,
		Receiver: message.Sender,
		Content: "Processing Request",
	}

	if !a.IsFriend(message.Sender) {
		messageResponse.Performative = aclmessage.REFUSE;
		a.env.SendMessage(environment.DEFAULT_CHANNEL, messageResponse)
	}

	// Send message Agree
	messageResponse.Performative = aclmessage.AGREE;
	a.env.SendMessage(environment.DEFAULT_CHANNEL, messageResponse)

	// Execute task
	a.logger("Executing tasks..")
}

func (a *agent) processAgree(message aclmessage.Message)  {
	// Request was accepted by the agent sender
	a.logger(fmt.Sprintf("Request was accepte by agent %s", message.Sender))
}

func (a *agent) process() {
	for {
		time.Sleep(1000 * time.Millisecond)
		a.logger("Fetching from channel ...")
		switch v := a.env.PubSubConn.Receive().(type) {
		case redis.Message:
			a.processMessage(v.Data)
		case redis.Subscription:
			log.Printf("[%s] Subscription: channel: %s type: %s count: %d\n", a.GetName(), v.Channel, v.Kind, v.Count)

		case error:
			a.logger("Error pub/sub, delivery has stopped \n")
			return
		}
	}
}

func (a *agent) logger(message string)  {
	defaultMessage := "[%s] " + message
	log.Printf(defaultMessage, a.GetName())
}