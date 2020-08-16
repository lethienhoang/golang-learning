package repository

import (
	"building-go-with-ddd-pattern/domain/entity"

	uuid "github.com/satori/go.uuid"
)

// FoodRepository interface
type FoodRepository interface {
	Insert(*entity.Food) (*entity.Food, map[string]string)
	GetFood(uuid.UUID) (*entity.Food, error)
	GetAllFood() ([]entity.Food, error)
	Update(uuid.UUID, *entity.Food) (*entity.Food, map[string]string)
	Delete(uuid.UUID) error
}
