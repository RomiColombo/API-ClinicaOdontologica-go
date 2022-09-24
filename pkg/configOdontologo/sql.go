package configodontologo

import (
	"Colombo-Romina/internal/domain"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type SqlStore struct {
	DB *sql.DB
}

func NewSqlStore(db *sql.DB) OdontologoInterface {
	return &SqlStore{
		DB: db,
	}
}

func (s *SqlStore) GetByID(id int) (*domain.Odontologo, error) {
	var odontologo domain.Odontologo
	query := "SELECT * FROM odontologos WHERE id = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&odontologo.Id, &odontologo.Nombre, &odontologo.Apellido, &odontologo.Matricula)
	if err != nil {
		return nil, err
	}

	return &odontologo, nil
}

func (s *SqlStore) GetAll() ([]*domain.Odontologo, error) {

	var odontogolos []*domain.Odontologo

	query := "SELECT * FROM odontologos;"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var odontologo domain.Odontologo
		err = rows.Scan(&odontologo.Id, &odontologo.Nombre, &odontologo.Apellido, &odontologo.Matricula)
		if err != nil {
			return nil, err
		}
		odontogolos = append(odontogolos, &odontologo)
	}

	return odontogolos, nil

}

func (s *SqlStore) Create(odontologo domain.Odontologo) (*domain.Odontologo, error) {

	query := "INSERT INTO odontologos (nombre, apellido, matricula) VALUES (?, ?, ?);"

	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(odontologo.Nombre, odontologo.Apellido, odontologo.Matricula)
	if err != nil {
		return nil, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return nil, err
	}

	lid, _ := res.LastInsertId()
	odontologo.Id = int(lid)
	return &odontologo, nil
}

func (s *SqlStore) Update(id int, odontologo domain.Odontologo) error {

	query := "UPDATE odontologos SET nombre = ?, apellido = ?, matricula = ? WHERE id = ?"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(odontologo.Nombre, odontologo.Apellido, odontologo.Matricula, id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (s *SqlStore) UpdateSome(id int, odontologo domain.OdontologoPartial) error {

	if !s.Exist(id) {
		return errors.New("The id does not exist")
	}

	var sb strings.Builder
	sb.WriteString("UPDATE odontologos SET")
	var attributes []interface{}

	if odontologo.Nombre != "" {
		sb.WriteString(" nombre = ?")
		attributes = append(attributes, odontologo.Nombre)
	}

	if odontologo.Apellido != "" {
		sb.WriteString(" apellido = ?")
		attributes = append(attributes, odontologo.Apellido)
	}

	if odontologo.Matricula != 0 {
		sb.WriteString(" matricula = ?")
		attributes = append(attributes, odontologo.Matricula)
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

	query := "SELECT id FROM odontologos WHERE id = ?;"
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

func (s *SqlStore) ExistMatricula(matricula int) bool {
	var exist bool
	var matriculaOd int

	query := "SELECT matricula FROM odontologos WHERE matricula = ?;"
	row := s.DB.QueryRow(query, matricula)
	err := row.Scan(&matriculaOd)
	if err != nil {
		return exist
	}
	if matriculaOd != 0 {
		exist = true
	}

	return exist
}

func (s *SqlStore) Delete(id int) error {

	query := "DELETE FROM odontologos WHERE id = ?"
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
