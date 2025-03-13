package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lyzsolar/ApiConsumer/citas/infrastructure"
	"github.com/lyzsolar/ApiConsumer/citas/infrastructure/routes"
	"log"
)

func main() {
	dependencies := infrastructure.InitCita()
	defer dependencies.RabbitMQAdapter.Close()

	r := gin.Default()

	r.Use(func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		context.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Content-Type")

		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(204)
			return
		}
		context.Next()
	})

	routes.ConfigureRoutesCita(r, dependencies.CreateCitaController)

	if err := r.Run(":8088"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
