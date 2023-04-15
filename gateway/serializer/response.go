package serializer

import (
	"errors"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

var (
	ErrorUnauthorized     = errors.New("unauthorized")
	ErrorPermissionDenied = errors.New("permission denied")
)

func ResponseSuccess(data interface{}) Response {
	return Response{
		Code: 0,
		Msg:  "success",
		Data: data,
	}
}

func ResponseError(err error) Response {
	return Response{
		Code: 1,
		Msg:  err.Error(),
	}
}

func ResponseErrorMsg(msg string) Response {
	return Response{
		Code: 1,
		Msg:  msg,
	}
}
