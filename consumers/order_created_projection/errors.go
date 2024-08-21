package order_created_projection

import "order-kafka-consumer/infrastructure/errors"

const prefix = "ORDER_CREATED_PROJECTION_CONSUMER"

var (
	ConvertEventError = errors.DefineError(prefix, 1, "Convert Event Error")
	SinopError        = errors.DefineError(prefix, 2, "Sinop Error")
)
