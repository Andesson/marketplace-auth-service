package migrations

import (
	"github.com/Andesson/marketplace-auth-service/database/schemas"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(&schemas.Users{}, &schemas.AuthCredential{}, &schemas.Sessions{}, &schemas.AuthProviders{})
	if err != nil {
		return err
	}
	return nil
}
