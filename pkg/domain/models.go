package domain

import (
	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

type RabbitMq struct {
	aztables.Entity
	Log string
}
