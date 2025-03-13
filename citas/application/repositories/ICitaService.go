package repositories

import "github.com/lyzsolar/ApiConsumer/citas/domain/entities"

type IMessageService interface {
	PublishEvent(eventType string, citas entities.Cita) error
}
