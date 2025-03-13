package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lyzsolar/ApiConsumer/citas/application"
	"github.com/lyzsolar/ApiConsumer/citas/domain"
	"github.com/lyzsolar/ApiConsumer/citas/domain/entities"
)

type CreateCitaController struct {
	useCase *application.CreateCita
	cita    domain.ICita
}

func NewCreateCitaController(useCase *application.CreateCita, cita domain.ICita) *CreateCitaController {
	return &CreateCitaController{useCase: useCase, cita: cita}
}

func (cs_a *CreateCitaController) Execute(c *gin.Context) {
	var cita entities.Cita
	if err := c.ShouldBindJSON(&cita); err != nil {
		c.JSON(400, gin.H{"error": "Datos inválidos"})
		return
	}

	err := cs_a.useCase.Execute(cita)
	if err != nil {
		c.JSON(500, gin.H{"error": "No se pudo crear la cita"})
		return
	}

	c.JSON(201, gin.H{"message": "Cita creada correctamente", "Cita": cita})
}
