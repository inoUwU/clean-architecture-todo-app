package model

import "time"

type User struct {
	ID       uint      `json:"id" gorm:"primaryKey"`
	Email    string    `json:"email " gorm:"unique"`
	Pssword  string    `json:"password"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type UserResponse struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Email string `json:"email " gorm:"unique"`
}
