package aclmessage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMessage(t *testing.T) {

	aclMessage := Message{Performative: REQUEST}

	assert.Equal(t, REQUEST, aclMessage.Performative)

}
