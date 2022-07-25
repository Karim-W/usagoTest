package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"rabbitmqtest/pkg/domain"
	"rabbitmqtest/pkg/infra/config"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
	"github.com/BetaLixT/usago"
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
	client := config.GetClient()
	_, err := client.CreateTable(context.TODO(), nil)
	if err != nil {
		fmt.Sprintln("Table already Created")
	}
	manager := usago.NewChannelManager("amqp://guest:guest@localhost:55001/", ctrl.logger)
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
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		msg := <-consumer
		log := domain.RabbitMq{
			Entity: aztables.Entity{
				PartitionKey: "1",
				RowKey:       "1",
			},
			Log: string(msg.Body),
		}
		marshalled, err := json.Marshal(log)
		if err != nil {
			panic(err)
		}
		fmt.Println("THIS IS IT", msg.Body)
		_, err = client.AddEntity(context.TODO(), marshalled, nil) // TODO: Check access policy, need Storage Table Data Contributor role
		if err != nil {
			panic(err)
		}
		if err != nil {
			panic(err)
		}
		ctrl.logger.Info(
			"message read",
			zap.String("body", string(msg.Body)),
		)
	}()
	wg.Wait()
	return "SUCCESS"
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
