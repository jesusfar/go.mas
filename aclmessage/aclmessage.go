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
	SUBSCRIBRE
)

// Message struct based on FIPA ACL message
type Message struct {
	performative   Performative
	sender         string
	receiver       string
	replyTo        string
	content        string
	language       string
	encoding       string
	ontology       string
	protocol       string
	conversationId string
	replyWith      string
	unReplyTo      string
	replyBy        string
}
