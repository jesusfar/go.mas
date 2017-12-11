package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"github.com/jesusfar/go.mas/agent"
	"github.com/jesusfar/go.mas/aclmessage"
)

func getRedisConn() (redis.Conn, error) {
	return redis.Dial("tcp", ":6379")
}

func main() {

	redisConn, err := getRedisConn()

	if err != nil {
		panic(err)
	}

	defer redisConn.Close()

	redisConnPubSub, err := getRedisConn()

	if err != nil {
		panic(err)
	}

	defer redisConnPubSub.Close()



	fmt.Println("Running...")

	pubSubConn := redis.PubSubConn{Conn: redisConnPubSub}
	defer pubSubConn.Close()

	channel := "mas-env"

	agent := agent.New("agent1", redisConn, &pubSubConn)

	agent.Subscribe(channel)

	agent.AddFriend("agent2")

	message := aclmessage.Message{
		Performative:aclmessage.REQUEST,
	}

	agent.SendMessage(channel, message)

	fmt.Println(message)

	agent.Run()
}
