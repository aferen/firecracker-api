package model

type CreateVMRequest struct {
	Cpu    int `json:"cpu"`
	Memory int `json:"memory"`
}

type ShutdownVMRequest struct {
	VmID string `json:"vmid"`
}
