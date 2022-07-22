package rabbitmqtest

import (
	"fmt"
	v1 "rabbitmqtest/pkg/app/controllers/v1"
	"rabbitmqtest/pkg/domain"
	"testing"

	"github.com/BetaLixT/usago"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestNewChannelManager(t *testing.T) {
	logger, _ := zap.NewProduction()
	manager := usago.NewChannelManager("amqp://guest:guest@localhost:55005/", logger)
	bldr := usago.NewChannelBuilder().WithQueue(
		"Notification",
		false,
		false,
		false,
		false,
		nil,
	).WithConfirms(true)
	chnl, err := manager.NewChannel(*bldr)
	if err != nil {
		fmt.Printf("failed to create channel")
		return
	}
	// consume
	fmt.Println("HELLOOO")
	body := "HI"
	_, err = chnl.Publish(
		"",
		"Notification",
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	for err != nil {
		_, err = chnl.Publish(
			"",
			"Notification",
			false, // mandatory
			false, // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			},
		)
	}
	fmt.Println("MESSAGE SENT")
	svc := domain.NewRabbitMqService(logger)
	res := v1.NewRabbitMqController(svc, logger).Csub("Notification")
	assert.Equal(t, res, "SUCCESS")
}
