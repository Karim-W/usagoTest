package repos

// type RabbitMqRepository struct {
// 	client *aztables.Client
// }

// var _ domain.IRabbitMqRepository = (*RabbitMqRepository)(nil)

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

// func (repo *RabbitMqRepository) Sub(logger *zap.Logger, topic string) string {
// 	_, err := repo.client.CreateTable(context.TODO(), nil)
// 	if err != nil {
// 		fmt.Sprintln("Table already Created")
// 	}
// 	manager := usago.NewChannelManager("amqp://guest:guest@localhost:55001/", logger)
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
// 		return "FAIL"
// 	}
// 	consumer, _ := chnl.RegisterConsumer(
// 		topic,
// 		"",
// 		true,
// 		false,
// 		false,
// 		false,
// 		nil,
// 	)
// 	fmt.Println("CONSUMER REGISTERED")
// 	wg := sync.WaitGroup{}
// 	wg.Add(1)

// 	go func() {
// 		msg := <-consumer
// 		log := domain.RabbitMq{
// 			Entity: aztables.Entity{
// 				PartitionKey: "1",
// 				RowKey:       "1",
// 			},
// 			Log: string(msg.Body),
// 		}
// 		marshalled, err := json.Marshal(log)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println("THIS IS IT", msg.Body)
// 		_, err = repo.client.AddEntity(context.TODO(), marshalled, nil) // TODO: Check access policy, need Storage Table Data Contributor role
// 		if err != nil {
// 			panic(err)
// 		}
// 		if err != nil {
// 			panic(err)
// 		}
// 		logger.Info(
// 			"message read",
// 			zap.String("body", string(msg.Body)),
// 		)
// 	}()
// 	wg.Wait()
// 	return "SUCCESS"

// }

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

// func Ptr[T any](v T) *T {
// 	return &v
// }

// func NewRabbitMqRepo(client *aztables.Client) *RabbitMqRepository {
// 	return &RabbitMqRepository{
// 		client: client,
// 	}
// }
