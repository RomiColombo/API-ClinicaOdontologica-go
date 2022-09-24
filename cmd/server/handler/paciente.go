package handler

import (
	"Colombo-Romina/internal/domain"
	"Colombo-Romina/internal/paciente"
	"Colombo-Romina/pkg/web"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type pacienteHandler struct {
	s paciente.IServicePaciente
}

func NewPacienteHandler(s paciente.IServicePaciente) *pacienteHandler {
	return &pacienteHandler{
		s: s,
	}
}

// GetAll
// @Summary      Listar todos los pacientes
// @Description  Listar todos los pacientes
// @Tags         paciente
// @Produce      json
// @Success      200 {object}  web.response
// @Router       /pacientes [get]
func (h *pacienteHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pacientes, err := h.s.GetAll()
		if err != nil {
			web.Failure(ctx, 400, errors.New("no se pudieron traer los pacientes"))
		}
		web.Success(ctx, 200, pacientes)
	}
}

// GetByID
// @Summary      Listar paciente por ID
// @Description  Listar paciente por ID
// @Tags         paciente
// @Produce      json
// @Param        id path int true "id"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /pacientes/{id} [get]
func (h *pacienteHandler) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("id inválido"))
			return
		}
		paciente, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(ctx, 404, errors.New("paciente no encontrado"))
			return
		}
		web.Success(ctx, 200, paciente)
	}
}

// Post
// @Summary      Crear un paciente
// @Description  Crear un paciente
// @Tags         paciente
// @Produce      json
// @Param        token header int true "token"
// @Param        body body domain.Paciente true "Paciente"
// @Success      201 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Router       /pacientes [post]
func (h *pacienteHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paciente domain.Paciente
		err := ctx.ShouldBindJSON(&paciente)
		if err != nil {
			web.Failure(ctx, 400, errors.New("json inválido"))
			return
		}
		newpaciente, err := h.s.Create(paciente)
		if err != nil {
			web.Failure(ctx, 400, errors.New("No se pudo crear el paciente"))
			return
		}
		web.Success(ctx, 201, newpaciente)
	}
}

// Put
// @Summary      Modificar completamente un paciente
// @Description  Modificar completamente un paciente
// @Tags         paciente
// @Produce      json
// @Param        token header int true "token"
// @Param        id path int true "id"
// @Param        body body domain.Paciente true "Paciente"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      409 {object}  web.errorResponse
// @Router       /pacientes/{id} [put]
func (h *pacienteHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("id inválido"))
			return
		}

		var paciente domain.Paciente
		err = ctx.ShouldBindJSON(&paciente)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid json"))
			return
		}
		err = h.s.Update(id, paciente)
		if err != nil {
			web.Failure(ctx, 409, errors.New("No se pudo modificar el paciente"))
			return
		}
		paciente.Id = id
		web.SuccessMsg(ctx, 200, "Paciente modificado exitosamente")
	}
}

// Patch
// @Summary      Modificar parcialmente un paciente
// @Description  Modificar parcialmente un paciente
// @Tags         paciente
// @Produce      json
// @Param        token header int true "token"
// @Param        id path int true "id"
// @Param        body body domain.Paciente true "Paciente"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Router       /pacientes/{id} [patch]
func (h *pacienteHandler) UpdateSome() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("id inválido"))
			return
		}
		var pacientePartial domain.PacientePartial
		err = ctx.ShouldBindJSON(&pacientePartial)
		if err != nil {
			web.Failure(ctx, 400, errors.New("json invalido"))
			return
		}
		err = h.s.UpdateSome(id, pacientePartial)
		if err != nil {
			web.Failure(ctx, 409, errors.New("No se pudo modificar el paciente"))
			return
		}
		pacientePartial.Id=id
		web.SuccessMsg(ctx, 200, "Paciente modificado exitosamente")
	}
}

// Delete
// @Summary      Borrar un paciente
// @Description  Borrar un paciente
// @Tags         paciente
// @Produce      json
// @Param        token header int true "token"
// @Param        id path int true "id"
// @Success      204 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /pacientes/{id} [delete]
func (h *pacienteHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("id inválido"))
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			web.Failure(ctx, 404, errors.New("No se pudo eliminar el paciente"))
			return
		}
		web.Success(ctx, 204, "Se eliminó correctamente el paciente")
	}
}
