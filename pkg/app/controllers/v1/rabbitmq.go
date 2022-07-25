package v1

import (
	"rabbitmqtest/pkg/domain"

	"go.uber.org/zap"
)

type RabbitMqController struct {
	svc    *domain.RabbitMqService
	logger *zap.Logger
	// channel *usago.ChannelContext
}

// func (ctrl *MqttController) cpub(ctx *gin.Context) {
// 	sp := ctx.MustGet("tx-context").(domain.IServiceProvider)
// 	topic := ctx.Param("topic")

// 	manager := usago.NewChannelManager("amqp://guest:guest@localhost:55005/", ctrl.logger)
// 	bldr := usago.NewChannelBuilder().WithQueue(
// 		topic,
// 		false,
// 		false,
// 		false,
// 		false,
// 		nil,
// 	).WithConfirms(true)

// 	chnl, err := manager.NewChannel(*bldr)
// 	if err != nil {
// 		fmt.Printf("failed to create channel")
// 		return
// 	}
// 	sp.GetMqttRepo().Pub(ctrl.logger, topic, chnl)
// 	ctx.JSON(200, "MESSAGE SENT")

// }

func (ctrl *RabbitMqController) Csub(topic string) string {
	return ctrl.svc.SubSvc(topic)
}

// func (ctrl *MqttController) RegisterRoutes(grp *gin.RouterGroup) {
// 	grp.POST("")
// }

// func (ctrl *MqttController) LMAOO() {
// }

func NewRabbitMqController(
	svc *domain.RabbitMqService,
	l *zap.Logger,
) *RabbitMqController {

	return &RabbitMqController{
		svc:    svc,
		logger: l,
	}
}
