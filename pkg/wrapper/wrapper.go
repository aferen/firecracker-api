package wrapper

import (
	"github.com/labstack/echo/v4"
)

type Props struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
}

func Data(code int, data interface{}, message string, c echo.Context) error {
	props := &Props{
		Code:    code,
		Data:    data,
		Message: message,
		Success: true,
	}
	return c.JSON(code, props)
}

func Error(code int, message string, c echo.Context) error {
	props := &Props{
		Code:    code,
		Data:    nil,
		Message: message,
		Success: false,
	}
	return c.JSON(code, props)
}
