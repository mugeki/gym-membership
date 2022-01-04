package controllers

import (
	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Meta struct {
		Status   int      `json:"status"`
		Message  string   `json:"message"`
		Messages []string `json:"messages,omitempty"`
	} `json:"meta"`
	Page interface{} `json:"page,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func NewSuccessResponse(c echo.Context, status int, data interface{}, args ...interface{}) error {
	res := BaseResponse{}
	res.Meta.Status = status
	res.Meta.Message = "Success"
	if data != "" && data != nil{
		res.Data = data
	}
	if len(args) > 0 {
		res.Page = args[0]
	}
	return c.JSON(status, res)
}

func NewErrorResponse(c echo.Context, status int, err error) error {
	res := BaseResponse{}
	res.Meta.Status = status
	res.Meta.Message = "Error"
	res.Meta.Messages = []string{err.Error()}

	return c.JSON(status, res)
}
