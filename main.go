package main

import (
	"github.com/Andesson/marketplace-auth-service/config"
	"github.com/Andesson/marketplace-auth-service/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("logs/app.log")

	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}
	router.Initialize()
}
