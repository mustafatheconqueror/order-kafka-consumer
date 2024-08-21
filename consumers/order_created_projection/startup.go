package order_created_projection

import (
	"context"
	"fmt"
	"github.com/mustafatheconqueror/karaca-kafka/config"
	"github.com/mustafatheconqueror/karaca-kafka/constants"
	"github.com/mustafatheconqueror/karaca-kafka/consumer"
	"github.com/spf13/cobra"
	"order-kafka-consumer/app_config"
	. "order-kafka-consumer/infrastructure/log"
)

func Init(cmd *cobra.Command, args []string) error {

	var logger = NewLogger()

	var kafkaBrokers = app_config.KafkaBrokers()

	var consumerConfig = config.ConsumerConfig{
		Brokers:             kafkaBrokers,
		AppName:             "order.created.consumer",
		Topics:              []string{"hepsiburada.oms.order.created.v1.main"},
		AutoOffsetResetType: constants.AutoOffsetResetTypeEarliest,
	}

	var kafkaMessageBus = consumer.NewKafkaConsumer(context.Background(), consumerConfig, kafkaBrokers)

	orderCreatedConsumer := NewOrderCreatedConsumer(&kafkaMessageBus)

	logger.Info(context.Background(), fmt.Sprintf("Kafka Order Created Consumer Started"))

	return kafkaMessageBus.StartConsume(orderCreatedConsumer.onConsume())
}
