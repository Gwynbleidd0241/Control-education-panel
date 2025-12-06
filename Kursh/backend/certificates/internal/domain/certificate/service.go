package certificate

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

func (s *Service) List(ctx context.Context) ([]Certificate, error) {
	return s.repo.List(ctx)
}

func (s *Service) GetByID(ctx context.Context, id string) (*Certificate, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) GetByCode(ctx context.Context, code string) (*Certificate, error) {
	if code == "" {
		return nil, errors.New("empty code")
	}
	return s.repo.GetByCode(ctx, code)
}

func (s *Service) Issue(ctx context.Context, in IssueInput) (*Certificate, error) {
	if in.StudentID == "" || in.CourseID == "" {
		return nil, errors.New("studentId and courseId are required")
	}
	return s.repo.Issue(ctx, in)
}

func (s *Service) ListByStudent(ctx context.Context, studentID string) ([]Certificate, error) {
	return s.repo.ListByStudent(ctx, studentID)
}
