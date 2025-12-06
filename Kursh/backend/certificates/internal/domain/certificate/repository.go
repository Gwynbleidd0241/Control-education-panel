package certificate

import "context"

type Repository interface {
	List(ctx context.Context) ([]Certificate, error)
	GetByID(ctx context.Context, id string) (*Certificate, error)
	GetByCode(ctx context.Context, code string) (*Certificate, error)
	Issue(ctx context.Context, in IssueInput) (*Certificate, error)
	ListByStudent(ctx context.Context, studentID string) ([]Certificate, error)
}
