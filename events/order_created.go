package events

type Created struct {
	OrderNumber string `json:"OrderNumber,omitempty"`
	MessageId   string `json:"MessageId,omitempty"`
}
