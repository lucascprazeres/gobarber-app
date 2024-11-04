package userRoutes

import (
	"github.com/gin-gonic/gin"
	"gobarber/internal/domain/validation"
	"gobarber/internal/schema"
	"gobarber/internal/service"
	"gobarber/pkg/errorwrapper"
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
		errorwrapper.SendError(c, validation.CustomValidationError(err))
		return
	}

	ctx := c.Request.Context()
	result, err := r.service.CreateUser(ctx, &input)
	if err != nil {
		errorwrapper.SendError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}
