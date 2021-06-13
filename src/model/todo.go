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
	Title       string     `gorm:"type:string;size:100" json:"title"`
	Description string     `gorm:"type:string;size:65535" json:"description"`
	Status      TodoStauts `gorm:"type:enum('idle', 'completed');default:'idle'" json:"status"`
	CreatedAt   time.Time  `sql:"DEFAULT:'current_timestamp'" json:"createAt"`
	UpdatedAt   time.Time  `sql:"DEFAULT:'current_timestamp'" json:"updateAt"`
	UserId      uuid.UUID  `gorm:"type:uuid;foreignKey;" json:"userId"`
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
