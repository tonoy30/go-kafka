package model

import "net/http"

type Response struct {
	Status     int         `json:"status"`
	StatusText string      `json:"statusText"`
	Message    string      `json:"message"`
	Payload    interface{} `json:"payload"`
}

func statusText(code int) string {
	return http.StatusText(code)
}
func NewResponse(status int, message string, payload interface{}) *Response {
	return &Response{
		status,
		statusText(status),
		message,
		payload,
	}
}
