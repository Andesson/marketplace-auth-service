package utils

import (
	"errors"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UUIDModel interface {
	SetID(uuid.UUID)
}

func BeforeCreate(tx *gorm.DB) (err error) {
	if model, ok := tx.Statement.Dest.(UUIDModel); ok {
		id := uuid.New()
		if id == uuid.Nil {
			return errors.New("can't save invalid data")
		}
		model.SetID(id)
		log.Printf("BeforeCreate: ID = %v", id)
	}
	return
}
