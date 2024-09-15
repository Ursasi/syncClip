package util

import "syncClip/server/service"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(code int, message string, data interface{}) Response {
	return Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

type RegisterResponse struct {
	ID     string          `json:"id"`
	Boards []service.Board `json:"boards"`
}

type ProbeResponse struct {
	Boards []service.Board `json:"boards"`
}
