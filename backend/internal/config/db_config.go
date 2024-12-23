package config

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// DBConfig is a struct that holds the database configuration
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// GetDBConfig creates a new DBConfig struct
func GetDBConfig() *DBConfig {
	return &DBConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   "users_db",
	}
}

// ConnectionString returns the connection string for the database
func GetDBConnectionString() string {
	cfg := GetDBConfig()
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
}