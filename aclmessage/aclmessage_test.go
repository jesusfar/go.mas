package aclmessage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewACLMessage(t *testing.T) {
	aclMessage := NewACLMessage(REQUEST)
	aclMessage.SetSender("Me")
	aclMessage.SetReceiver("MyFriend")

	assert.Equal(t, REQUEST, aclMessage.GetPerformative())
	assert.NotNil(t, aclMessage.ConversationId)
}
