package agent

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/garyburd/redigo/redis"
)

func TestNew(t *testing.T) {
    conn := redis.PubSubConn{}
    agent := New("MyAgent", &conn)

    assert.Equal(t, "MyAgent", agent.GetName(), "Agent name should be: MyAgent")
    assert.Equal(t, CREATED, agent.GetStatus(), "Agent status should be: CREATED")
}
