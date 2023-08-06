package models

type User struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
}
