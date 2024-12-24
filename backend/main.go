package main

import (
	"fmt"

	"github.com/RohithReddy35/go-angular/internal/db"
	"github.com/RohithReddy35/go-angular/internal/repository"
	"github.com/RohithReddy35/go-angular/internal/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/RohithReddy35/go-angular/docs" // Import Swagger docs
	"github.com/swaggo/echo-swagger"             // Swagger handler
)


// @title User Management API
// @version 1.0
// @description This is a sample CRUD application with Go and Echo framework.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	fmt.Println("Hello, World!")

	db.InitDB()
	defer db.DB.Close()

	repo := repository.NewUserRepository(db.DB)

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Add CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4200"}, // Allow only your frontend's origin
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE}, // Allow specific HTTP methods
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Swagger documentation route
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Register API routes
	api.RegisterRoutes(e, repo)

	// Serve Angular static files
	e.Static("/", "frontend/dist/user-management")

	// Handle Angular routes
	e.GET("/*", func(c echo.Context) error {
		return c.File("frontend/dist/user-management/index.html")
	})	

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}