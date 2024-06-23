package database

import (
	"agendahora/domain/entities"
	"os"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestNew(t *testing.T) {
	origDatabase := os.Getenv("DB_DATABASE")
	origPassword := os.Getenv("DB_PASSWORD")
	origUsername := os.Getenv("DB_USERNAME")
	origPort := os.Getenv("DB_PORT")
	origHost := os.Getenv("DB_HOST")
	origSchema := os.Getenv("DB_SCHEMA")

	os.Setenv("DB_DATABASE", ":memory:")
	os.Setenv("DB_PASSWORD", "")
	os.Setenv("DB_USERNAME", "")
	os.Setenv("DB_PORT", "")
	os.Setenv("DB_HOST", "")
	os.Setenv("DB_SCHEMA", "")

	defer func() {
		os.Setenv("DB_DATABASE", origDatabase)
		os.Setenv("DB_PASSWORD", origPassword)
		os.Setenv("DB_USERNAME", origUsername)
		os.Setenv("DB_PORT", origPort)
		os.Setenv("DB_HOST", origHost)
		os.Setenv("DB_SCHEMA", origSchema)
	}()

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to the database: %v", err)
	}

	err = db.AutoMigrate(&entities.WaitlistEntry{})

	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	// Teste adicional para verificar se a tabela foi criada corretamente
	var count int64
	db.Model(&entities.WaitlistEntry{}).Count(&count)
	if count != 0 {
		t.Fatalf("expected 0 entries in the waitlist, got %d", count)
	}
}
