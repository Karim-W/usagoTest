package domain

import (
	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

type RabbitMq struct {
	aztables.Entity
	Level      string
	Ts         string
	Caller     string
	Msg        string
	Body       string
	RoutingKey string
}
