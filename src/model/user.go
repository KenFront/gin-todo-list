package model

import (
	"time"

	"github.com/google/uuid"
)

type UserStauts string

const (
	USER_ACTIVE    UserStauts = "active"
	USER_INACTIVE  UserStauts = "inactive"
	USER_FORBIDDEN UserStauts = "forbidden"
)

type User struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;" json:"id"`
	Name      string     `gorm:"type:string;size:100" json:"name"`
	Account   string     `gorm:"type:string;size:100;unique" json:"account"`
	Password  string     `gorm:"type:string;size:100" json:"-"`
	Email     string     `gorm:"type:string;size:255" json:"email"`
	Status    UserStauts `gorm:"type:string;size:32;default:'active';check:status IN ('active', 'inactive', 'forbidden')" json:"status"`
	CreatedAt time.Time  `sql:"DEFAULT:'current_timestamp'" json:"createAt"`
	UpdatedAt time.Time  `sql:"DEFAULT:'current_timestamp'" json:"updateAt"`
}

type AddUser struct {
	Name  string `binding:"min=3,max=100" json:"name"`
	Email string `binding:"email,min=5,max=255" json:"email"`
	SignIn
}

type SignIn struct {
	Account  string `binding:"min=3,max=100" json:"account"`
	Password string `binding:"min=3,max=100" json:"password"`
}
