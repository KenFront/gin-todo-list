package model

import (
	"time"

	"github.com/google/uuid"
)

type UserStauts string

const (
	UserActive    UserStauts = "active"
	UserInactive  UserStauts = "inactive"
	UserForbidden UserStauts = "forbidden"
)

type User struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;" json:"id"`
	Name      string     `gorm:"type:string;size:100" json:"name"`
	Account   string     `gorm:"type:string;size:100;unique" json:"account"`
	Password  string     `gorm:"type:string;size:100" json:"password"`
	Email     string     `gorm:"type:string;size:255" json:"email"`
	Status    UserStauts `gorm:"type:enum('active', 'inactive', 'forbidden');default:'active'" json:"status"`
	CreatedAt time.Time  `sql:"DEFAULT:'current_timestamp'" json:"createAt"`
	UpdatedAt time.Time  `sql:"DEFAULT:'current_timestamp'" json:"updateAt"`
}

type AddUser struct {
	Name     string
	Account  string
	Password string
	Email    string
}

type SignIn struct {
	Account  string
	Password string
}
