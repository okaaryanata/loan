package user

import (
	"github.com/gin-gonic/gin"
	"github.com/okaaryanata/loan/internal/service"
)

type (
	Controller struct {
		userService *service.UserService
	}
)

func NewUserController(userService *service.UserService) *Controller {
	return &Controller{
		userService: userService,
	}
}

func (c *Controller) RegisterRoutes(router *gin.RouterGroup) {
	UserRouter := router.Group("/user")
	{
		UserRouter.POST(``, c.CreateUser)
		UserRouter.GET(``, c.GetUsers)
		UserRouter.GET(`/:userID`, c.GetUserByID)
	}
}
