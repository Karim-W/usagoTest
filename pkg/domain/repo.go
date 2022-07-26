package domain

import (
	"context"
	"encoding/json"
	"fmt"
	"rabbitmqtest/pkg/infra/config"
	"runtime"
	"strconv"

	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
	"github.com/BetaLixT/usago"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

type RabbitMqRepository struct {
	client *aztables.Client
}

func (repo *RabbitMqRepository) Channel(logger *zap.Logger, topic string) {
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
		_, src, line, _ := runtime.Caller(0)
		repo.Log("", "", src, line, logger, "failed to create channel")
		fmt.Printf("failed to create channel")
		return
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
	msg := <-consumer
	_, src, line, _ := runtime.Caller(0)
	repo.Log(string(msg.Body), msg.RoutingKey, src, line, logger, "message read")
}

func (repo *RabbitMqRepository) Log(body string, routingKey string, src string, line int, logger *zap.Logger, message string) {
	count := 1

	observedZapCore, observedLogs := observer.New(zap.InfoLevel)
	observedLogger := zap.New(observedZapCore)
	observedLogger.Info(
		message,
		zap.String("body", string(body)),
	)
	logbody := observedLogs.All()[0]

	log := RabbitMq{
		Entity: aztables.Entity{
			PartitionKey: "1",
			RowKey:       strconv.Itoa(count),
		},
		Level:      logbody.Level.String(),
		Ts:         strconv.Itoa(int(logbody.Time.Unix())),
		Caller:     src + ":" + strconv.Itoa(line),
		Msg:        logbody.Message,
		Body:       string(body),
		RoutingKey: string(routingKey),
	}
	fmt.Println(log)
	marshalled, err := json.Marshal(log)
	if err != nil {
		panic(err)
	}
	_, err = repo.client.AddEntity(context.TODO(), marshalled, nil) // TODO: Check access policy, need Storage Table Data Contributor role
	if err != nil {
		panic(err)
	}

}

func (repo *RabbitMqRepository) Table(logger *zap.Logger, topic string) {
	_, err := repo.client.CreateTable(context.TODO(), nil)
	if err != nil {
		fmt.Println("Table already Created")
	} else {
		fmt.Println("Table created")
	}
}

func (repo *RabbitMqRepository) Sub(logger *zap.Logger, topic string) string {
	repo.Table(logger, topic)
	repo.Channel(logger, topic)
	// for err2 != nil {
	// 	count = count + 1
	// 	log.RowKey = strconv.Itoa(count)
	// 	marshalled, err := json.Marshal(log)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	_, err2 = repo.client.AddEntity(context.TODO(), marshalled, nil) // TODO: Check access policy, need Storage Table Data Contributor role
	// }
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
