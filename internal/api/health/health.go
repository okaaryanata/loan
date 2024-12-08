package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Home(ctx *gin.Context) {
	responseBody := map[string]interface{}{
		"hello": "svaha-loan is here",
	}
	ctx.IndentedJSON(http.StatusOK, responseBody)
}

func (c *Controller) HealthCheck(ctx *gin.Context) {
	responseBody := map[string]interface{}{
		"status": "ready",
	}
	ctx.IndentedJSON(http.StatusOK, responseBody)
}
