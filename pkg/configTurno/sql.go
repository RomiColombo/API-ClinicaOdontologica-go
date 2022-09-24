package configturno

import (
	"Colombo-Romina/internal/domain"
	"database/sql"
	"errors"
	"strings"
)

type SqlStore struct {
	DB *sql.DB
}

func NewSqlStore(db *sql.DB) TurnoInterface {
	return &SqlStore{
		DB: db,
	}
}

func (s *SqlStore) GetByID(id int) (*domain.Turno, error) {

	var turno domain.Turno

	query := "SELECT * FROM turnos WHERE id = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&turno.Id, &turno.Paciente.DNI, &turno.Odontologo.Matricula, &turno.Fecha, &turno.Hora, &turno.Descripcion)
	if err != nil {
		return nil, err
	}

	queryPaciente := "SELECT * FROM pacientes WHERE dni = ?"
	row = s.DB.QueryRow(queryPaciente, turno.Paciente.DNI)
	err = row.Scan(&turno.Paciente.Id, &turno.Paciente.Nombre, &turno.Paciente.Apellido, &turno.Paciente.Domicilio, &turno.Paciente.DNI, &turno.Paciente.FechaAlta)
	if err != nil {
		return nil, err
	}

	queryOdontologo := "SELECT * FROM odontologos WHERE matricula = ?"
	row = s.DB.QueryRow(queryOdontologo, turno.Odontologo.Matricula)
	err = row.Scan(&turno.Odontologo.Id, &turno.Odontologo.Nombre, &turno.Odontologo.Apellido, &turno.Odontologo.Matricula)
	if err != nil {
		return nil, err
	}
	return &turno, nil
}

func (s *SqlStore) GetAll() ([]*domain.Turno, error) {

	var turnos []*domain.Turno

	query := "SELECT * FROM turnos;"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var turno domain.Turno
		err = rows.Scan(&turno.Id, &turno.Paciente.DNI, &turno.Odontologo.Matricula, &turno.Fecha, &turno.Hora, &turno.Descripcion)
		if err != nil {
			return nil, err
		}

		queryPaciente := "SELECT * FROM pacientes WHERE dni = ?"
		row := s.DB.QueryRow(queryPaciente, turno.Paciente.DNI)
		err = row.Scan(&turno.Paciente.Id, &turno.Paciente.Nombre, &turno.Paciente.Apellido, &turno.Paciente.Domicilio, &turno.Paciente.DNI, &turno.Paciente.FechaAlta)
		if err != nil {
			return nil, err
		}

		queryOdontologo := "SELECT * FROM odontologos WHERE matricula = ?"
		row = s.DB.QueryRow(queryOdontologo, turno.Odontologo.Matricula)
		err = row.Scan(&turno.Odontologo.Id, &turno.Odontologo.Nombre, &turno.Odontologo.Apellido, &turno.Odontologo.Matricula)
		if err != nil {
			return nil, err
		}
		turnos = append(turnos, &turno)
	}
	return turnos, nil
}

func (s *SqlStore) GetByDNI(dni int) ([]*domain.Turno, error) {

	var turnos []*domain.Turno

	query := "SELECT * FROM turnos WHERE paciente = ?;"
	rows, err := s.DB.Query(query,dni)
	if err != nil {
		return nil, err
	}	

	for rows.Next() {
		var turno domain.Turno
		err = rows.Scan(&turno.Id, &turno.Paciente.DNI, &turno.Odontologo.Matricula, &turno.Fecha, &turno.Hora, &turno.Descripcion)
		if err != nil {
			return nil, err
		}

		queryPaciente := "SELECT * FROM pacientes WHERE dni = ?"
		row := s.DB.QueryRow(queryPaciente, turno.Paciente.DNI)
		err = row.Scan(&turno.Paciente.Id, &turno.Paciente.Nombre, &turno.Paciente.Apellido, &turno.Paciente.Domicilio, &turno.Paciente.DNI, &turno.Paciente.FechaAlta)
		if err != nil {
			return nil, err
		}

		queryOdontologo := "SELECT * FROM odontologos WHERE matricula = ?"
		row = s.DB.QueryRow(queryOdontologo, turno.Odontologo.Matricula)
		err = row.Scan(&turno.Odontologo.Id, &turno.Odontologo.Nombre, &turno.Odontologo.Apellido, &turno.Odontologo.Matricula)
		if err != nil {
			return nil, err
		}

		turnos = append(turnos, &turno)
	}
	return turnos, nil
}

func (s *SqlStore) Create(turno domain.TurnoAdd) (*domain.Turno, error) {

	query := "INSERT INTO turnos (paciente, odontologo, fecha, hora, descripcion) VALUES (?, ?, ?, ?, ?);"
	
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	
	res, err := stmt.Exec(&turno.Paciente, &turno.Odontologo, &turno.Fecha, &turno.Hora, &turno.Descripcion)
	if err != nil {
		return nil, err
	}
	
	_, err = res.RowsAffected()
	if err != nil {
		return nil, err
	}
	
	var turnoF = domain.Turno{Fecha:turno.Fecha, Hora: turno.Hora, Descripcion: turno.Descripcion}
	lid, _ := res.LastInsertId()
	turnoF.Id = int(lid)
	
	queryPaciente := "SELECT * FROM pacientes WHERE dni = ?"
	row := s.DB.QueryRow(queryPaciente, turno.Paciente)
	err = row.Scan(&turnoF.Paciente.Id, &turnoF.Paciente.Nombre, &turnoF.Paciente.Apellido, &turnoF.Paciente.Domicilio, &turnoF.Paciente.DNI, &turnoF.Paciente.FechaAlta)
	if err != nil {
		return nil, err
	}

	queryOdontologo := "SELECT * FROM odontologos WHERE matricula = ?"
	row = s.DB.QueryRow(queryOdontologo, turno.Odontologo)
	err = row.Scan(&turnoF.Odontologo.Id, &turnoF.Odontologo.Nombre, &turnoF.Odontologo.Apellido, &turnoF.Odontologo.Matricula)
	if err != nil {
		return nil, err
	}

	return &turnoF, nil
}

func (s *SqlStore) Update(id int, turno domain.TurnoAdd) error {

	query := "UPDATE turnos SET paciente = ?, odontologo = ?, fecha = ?, hora = ?, descripcion = ? WHERE id = ?"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(turno.Paciente, turno.Odontologo, turno.Fecha, turno.Hora, turno.Descripcion, id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (s *SqlStore) UpdateSome(id int, turno domain.TurnoPartial) error {

	if !s.Exist(id) {
		return errors.New("The id does not exist")
	}

	var sb strings.Builder
	sb.WriteString("UPDATE turnos SET")
	var attributes []interface{}

	if turno.Paciente != 0 {
		sb.WriteString(" paciente = ?")
		attributes = append(attributes, turno.Paciente)
	}

	if turno.Odontologo != 0 {
		sb.WriteString(" odontologo = ?")
		attributes = append(attributes, turno.Odontologo)
	}

	if turno.Fecha != "" {
		sb.WriteString(" fecha = ?")
		attributes = append(attributes, turno.Fecha)
	}

	if turno.Hora != "" {
		sb.WriteString(" hora = ?")
		attributes = append(attributes, turno.Hora)
	}

	if turno.Descripcion != "" {
		sb.WriteString(" descripcion = ?")
		attributes = append(attributes, turno.Descripcion)
	}

	sb.WriteString(" WHERE id = ?")
	attributes = append(attributes, id)
	query := sb.String()

	_, err := s.DB.Exec(query, attributes...)
	if err != nil {
		return err
	}

	return nil
}

func (s *SqlStore) Exist(id int) bool {
	var exist bool
	var idCode int

	query := "SELECT id FROM turnos WHERE id = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&idCode)
	if err != nil {
		return exist
	}

	if idCode != 0 {
		exist = true
	}

	return exist
}

func (s *SqlStore) Delete(id int) error {

	query := "DELETE FROM turnos WHERE id = ?"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil

}