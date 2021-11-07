package utils

import (
	"github.com/labstack/echo"
)

type response struct {
	StatusCode int `json:"-"`
	Status     struct {
		Error        bool   `json:"error"`
		ErrorMessage string `json:"errorMessage,omitempty"`
		Code         string `json:"code,omitempty"`
	} `json:"status"`
	Data interface{} `json:"data,omitempty"`
}

type ResponseInterface interface {
	JSON(c echo.Context) error
}

func (r *response) JSON(c echo.Context) error {
	return c.JSON(r.StatusCode, r)
}

func Success(statusCode int, systemCode string, data interface{}) ResponseInterface {
	r := response{}
	r.StatusCode = statusCode
	r.Status.Code = systemCode
	r.Status.Error = false
	r.Data = data

	return &r
}

func Error(statusCode int, systemCode, errorMessage string, err error) ResponseInterface {
	r := response{}
	r.StatusCode = statusCode
	r.Status.Error = true
	r.Status.ErrorMessage = errorMessage
	r.Status.Code = systemCode

	return &r
}
