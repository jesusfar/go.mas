package agent

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestNew(t *testing.T) {
    agent := New("MyAgent")

    assert.Equal(t, "MyAgent", agent.GetName(), "Agent name should be: MyAgent")
    assert.Equal(t, CREATED, agent.GetStatus(), "Agent status should be: CREATED")
}
