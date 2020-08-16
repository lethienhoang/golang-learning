package infrastructure

import (
	"building-go-with-ddd-pattern/domain/entity"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/jinzhu/gorm"
)

// DbContext is struct
type DbContext struct {
	db *gorm.DB
}

// NewDbContext - initialize
func NewDbContext(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*DbContext, error) {
	DBURL := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbUser, DbName, DbPassword)
	db, err := gorm.Open(Dbdriver, DBURL)

	if err != nil {
		return nil, err
	}

	db.LogMode(true)

	return &DbContext{
		db: db,
	}, nil
}

// GetDbContext is context of init db conenction
func (d *DbContext) GetDbContext() *gorm.DB {
	return d.db
}

// Close is closing connection to db
func (d *DbContext) Close() error {
	return d.db.Close()
}

// Automigrate migariton model schema
func (d *DbContext) Automigrate() error {
	return d.db.AutoMigrate(&entity.Food{}, &entity.User{}).Error
}
