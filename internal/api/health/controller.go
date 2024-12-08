package health

import "github.com/gin-gonic/gin"

type (
	Controller struct{}
)

func NewHealthController() *Controller {
	return &Controller{}
}

func (c *Controller) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/", c.Home)
	router.GET("/health", c.HealthCheck)
}
