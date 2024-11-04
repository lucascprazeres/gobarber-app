package routes

import (
	"github.com/gin-gonic/gin"
	"gobarber/internal/infra/http/routes/user"
)

func Register(app *gin.Engine) {
	v1 := app.Group("/v1")

	v1.GET("/health-check", healthCheck)

	userRoutes.Register(v1.Group("/users"))
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"ok": true})
}
