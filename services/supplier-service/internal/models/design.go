package models

type Design struct {
	ID          string `gorm:"primaryKey"`
	DesignerID  string `gorm:"not null"`
	Title       string `gorm:"not null"`
	Description string
	Category    string `gorm:"not null"`
	FilePath    string `gorm:"not null"`
	Status      string `gorm:"not null"`
}