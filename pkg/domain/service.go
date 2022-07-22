package domain

import (
	"github.com/BetaLixT/usago"
	"go.uber.org/zap"
)

type RabbitMqService struct {
	logger *zap.Logger
}

func (svc *RabbitMqService) SubSvc(prov IServiceProvider, topic string, channel *usago.ChannelContext) {
	prov.GetRabbitMqRepo().Sub(svc.logger, topic, channel)
	// prov.GetMqttRepo().PubSub(svc.logger)
}

func NewRabbitMqService(logger *zap.Logger) *RabbitMqService {
	return &RabbitMqService{
		logger: logger,
	}

}
