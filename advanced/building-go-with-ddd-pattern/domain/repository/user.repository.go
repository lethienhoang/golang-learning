package repository

import (
	"building-go-with-ddd-pattern/domain/entity"

	uuid "github.com/satori/go.uuid"
)

// UserRepository interface
type UserRepository interface {
	Insert(*entity.User) (*entity.User, map[string]string)
	GetUser(uuid.UUID) (*entity.Food, error)
	GetUsers() ([]entity.User, error)
	GetUserByEmailAndPassword(*entity.User) (*entity.User, map[string]string)
}
