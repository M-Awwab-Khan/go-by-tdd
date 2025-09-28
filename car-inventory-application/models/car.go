package models

type Car struct {
	ID    int    `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Brand string `json:"brand"`
}
