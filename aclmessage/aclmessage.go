package aclmessage

import "github.com/satori/go.uuid"

type Performative int

// FIPA http://www.fipa.org/specs/fipa00037/SC00037J.html
const (
	ACCEPT_PROPOSAL Performative = iota + 1
	AGREE
	CANCEL
	CFP
	CONFIRM
	DISCONFIRM
	FAILURE
	INFORM
	INFORM_IF
	INFORM_REF
	NOT_UNDERSTOOD
	PROPAGATE
	PROPOSE
	PROXY
	QUERY_IF
	QUERY_REF
	REFUSE
	REJECT_PROPOSAL
	REQUEST
	REQUEST_WHEN
	REQUEST_WHENEVER
	SUBSCRIBE
)

// Message struct based on FIPA ACL Message
type Message struct {
	ConversationId string `json:"conversation_id"`
	Performative   Performative `json:"performative"`
	Sender         string `json:"sender"`
	Receiver       string `json:"receiver"`
	ReplyTo        string `json:"reply_to"`
	Content        string `json:"content"`
	Language       string `json:"language"`
	Encoding       string `json:"encoding"`
	Ontology       string `json:"ontology"`
	Protocol       string `json:"protocol"`
	ReplyWith      string `json:"reply_with"`
	UnReplyTo      string `json:"un_reply_to"`
	ReplyBy        string `json:"reply_by"`
}

func NewACLMessage(performative Performative) Message {
	message := Message{
		ConversationId: uuid.NewV4().String(),
		Performative: performative,
	}

	return message
}

func (m Message) GetPerformative() Performative {
	return m.Performative
}

func (m Message) GetConversationId() string {
	return m.ConversationId
}

func (m Message) GetSender() string {
	return m.Sender
}

func (m Message) GetReceiver() string {
	return m.Receiver
}

func (m Message) GetReplyTo() string {
	return m.ReplyTo
}

func (m Message) GetContent() string {
	return m.Content
}

func (m Message) GetLanguage() string {
	return m.Language
}

func (m Message) GetOntology() string {
	return m.Ontology
}

func (m Message) GetProtocol() string {
	return m.Protocol
}

func (m Message) GetReplyWith() string {
	return m.ReplyWith
}

func (m Message) GetUnReplyTo() string {
	return m.UnReplyTo
}

func (m Message) GetReplyBy() string {
	return m.ReplyBy
}