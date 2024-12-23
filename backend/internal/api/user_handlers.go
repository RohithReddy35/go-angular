package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/RohithReddy35/go-angular/internal/models"
	"github.com/RohithReddy35/go-angular/internal/repository"
)

// UserAPI is a struct that holds the UserRepository
type UserAPI struct {
	UserRepository *repository.UserRepository
}

// NewUserAPI creates a new UserAPI
func NewUserAPI(ur *repository.UserRepository) *UserAPI {
	return &UserAPI{
		UserRepository: ur,
	}
}

// GetUsers returns all users
func (api *UserAPI) GetUsers(c echo.Context) error {
	users, err := api.UserRepository.GetAllUsers()
	fmt.Printf("Users: %v\n", users)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, users)
}

// CreateUser creates a new user
func (api *UserAPI) CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := api.UserRepository.CreateUser(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "User created successfully"})
}

// UpdateUser updates a user
func (api *UserAPI) UpdateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	id, err := strconv.Atoi(c.Param("id"))
	user.ID = uint(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := api.UserRepository.UpdateUser(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User updated successfully"})
}


// DeleteUser deletes a user
func (api *UserAPI) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := api.UserRepository.DeleteUser(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User deleted successfully"})
}
