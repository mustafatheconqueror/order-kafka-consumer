package cmd

import (
	"github.com/spf13/cobra"
	"order-kafka-consumer/consumers/order_created_projection"
)

func init() {
	RootCmd.AddCommand(&cobra.Command{
		Use:   "order_created_projection_consumer",
		Short: "Running Order Created Projection",
		Long:  `"Running Order Created Projection`,
		RunE:  order_created_projection.Init,
	})
}
