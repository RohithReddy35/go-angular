package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/go-playground/validator/v10"

	"github.com/RohithReddy35/go-angular/internal/models"
	"github.com/RohithReddy35/go-angular/internal/repository"
)

var validate = validator.New()

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


// GetUsers godoc
// @Summary Get all users
// @Description Get a list of all users
// @Tags users
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} map[string]string
// @Router /users [get]
func (api *UserAPI) GetUsers(c echo.Context) error {
	users, err := api.UserRepository.GetAllUsers()
	fmt.Printf("Users: %v\n", users)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, users)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the given details
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User object"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
func (api *UserAPI) CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Validate the input
	if err := validate.Struct(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := api.UserRepository.CreateUser(&user); err != nil {
		if err.Error() == "user_name already exists" {
			return c.JSON(http.StatusConflict, map[string]string{"error": "user_name already exists"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "User created successfully"})
}

// UpdateUser godoc
// @Summary Update an existing user
// @Description Update an existing user with the given details
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.User true "User object"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [put]
func (api *UserAPI) UpdateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Validate the input
	if err := validate.Struct(user); err != nil {
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


// DeleteUser godoc
// @Summary Delete an existing user
// @Description Delete an existing user with the given ID
// @Tags users
// @Param id path int true "User ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [delete]
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
