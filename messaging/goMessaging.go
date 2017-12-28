package messaging

import "github.com/jesusfar/go.mas/aclmessage"

// GoMessaging struct implements Messaging
type GoMessaging struct {
	msgChannel chan []byte
}

func NewGoMessaging() *GoMessaging  {
	channel := make(chan []byte, 10)
	goMessaging := GoMessaging{
		msgChannel: channel,
	}

	return &goMessaging
}

func (g *GoMessaging) Subscribe(channel string) error {
	// Not implemented yet
	return nil
}

func (g *GoMessaging) Publish(channel string, message aclmessage.Message) error {
	// Not implemented yet
	return nil
}

func (g *GoMessaging) GetMessageChannel() chan []byte {
	// Not implemented yet
	return g.msgChannel
}

func (g *GoMessaging) Run() {

}