package entity

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// BaseEntity define default property
type BaseEntity struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

// BeforeCreate will be execute before BeforeSave
func (base *BaseEntity) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid)
}

// BeforeSave will be execute before saving
func (base *BaseEntity) BeforeSave() {
	base.CreatedAt = time.Now()
	base.UpdatedAt = time.Now()
}
