package messaging

import (
	"github.com/jesusfar/go.mas/aclmessage"
	"github.com/garyburd/redigo/redis"
	"log"
	"encoding/json"
	"flag"
)

var (
	redisAddress = flag.String("redis-address", ":6379", "Redis address")
	maxConn      = flag.Int("max-conn", 10, "Max connection")
)

type RedisMessaging struct {
	conn       redis.Conn
	poolConn   *redis.Pool
	pubSubConn *redis.PubSubConn
	msgChannel chan []byte
}

func NewRedisMessaging() *RedisMessaging {
	channel := make(chan []byte, 10)

	redisPool := GetRedisPool()
	pubSubConn := redis.PubSubConn{Conn: redisPool.Get()}

	messaging := RedisMessaging{
		conn: redisPool.Get(),
		poolConn: redisPool,
		pubSubConn: &pubSubConn,
		msgChannel: channel,
	}

	messaging.Run()

	return &messaging
}


func  GetRedisPool() *redis.Pool {
	flag.Parse()

	redisPool := redis.NewPool(func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", *redisAddress)

		if err != nil {
			return nil, err
		}
		return c, err
	}, *maxConn)

	return redisPool
}

func (r *RedisMessaging) GetMessageChannel() chan []byte {
	return r.msgChannel
}

func (r *RedisMessaging) Publish(channel string, message aclmessage.Message) error {
	serializedMessage, err := json.Marshal(message)

	if err != nil {
		log.Println("[RedisMessaging] Error on sealization of message.")
		return err
	}

	_, err = r.conn.Do("PUBLISH", channel, serializedMessage)

	if err != nil {
		log.Println("[RedisMessaging] Error sending message.")
		return err
	}
	return nil
}

func (r *RedisMessaging) Subscribe(channel string) error {
	err := r.pubSubConn.Subscribe(channel)

	if err != nil {
		log.Println("[RedisMessaging] Error subscribing to channel: " + channel)
		return err
	}

	return nil
}

func (r *RedisMessaging) Run()  {
	go func() {
		for {
			switch v := r.pubSubConn.Receive().(type) {
			case redis.Message:
				log.Println("[RedisMessaging] Message received")
				r.msgChannel <- v.Data

			case redis.Subscription:
				log.Printf("[RedisMessaging] Subscription: channel: %s type: %s count: %d\n", v.Channel, v.Kind, v.Count)
			case error:
				log.Println("[RedisMessaging] Error pub/sub, delivery has stopped.")
				log.Println(v)
				return
			}
		}
	}()
}