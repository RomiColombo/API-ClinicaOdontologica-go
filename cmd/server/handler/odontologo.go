package handler

import (
	"Colombo-Romina/internal/domain"
	"Colombo-Romina/internal/odontologo"
	"Colombo-Romina/pkg/web"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type odontologoHandler struct {
	s odontologo.IServiceOdontologo
}

func NewOdontologoHandler(s odontologo.IServiceOdontologo) *odontologoHandler {
	return &odontologoHandler{
		s: s,
	}
}

// GetAll
// @Summary      Listar todos los odontologos
// @Description  Listar todos los odontologos
// @Tags         odontologo
// @Produce      json
// @Success      200 {object}  web.response
// @Router       /odontologos [get]
func (h *odontologoHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		odontologos, err := h.s.GetAll()
		if err != nil {
			web.Failure(ctx, 400, errors.New("No se pudieron traer los odontologos"))
		}
		web.Success(ctx, 200, odontologos)
	}
}

// GetByID
// @Summary      Listar odontologo por ID
// @Description  Listar odontologo por ID
// @Tags         odontologo
// @Produce      json
// @Param        id path int true "id"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /odontologos/{id} [get]
func (h *odontologoHandler) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("id inválido"))
			return
		}
		odontologo, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(ctx, 404, errors.New("Odontólogo no encontrado"))
			return
		}
		web.Success(ctx, 200, odontologo)
	}
}

// Post
// @Summary      Crear un odontologo
// @Description  Crear un odontologo
// @Tags         odontologo
// @Produce      json
// @Param        token header int true "token"
// @Param        body body domain.Odontologo true "Odontologo"
// @Success      201 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Router       /odontologos [post]
func (h *odontologoHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var odontologo domain.Odontologo
		err := ctx.ShouldBindJSON(&odontologo)
		if err != nil {
			web.Failure(ctx, 400, errors.New("json inválido"))
			return
		}
		newOdontologo, err := h.s.Create(odontologo)
		if err != nil {
			web.Failure(ctx, 400, errors.New("No se pudo crear el odontologo"))
			return
		}
		web.Success(ctx, 201, newOdontologo)
	}
}

// Put
// @Summary      Modificar completamente un odontologo
// @Description  Modificar completamente un odontologo
// @Tags         odontologo
// @Produce      json
// @Param        token header int true "token"
// @Param        id path int true "id"
// @Param        body body domain.Odontologo true "Odontologo"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Router       /odontologos/{id} [put]
func (h *odontologoHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("id inválido"))
			return
		}
		var odontologo domain.Odontologo
		err = ctx.ShouldBindJSON(&odontologo)
		if err != nil {
			web.Failure(ctx, 400, errors.New("Json inválido"))
			return
		}
		err = h.s.Update(id, odontologo)
		if err != nil {
			web.Failure(ctx, 409, errors.New("No se pudo modificar el odontologo"))
			return
		}
		odontologo.Id = id
		web.SuccessMsg(ctx, 200, "Odontologo modificado exitosamente")
	}
}

// Patch
// @Summary      Modificar parcialmente un odontologo
// @Description  Modificar parcialmente un odontologo
// @Tags         odontologo
// @Produce      json
// @Param        token header int true "token"
// @Param        id path int true "id"
// @Param        body body domain.Odontologo true "Odontologo"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Router       /odontologos/{id} [patch]
func (h *odontologoHandler) UpdateSome() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("id inválido"))
			return
		}
		var odontologoPartial domain.OdontologoPartial
		err = ctx.ShouldBindJSON(&odontologoPartial)
		if err != nil {
			web.Failure(ctx, 400, errors.New("json invalido"))
			return
		}
		err = h.s.UpdateSome(id, odontologoPartial)
		if err != nil {
			web.Failure(ctx, 409, errors.New("No se pudo modificar el odontologo"))
			return
		}
		odontologoPartial.Id = id
		web.Success(ctx, 200, "Odontologo modificado exitosamente")
	}
}

// Delete
// @Summary      Borrar un odontologo
// @Description  Borrar un odontologo
// @Tags         odontologo
// @Produce      json
// @Param        token header int true "token"
// @Param        id path int true "id"
// @Success      204 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /odontologos/{id} [delete]
func (h *odontologoHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("id inválido"))
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			web.Failure(ctx, 404, errors.New("No se pudo eliminar el odontologo"))
			return
		}
		web.SuccessMsg(ctx, 204, "Se eliminó correctamente el odontólogo")
	}
}
