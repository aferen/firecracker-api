package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aferen/firecracker-api/internal/config"
	"github.com/aferen/firecracker-api/internal/handler"
	"github.com/aferen/firecracker-api/internal/repository"
	"github.com/aferen/firecracker-api/internal/service"
	mongodb "github.com/aferen/firecracker-api/pkg/mongo"
	"github.com/labstack/echo/v4"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

var Version = "1.0.0"

var flagConfig = flag.String("config", "./config/local.yaml", "path to the config file")

func main() {
	flag.Parse()
	logger := log.New().With(nil, "version", Version)

	cfg, err := config.Load(*flagConfig, logger)
	if err != nil {
		logger.Errorf("failed to load application configuration: %s", err)
		os.Exit(-1)
	}
	e := echo.New()
	db := mongodb.ConnectDB(cfg.MONGOURI)
	vmRepo := repository.NewVMepository(db)
	vmSvc := service.NewVMService(vmRepo)
	handler.NewVMHandler(e, vmSvc)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.ServerPort)))
}
