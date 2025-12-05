package enrollment

import "context"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) List(ctx context.Context) ([]Enrollment, error) {
	return s.repo.List(ctx)
}

func (s *Service) GetByID(ctx context.Context, id string) (*Enrollment, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) Create(ctx context.Context, in CreateInput) (*Enrollment, error) {
	// здесь можно добавить проверки: существует ли студент/курс, нет ли дубля
	return s.repo.Create(ctx, in)
}

func (s *Service) SetStatus(ctx context.Context, id string, status Status) error {
	// можно добавить ограничения переходов статусов
	return s.repo.SetStatus(ctx, id, status)
}
