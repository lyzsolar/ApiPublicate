package domain

import "time"

type Cita struct {
	ID        int       `json:"id"`
	Fecha     time.Time `json:"fecha"`
	MascotaID int       `json:"mascota_id"`
}
