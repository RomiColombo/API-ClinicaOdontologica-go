package odontologo

import "Colombo-Romina/internal/domain"

type IServiceOdontologo interface {
	GetAll() ([]*domain.Odontologo, error)
	GetByID(id int) (*domain.Odontologo, error)
	Create(odontologo domain.Odontologo) (*domain.Odontologo, error)
	Update(id int, odontologo domain.Odontologo) error
	UpdateSome(id int, odontologo domain.OdontologoPartial) error
	Delete(id int) error
}

type serviceOdontologo struct {
	r IOdontologoRepository
}

func NewService(r IOdontologoRepository) IServiceOdontologo {
	return &serviceOdontologo{r}
}

func (s *serviceOdontologo) GetAll() ([]*domain.Odontologo, error) {
	odontologos, err := s.r.GetAll()
	if err != nil {
		return nil, err
	}
	return odontologos, nil
}

func (s *serviceOdontologo) GetByID(id int) (*domain.Odontologo, error) {
	odontologo, err := s.r.GetByID(id)
	if err != nil {
		return nil, err
	}
	return odontologo, nil
}

func (s *serviceOdontologo) Create(odontologo domain.Odontologo) (*domain.Odontologo, error) {
	newOdontologo, err := s.r.Create(odontologo)
	if err != nil {
		return &domain.Odontologo{}, err
	}

	return newOdontologo, nil
}

func (s *serviceOdontologo) Update(id int, odontologo domain.Odontologo) error {
	err := s.r.Update(id, odontologo)
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceOdontologo) UpdateSome(id int, odontologo domain.OdontologoPartial) error {
	err := s.r.UpdateSome(id, odontologo)
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceOdontologo) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
