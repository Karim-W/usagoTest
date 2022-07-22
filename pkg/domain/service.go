package domain

import (
	"github.com/BetaLixT/usago"
	"go.uber.org/zap"
)

type MqttService struct {
	logger *zap.Logger
}

func (svc *MqttService) SubSvc(prov IServiceProvider, topic string, channel *usago.ChannelContext) {
	prov.GetMqttRepo().Sub(svc.logger, topic, channel)
	// prov.GetMqttRepo().PubSub(svc.logger)
}

func NewMqttService(logger *zap.Logger) *MqttService {
	return &MqttService{
		logger: logger,
	}

}
