package service

import (
	"context"

	"github.com/aferen/firecracker-api/internal/model"
	"github.com/aferen/firecracker-api/internal/repository"
	"github.com/aferen/firecracker-api/pkg/vmm"
	log "github.com/sirupsen/logrus"
)

type Service interface {
	Create(ctx context.Context, input *model.CreateVMRequest) (VM, error)
}

type VM struct {
	model.VM
}

type VMService struct {
	repo repository.VMRepositoryI
}

func NewVMService(repo repository.VMRepositoryI) Service {
	return VMService{repo}
}

func (s VMService) Get(ctx context.Context, id string) (VM, error) {
	return VM{}, nil
}

func (s VMService) Create(ctx context.Context, req *model.CreateVMRequest) (VM, error) {
	ctx2 := context.Background()
	vm, err := vmm.CreateAndStartVM(ctx2)
	if err != nil {
		log.Error("failed to create VMM")
		return VM{}, err
	}
	log.WithField("ip", vm.Ip).Info("New VM created and started")
	err = s.repo.Create(ctx, model.VM{
		Ip:      vm.Ip.String(),
		VmID:    vm.VmID,
		Machine: *&vm.Machine,
	})
	if err != nil {
		return VM{}, err
	}
	return s.Get(ctx, "0")
}
