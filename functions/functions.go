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
		return &Response{
			StatusCode: 500,
			Headers:    Headers,
			Body:       err.Error() + ", original error: " + message,
		}
	}
	return &Response{
		StatusCode: statusCode,
		Headers:    Headers,
		Body:       string(eb),
	}
}

func StatusCodeTitle(statusCode int) string {
	switch statusCode {
	case 200:
		return "OK"
	case 201:
		return "Created"
	case 204:
		return "No Content"
	case 400:
		return "Bad Request"
	case 401:
		return "Unauthorized"
	case 403:
		return "Forbidden"
	case 404:
		return "Not Found"
	case 405:
		return "Method Not Allowed"
	case 406:
		return "Not Acceptable"
	case 409:
		return "Conflict"
	case 415:
		return "Unsupported Media Type"
	case 500:
		return "Internal Server Error"
	case 501:
		return "Not Implemented"
	case 502:
		return "Bad Gateway"
	case 503:
		return "Service Unavailable"
	case 504:
		return "Gateway Timeout"
	default:
		return "Unknown"
	}
}
