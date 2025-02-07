package handler

import (
	"github.com/Andesson/marketplace-auth-service/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetPostgree()
}
