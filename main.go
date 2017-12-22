package main

import (
	"github.com/garyburd/redigo/redis"
	"github.com/jesusfar/go.mas/agent"
	"log"
	"github.com/jesusfar/go.mas/environment"
	"fmt"
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

	log.Println("Start running ...")

	pubSubConn := redis.PubSubConn{Conn: redisConnPubSub}
	defer pubSubConn.Close()

	// Set environment
	env := environment.Environment{
		Conn: redisConn,
		PubSubConn: &pubSubConn,
	}

	agent1 := agent.New("agent1", &env)

	// Agents subscribes to default channel
	agent1.Subscribe(environment.DEFAULT_CHANNEL)

	agent1.AddFriend("agent2")

	go agent1.Run()

	agent2 := agent.New("agent2", &env)
	agent2.Subscribe(environment.DEFAULT_CHANNEL)
	agent2.AddFriend("agent1")

	go agent2.Run()

	var input string
	fmt.Scanln(&input)
}
