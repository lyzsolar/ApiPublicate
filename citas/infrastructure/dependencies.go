package infrastructure

import (
	"github.com/lyzsolar/ApiConsumer/citas/application"
	"github.com/lyzsolar/ApiConsumer/citas/application/repositories"
	"github.com/lyzsolar/ApiConsumer/citas/infrastructure/adapter"
	"github.com/lyzsolar/ApiConsumer/citas/infrastructure/controllers"
	"log"
)

type DependenciesCita struct {
	CreateCitaController *controllers.CreateCitaController
	RabbitMQAdapter      *adapter.RabbitMQAdapter
}

func InitCita() *DependenciesCita {
	rmqClient, err := adapter.NewRabbitMQAdapter()
	if err != nil {
		log.Fatalf("Error creating RabbitMQ client: %v", err)
	}

	messageService := repositories.NewServiceNotification(rmqClient)

	createCitaUseCase := application.NewCreateCita(rmqClient, messageService)

	return &DependenciesCita{
		CreateCitaController: controllers.NewCreateCitaController(createCitaUseCase, rmqClient),
		RabbitMQAdapter:      rmqClient,
	}
}
