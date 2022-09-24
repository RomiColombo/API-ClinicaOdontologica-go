package configpaciente

import (
	"Colombo-Romina/internal/domain"
	"database/sql"
	"errors"
	"strings"
)

type SqlStore struct {
	DB *sql.DB
}

func NewSqlStore(db *sql.DB) PacienteInterface {
	return &SqlStore{
		DB: db,
	}
}

func (s *SqlStore) GetByID(id int) (*domain.Paciente, error) {
	var paciente domain.Paciente
	query := "SELECT * FROM pacientes WHERE id = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&paciente.Id, &paciente.Nombre, &paciente.Apellido, &paciente.Domicilio, &paciente.DNI, &paciente.FechaAlta)
	if err != nil {
		return nil, err
	}

	return &paciente, nil
}

func (s *SqlStore) GetAll() ([]*domain.Paciente, error) {

	var pacientes []*domain.Paciente

	query := "SELECT * FROM pacientes;"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var paciente domain.Paciente
		err = rows.Scan(&paciente.Id, &paciente.Nombre, &paciente.Apellido, &paciente.Domicilio, &paciente.DNI, &paciente.FechaAlta)
		if err != nil {
			return nil, err
		}
		pacientes = append(pacientes, &paciente)
	}

	return pacientes, nil

}

func (s *SqlStore) Create(paciente domain.Paciente) (*domain.Paciente, error) {

	query := "INSERT INTO pacientes (nombre, apellido, domicilio, dni, fechaAlta) VALUES (?, ?, ?, ?, ?);"

	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(&paciente.Nombre, &paciente.Apellido, &paciente.Domicilio, &paciente.DNI, &paciente.FechaAlta)
	if err != nil {
		return nil, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return nil, err
	}

	lid, _ := res.LastInsertId()
	paciente.Id = int(lid)
	return &paciente, nil
}

func (s *SqlStore) Update(id int, paciente domain.Paciente) error {

	query := "UPDATE pacientes SET nombre = ?, apellido = ?, domicilio = ?, dni = ?, fechaAlta = ? WHERE id = ?"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(paciente.Nombre, paciente.Apellido, paciente.Domicilio, paciente.DNI, paciente.FechaAlta, id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (s *SqlStore) UpdateSome(id int, paciente domain.PacientePartial) error {

	if !s.Exist(id) {
		return errors.New("The id does not exist")
	}

	var sb strings.Builder
	sb.WriteString("UPDATE pacientes SET")
	var attributes []interface{}

	if paciente.Nombre != "" {
		sb.WriteString(" nombre = ?")
		attributes = append(attributes, paciente.Nombre)
	}

	if paciente.Apellido != "" {
		sb.WriteString(" apellido = ?")
		attributes = append(attributes, paciente.Apellido)
	}

	if paciente.Domicilio != "" {
		sb.WriteString(" domicilio = ?")
		attributes = append(attributes, paciente.Domicilio)
	}

	if paciente.DNI != 0 {
		sb.WriteString(" dni = ?")
		attributes = append(attributes, paciente.DNI)
	}

	if paciente.FechaAlta != "" {
		sb.WriteString(" fechaAlta = ?")
		attributes = append(attributes, paciente.FechaAlta)
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

	query := "SELECT id FROM pacientes WHERE id = ?;"
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

func (s *SqlStore) ExistDni(dni int) bool {
	var exist bool
	var dniPac int

	query := "SELECT dni FROM pacientes WHERE dni = ?;"
	row := s.DB.QueryRow(query, dni)
	err := row.Scan(&dniPac)
	if err != nil {
		return exist
	}
	if dniPac != 0 {
		exist = true
	}

	return exist
}

func (s *SqlStore) Delete(id int) error {

	query := "DELETE FROM pacientes WHERE id = ?"
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
