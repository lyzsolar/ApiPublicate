package ports

import "github.com/lyzsolar/ApiConsumer/application/domain"

type CitaService interface {
	CrearCita(cita domain.Cita) error
}

type CitaRepository interface {
	GuardarCita(cita domain.Cita) error
}
