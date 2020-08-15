package entity

// Food information
type Food struct {
	BaseEntity
	Title       string `gorm:"size:50;not null;"`
	Description string `gorm:"size:150;"`
	ImageURL    string `gorm:"size:250;"`
}
