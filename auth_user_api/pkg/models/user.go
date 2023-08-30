package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"unique"`
	Password string
	Name     string
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
