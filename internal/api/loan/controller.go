package loan

import (
	"github.com/gin-gonic/gin"
	"github.com/okaaryanata/loan/internal/service"
)

type (
	Controller struct {
		loanService *service.LoanService
	}
)

func NewLoanController(loanService *service.LoanService) *Controller {
	return &Controller{
		loanService: loanService,
	}
}

func (c *Controller) RegisterRoutes(router *gin.RouterGroup) {
	loanRouter := router.Group("/loan")
	{
		loanRouter.POST(``, c.CreateLoan)
		loanRouter.GET(`/:loanID/user/:userID`, c.GetLoanByIDAndUserID)
		loanRouter.GET(`/user/:userID`, c.GetLoansByUserID)
	}
}
