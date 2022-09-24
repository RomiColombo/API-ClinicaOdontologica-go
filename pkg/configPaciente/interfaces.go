package configpaciente

import "Colombo-Romina/internal/domain"

type PacienteInterface interface {
	GetAll() ([]*domain.Paciente, error)
	GetByID(id int) (*domain.Paciente, error)
	Create(paciente domain.Paciente) (*domain.Paciente, error)
	Update(id int, paciente domain.Paciente) error
	UpdateSome(id int, paciente domain.PacientePartial) error
	Delete(id int) error
	Exist(id int) bool
	ExistDni(dni int) bool
}
