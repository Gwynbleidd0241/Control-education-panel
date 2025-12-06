package postgres

import (
	"context"
	"math/rand"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"certificates/internal/domain/certificate"
)

type CertificateRepository struct {
	pool *pgxpool.Pool
}

func NewCertificateRepository(pool *pgxpool.Pool) *CertificateRepository {
	return &CertificateRepository{pool: pool}
}

func (r *CertificateRepository) List(ctx context.Context) ([]certificate.Certificate, error) {
	const q = `
SELECT id, code, student_id, course_id,
       issued_at, COALESCE(grade,''), created_at
FROM certificates
ORDER BY issued_at DESC;
`
	rows, err := r.pool.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []certificate.Certificate
	for rows.Next() {
		var c certificate.Certificate
		if err = rows.Scan(
			&c.ID, &c.Code, &c.StudentID, &c.CourseID,
			&c.IssuedAt, &c.Grade, &c.CreatedAt,
		); err != nil {
			return nil, err
		}
		res = append(res, c)
	}
	return res, rows.Err()
}

func (r *CertificateRepository) ListByStudent(ctx context.Context, studentID string) ([]certificate.Certificate, error) {
	const q = `
SELECT id, code, student_id, course_id,
       issued_at, COALESCE(grade,''), created_at
FROM certificates
WHERE student_id = $1
ORDER BY issued_at DESC;
`
	rows, err := r.pool.Query(ctx, q, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []certificate.Certificate
	for rows.Next() {
		var c certificate.Certificate
		if err = rows.Scan(
			&c.ID, &c.Code, &c.StudentID, &c.CourseID,
			&c.IssuedAt, &c.Grade, &c.CreatedAt,
		); err != nil {
			return nil, err
		}
		res = append(res, c)
	}
	return res, rows.Err()
}

func (r *CertificateRepository) GetByID(ctx context.Context, id string) (*certificate.Certificate, error) {
	const q = `
SELECT id, code, student_id, course_id,
       issued_at, COALESCE(grade,''), created_at
FROM certificates
WHERE id = $1;
`
	row := r.pool.QueryRow(ctx, q, id)
	var c certificate.Certificate
	if err := row.Scan(
		&c.ID, &c.Code, &c.StudentID, &c.CourseID,
		&c.IssuedAt, &c.Grade, &c.CreatedAt,
	); err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *CertificateRepository) GetByCode(ctx context.Context, code string) (*certificate.Certificate, error) {
	const q = `
SELECT id, code, student_id, course_id,
       issued_at, COALESCE(grade,''), created_at
FROM certificates
WHERE code = $1;
`
	row := r.pool.QueryRow(ctx, q, code)
	var c certificate.Certificate
	if err := row.Scan(
		&c.ID, &c.Code, &c.StudentID, &c.CourseID,
		&c.IssuedAt, &c.Grade, &c.CreatedAt,
	); err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *CertificateRepository) Issue(ctx context.Context, in certificate.IssueInput) (*certificate.Certificate, error) {
	const q = `
INSERT INTO certificates (
    id, code, student_id, course_id, grade
)
VALUES ($1,$2,$3,$4,$5)
RETURNING issued_at, created_at;
`
	id := generateCode()
	code := generateCode()

	var issuedAt, createdAt time.Time
	if err := r.pool.QueryRow(
		ctx, q,
		id, code, in.StudentID, in.CourseID, in.Grade,
	).Scan(&issuedAt, &createdAt); err != nil {
		return nil, err
	}

	return &certificate.Certificate{
		ID: id, Code: code, StudentID: in.StudentID, CourseID: in.CourseID,
		IssuedAt: issuedAt, Grade: in.Grade, CreatedAt: createdAt,
	}, nil
}

func generateCode() string {
	const letters = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"
	b := make([]byte, 10)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return "CERT-" + string(b)
}
