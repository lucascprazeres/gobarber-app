package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gobarber/internal/infra/http/routes"
	"gobarber/internal/settings"
	"net/http"
)

func NewServer() *http.Server {
	app := gin.Default()

	routes.Register(app)

	addr := fmt.Sprintf("%s:%d", settings.GetEnvs().AppHost, settings.GetEnvs().AppPort)
	srv := &http.Server{
		Addr:    addr,
		Handler: app,
	}

	return srv
}
