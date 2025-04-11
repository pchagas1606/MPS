package model

import (
	"time"
)

// TaskDAO define a estrutura de dados das tarefas
type TaskDAO struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   string    `json:"start-date"`
	EndDate     string    `json:"end-date"`
	CreatedAt   time.Time `json:"created_at"` // Data de criação do registro
}

// NewTaskInput contém os dados necessários para criar uma nova tarefa - DTO
type NewTaskInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	StartDate   string `json:"start-date"`
	EndDate     string `json:"end-date"`
}
