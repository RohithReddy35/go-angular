package api

import (
	"github.com/labstack/echo/v4"
	"fmt"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/users", GetUsers)
	e.POST("/users", CreateUser)
	e.PUT("/users/:id", UpdateUser)
	e.DELETE("/users/:id", DeleteUser)
}