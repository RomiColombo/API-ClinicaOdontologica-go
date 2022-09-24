package domain

type Odontologo struct {
	Id        int    `json:"id"`
	Nombre    string `json:"nombre" binding:"required"`
	Apellido  string `json:"apellido" binding:"required"`
	Matricula int    `json:"matricula" binding:"required"`
}

type OdontologoPartial struct {
	Id        int    `json:"id"`
	Nombre    string `json:"nombre,omitempty"`
	Apellido  string `json:"apellido,omitempty"`
	Matricula int    `json:"matricula,omitempty"`
}

