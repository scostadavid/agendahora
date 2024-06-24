package database

import (
	"agendahora/domain/entities"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	database = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	port     = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
	schema   = os.Getenv("DB_SCHEMA")
	ssl      = os.Getenv("DB_SSL")
)

func New() *gorm.DB {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, username, password, database, port, ssl)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to the database:", err)
	}

	db.AutoMigrate(&entities.WaitlistEntry{})

	return db
}
