package order_created_projection

import (
	"context"
	karacakafka "github.com/mustafatheconqueror/karaca-kafka"
	"time"

	"github.com/spf13/cobra"
)

func Init(cmd *cobra.Command, args []string) error {

	//var logger = NewLogger()

	//var kafkaBrokers = app_config.KafkaBrokers()

	var consumerConfig = karacakafka.ConsumerConfig{
		Brokers:             []string{"localhost:9092", "localhost:9093", "localhost:9094"},
		AppName:             "order.created.consumer",
		Topics:              []string{"hepsiburada.oms.order.created.v1.main"},
		TopicDomainName:     "hepsiburada",
		TopicSubDomainName:  "oms",
		AutoOffsetResetType: karacakafka.AutoOffsetResetTypeLatest,
	}

	var producerConfig = karacakafka.ProducerConfig{
		Brokers:            []string{"localhost:9092", "localhost:9093", "localhost:9094"},
		AcknowledgeType:    karacakafka.AcknowledgeTypeAll,
		CompressionType:    karacakafka.CompressionTypeGzip,
		TopicDomainName:    "hepsiburada",
		TopicSubDomainName: "oms",
		DeliveryTimeoutMs:  karacakafka.DefaultDeliveryTimeoutMs * time.Second,
	}

	var readerConfig = karacakafka.ReaderConfig{
		Brokers:               []string{"localhost:9092", "localhost:9093", "localhost:9094"},
		GroupID:               "order.created.consumer",
		AutoOffsetResetType:   karacakafka.AutoOffsetResetTypeLatest,
		AllowAutoCreateTopics: false,
		EnableAutoCommit:      false,
		TopicDomainName:       "hepsiburada",
		TopicSubDomainName:    "oms",
		FetchMaxBytes:         6428800,
		SessionTimeout:        10 * time.Second,
		Debug:                 "consumer",
		ClientID:              "",
	}

	var karacaKafkaConfig = karacakafka.KaracaKafkaConfig{
		ConsumerConfig: consumerConfig,
		ReaderConfig:   readerConfig,
		ProducerConfig: producerConfig,
	}

	var kafkaMessageBus = karacakafka.NewKafkaConsumer(context.Background(), karacaKafkaConfig)

	orderCreatedConsumer := NewOrderCreatedConsumer(kafkaMessageBus)

	return kafkaMessageBus.StartConsume(orderCreatedConsumer.onConsume())
}
