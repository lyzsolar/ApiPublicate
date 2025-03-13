package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lyzsolar/ApiConsumer/citas/infrastructure/controllers"
)

func ConfigureRoutesCita(
	r *gin.Engine,
	createCitaController *controllers.CreateCitaController,

) {
	r.POST("/send-cita", createCitaController.Execute)
}
