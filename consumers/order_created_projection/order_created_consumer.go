package order_created_projection

import (
	"encoding/json"
	"github.com/mustafatheconqueror/karaca-kafka/consumer"
	"github.com/mustafatheconqueror/karaca-kafka/kafka_message"
	"order-kafka-consumer/events"
	"order-kafka-consumer/infrastructure/errors"
)

type OrderCreatedConsumer struct {
	messageBus *consumer.Consumer
}

type Order struct {
	OrderNumber string
}

func NewOrderCreatedConsumer(messageBus *consumer.Consumer) *OrderCreatedConsumer {
	return &OrderCreatedConsumer{
		messageBus: messageBus,
	}
}

func (ol *OrderCreatedConsumer) onConsume() func(message kafka_message.KafkaMessage) error {
	return func(message kafka_message.KafkaMessage) error {

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

		return nil
	}
}
