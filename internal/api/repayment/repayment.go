package repayment

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/okaaryanata/loan/internal/domain"
	"github.com/okaaryanata/loan/internal/helper"
)

func (c *Controller) MakePayment(ctx *gin.Context) {
	var args domain.MakePaymentRequest
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

	errx := c.repaymentService.MakePayment(ctx, &args)
	if errx != nil {
		helper.CreateResponseBody(ctx, &domain.ResponseBodyArgs{
			Status: errx.GetStatusCode(),
			Error:  errx,
		})
		return
	}

	helper.CreateResponseBody(ctx, &domain.ResponseBodyArgs{
		Status: http.StatusOK,
	})
}

func (c *Controller) PrintSchedule(ctx *gin.Context) {
	id := ctx.Param("userID")
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		helper.CreateResponseBody(ctx, &domain.ResponseBodyArgs{
			Status: http.StatusBadRequest,
			Error:  err,
		})
		return
	}

	var loanID int64
	id = ctx.Query("loanID")
	if id != "" {
		loanID, err = strconv.ParseInt(id, 10, 64)
		if err != nil {
			helper.CreateResponseBody(ctx, &domain.ResponseBodyArgs{
				Status: http.StatusBadRequest,
				Error:  err,
			})
			return
		}
	}

	schedule, errx := c.repaymentService.PrintSchedule(ctx, userID, loanID)
	if errx != nil {
		helper.CreateResponseBody(ctx, &domain.ResponseBodyArgs{
			Status: errx.GetStatusCode(),
			Error:  errx,
		})
		return
	}

	helper.CreateResponseBody(ctx, &domain.ResponseBodyArgs{
		Status: http.StatusOK,
		Data:   schedule,
	})
}
