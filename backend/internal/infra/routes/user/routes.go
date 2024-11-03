package userRoutes

import (
	"github.com/gin-gonic/gin"
	"gobarber/internal/schema"
	"gobarber/internal/service"
	"net/http"
)

type Router struct {
	service service.UserService
}

func Register(router *gin.RouterGroup) {
	r := Router{
		service: service.NewService(),
	}

	router.POST("", r.createUser)
}

func (r *Router) createUser(c *gin.Context) {
	var input schema.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	result, err := r.service.CreateUser(ctx, &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}
