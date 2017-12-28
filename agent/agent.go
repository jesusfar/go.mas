package agent

import (
	"github.com/satori/go.uuid"
	"time"
	"log"
	"github.com/jesusfar/go.mas/aclmessage"
	"encoding/json"
	"fmt"
	"github.com/jesusfar/go.mas/messaging"
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

type Agent struct {
	id      uuid.UUID
	name    string
	status  Status

	friends []string

	msgConn messaging.Messaging
}

// New creates a valid agent
func New(name string, msgConn messaging.Messaging) *Agent {
	agent := &Agent{}
	agent.id = uuid.NewV4()
	agent.name = name
	agent.status = CREATED
	agent.msgConn = msgConn

	// By default agent subscribe to broadcast channel
	agent.Subscribe(messaging.BROADCAST_CHANNEL)

	return agent
}

// GetId returns uuid agent
func (a *Agent) GetId() uuid.UUID {
	return a.id
}

// GetName returns agent's name
func (a *Agent) GetName() string {
	return a.name
}

// GetStatus returns agent's name
func (a *Agent) GetStatus() Status {
	return a.status
}

func (a *Agent) ChangeStatus(status Status) {
	a.status = status
}

// Run the agent
func (a *Agent) Run() {
	a.ChangeStatus(RUNNING)

	message := aclmessage.Message{
		Performative: aclmessage.PROPAGATE,
		Sender: a.GetName(),
		Receiver: "ALL",
	}

	a.SendMessage(messaging.BROADCAST_CHANNEL, message)

	a.process()
}

func (a *Agent) SendMessage(channel string, message aclmessage.Message)  {
	a.msgConn.Publish(channel, message)
}

// Subscribe agent to channel
func (a *Agent) Subscribe(channel string)  {

	err := a.msgConn.Subscribe(channel)

	// TODO handle error subscribe
	if err != nil {
		panic(err)
	}
}

func (a *Agent) IsFriend(agentName string) bool {
	for _,friend := range a.friends {
		if friend == agentName {
			return true
		}
	}

	return false
}

func (a *Agent) AddFriend(agentName string)  {
	a.logger(fmt.Sprintf("Add friend to %s", agentName))
	a.friends = append(a.friends, agentName)
}

func (a *Agent) processMessage(messageData []byte)  {
	a.logger(fmt.Sprintf("Message received: %s", string(messageData)))
	var message aclmessage.Message

	err := json.Unmarshal(messageData, &message)

	if err != nil {
		a.logger("Error parsing message. Message should be ACL FIPA")
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

func (a *Agent) processRequest(message aclmessage.Message)  {
	messageResponse := aclmessage.Message{
		Sender: a.name,
		Receiver: message.Sender,
		Content: "Processing Request",
	}

	if !a.IsFriend(message.Sender) {
		messageResponse.Performative = aclmessage.REFUSE;
		a.SendMessage(message.Sender, messageResponse)
	}

	// Send message Agree
	messageResponse.Performative = aclmessage.AGREE;
	a.SendMessage(message.Sender, messageResponse)

	// Execute task
	a.logger("Executing tasks..")
}

func (a *Agent) processAgree(message aclmessage.Message)  {
	// Request was accepted by the agent sender
	a.logger(fmt.Sprintf("Request was accepte by agent %s", message.Sender))
}

func (a *Agent) process() {
	go func() {
		for {
			time.Sleep(1000 * time.Millisecond)
			a.logger("Running...")
			messageData := <- a.msgConn.GetMessageChannel()
			a.processMessage(messageData)
		}
	}()
}

func (a *Agent) logger(message string)  {
	defaultMessage := "[%s] " + message
	log.Printf(defaultMessage, a.GetName())
}