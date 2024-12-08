package repayment

import (
	"github.com/gin-gonic/gin"
	"github.com/okaaryanata/loan/internal/service"
)

type (
	Controller struct {
		repaymentService *service.RepaymentService
	}
)

func NewRepaymentController(repaymentService *service.RepaymentService) *Controller {
	return &Controller{
		repaymentService: repaymentService,
	}
}

func (c *Controller) RegisterRoutes(router *gin.RouterGroup) {
	RepaymentRouter := router.Group("/repayment")
	{
		RepaymentRouter.POST(``, c.MakePayment)
		RepaymentRouter.GET(`/schedule/user/:userID`, c.PrintSchedule)
	}
}
