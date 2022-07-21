package main

import (
	"fmt"
	"sync"

	"github.com/BetaLixT/usago"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	manager := usago.NewChannelManager("amqp://guest:guest@localhost:55001/", logger)
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

	consumer, err := chnl.RegisterConsumer(
		"Notification",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	// publish
	body := "Hello World!"
	messageCount := 10
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < messageCount; i++ {
			msg := <-consumer
			logger.Info(
				"message read",
				zap.String("body", string(msg.RoutingKey)),
			)
		}
	}()

	for i := 0; i < messageCount; i++ {
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
	}
	wg.Wait()
}
