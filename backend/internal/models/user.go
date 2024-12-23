package models

import (
	"time"
)

// User is a struct that holds the user data
type User struct {
	ID        uint   `json:"id" sql:"id"`
	UserName  string `json:"user_name" sql:"user_name"`
	Email     string `json:"email" sql:"email"`
	// Password  string `json:"password" sql:"password"`
	CreatedAt time.Time `json:"created_at" sql:"created_at"`
	UpdatedAt time.Time `json:"updated_at" sql:"updated_at"`
}