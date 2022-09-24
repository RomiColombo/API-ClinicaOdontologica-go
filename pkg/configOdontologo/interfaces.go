package configodontologo

import (
	"Colombo-Romina/internal/domain"
)

type OdontologoInterface interface{
	GetAll() ([]*domain.Odontologo, error)
	GetByID(id int) (*domain.Odontologo, error)
	Create(odontologo domain.Odontologo) (*domain.Odontologo, error)
	Update(id int, odontologo domain.Odontologo) error
	UpdateSome(id int, odontologo domain.OdontologoPartial) error
	Delete(id int) error
	Exist(id int) bool
	ExistMatricula(matricula int) bool
}
