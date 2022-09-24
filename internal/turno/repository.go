package turno

import (
	"Colombo-Romina/internal/domain"
	configodontologo "Colombo-Romina/pkg/configOdontologo"
	configpaciente "Colombo-Romina/pkg/configPaciente"
	configturno "Colombo-Romina/pkg/configTurno"
	"errors"
)

type ITurnoRepository interface {
	GetAll() ([]*domain.Turno, error)
	GetByID(id int) (*domain.Turno, error)
	GetByDNI(dni int) ([]*domain.Turno, error)
	Create(turno domain.Turno) (*domain.Turno, error)
	CreatePartial(turno domain.TurnoAdd) (*domain.Turno, error)
	Update(id int, turno domain.TurnoAdd) error
	UpdateSome(id int, turno domain.TurnoPartial) error
	Delete(id int) error
}

type TurnoRepository struct {
	iTurno configturno.TurnoInterface
	iOdontologo configodontologo.OdontologoInterface
	iPaciente configpaciente.PacienteInterface
}


// NewRepository crea un nuevo repositorio
func NewRepository(iTurno configturno.TurnoInterface, iOdontologo configodontologo.OdontologoInterface, iPaciente configpaciente.PacienteInterface ) ITurnoRepository {
	return &TurnoRepository{iTurno, iOdontologo, iPaciente}
}

func (r *TurnoRepository) GetAll() ([]*domain.Turno, error) {
	turnos, err := r.iTurno.GetAll()
	if err != nil {
		return nil, err
	}
	return turnos, nil
}

func (r *TurnoRepository) GetByID(id int) (*domain.Turno, error) {
	if !r.iTurno.Exist(id) {
		return nil, errors.New("el id ingresado no existe")
	}
	
	paciente, err := r.iTurno.GetByID(id)
	if err != nil {
		return &domain.Turno{}, errors.New("paciente no encontradp")
	}
	return paciente, nil
}

func (r *TurnoRepository) GetByDNI(dni int) ([]*domain.Turno, error) {
	turnos, err := r.iTurno.GetByDNI(dni)
	if err != nil {
		return nil, err
	}
	return turnos, nil
}

func (r *TurnoRepository) Create(turno domain.Turno) (*domain.Turno, error) {
	
	paciente := turno.Paciente
	odontologo := turno.Odontologo

	if !r.iOdontologo.ExistMatricula(odontologo.Matricula) {
		r.iOdontologo.Create(odontologo)
	}

	if !r.iPaciente.ExistDni(paciente.DNI) {
		r.iPaciente.Create(paciente)
	}

	var newTurno = domain.TurnoAdd{Paciente: paciente.DNI, Odontologo: odontologo.Matricula, Fecha: turno.Fecha, Hora: turno.Hora, Descripcion: turno.Descripcion}

	turnoCreated, err := r.iTurno.Create(newTurno)
	if err != nil {
		return &domain.Turno{}, errors.New("se produjo un error cargando el turno")
	}

	return turnoCreated, nil
}

func (r *TurnoRepository) CreatePartial(turno domain.TurnoAdd) (*domain.Turno, error) {

	if !r.iOdontologo.ExistMatricula(turno.Odontologo) {
		return nil, errors.New("la matricula ingresada no existe")
	}

	if !r.iPaciente.ExistDni(turno.Paciente) {
		return nil, errors.New("el dni ingresado no existe")
	}

	newTurno, err := r.iTurno.Create(turno)
	if err != nil {
		return &domain.Turno{}, errors.New("se produjo un error cargando el turno")
	}
	return newTurno, nil
}

func (r *TurnoRepository) Update(id int, paciente domain.TurnoAdd) error {
	if !r.iTurno.Exist(id) {
		return errors.New("el id ingresado no existe")
	}
	err := r.iTurno.Update(id, paciente)
	if err != nil {
		return err
	}
	return nil
}

func (r *TurnoRepository) UpdateSome(id int, paciente domain.TurnoPartial) error {
	if !r.iTurno.Exist(id) {
		return errors.New("el id ingresado no existe")
	}
	err := r.iTurno.UpdateSome(id, paciente)
	if err != nil {
		return err
	}
	return nil
}

func (r *TurnoRepository) Delete(id int) error {
	if !r.iTurno.Exist(id) {
		return errors.New("el id ingresado no existe")
	}
	err := r.iTurno.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

