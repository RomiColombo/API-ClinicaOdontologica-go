package turno

import "Colombo-Romina/internal/domain"

type IServiceTurno interface {
	GetAll() ([]*domain.Turno, error) 
	GetByID(id int) (*domain.Turno, error) 
	GetByDNI(dni int) ([]*domain.Turno, error) 
	Create(turno domain.Turno) (*domain.Turno, error) 
	CreatePartial(turno domain.TurnoAdd) (*domain.Turno, error) 
	Update(id int, turno domain.TurnoAdd) error 
	UpdateSome(id int, turno domain.TurnoPartial) error
	Delete(id int) error
}

type serviceTurno struct {
	r ITurnoRepository
}

func NewService(r ITurnoRepository) IServiceTurno {
	return &serviceTurno{r}
}

func (s *serviceTurno) GetAll() ([]*domain.Turno, error) {
	turnos, err := s.r.GetAll()
	if err != nil {
		return nil, err
	}
	return turnos, nil
}

func (s *serviceTurno) GetByID(id int) (*domain.Turno, error) {
	turno, err := s.r.GetByID(id)
	if err != nil {
		return nil, err
	}
	return turno, nil
}

func (s *serviceTurno) GetByDNI(dni int) ([]*domain.Turno, error) {
	turnos, err := s.r.GetByDNI(dni)
	if err != nil {
		return nil, err
	}
	return turnos, nil
}

func (s *serviceTurno) Create(turno domain.Turno) (*domain.Turno, error) {
	newturno, err := s.r.Create(turno)
	if err != nil {
		return &domain.Turno{}, err
	}

	return newturno, nil
}

func (s *serviceTurno) CreatePartial(turno domain.TurnoAdd) (*domain.Turno, error) {
	newturno, err := s.r.CreatePartial(turno)
	if err != nil {
		return &domain.Turno{}, err
	}
	return newturno, nil
}

func (s *serviceTurno) Update(id int, turno domain.TurnoAdd) error {
	err := s.r.Update(id, turno)
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceTurno) UpdateSome(id int, turno domain.TurnoPartial) error {
	err := s.r.UpdateSome(id, turno)
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceTurno) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
