package domain

import "github.com/google/uuid"

type User struct {
	ID   uuid.UUID `json:"id" gorm:"primaryKey"`
	Name string    `json:"name"`
	Age  int       `json:"age"`
}

type UserService interface {
	Create(user *User) error
}

type UserRepository interface {
	Save(user *User) error
}
