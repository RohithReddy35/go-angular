package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/RohithReddy35/go-angular/internal/models"
)

// GetAllUsers returns all users from the database
func (ur *UserRepository) GetAllUsers() ([]models.User, error) {
	ur.DB.Exec("SET search_path TO public")
	query := ur.QueryBuilder.Select("id", "username", "email", "created_at", "updated_at").From("users")

	rows, err := query.RunWith(ur.DB).Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.UserName, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("no users found")
	}

	return users, nil
}

// GetUserByID returns a user by ID from the database
func (ur *UserRepository) GetUserByID(id int) (*models.User, error) {
	query := ur.QueryBuilder.Select("id", "username", "email").From("users").Where("id = ?", id)

	row := query.RunWith(ur.DB).QueryRow()

	var user models.User
	if err := row.Scan(&user.ID, &user.UserName, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

// CreateUser creates a new user in the database
func (ur *UserRepository) CreateUser(user *models.User) error {
	query := ur.QueryBuilder.Insert("users").Columns("username", "email").Values(user.UserName, user.Email).Suffix("RETURNING id, created_at, updated_at")

	_, err := query.RunWith(ur.DB).Exec()
	if err != nil {
		return err
	}

	return nil
}

// UpdateUser updates a user in the database
func (ur *UserRepository) UpdateUser(user *models.User) error {
	query := ur.QueryBuilder.Update("users").Set("username", user.UserName).Set("email", user.Email).Set("updated_at", time.Now()).Where("id = ?", user.ID).Suffix("RETURNING id, created_at, updated_at")

	_, err := query.RunWith(ur.DB).Exec()
	if err != nil {
		return err
	}

	return nil	
}

// DeleteUser deletes a user from the database
func (ur *UserRepository) DeleteUser(id int) error {
	query := ur.QueryBuilder.Delete("users").Where("id = ?", id)

	_, err := query.RunWith(ur.DB).Exec()
	if err != nil {
		return err
	}

	return nil
}