package domain

import (
	"go.uber.org/zap"
)

type IRabbitMqRepository interface {
	Sub(logger *zap.Logger, topic string)
}
