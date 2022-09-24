package middleware

import (
	"Colombo-Romina/pkg/web"
	"errors"
	"os"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("Para continuar debe ingresar un token"))
			c.Abort()
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("Token inválido"))
			c.Abort()
			return
		}
		c.Next()
	}
}
