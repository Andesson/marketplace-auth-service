package model

import (
	"time"

	"github.com/google/uuid"
)

type AuthCredential struct {
	ID            uuid.UUID
	UserID        uuid.UUID
	PassHash      string
	HashAlgorithm string
	Salt          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
}

func (a *AuthCredential) SetID(id uuid.UUID) {
	a.ID = id
}
