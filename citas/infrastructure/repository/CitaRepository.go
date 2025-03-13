package repository

import (
	"github.com/lyzsolar/ApiConsumer/citas/domain/entities"
	"github.com/lyzsolar/ApiConsumer/citas/infrastructure/adapter"
)

type CitaRepository struct {
	rmqClient *adapter.RabbitMQAdapter
}

func NewNotificationRepository(rmqClient *adapter.RabbitMQAdapter) *CitaRepository {
	return &CitaRepository{rmqClient: rmqClient}
}

func (nr *CitaRepository) Send(cita entities.Cita) error {
	return nr.rmqClient.Send(cita)
}
