package domain

type Turno struct {
	Id          int        `json:"id"`
	Paciente    Paciente   `json:"paciente" binding:"required"`
	Odontologo  Odontologo `json:"odontologo" binding:"required"`
	Fecha       string     `json:"fecha" binding:"required"`
	Hora        string     `json:"hora" binding:"required"`
	Descripcion string     `json:"descripcion,omitempty"`
}

type TurnoAdd struct {
	Id          int    `json:"id"`
	Paciente    int    `json:"paciente" binding:"required"`
	Odontologo  int    `json:"odontologo" binding:"required"`
	Fecha       string `json:"fecha" binding:"required"`
	Hora        string `json:"hora" binding:"required"`
	Descripcion string `json:"descripcion,omitempty"`
}

type TurnoPartial struct {
	Id          int    `json:"id"`
	Paciente    int    `json:"paciente"`
	Odontologo  int    `json:"odontologo"`
	Fecha       string `json:"fecha"`
	Hora        string `json:"hora"`
	Descripcion string `json:"descripcion,omitempty"`
}
