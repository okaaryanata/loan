package domain

type (
	TypeStatusCode      string
	TypeStatusCodeValue int
)

var (
	MainRoute                 = "/svaha-loan"
	StatusCode TypeStatusCode = "statusCode"
)

type (
	ResponseBody struct {
		Message Message     `json:"message"`
		Data    interface{} `json:"data"`
		Error   interface{} `json:"error"`
	}

	Message struct {
		ID string `json:"id"`
		EN string `json:"en"`
	}

	ResponseBodyArgs struct {
		Status  int
		Message *Message
		Data    interface{}
		Error   error
	}
)
