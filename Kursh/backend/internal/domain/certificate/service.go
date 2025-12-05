package certificate

import (
	"context"
	"time"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) List(ctx context.Context) ([]Certificate, error) {
	return s.repo.List(ctx)
}

func (s *Service) ListTemplates(ctx context.Context) ([]Template, error) {
	return s.repo.ListTemplates(ctx)
}

func (s *Service) Issue(ctx context.Context, in IssueInput) (*Certificate, error) {
	// можно добавить проверки: не выдан ли уже, прошёл ли курс и т.д.
	return s.repo.Issue(ctx, in)
}

func (s *Service) Revoke(ctx context.Context, id string) error {
	return s.repo.SetStatus(ctx, id, StatusRevoked)
}

func (s *Service) VerifyByCode(ctx context.Context, code string) (*Certificate, error) {
	cert, err := s.repo.GetByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	// бизнес-логика истечения
	if cert.ExpiresAt != nil && cert.ExpiresAt.Before(time.Now()) && cert.Status == StatusValid {
		// можно автоматически пометить как expired, если хочешь
		_ = s.repo.SetStatus(ctx, cert.ID, StatusExpired)
		cert.Status = StatusExpired
	}

	return cert, nil
}
