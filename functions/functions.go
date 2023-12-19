package functions

import "encoding/json"

type Response struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       string            `json:"body,omitempty"`
}

type ErrorResponse struct {
	Title      string `json:"title,omitempty"`
	StatusCode int    `json:"status,omitempty"`
	Detail     string `json:"detail,omitempty"`
}

var Headers = map[string]string{
	"Content-type": "application/json",
}

func ErrorHandler(title string, statusCode int, message string) *Response {
	eb, err := json.Marshal(ErrorResponse{
		Title:      title,
		StatusCode: statusCode,
		Detail:     message,
	})
	if err != nil {
		panic(err)
	}
	return &Response{
		StatusCode: statusCode,
		Headers:    Headers,
		Body:       string(eb),
	}
}
