package aclmessage

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

// Message struct based on FIPA ACL message
type Message struct {
	Performative   Performative `json:"performative"`
	Sender         string `json:"sender"`
	Receiver       string `json:"receiver"`
	ReplyTo        string `json:"reply_to"`
	Content        string `json:"content"`
	Language       string `json:"language"`
	Encoding       string `json:"encoding"`
	Ontology       string `json:"ontology"`
	Protocol       string `json:"protocol"`
	ConversationId string `json:"conversation_id"`
	ReplyWith      string `json:"reply_with"`
	UnReplyTo      string `json:"un_reply_to"`
	ReplyBy        string `json:"reply_by"`
}
