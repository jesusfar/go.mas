package messaging

import "github.com/jesusfar/go.mas/aclmessage"

const BROADCAST_CHANNEL string  = "broadcast-mas"

// Messaging Interface
type Messaging interface {
	Publish(channel string, message aclmessage.Message) error
	Subscribe(channel string) error
	GetMessageChannel() chan []byte
	Run()
}
