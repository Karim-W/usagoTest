package domain

import (
	"go.uber.org/zap"
)

type RabbitMqService struct {
	repo   *RabbitMqRepository
	logger *zap.Logger
}

func (svc *RabbitMqService) SubSvc(topic string) string {
	return svc.repo.Sub(svc.logger, topic)
	// svc.repo.Sub(svc.logger, topic)
	// prov.GetMqttRepo().PubSub(svc.logger)
}

func NewRabbitMqService(logger *zap.Logger, repo *RabbitMqRepository) *RabbitMqService {
	return &RabbitMqService{
		logger: logger,
		repo:   repo,
	}

}
