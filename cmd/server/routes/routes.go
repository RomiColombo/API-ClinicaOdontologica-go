package routes

import (
	"Colombo-Romina/cmd/server/handler"
	"Colombo-Romina/docs"
	"Colombo-Romina/internal/odontologo"
	"Colombo-Romina/internal/paciente"
	"Colombo-Romina/internal/turno"
	configodontologo "Colombo-Romina/pkg/configOdontologo"
	configpaciente "Colombo-Romina/pkg/configPaciente"
	configturno "Colombo-Romina/pkg/configTurno"
	"Colombo-Romina/pkg/middleware"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Routes() *gin.Engine {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al intentar cargar archivo .env")
	}

	bddconection := os.Getenv("BDDCONECTION")

	bdd, err := sql.Open("mysql", bddconection)
	if err != nil {
		panic(err.Error())
	}

	errPing := bdd.Ping()
	if errPing != nil {
		panic(errPing)
	}

	odontologoSQL := configodontologo.NewSqlStore(bdd)
	odontologoRepo := odontologo.NewRepository(odontologoSQL)
	odontologoService := odontologo.NewService(odontologoRepo)
	odontologoHandler := handler.NewOdontologoHandler(odontologoService)

	pacienteSQL := configpaciente.NewSqlStore(bdd)
	pacienteRepo := paciente.NewRepository(pacienteSQL)
	pacienteService := paciente.NewService(pacienteRepo)
	pacienteHandler := handler.NewPacienteHandler(pacienteService)

	turnoSQL := configturno.NewSqlStore(bdd)
	turnoRepo := turno.NewRepository(turnoSQL, odontologoSQL, pacienteSQL)
	turnoService := turno.NewService(turnoRepo)
	turnoHandler := handler.NewTurnoHandler(turnoService)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	odontologoGroup := r.Group("/odontologos")
	{
		odontologoGroup.GET(":id", odontologoHandler.GetByID())
		odontologoGroup.GET("", odontologoHandler.GetAll())
		odontologoGroup.POST("", middleware.Authentication(), odontologoHandler.Create())
		odontologoGroup.PUT(":id", middleware.Authentication(), odontologoHandler.Update())
		odontologoGroup.PATCH(":id", middleware.Authentication(), odontologoHandler.UpdateSome())
		odontologoGroup.DELETE(":id", middleware.Authentication(), odontologoHandler.Delete())
	}

	pacienteGroup := r.Group("/pacientes")
	{
		pacienteGroup.GET(":id", pacienteHandler.GetByID())
		pacienteGroup.GET("", pacienteHandler.GetAll())
		pacienteGroup.POST("", middleware.Authentication(), pacienteHandler.Create())
		pacienteGroup.PUT(":id", middleware.Authentication(), pacienteHandler.Update())
		pacienteGroup.PATCH(":id", middleware.Authentication(), pacienteHandler.UpdateSome())
		pacienteGroup.DELETE(":id", middleware.Authentication(), pacienteHandler.Delete())
	}

	turnoGroup := r.Group("/turnos")
	{
		turnoGroup.GET(":id", turnoHandler.GetByID())
		turnoGroup.GET("", turnoHandler.GetAll())
		turnoGroup.GET("/dni/:dni", turnoHandler.GetByDni())
		turnoGroup.POST("", middleware.Authentication(), turnoHandler.Create())
		turnoGroup.POST("/partial", middleware.Authentication(), turnoHandler.CreatePartial())
		turnoGroup.PUT(":id", middleware.Authentication(), turnoHandler.Update())
		turnoGroup.PATCH(":id", middleware.Authentication(), turnoHandler.UpdateSome())
		turnoGroup.DELETE(":id", middleware.Authentication(), turnoHandler.Delete())
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Not found"})
	})

	return r
}
