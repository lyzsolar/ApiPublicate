package service

import (
	"github.com/lyzsolar/ApiConsumer/application/domain"
	"github.com/lyzsolar/ApiConsumer/infrastructure/repository"
)

// CitaService es la interfaz que define los métodos para gestionar citas.
type CitaService interface {
	CrearCita(cita domain.Cita) error
}

// citaServiceImpl es la implementación concreta de la interfaz CitaService.
type citaServiceImpl struct {
	repo repository.CitaRepository
}

// NewCitaService crea una nueva instancia de citaServiceImpl.
func NewCitaService(repo repository.CitaRepository) CitaService {
	return &citaServiceImpl{repo: repo}
}

// CrearCita crea una nueva cita a través del repositorio.
func (s *citaServiceImpl) CrearCita(cita domain.Cita) error {
	return s.repo.GuardarCita(cita)
}
