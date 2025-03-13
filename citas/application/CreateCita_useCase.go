package application

import (
	"github.com/lyzsolar/ApiConsumer/citas/application/repositories"
	"github.com/lyzsolar/ApiConsumer/citas/domain"
	"github.com/lyzsolar/ApiConsumer/citas/domain/entities"
	"log"
)

type CreateCita struct {
	citaRepo    domain.ICita
	serviceCita repositories.IMessageService
}

func NewCreateCita(citaRepo domain.ICita, serviceCita repositories.IMessageService) *CreateCita {
	return &CreateCita{
		citaRepo:    citaRepo,
		serviceCita: serviceCita,
	}
}

func (c *CreateCita) Execute(cita entities.Cita) error {
	err := c.citaRepo.Send(cita)
	if err != nil {
		return err
	}

	err = c.serviceCita.PublishEvent("CitaCreated", cita)
	if err != nil {
		log.Printf("Error notificando sobre la cita creada: %v", err)
		return err
	}

	return nil
}
