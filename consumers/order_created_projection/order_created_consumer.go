package order_created_projection

import (
	"encoding/json"
	karacaKafka "github.com/mustafatheconqueror/karaca-kafka"
	"order-kafka-consumer/events"
	"order-kafka-consumer/infrastructure/errors"
)

type OrderCreatedConsumer struct {
	messageBus karacaKafka.KaracaConsumer
}

type Order struct {
	OrderNumber string
}

func NewOrderCreatedConsumer(messageBus karacaKafka.KaracaConsumer) *OrderCreatedConsumer {
	return &OrderCreatedConsumer{
		messageBus: messageBus,
	}
}

func (ol *OrderCreatedConsumer) onConsume() func(message karacaKafka.KaracaMessage) error {
	return func(message karacaKafka.KaracaMessage) error {

		var (
			kafkaEvent      events.KafkaEvent
			err             error
			messageAsString []byte
		)

		kafkaEvent.CorrelationId = message.CorrelationId
		kafkaEvent.Headers = events.MapHeaders(message.Headers)

		err = json.Unmarshal(message.Payload, &kafkaEvent.Message)
		if err != nil {
			return errors.NewWithCause(ConvertEventError, err)
		}

		if kafkaEvent.Headers.IdentityType == "Something" {
			return nil
		}

		kafkaEvent.CheckValidGuidCorrelationId()

		messageAsString, err = json.Marshal(kafkaEvent.Message)
		if err != nil {
			return err
		}

		var (
			orderCreatedEvent *events.Created
		)

		err = json.Unmarshal(messageAsString, &orderCreatedEvent)
		if err != nil {
			return errors.NewWithCause(ConvertEventError, err, orderCreatedEvent)
		}

		if orderCreatedEvent.OrderNumber == "57" {
			return errors.NewWithCause(SinopError, err)
		}
		if orderCreatedEvent.OrderNumber == "58" {
			message.Headers.IsRetryable = "false"
			message.CorrelationId = "5757"
			message.Partition = 2
			return errors.NewWithCause(SinopError, err)
		}

		return nil
	}
}
