package migrations

import (
	"github.com/Andesson/marketplace-auth-service/schemas"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(&schemas.Users{}, &schemas.AuthCredential{})
	if err != nil {
		return err
	}
	return nil
}
