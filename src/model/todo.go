package model

import (
	"time"

	"github.com/google/uuid"
)

type TodoStauts string

const (
	TODO_IDLE      TodoStauts = "idle"
	TODO_COMPLETED TodoStauts = "completed"
)

type Todo struct {
	ID          uuid.UUID  `gorm:"type:uuid;primary_key;" json:"id"`
	Title       string     `gorm:"type:string;size:100" json:"title"`
	Description string     `gorm:"type:string;size:65535" json:"description"`
	Status      TodoStauts `gorm:"type:string;size:32;default:'idle';check:status IN ('idle', 'completed')" json:"status"`
	CreatedAt   time.Time  `sql:"DEFAULT:'current_timestamp'" json:"createdAt"`
	UpdatedAt   time.Time  `sql:"DEFAULT:'current_timestamp'" json:"updatedAt"`
	UserId      uuid.UUID  `gorm:"type:uuid;foreignKey:User;" json:"userId"`
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

type TodoUri struct {
	TodoId string `uri:"todoId" binding:"required,uuid"`
}
