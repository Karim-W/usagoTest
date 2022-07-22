package v1

import (
	"fmt"
	"rabbitmqtest/pkg/domain"

	"github.com/BetaLixT/usago"
	"go.uber.org/zap"
)

type MqttController struct {
	svc    *domain.MqttService
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

func (ctrl *MqttController) Csub(topic string) string {

	manager := usago.NewChannelManager("amqp://guest:guest@localhost:55005/", ctrl.logger)
	bldr := usago.NewChannelBuilder().WithQueue(
		topic,
		false,
		false,
		false,
		false,
		nil,
	).WithConfirms(true)
	chnl, err := manager.NewChannel(*bldr)
	if err != nil {
		fmt.Printf("failed to create channel")
		return "FAIL"
	}
	consumer, _ := chnl.RegisterConsumer(
		topic,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	fmt.Println("CONSUMER REGISTERED")
	go func() {
		msg := <-consumer
		fmt.Println("THIS IS THE MESSAGE", msg.Body)
		ctrl.logger.Info(
			"message read",
			zap.String("body", string(msg.RoutingKey)),
		)
	}()
	return "SUCCESS"
}

// func (ctrl *MqttController) RegisterRoutes(grp *gin.RouterGroup) {
// 	grp.POST("")
// }

// func (ctrl *MqttController) LMAOO() {
// }

func NewMqttController(
	svc *domain.MqttService,
	l *zap.Logger,
) *MqttController {

	return &MqttController{
		svc:    svc,
		logger: l,
	}
}
