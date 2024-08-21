package order_created_projection

import (
	"github.com/spf13/cobra"
	. "order-kafka-consumer/infrastructure/log"
)

func Init(cmd *cobra.Command, args []string) error {

	var logger = NewLogger()

	var messageKafkaBus = kafka.NewKafkaClient(kafka.ClientConfig{
		Brokers: KafkaBrokers(),
		AppName: "Com.Donationcounter.Projection",
	}, logger)

	return messageKafkaBus.RunConsumers()
}
