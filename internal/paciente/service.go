package paciente

import "Colombo-Romina/internal/domain"

type IServicePaciente interface {
	GetAll() ([]*domain.Paciente, error)
	GetByID(id int) (*domain.Paciente, error)
	Create(paciente domain.Paciente) (*domain.Paciente, error)
	Update(id int, paciente domain.Paciente) error
	UpdateSome(id int, paciente domain.PacientePartial) error
	Delete(id int) error
}

type servicePaciente struct {
	r IPacienteRepository
}

func NewService(r IPacienteRepository) IServicePaciente {
	return &servicePaciente{r}
}

func (s *servicePaciente) GetAll() ([]*domain.Paciente, error) {
	pacientes, err := s.r.GetAll()
	if err != nil {
		return nil, err
	}
	return pacientes, nil
}

func (s *servicePaciente) GetByID(id int) (*domain.Paciente, error) {
	paciente, err := s.r.GetByID(id)
	if err != nil {
		return nil, err
	}
	return paciente, nil
}

func (s *servicePaciente) Create(paciente domain.Paciente) (*domain.Paciente, error) {
	newpaciente, err := s.r.Create(paciente)
	if err != nil {
		return &domain.Paciente{}, err
	}

	return newpaciente, nil
}

func (s *servicePaciente) Update(id int, paciente domain.Paciente) error {
	err := s.r.Update(id, paciente)
	if err != nil {
		return err
	}
	return nil
}

func (s *servicePaciente) UpdateSome(id int, paciente domain.PacientePartial) error {
	err := s.r.UpdateSome(id, paciente)
	if err != nil {
		return err
	}
	return nil
}

func (s *servicePaciente) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
