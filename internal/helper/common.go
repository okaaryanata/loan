package helper

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/okaaryanata/loan/internal/domain"
)

var (
	StatusMessageMapper = map[int]*domain.Message{
		http.StatusOK: {
			ID: "OK",
			EN: "OK",
		},
		http.StatusCreated: {
			ID: "Created",
			EN: "Created",
		},
		http.StatusBadRequest: {
			ID: "Bad Request",
			EN: "Bad Request",
		},
		http.StatusNotFound: {
			ID: "Not Found",
			EN: "Not Found",
		},
		http.StatusInternalServerError: {
			ID: "Internal Server Error",
			EN: "Internal Server Error",
		},
		http.StatusServiceUnavailable: {
			ID: "Service Unavailable",
			EN: "Service Unavailable",
		},
	}
)

func JakartaTime() (*time.Time, error) {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return nil, err
	}

	jakartaTime := time.Now().In(location)
	return &jakartaTime, nil
}

func CreateResponseBody(ctx *gin.Context, args *domain.ResponseBodyArgs) {
	resp := domain.ResponseBody{
		Message: *StatusMessageMapper[http.StatusInternalServerError],
	}
	if args == nil {
		ctx.JSON(http.StatusInternalServerError, resp)

		return
	}

	var code int
	if args.Status >= 200 {
		code = args.Status
	} else {
		code = http.StatusInternalServerError
	}
	args.Status = code

	if args.Message == nil {
		args.Message = StatusMessageMapper[args.Status]
	}

	resp.Message = *args.Message
	switch args.Error {
	case nil:
		if args.Data != nil {
			resp.Data = args.Data
		}
	default:
		resp.Error = args.Error.Error()
	}

	ctx.JSON(args.Status, resp)
}

func Chains(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}
