package course

import "context"

type Repository interface {
	List(ctx context.Context) ([]Course, error)
	GetByID(ctx context.Context, id string) (*Course, error)
	Create(ctx context.Context, in CreateCourseInput) (*Course, error)
	Update(ctx context.Context, id string, in CreateCourseInput) (*Course, error)
	Delete(ctx context.Context, id string) error
}
