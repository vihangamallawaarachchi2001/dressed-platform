package models

type Design struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	DesignerID  string    `gorm:"not null" json:"designerId"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `json:"description"`
	Category    string    `gorm:"not null" json:"category"`
	FilePath    string    `gorm:"not null" json:"filePath"`
	Status      string    `gorm:"not null" json:"status"`
}