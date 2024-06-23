package repositories

import (
	"agendahora/domain/entities"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type WaitListRepository struct {
	db *gorm.DB
}

func NewWaitlistRepository(db *gorm.DB) *WaitListRepository {
	return &WaitListRepository{db: db}
}

func (r *WaitListRepository) Create(ctx context.Context, email string) error {
	fmt.Print("create", email)
	data := &entities.WaitlistEntry{
		Email: email,
	}
	record := r.db.Create(data)
	return record.Error
}

func (r *WaitListRepository) FindByEmail(ctx context.Context, email string) *entities.WaitlistEntry {

	record := &entities.WaitlistEntry{}
	r.db.Where(&entities.WaitlistEntry{Email: email}).First(&record)

	return record
}
