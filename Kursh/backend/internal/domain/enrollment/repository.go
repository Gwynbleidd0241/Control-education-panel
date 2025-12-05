package enrollment

import "context"

type Repository interface {
	List(ctx context.Context) ([]Enrollment, error)
	GetByID(ctx context.Context, id string) (*Enrollment, error)
	Create(ctx context.Context, in CreateInput) (*Enrollment, error)
	SetStatus(ctx context.Context, id string, status Status) error
}
