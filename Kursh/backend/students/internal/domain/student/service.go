package student

import (
	"context"
	"errors"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) List(ctx context.Context) ([]Student, error) {
	return s.repo.List(ctx)
}

func (s *Service) GetByID(ctx context.Context, id string) (*Student, error) {
	if id == "" {
		return nil, errors.New("empty id")
	}
	return s.repo.GetByID(ctx, id)
}

func (s *Service) Update(ctx context.Context, id string, in UpdateInput) (*Student, error) {
	if id == "" {
		return nil, errors.New("empty id")
	}
	if in.Progress < 0 || in.Progress > 100 {
		return nil, errors.New("progress must be 0..100")
	}
	if in.Format == "" {
		in.Format = "Онлайн"
	}
	return s.repo.Update(ctx, id, in)
}
