package handler

import (
	"Colombo-Romina/internal/domain"
	"Colombo-Romina/internal/turno"
	"Colombo-Romina/pkg/web"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type turnoHandler struct {
	s turno.IServiceTurno
}

func NewTurnoHandler(s turno.IServiceTurno) *turnoHandler {
	return &turnoHandler{
		s: s,
	}
}

// GetAll
// @Summary      Listar todos los turnos
// @Description  Listar todos los turnos
// @Tags         turno
// @Produce      json
// @Success      200 {object}  web.response
// @Router       /turnos [get]
func (h *turnoHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		turnos, err := h.s.GetAll()
		if err != nil {
			web.Failure(ctx, 400, errors.New("no se pudieron traer los turnos"))
		}
		web.Success(ctx, 200, turnos)
	}
}

// GetByID
// @Summary      Traer turno por DNI del turno
// @Description  Traer turno por DNI del turno
// @Tags         turno
// @Produce      json
// @Param        id path int true "id"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /turnos/{id} [get]
func (h *turnoHandler) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("id inválido"))
			return
		}
		turno, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(ctx, 404, errors.New("turno no encontrado"))
			return
		}
		web.Success(ctx, 200, turno)
	}
}

// GetByDNI
// @Summary      Traer turno por DNI del turno
// @Description  Traer turno por DNI del turno
// @Tags         turno
// @Produce      json
// @Param        id path int true "dni"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /turnos/dni/{dni} [get]
func (h *turnoHandler) GetByDni() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dniParam := ctx.Param("dni")
		dni, err := strconv.Atoi(dniParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("DNI inválido"))
			return
		}

		turnos, err := h.s.GetByDNI(dni)
		if err != nil {
			web.Failure(ctx, 404, errors.New("no se pudieron traer los turnos"))
		}
		web.Success(ctx, 200, turnos)
	}
}

// Post
// @Summary      Crear un turno con estructura de turno y odontologo
// @Description  Crear un turno con estructura de turno y odontologo
// @Tags         turno
// @Produce      json
// @Param        token header int true "token"
// @Param        body body domain.Turno true "Turno"
// @Success      201 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Router       /turnos [post]
func (h *turnoHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var turno domain.Turno
		err := ctx.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(ctx, 400, errors.New("json inválido"))
			return
		}
		newTurno, err := h.s.Create(turno)
		if err != nil {
			web.Failure(ctx, 400, errors.New("No se pudo crear el turno"))
			return
		}
		web.Success(ctx, 201, newTurno)
	}
}

// Post
// @Summary      Crear un turno con dni de turno y matricula de odontologo
// @Description  Crear un turno con dni de turno y matricula de odontologo
// @Tags         turno
// @Produce      json
// @Param        token header int true "token"
// @Param        body body domain.Turno true "Turno"
// @Success      201 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Router       /turnos/partial [post]
func (h *turnoHandler) CreatePartial() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var turno domain.TurnoAdd
	
		err := ctx.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(ctx, 400, errors.New("json inválido"))
			return
		}
		newTurno, err := h.s.CreatePartial(turno)
		if err != nil {
			web.Failure(ctx, 400, errors.New("No se pudo crear el turno"))
			return
		}
		web.Success(ctx, 201, newTurno)
	}
}

// Put
// @Summary      Modificar completamente un turno
// @Description  Modificar completamente un turno
// @Tags         turno
// @Produce      json
// @Param        token header int true "token"
// @Param        id path int true "id"
// @Param        body body domain.Turno true "Turno"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      409 {object}  web.errorResponse
// @Router       /turnos/{id} [put]
func (h *turnoHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("id inválido"))
			return
		}
		var turno domain.TurnoAdd
		err = ctx.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid json"))
			return
		}
		err = h.s.Update(id, turno)
		if err != nil {
			web.Failure(ctx, 409, errors.New("No se pudo modificar el turno"))
			return
		}
		turno.Id = id
		web.SuccessMsg(ctx, 200, "Turno modificado exitosamente")
	}
}

// Patch
// @Summary      Modificar parcialmente un turno
// @Description  Modificar parcialmente un turno
// @Tags         turno
// @Produce      json
// @Param        token header int true "token"
// @Param        id path int true "id"
// @Param        body body domain.Turno true "Turno"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Router       /turnos/{id} [patch]
func (h *turnoHandler) UpdateSome() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("id inválido"))
			return
		}
		var turnoPartial domain.TurnoPartial
		err = ctx.ShouldBindJSON(&turnoPartial)
		if err != nil {
			web.Failure(ctx, 400, errors.New("json invalido"))
			return
		}
		err = h.s.UpdateSome(id, turnoPartial)
		if err != nil {
			web.Failure(ctx, 409, errors.New("No se pudo modificar el turno"))
			return
		}
		turnoPartial.Id = id
		web.SuccessMsg(ctx, 200, "Turno modificado exitosamente")
	}
}

// Delete
// @Summary      Borrar un turno
// @Description  Borrar un turno
// @Tags         turno
// @Produce      json
// @Param        token header int true "token"
// @Param        id path int true "id"
// @Success      204 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /turnos/{id} [delete]
func (h *turnoHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("id inválido"))
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			web.Failure(ctx, 404, errors.New("No se pudo eliminar el turno"))
			return
		}
		web.Success(ctx, 204, "Se eliminó correctamente el turno")
	}
}
