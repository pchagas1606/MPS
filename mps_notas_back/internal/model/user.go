package model

import "time"

// User define a estrutura de dados do usuário
type User struct {
	ID        int       `json:"id"`         // Identificador único do usuário
	Name      string    `json:"name"`       // Nome do usuário
	Email     string    `json:"email"`      // Email do usuário
	CreatedAt time.Time `json:"created_at"` // Data de criação do registro
}

// NewUserInput contém os dados necessários para criar um novo usuário, DTO
type NewUserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}