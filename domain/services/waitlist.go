package services

import (
	"agendahora/domain/entities"
	"agendahora/domain/repositories"
	"context"
)

type WaitlistService struct {
	repository repositories.WaitListRepository
}

func NewWaitlistService(repository repositories.WaitListRepository) *WaitlistService {
	return &WaitlistService{
		repository: repository,
	}
}

func (s *WaitlistService) Create(ctx context.Context, email string) error {
	err := s.repository.Create(ctx, email)

	if err != nil {
		return err
	}

	return nil
}

func (s *WaitlistService) FindByEmail(ctx context.Context, email string) *entities.WaitlistEntry {
	record := s.repository.FindByEmail(ctx, email)

	if record != nil {
		return record
	}

	return nil
}
