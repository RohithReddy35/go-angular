package main

import (
	"fmt"

	"github.com/RohithReddy35/go-angular/internal/db"
	"github.com/RohithReddy35/go-angular/internal/repository"
	"github.com/RohithReddy35/go-angular/internal/api"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Hello, World!")

	db.InitDB()
	defer db.DB.Close()

	repo := repository.NewUserRepository(db.DB)

	e := echo.New()
	api.RegisterRoutes(e, repo)
	
	e.Logger.Fatal(e.Start(":8080"))
}