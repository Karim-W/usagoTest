package domain

import (
	"context"
	"encoding/json"
	"fmt"
	"rabbitmqtest/pkg/infra/config"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
	"github.com/BetaLixT/usago"
	"go.uber.org/zap"
)

type RabbitMqRepository struct {
	client *aztables.Client
}

func (repo *RabbitMqRepository) Sub(logger *zap.Logger, topic string) string {
	_, err := repo.client.CreateTable(context.TODO(), nil)
	if err != nil {
		fmt.Sprintln("Table already Created")
	}
	manager := usago.NewChannelManager("amqp://guest:guest@localhost:55001/", logger)
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
		log := RabbitMq{
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
		_, err = repo.client.AddEntity(context.TODO(), marshalled, nil) // TODO: Check access policy, need Storage Table Data Contributor role
		if err != nil {
			panic(err)
		}
		if err != nil {
			panic(err)
		}
		logger.Info(
			"message read",
			zap.String("body", string(msg.Body)),
		)
	}()
	wg.Wait()
	return "SUCCESS"

}

func Ptr[T any](v T) *T {
	return &v
}

func NewRabbitMqRepo() *RabbitMqRepository {
	return &RabbitMqRepository{
		client: config.GetClient(),
	}
}
