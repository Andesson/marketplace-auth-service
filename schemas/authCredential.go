package schemas

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthCredential struct {
	gorm.Model
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID        uuid.UUID `gorm:"not null"` // Defina UserID como uint (ou int) e not null
	PassHash      string    `gorm:"column:pass_hash"`
	HashAlgorithm string    `gorm:"column:hash_algorithm"`
	Salt          string    `gorm:"column:salt"`
	User          Users     `gorm:"foreignKey:UserID;references:ID"` // Defina a relação correta
}
