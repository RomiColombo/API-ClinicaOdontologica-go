package odontologo

import (
	"Colombo-Romina/internal/domain"
	configodontologo "Colombo-Romina/pkg/configOdontologo"
	"errors"
)

type IOdontologoRepository interface {
	GetAll() ([]*domain.Odontologo, error)
	GetByID(id int) (*domain.Odontologo, error)
	Create(odontologo domain.Odontologo) (*domain.Odontologo, error)
	Update(id int, odontologo domain.Odontologo) error
	UpdateSome(id int, odontologo domain.OdontologoPartial) error
	Delete(id int) error
}

type OdontologoRepository struct {
	iOdontologo configodontologo.OdontologoInterface
}

// NewRepository crea un nuevo repositorio
func NewRepository(iOdontologo configodontologo.OdontologoInterface) IOdontologoRepository {
	return &OdontologoRepository{iOdontologo}
}

func (r *OdontologoRepository) GetAll() ([]*domain.Odontologo, error) {
	odontologos, err := r.iOdontologo.GetAll()
	if err != nil {
		return nil, err
	}
	return odontologos, nil
}

func (r *OdontologoRepository) GetByID(id int) (*domain.Odontologo, error) {
	if !r.iOdontologo.Exist(id) {
		return nil, errors.New("el id ingresado no existe")
	}

	odontologo, err := r.iOdontologo.GetByID(id)
	if err != nil {
		return &domain.Odontologo{}, errors.New("odontologo not found")
	}
	return odontologo, nil
}

func (r *OdontologoRepository) Create(odontologo domain.Odontologo) (*domain.Odontologo, error) {

	if r.iOdontologo.ExistMatricula(odontologo.Matricula) {
		return &domain.Odontologo{}, errors.New("la matricula ingresada ya se encuentra en nuestro sistema")
	}

	newOdontologo, err := r.iOdontologo.Create(odontologo)
	if err != nil {
		return &domain.Odontologo{}, errors.New("se produjo un error cargando el odontologo")
	}
	
	return newOdontologo, nil
}

func (r *OdontologoRepository) Update(id int, odontologo domain.Odontologo) error {
	if !r.iOdontologo.Exist(id) {
		return errors.New("el id ingresado no existe")
	}
	err := r.iOdontologo.Update(id, odontologo)
	if err != nil {
		return err
	}
	return nil
}

func (r *OdontologoRepository) UpdateSome(id int, odontologo domain.OdontologoPartial) error {
	if !r.iOdontologo.Exist(id) {
		return errors.New("el id ingresado no existe")
	}
	err := r.iOdontologo.UpdateSome(id, odontologo)
	if err != nil {
		return err
	}
	return nil
}

func (r *OdontologoRepository) Delete(id int) error {
	if !r.iOdontologo.Exist(id) {
		return errors.New("el id ingresado no existe")
	}
	err := r.iOdontologo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
