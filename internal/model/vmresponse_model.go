package model

import (
	"github.com/labstack/echo/v4"
)

type VMResponse struct {
	IpAddress string    `json:"ip_address"`
	ID        string    `json:"id"`
	Data      *echo.Map `json:"data"`
}
