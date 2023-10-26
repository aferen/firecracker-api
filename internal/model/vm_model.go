package model

import (
	"github.com/firecracker-microvm/firecracker-go-sdk"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VM struct {
	Id      primitive.ObjectID   `json:"id,omitempty"`
	Cpu     int                  `json:"cpu,omitempty" validate:"required"`
	Memory  int                  `json:"memory,omitempty" validate:"required"`
	Ip      string               `json:"ip"`
	VmID    string               `json:"vmid"`
	Machine *firecracker.Machine `json:"machine"`
}
