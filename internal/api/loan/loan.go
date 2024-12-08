package loan

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/okaaryanata/loan/internal/domain"
	"github.com/okaaryanata/loan/internal/helper"
)

func (c *Controller) CreateLoan(ctx *gin.Context) {
	var args domain.LoanRequest
	if err := ctx.ShouldBindJSON(&args); err != nil {
		helper.CreateResponseBody(ctx, &domain.ResponseBodyArgs{
			Status: http.StatusBadRequest,
			Message: &domain.Message{
				ID: "Bad Request",
				EN: "Bad Request",
			},
			Error: err,
		})
		return
	}

	loan, errx := c.loanService.CreateLoan(ctx, &args)
	if errx != nil {
		helper.CreateResponseBody(ctx, &domain.ResponseBodyArgs{
			Status: errx.GetStatusCode(),
			Error:  errx,
		})
		return
	}

	helper.CreateResponseBody(ctx, &domain.ResponseBodyArgs{
		Status: http.StatusOK,
		Data:   loan,
	})
}

func (c *Controller) GetLoanByIDAndUserID(ctx *gin.Context) {
	id := ctx.Param("loanID")
	loanID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		helper.CreateResponseBody(ctx, &domain.ResponseBodyArgs{
			Status: http.StatusBadRequest,
			Error:  err,
		})
		return
	}

	id = ctx.Param("userID")
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		helper.CreateResponseBody(ctx, &domain.ResponseBodyArgs{
			Status: http.StatusBadRequest,
			Error:  err,
		})
		return
	}

	loan, errx := c.loanService.GetLoanByIDandUserID(ctx, loanID, userID)
	if errx != nil {
		helper.CreateResponseBody(ctx, &domain.ResponseBodyArgs{
			Status: errx.GetStatusCode(),
			Error:  errx,
		})
		return
	}

	helper.CreateResponseBody(ctx, &domain.ResponseBodyArgs{
		Status: http.StatusOK,
		Data:   loan,
	})
}

func (c *Controller) GetLoansByUserID(ctx *gin.Context) {
	id := ctx.Param("userID")
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		helper.CreateResponseBody(ctx, &domain.ResponseBodyArgs{
			Status: http.StatusBadRequest,
			Error:  err,
		})
		return
	}

	loans, errx := c.loanService.GetLoansByUserID(ctx, userID)
	if errx != nil {
		helper.CreateResponseBody(ctx, &domain.ResponseBodyArgs{
			Status: http.StatusInternalServerError,
			Error:  errx,
		})
		return
	}

	helper.CreateResponseBody(ctx, &domain.ResponseBodyArgs{
		Status: http.StatusOK,
		Data:   loans,
	})
}
