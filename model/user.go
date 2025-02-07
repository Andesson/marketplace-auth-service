package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  // Identificador único
	Email     string     // E-mail do usuário
	FullName  string     // Nome completo do usuário
	CreatedAt time.Time  // Data de criação
	UpdatedAt time.Time  // Data de atualização
	DeletedAt *time.Time // Data de exclusão, se aplicável
}

func (u *User) SetID(id uuid.UUID) {
	u.ID = id
}
