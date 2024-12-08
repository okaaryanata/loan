package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/okaaryanata/loan/internal/domain"
	"github.com/okaaryanata/loan/internal/helper"
)

func (c *Controller) CreateUser(ctx *gin.Context) {
	var newUser domain.UserRequest
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
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

	user, errx := c.userService.CreateUser(ctx, &domain.UserRequest{
		Username:   newUser.Username,
		OperatedBy: newUser.OperatedBy,
	})
	if errx != nil {
		helper.CreateResponseBody(ctx, &domain.ResponseBodyArgs{
			Status: errx.GetStatusCode(),
			Error:  errx,
		})
		return
	}

	helper.CreateResponseBody(ctx, &domain.ResponseBodyArgs{
		Status: http.StatusCreated,
		Data:   user,
	})
}

func (c *Controller) GetUsers(ctx *gin.Context) {
	var listUsername []string
	username := ctx.Query("username")
	if username != "" {
		listUsername = append(listUsername, username)
	}

	users, errx := c.userService.GetUsers(ctx, listUsername...)
	if errx != nil {
		helper.CreateResponseBody(ctx, &domain.ResponseBodyArgs{
			Status: errx.GetStatusCode(),
			Error:  errx,
		})
	}

	helper.CreateResponseBody(ctx, &domain.ResponseBodyArgs{
		Status: http.StatusOK,
		Data:   users,
	})
}

func (c *Controller) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("userID")
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		helper.CreateResponseBody(ctx, &domain.ResponseBodyArgs{
			Status: http.StatusBadRequest,
			Error:  err,
		})
		return
	}

	user, errx := c.userService.GetUserByID(ctx, userID)
	if errx != nil {
		helper.CreateResponseBody(ctx, &domain.ResponseBodyArgs{
			Status: errx.GetStatusCode(),
			Error:  errx,
		})
		return
	}

	helper.CreateResponseBody(ctx, &domain.ResponseBodyArgs{
		Status: http.StatusOK,
		Data:   user,
	})
}
