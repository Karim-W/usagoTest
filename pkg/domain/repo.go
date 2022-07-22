package domain

import (
	"github.com/BetaLixT/usago"
	"go.uber.org/zap"
)

type IRabbitMqRepository interface {
	Sub(logger *zap.Logger, topic string, chnl *usago.ChannelContext)
}
