package domain

type IServiceProvider interface {
	GetMqttRepo() IMqttRepository
}
