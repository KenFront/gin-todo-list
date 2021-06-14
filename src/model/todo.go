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
	Status      TodoStauts `gorm:"default:'idle';" sql:"type:ENUM('idle', 'completed');" json:"status"`
	CreatedAt   time.Time  `sql:"DEFAULT:'current_timestamp'" json:"createAt"`
	UpdatedAt   time.Time  `sql:"DEFAULT:'current_timestamp'" json:"updateAt"`
	UserId      uuid.UUID  `gorm:"type:uuid;foreignKey;" json:"userId"`
}

type AddTodo struct {
	Title       string `binding:"min=1,max=100" json:"title"`
	Description string `binding:"min=1,max=65535" json:"description"`
}

type PatchTodo struct {
	Title       string     `binding:"omitempty,max=100" json:"title"`
	Description string     `binding:"omitempty,max=65535" json:"description"`
	Status      TodoStauts `binding:"omitempty,oneof=idle completed" json:"status"`
}
