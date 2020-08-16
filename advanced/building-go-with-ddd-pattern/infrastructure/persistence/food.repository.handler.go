package persistence

import (
	"building-go-with-ddd-pattern/domain/entity"
	"building-go-with-ddd-pattern/domain/repository"
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// FoodRepository infacstructure
type FoodRepository struct {
	db *gorm.DB
}

// NewFoodRepository - initialize
func NewFoodRepository(db *gorm.DB) *FoodRepository {
	return &FoodRepository{db}
}

// FoodRepo implements the repository.FoodRepository interface
var _ repository.FoodRepository = &FoodRepository{}

// Insert is function that will insert data from user to database
func (f *FoodRepository) Insert(food *entity.Food) (*entity.Food, map[string]string) {
	dbErr := map[string]string{}

	err := f.db.Create(&food).Error
	if err != nil {
		//since our title is unique
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["unique_title"] = "food title already taken"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}

	return food, dbErr
}

// GetFood - get food data base on id
func (f *FoodRepository) GetFood(id uuid.UUID) (*entity.Food, error) {
	var food entity.Food

	err := f.db.Where("id=?", id).Take(&food).Error
	if err != nil {
		return nil, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("food not found")
	}

	return &food, nil
}

// GetAllFood - get all records
func (f *FoodRepository) GetAllFood() ([]entity.Food, error) {
	var foods []entity.Food

	err := f.db.Limit(100).Find(&foods).Error
	if err != nil {
		return nil, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("food not found")
	}

	return foods, nil
}

// Update - update new data from exisiting old data
func (f *FoodRepository) Update(id uuid.UUID, food *entity.Food) (*entity.Food, map[string]string) {
	dbErr := map[string]string{}

	err := f.db.Model(&food).Where("id=?", id).Update("Title", "Description", "ImageURL").Error
	if err != nil {
		//since our title is unique
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["unique_title"] = "food title already taken"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}

	return food, nil
}

//Delete - delete file
func (f *FoodRepository) Delete(id uuid.UUID) error {
	var food entity.Food

	err := f.db.Where("id=?", id).Delete(&food).Error
	if err != nil {
		return errors.New("database error, please try again")
	}

	return nil
}
