package schemas

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthProviders struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID     uuid.UUID `gorm:"not null;onDelete:CASCADE"`
	Provider   string    `gorm:"type:text;not null"`
	ProviderID string    `gorm:"type:text;not null;unique"`
	Email      string    `gorm:"type:text;not null"`
	User       Users     `gorm:"foreignKey:UserID;references:ID"`
}
