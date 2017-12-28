package messaging

import "github.com/jesusfar/go.mas/aclmessage"

const BROADCAST_CHANNEL string  = "broadcast-mas"

// Messaging Interface
type Messaging interface {
	// Publish a message in a channel
	Publish(channel string, message aclmessage.Message) error

	// Subscribe to channel
	Subscribe(channel string) error

	// GetMessageChannel returns a msgChannel
	GetMessageChannel() chan []byte

	// Run Messaging System
	Run()
}
