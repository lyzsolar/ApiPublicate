package entities

type Cita struct {
	ID      int
	Cita    string
	Message string
}

func NewNotification(id int, cita string) *Cita {
	return &Cita{
		ID:      id,
		Cita:    cita,
		Message: "La cita se ha registrado: " + cita,
	}
}
