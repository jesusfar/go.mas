package agent

import (
	"fmt"
	"github.com/satori/go.uuid"
	"time"
	"github.com/garyburd/redigo/redis"
	"log"
	"github.com/jesusfar/go.mas/aclmessage"
	"encoding/json"
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

	conn redis.Conn
	pubSubConn *redis.PubSubConn
}

// New creates a valid agent
func New(name string, conn redis.Conn, pubSubConn *redis.PubSubConn) *agent {
	agent := &agent{}
	agent.id = uuid.NewV4()
	agent.name = name
	agent.status = CREATED

	agent.conn = conn
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

	log.Printf("Agent: %s status: %s \n", a.name, a.status)

	a.process()
}

// Subscribe agent to channel
func (a *agent) Subscribe(channel string)  {

	err := a.pubSubConn.Subscribe(channel)

	// TODO handle error subscribe
	if err != nil {
		panic(err)
	}
}

func (a *agent) SendMessage(channel string, message aclmessage.Message)  {

	serializedMessage, err := json.Marshal(message)

	if err != nil {
		log.Printf("Error on sealization of message \n")
		return
	}

	_, err = a.conn.Do("PUBLISH", channel, serializedMessage)

	if err != nil {
		log.Printf("error sending message \n")
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
	log.Printf("Agent % add friend to %s", a.GetName(), agentName)
	a.friends = append(a.friends, agentName)
}

func (a *agent) processMessage(messageInput []byte)  {
	var message aclmessage.Message

	err := json.Unmarshal(messageInput, &message)

	if err != nil {
		log.Println("Error parsing message")
		return
	}

	// Only process message if agent name is equal to receiver.
	if message.Receiver == a.name && a.IsFriend(message.Sender) {
		log.Println("Message accepted")
	}
}

func (a *agent) process() {
	for {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("Fetch from channel")
		switch v := a.pubSubConn.Receive().(type) {
		case redis.Message:
			log.Printf(string(v.Data))
			a.processMessage(v.Data)
		case redis.Subscription:
			log.Printf("Subscription: channel: %s type: %s count: %d\n", v.Channel, v.Kind, v.Count)

		case error:
			log.Println("error pub/sub, delivery has stopped \n")
			return
		}
	}
}