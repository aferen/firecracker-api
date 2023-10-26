package handler

import (
	"fmt"
	"net/http"

	"github.com/aferen/firecracker-api/internal/model"
	"github.com/aferen/firecracker-api/internal/service"
	"github.com/aferen/firecracker-api/pkg/wrapper"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type VMHandler struct {
	uUcase          service.Service
	createVMRequest model.CreateVMRequest
}

func NewVMHandler(e *echo.Echo, ur service.Service) {
	vmh := &VMHandler{
		uUcase: ur,
	}
	e.POST("/vm", vmh.Create)
}

func isRequestValid(m *model.VM) (bool, error) {

	validate := validator.New()

	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (h *VMHandler) Create(c echo.Context) error {
	err := c.Bind(&h.createVMRequest)
	if err != nil {
		return wrapper.Error(http.StatusUnprocessableEntity, err.Error(), c)
	}
	if err != nil {
		return wrapper.Error(http.StatusUnprocessableEntity, err.Error(), c)
	}

	vm, err := h.uUcase.Create(c.Request().Context(), &h.createVMRequest)
	if err != nil {
		fmt.Println(err.Error())
		return wrapper.Error(http.StatusConflict, "vm is already created", c)
	}
	return wrapper.Data(http.StatusCreated, vm, "vm is created", c)
}
