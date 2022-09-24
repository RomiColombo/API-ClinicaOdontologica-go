package paciente

import (
	"Colombo-Romina/internal/domain"
	configpaciente "Colombo-Romina/pkg/configPaciente"
	"errors"
)

type IPacienteRepository interface {
	GetAll() ([]*domain.Paciente, error)
	GetByID(id int) (*domain.Paciente, error)
	Create(paciente domain.Paciente) (*domain.Paciente, error)
	Update(id int, paciente domain.Paciente) error
	UpdateSome(id int, paciente domain.PacientePartial) error
	Delete(id int) error
}

type PacienteRepository struct {
	iPaciente configpaciente.PacienteInterface
}

// NewRepository crea un nuevo repositorio
func NewRepository(iPaciente configpaciente.PacienteInterface) IPacienteRepository {
	return &PacienteRepository{iPaciente}
}

func (r *PacienteRepository) GetAll() ([]*domain.Paciente, error) {
	pacientes, err := r.iPaciente.GetAll()
	if err != nil {
		return nil, err
	}
	return pacientes, nil
}

func (r *PacienteRepository) GetByID(id int) (*domain.Paciente, error) {
	if !r.iPaciente.Exist(id) {
		return nil, errors.New("el id ingresado no existe")
	}

	paciente, err := r.iPaciente.GetByID(id)
	if err != nil {
		return &domain.Paciente{}, errors.New("paciente no encontrado")
	}
	return paciente, nil
}

func (r *PacienteRepository) Create(paciente domain.Paciente) (*domain.Paciente, error) {

	if r.iPaciente.ExistDni(paciente.DNI) {
		return &domain.Paciente{}, errors.New("el dni ingresado ya se encuentra en nuestro sistema")
	}

	newpaciente, err := r.iPaciente.Create(paciente)
	if err != nil {
		return &domain.Paciente{}, errors.New("se produjo un error cargando el paciente")
	}
	
	return newpaciente, nil
}

func (r *PacienteRepository) Update(id int, paciente domain.Paciente) error {
	if !r.iPaciente.Exist(id) {
		return errors.New("el id ingresado no existe")
	}
	err := r.iPaciente.Update(id, paciente)
	if err != nil {
		return err
	}
	return nil
}

func (r *PacienteRepository) UpdateSome(id int, paciente domain.PacientePartial) error {
	if !r.iPaciente.Exist(id) {
		return errors.New("el id ingresado no existe")
	}
	err := r.iPaciente.UpdateSome(id, paciente)
	if err != nil {
		return err
	}
	return nil
}

func (r *PacienteRepository) Delete(id int) error {
	if !r.iPaciente.Exist(id) {
		return errors.New("el id ingresado no existe")
	}
	err := r.iPaciente.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
