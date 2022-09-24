package configturno

import "Colombo-Romina/internal/domain"

type TurnoInterface interface {
	GetAll() ([]*domain.Turno, error)
	GetByID(id int) (*domain.Turno, error)
	Create(Turno domain.TurnoAdd) (*domain.Turno, error)
	Update(id int, Turno domain.TurnoAdd) error
	UpdateSome(id int, Turno domain.TurnoPartial) error
	Delete(id int) error
	Exist(id int) bool
	GetByDNI(dni int) ([]*domain.Turno, error)
}
