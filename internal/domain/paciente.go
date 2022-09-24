package domain

type Paciente struct {
	Id        int    `json:"id"`
	Nombre    string `json:"nombre" binding:"required"`
	Apellido  string `json:"apellido" binding:"required"`
	Domicilio string `json:"domicilio" binding:"required"`
	DNI       int    `json:"dni" binding:"required"`
	FechaAlta string `json:"fechaAlta" binding:"required"`
}

type PacientePartial struct {
	Id        int    `json:"id"`
	Nombre    string `json:"nombre,omitempty"`
	Apellido  string `json:"apellido,omitempty"`
	Domicilio string `json:"domicilio,omitempty"`
	DNI       int    `json:"dni,omitempty"`
	FechaAlta string `json:"fechaAlta,omitempty"`
}
