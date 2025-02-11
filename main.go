package main

import (
	"github.com/Andesson/marketplace-auth-service/config"
	"github.com/Andesson/marketplace-auth-service/router"
	"github.com/joho/godotenv"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("logs/app.log")
	logger = config.GetLogger("main")
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		logger.Errorf("Erro ao carregar o arquivo .env")
	}
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}
	router.Initialize()
}
