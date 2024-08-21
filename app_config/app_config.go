package app_config

import config "order-kafka-consumer/infrastructure/config_management"

var (
	CurrentEnvironment = config.From(config.Environment()).StringParam("ENV").Build()
)

func IsProductionEnvironment() bool {
	return CurrentEnvironment() == "production"
}
