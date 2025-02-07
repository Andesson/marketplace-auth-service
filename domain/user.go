// domain/user.go
package model

import "time"

type User struct {
	ID        uint       // Identificador único
	Email     string     // E-mail do usuário
	FullName  string     // Nome completo do usuário
	Password  string     // Senha do usuário
	CreatedAt time.Time  // Data de criação
	UpdatedAt time.Time  // Data de atualização
	DeletedAt *time.Time // Data de exclusão, se aplicável
}
