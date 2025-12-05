package certificate

import "context"

type Repository interface {
	List(ctx context.Context) ([]Certificate, error)
	ListTemplates(ctx context.Context) ([]Template, error)
	Issue(ctx context.Context, in IssueInput) (*Certificate, error)
	SetStatus(ctx context.Context, id string, status Status) error
	GetByCode(ctx context.Context, code string) (*Certificate, error)
}
