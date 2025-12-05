package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"Gwynbleidd/internal/domain/certificate"
)

type CertificateRepository struct {
	pool *pgxpool.Pool
}

func NewCertificateRepository(pool *pgxpool.Pool) *CertificateRepository {
	return &CertificateRepository{pool: pool}
}

func (r *CertificateRepository) List(ctx context.Context) ([]certificate.Certificate, error) {
	const q = `
SELECT cert.id,
       cert.code,
       cert.student_id,
       s.full_name,
       cert.course_id,
       c.title,
       cert.issued_at,
       cert.expires_at,
       cert.status,
       cert.grade
FROM certificates cert
JOIN students s ON s.id = cert.student_id
JOIN courses c ON c.id = cert.course_id
ORDER BY cert.issued_at DESC;
`
	rows, err := r.pool.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []certificate.Certificate
	for rows.Next() {
		var c certificate.Certificate
		if err := rows.Scan(
			&c.ID,
			&c.Code,
			&c.StudentID,
			&c.StudentName,
			&c.CourseID,
			&c.CourseTitle,
			&c.IssuedAt,
			&c.ExpiresAt,
			&c.Status,
			&c.Grade,
		); err != nil {
			return nil, err
		}
		res = append(res, c)
	}
	return res, rows.Err()
}

func (r *CertificateRepository) ListTemplates(ctx context.Context) ([]certificate.Template, error) {
	const q = `
SELECT t.id,
       t.course_id,
       c.title,
       t.name,
       t.validity_days,
       t.created_at
FROM certificate_templates t
JOIN courses c ON c.id = t.course_id
ORDER BY t.created_at DESC;
`
	rows, err := r.pool.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []certificate.Template
	for rows.Next() {
		var t certificate.Template
		if err := rows.Scan(
			&t.ID,
			&t.CourseID,
			&t.CourseTitle,
			&t.Name,
			&t.ValidityDays,
			&t.CreatedAt,
		); err != nil {
			return nil, err
		}
		res = append(res, t)
	}
	return res, rows.Err()
}

func (r *CertificateRepository) Issue(ctx context.Context, in certificate.IssueInput) (*certificate.Certificate, error) {
	const q = `
INSERT INTO certificates (id, code, student_id, course_id, template_id, grade, status)
VALUES ($1, $2, $3, $4, $5, $6, 'valid')
RETURNING issued_at, expires_at;
`
	id := uuid.NewString()

	var issuedAt time.Time
	var expiresAt *time.Time
	if err := r.pool.QueryRow(
		ctx,
		q,
		id,
		in.Code,
		in.StudentID,
		in.CourseID,
		in.TemplateID,
		in.Grade,
	).Scan(&issuedAt, &expiresAt); err != nil {
		return nil, err
	}

	// подтягиваем доп.данные для фронта (имя студента, название курса)
	const q2 = `
SELECT s.full_name, c.title
FROM students s
JOIN courses c ON c.id = $2
WHERE s.id = $1;
`
	var studentName, courseTitle string
	if err := r.pool.QueryRow(ctx, q2, in.StudentID, in.CourseID).Scan(&studentName, &courseTitle); err != nil {
		return nil, err
	}

	return &certificate.Certificate{
		ID:          id,
		Code:        in.Code,
		StudentID:   in.StudentID,
		StudentName: studentName,
		CourseID:    in.CourseID,
		CourseTitle: courseTitle,
		IssuedAt:    issuedAt,
		ExpiresAt:   expiresAt,
		Status:      certificate.StatusValid,
		Grade:       in.Grade,
	}, nil
}

func (r *CertificateRepository) SetStatus(ctx context.Context, id string, status certificate.Status) error {
	const q = `UPDATE certificates SET status = $2 WHERE id = $1;`
	_, err := r.pool.Exec(ctx, q, id, status)
	return err
}

func (r *CertificateRepository) GetByCode(ctx context.Context, code string) (*certificate.Certificate, error) {
	const q = `
SELECT cert.id,
       cert.code,
       cert.student_id,
       s.full_name,
       cert.course_id,
       c.title,
       cert.issued_at,
       cert.expires_at,
       cert.status,
       cert.grade
FROM certificates cert
JOIN students s ON s.id = cert.student_id
JOIN courses c ON c.id = cert.course_id
WHERE cert.code = $1;
`
	row := r.pool.QueryRow(ctx, q, code)

	var c certificate.Certificate
	if err := row.Scan(
		&c.ID,
		&c.Code,
		&c.StudentID,
		&c.StudentName,
		&c.CourseID,
		&c.CourseTitle,
		&c.IssuedAt,
		&c.ExpiresAt,
		&c.Status,
		&c.Grade,
	); err != nil {
		return nil, err
	}
	return &c, nil
}
