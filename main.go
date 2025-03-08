package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lyzsolar/ApiConsumer/application/domain"
	"github.com/lyzsolar/ApiConsumer/application/service"
	"github.com/lyzsolar/ApiConsumer/config"
	"github.com/lyzsolar/ApiConsumer/infrastructure/repository"
	"log"
	"net/http"
)

func main() {
	// Crear servidor Gin
	r := gin.Default()

	// Conectar a la base de datos
	db := config.ConectarDB()
	defer db.Close()

	// Crear repositorio y servicio
	repo := repository.NewCitaRepository(db)
	citaService := service.NewCitaService(repo)

	// Definir el endpoint POST /citas
	r.POST("/citas", func(c *gin.Context) {
		var nuevaCita domain.Cita
		// Validar que la solicitud tenga los datos correctos
		if err := c.ShouldBindJSON(&nuevaCita); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
			return
		}

		// Crear la cita
		if err := citaService.CrearCita(nuevaCita); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error guardando la cita"})
			return
		}

		// Respuesta de éxito
		c.JSON(http.StatusOK, gin.H{"message": "Cita creada correctamente"})
	})

	// Iniciar el servidor en el puerto 8081
	log.Println("Servidor en el puerto 8081")
	r.Run(":8081")
}
