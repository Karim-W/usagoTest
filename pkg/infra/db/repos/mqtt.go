package repos

import (
	"rabbitmqtest/pkg/domain"

	"github.com/BetaLixT/usago"
	"go.uber.org/zap"
)

type MqttRepository struct {
}

var _ domain.IMqttRepository = (*MqttRepository)(nil)

// func (repo *MqttRepository) Pub(logger *zap.Logger, topic string, chnl *usago.ChannelContext) {
// 	// cnfrms, err := chnl.GetConfirmsChannel()
// 	// if err != nil {
// 	// 	logger.Error(
// 	// 		"failed to get confirms channel",
// 	// 		zap.Error(err),
// 	// 	)
// 	// } else {
// 	// 	fmt.Println(cnfrms)
// 	// 	logger.Info("CHECK")
// 	// }
// 	// go func() {
// 	// 	for i := 0; i < 10; i++ {
// 	// 		ack := <-cnfrms
// 	// 		logger.Info("confirm recieved")
// 	// 		if ack.Ack {
// 	// 		} else {
// 	// 			logger.Error("failed delivery")
// 	// 		}
// 	// 	}
// 	// }()
// 	body := "Hello World"
// 	_, err := chnl.Publish(
// 		"",
// 		topic,
// 		false, // mandatory
// 		false, // immediate
// 		amqp.Publishing{
// 			ContentType: "text/plain",
// 			Body:        []byte(body),
// 		},
// 	)
// 	for err != nil {
// 		_, err = chnl.Publish(
// 			"",
// 			topic,
// 			false, // mandatory
// 			false, // immediate
// 			amqp.Publishing{
// 				ContentType: "text/plain",
// 				Body:        []byte(body),
// 			},
// 		)
// 	}

// }

func (repo *MqttRepository) Sub(logger *zap.Logger, topic string, chnl *usago.ChannelContext) {

}

// func (repo *MqttRepository) PubSub(logger *zap.Logger, topic string) {
// 	manager := usago.NewChannelManager("amqp://guest:guest@localhost:55005/", logger)
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
// 	wg := sync.WaitGroup{}
// 	wg.Add(1)
// 	consumer, _ := chnl.RegisterConsumer(
// 		topic,
// 		"",
// 		true,
// 		false,
// 		false,
// 		false,
// 		nil,
// 	)
// 	go func() {
// 		defer wg.Done()
// 		msg := <-consumer
// 		fmt.Println("THIS IS THE MESSAGE", msg.Body)
// 		logger.Info(
// 			"message read",
// 			zap.String("body", string(msg.RoutingKey)),
// 		)
// 	}()
// 	body := "Hello World"
// 	_, err = chnl.Publish(
// 		"",
// 		topic,
// 		false, // mandatory
// 		false, // immediate
// 		amqp.Publishing{
// 			ContentType: "text/plain",
// 			Body:        []byte(body),
// 		},
// 	)
// 	for err != nil {
// 		_, err = chnl.Publish(
// 			"",
// 			topic,
// 			false, // mandatory
// 			false, // immediate
// 			amqp.Publishing{
// 				ContentType: "text/plain",
// 				Body:        []byte(body),
// 			},
// 		)
// 	}
// 	wg.Wait()

// }

func Ptr[T any](v T) *T {
	return &v
}

func NewMqttRepo() *MqttRepository {
	return &MqttRepository{}
}
