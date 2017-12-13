package environment

import (
	"github.com/garyburd/redigo/redis"
	"github.com/jesusfar/go.mas/aclmessage"
	"encoding/json"
	"log"
)

const DEFAULT_CHANNEL string  = "mas-env"

type Environment struct {
	Conn redis.Conn
	PoolConn *redis.Pool
	PubSubConn *redis.PubSubConn
}

func (e *Environment) SendMessage(channel string, message aclmessage.Message)  {
	serializedMessage, err := json.Marshal(message)

	if err != nil {
		log.Printf("[Environment] Error on sealization of message\n")
		return
	}

	_, err = e.Conn.Do("PUBLISH", channel, serializedMessage)

	if err != nil {
		log.Printf("[Environment] Error sending message \n")
	}
}
