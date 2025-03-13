package domain

import "github.com/lyzsolar/ApiConsumer/citas/domain/entities"

type ICita interface {
	Send(cita entities.Cita) error
}
