package models


// User is a struct that holds the user data
type User struct {
	ID        uint   `json:"id" sql:"id"`
	UserName  string `json:"user_name" sql:"user_name" validate:"required,min=3,max=50,alphanum"`
	Email     string `json:"email" sql:"email" validate:"required,email"`
	// Password  string `json:"password" sql:"password" validate:"required,min=6,max=50"`
	CreatedAt string `json:"created_at" sql:"created_at"`
	UpdatedAt string `json:"updated_at" sql:"updated_at"`
}