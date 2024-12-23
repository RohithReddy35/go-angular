package api

import (
	"github.com/labstack/echo/v4"
	"github.com/RohithReddy35/go-angular/internal/repository"
)

func RegisterRoutes(e *echo.Echo, UserRepository *repository.UserRepository) {
	userAPI := NewUserAPI(UserRepository)

	e.GET("/users", userAPI.GetUsers)
	e.POST("/users", userAPI.CreateUser)
	e.PUT("/users/:id", userAPI.UpdateUser)
	e.DELETE("/users/:id", userAPI.DeleteUser)
}