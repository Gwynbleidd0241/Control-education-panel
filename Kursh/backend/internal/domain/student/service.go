package student

import "context"

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
	return s.repo.GetByID(ctx, id)
}

// Создание студента по полному инпуту
func (s *Service) Create(ctx context.Context, in CreateInput) (*Student, error) {
	return s.repo.Create(ctx, in)
}

// Обновление студента по id и полному инпуту
func (s *Service) Update(ctx context.Context, id string, in UpdateInput) (*Student, error) {
	return s.repo.Update(ctx, id, in)
}
