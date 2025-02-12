package schemas

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Sessions struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID `gorm:"not null;onDelete:CASCADE"`
	Token     string    `gorm:"type:text;not null"`
	ExpiresAt time.Time `gorm:"type:timestamp;not null"`
	User      Users     `gorm:"foreignKey:UserID;references:ID"`
}
