package repositories

import (
	"github.com/lyzsolar/ApiConsumer/citas/domain/entities"
	"log"
)

type ServiceNotification struct {
	imageService IMessageService
}

// NewServiceNotification creates a new ServiceNotification instance
func NewServiceNotification(imageService IMessageService) *ServiceNotification {
	return &ServiceNotification{imageService: imageService}
}

func (sn *ServiceNotification) PublishEvent(eventType string, cita entities.Cita) error {
	log.Println("Publishing event:", eventType)

	err := sn.imageService.PublishEvent(eventType, cita)
	if err != nil {
		log.Printf("Error publishing event: %v", err)
		return err
	}
	return nil
}
