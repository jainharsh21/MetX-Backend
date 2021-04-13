package models

import (
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	UserType  string    `json:"user_type"`
	ImgUrl    string    `json:"img_url"`
	CreatedAt time.Time `json:"created_at"`
}