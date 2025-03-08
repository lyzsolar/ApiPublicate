package repository

import (
	"database/sql"
	"github.com/lyzsolar/ApiConsumer/application/domain"
	"log"
)

// CitaRepository es la interfaz que define el comportamiento de los repositorios de citas.
type CitaRepository interface {
	GuardarCita(cita domain.Cita) error
}

// CitaRepositoryImpl es la implementación concreta de la interfaz CitaRepository.
type CitaRepositoryImpl struct {
	db *sql.DB
}

// NewCitaRepository crea una nueva instancia de CitaRepositoryImpl.
func NewCitaRepository(db *sql.DB) CitaRepository {
	return &CitaRepositoryImpl{db: db}
}

// GuardarCita guarda una nueva cita en la base de datos.
func (r *CitaRepositoryImpl) GuardarCita(cita domain.Cita) error {
	fecha := cita.Fecha.Format("2006-01-02 15:04:05") // Convertir fecha al formato adecuado
	_, err := r.db.Exec("INSERT INTO citas (fecha, mascota_id) VALUES (?, ?)", fecha, cita.MascotaID)
	if err != nil {
		log.Println("Error guardando la cita:", err)
		return err
	}
	return nil
}
