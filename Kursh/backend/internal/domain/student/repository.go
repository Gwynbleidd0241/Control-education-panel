package student

import "context"

type Repository interface {
	List(ctx context.Context) ([]Student, error)
	GetByID(ctx context.Context, id string) (*Student, error)
	Create(ctx context.Context, in CreateInput) (*Student, error)
	Update(ctx context.Context, id string, in UpdateInput) (*Student, error)
}
