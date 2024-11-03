package database

import (
	"fmt"
	"gobarber/internal/infra/database/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

func GetConnection() *gorm.DB {
	return db
}

func Connect() error {
	conn, err := gorm.Open(postgres.Open(connectionString()), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return err
	}
	_ = conn.AutoMigrate(&entity.User{})

	db = conn
	return nil
}

func connectionString() string {
	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"),
	)
}
