package domain

import (
	"go.uber.org/zap"
)

type RabbitMqService struct {
	logger *zap.Logger
}

func (svc *RabbitMqService) SubSvc(topic string) {
	// svc.repo.Sub(svc.logger, topic)
	// prov.GetMqttRepo().PubSub(svc.logger)
}

func NewRabbitMqService(logger *zap.Logger) *RabbitMqService {
	return &RabbitMqService{
		logger: logger,
	}

}
