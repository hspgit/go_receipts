package models

// Item model for database operations
type Item struct {
	ID    uint    `gorm:"primaryKey;autoIncrement"`
	Name  string  `gorm:"column:name;not null"`
	Price float64 `gorm:"column:price;not null"`
}
