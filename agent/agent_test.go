package agent

import (
    "testing"
	"github.com/jesusfar/go.mas/messaging"
	"github.com/stretchr/testify/assert"
	"github.com/jesusfar/go.mas/aclmessage"
)

func Test_New(t *testing.T) {

	agentName := "agent1"
	agent := New(agentName, messaging.NewGoMessaging())
	assert.Equal(t, agentName, agent.GetName())
	assert.Equal(t, CREATED, agent.GetStatus())
}

func TestAgent_Run(t *testing.T) {
	agent := New("agent", messaging.NewGoMessaging())
	assert.Equal(t, "agent", agent.GetName())
	agent.Run()
	assert.Equal(t, RUNNING, agent.GetStatus())
}

func TestAgent_Subscribe(t *testing.T) {
	agent1 := New("agent1", messaging.NewGoMessaging())
	agent1.Subscribe("agent2")
}

func TestAgent_SendMessage(t *testing.T) {
	agent := New("agent", messaging.NewGoMessaging())

	msg := aclmessage.Message{
		Performative: aclmessage.REQUEST,
		Sender: agent.GetName(),
		Receiver: "ALL",
		Content: "Help me!",
	}

	agent.SendMessage(messaging.BROADCAST_CHANNEL, msg)
}

func TestAgent_Publish(t *testing.T) {
	agent2 := New("agent2", messaging.NewGoMessaging())

	msg := aclmessage.Message{
		Performative: aclmessage.REQUEST,
		Sender: agent2.GetName(),
		Receiver: "agent1",
		Content: "Help me!",
	}

	agent2.Publish(msg)
}
