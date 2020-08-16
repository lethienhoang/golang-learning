package services

import (
	contracts "building-go-with-ddd-pattern/application/contracts/foods"
	"building-go-with-ddd-pattern/domain/entity"
	"building-go-with-ddd-pattern/infrastructure/persistence"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// FoodService
type FoodService struct {
	// db *gorm.DB
	FoodRepo *persistence.FoodRepository
}

// NewFoodService
func NewFoodService(db *gorm.DB) *FoodService {
	return &FoodService{
		FoodRepo: persistence.NewFoodRepository(db),
	}
}

// InsertFood
func (f *FoodService) InsertFood(foodContract contracts.UpsertFoodContract) (*entity.Food, map[string]string) {

	errs := foodContract.Validate("insert")
	if errs != nil {
		return nil, errs
	}

	foodEntity := entity.Food{
		Title:       foodContract.Title,
		Description: foodContract.Description,
	}

	result, errs := f.FoodRepo.Insert(&foodEntity)

	return result, errs
}

func (f *FoodService) UpdateFood(foodContract contracts.UpsertFoodContract) (*entity.Food, map[string]string) {

	errs := foodContract.Validate("update")
	if errs != nil {
		return nil, errs
	}

	foodEntity := entity.Food{
		Title:       foodContract.Title,
		Description: foodContract.Description,
	}

	id := foodContract.ID
	result, errs := f.FoodRepo.Update(id, &foodEntity)

	return result, errs
}

func (f *FoodService) GetFood(id uuid.UUID) (*entity.Food, error) {
	result, err := f.FoodRepo.GetFood(id)

	return result, err
}

func (f *FoodService) GetAll() ([]entity.Food, error) {
	result, err := f.FoodRepo.GetAllFood()

	return result, err
}

func (f *FoodService) Delete(id uuid.UUID) error {
	err := f.FoodRepo.Delete(id)

	return err
}
