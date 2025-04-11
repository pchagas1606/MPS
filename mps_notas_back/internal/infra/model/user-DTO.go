package model

import "time"

type UserDTO struct {
	ID        int       `json:"id"`         // Identificador único do usuário
	Name      string    `json:"name"`       // Nome do usuário
	Email     string    `json:"email"`      // Email do usuário
	CreatedAt time.Time `json:"created_at"` // Data de criação do registro
}
