package domain

type IServiceProvider interface {
	GetRabbitMqRepo() IRabbitMqRepository
}
