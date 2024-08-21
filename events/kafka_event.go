package events

import (
	"errors"
	"github.com/mustafatheconqueror/karaca-kafka/kafka_message"
	uuid "github.com/satori/go.uuid"
	"strings"
	"time"
)

type KafkaEvent struct {
	CorrelationId string `json:"correlationId"`
	Message       map[string]interface{}
	Headers       KafkaHeaders `json:"headers"`
}

type KafkaHeaders struct {
	TimeStamp    time.Time `json:"timeStamp"`
	UserName     string    `json:"userName"`
	IdentityName string    `json:"identityName"`
	IdentityType string    `json:"identityType"`
	Type         string    `json:"type"`
	Version      int       `json:"version"`
}

func (self *KafkaEvent) CheckValidGuidCorrelationId() {

	if self.CorrelationId == "00000000-0000-0000-0000-000000000000" || self.CorrelationId == "" {
		guid := uuid.NewV4()
		self.CorrelationId = guid.String()
	}
}

func (self *KafkaEvent) GetEventType() (string, error) {
	if len(self.Headers.Type) == 0 {
		return "", errors.New("MissingMessageType")
	}

	return strings.Split(self.Headers.Type, ",")[0], nil
}

func MapHeaders(headers kafka_message.KafkaHeaders) KafkaHeaders {

	var (
		kafkaHeaders KafkaHeaders
	)

	kafkaHeaders.Version = headers.Version
	kafkaHeaders.UserName = headers.UserName
	kafkaHeaders.TimeStamp = headers.TimeStamp
	kafkaHeaders.IdentityType = headers.IdentityType
	kafkaHeaders.IdentityName = headers.IdentityName
	kafkaHeaders.Type = headers.MessageType

	return kafkaHeaders
}
