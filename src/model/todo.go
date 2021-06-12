package model

import (
	"time"

	"github.com/google/uuid"
)

type TodoStauts string

const (
	TodoIdle      TodoStauts = "idle"
	TodoCompleted TodoStauts = "completed"
)

type Todo struct {
	ID          uuid.UUID  `gorm:"type:uuid;primary_key;" json:"id"`
	Title       string     `gorm:"type:string" json:"title"`
	Description string     `gorm:"type:string" json:"description"`
	Status      TodoStauts `gorm:"type:enum('idle', 'completed');default:'idle'" json:"status"`
	CreatedAt   time.Time  `sql:"DEFAULT:'current_timestamp'" json:"createAt"`
	UpdatedAt   time.Time  `sql:"DEFAULT:'current_timestamp'" json:"updateAt"`
}

type AddTodo struct {
	Title       string
	Description string
}

type PatchTodo struct {
	Title       string
	Description string
	Status      TodoStauts
}
