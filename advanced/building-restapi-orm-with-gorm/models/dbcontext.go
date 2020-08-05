package models

import "github.com/jinzhu/gorm"

type DbContext struct {
	Db *gorm.DB
}
