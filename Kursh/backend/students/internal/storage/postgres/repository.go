package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"students/internal/domain/student"
)

type StudentRepository struct {
	pool *pgxpool.Pool
}

func NewStudentRepository(pool *pgxpool.Pool) *StudentRepository {
	return &StudentRepository{pool: pool}
}

func (r *StudentRepository) List(ctx context.Context) ([]student.Student, error) {
	const q = `
SELECT id, full_name, email, age, performance,
       COALESCE(photo_url, ''), COALESCE(city, ''), COALESCE(phone, ''),
       format, progress, course_id, created_at, updated_at
FROM students
ORDER BY created_at DESC;
`
	rows, err := r.pool.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []student.Student
	for rows.Next() {
		var s student.Student
		var courseID *string
		if err = rows.Scan(
			&s.ID, &s.FullName, &s.Email, &s.Age, &s.Performance,
			&s.PhotoURL, &s.City, &s.Phone,
			&s.Format, &s.Progress, &courseID, &s.CreatedAt, &s.UpdatedAt,
		); err != nil {
			return nil, err
		}
		if courseID != nil {
			s.CourseID = courseID
		}
		res = append(res, s)
	}
	return res, rows.Err()
}

func (r *StudentRepository) GetByID(ctx context.Context, id string) (*student.Student, error) {
	const q = `
SELECT id, full_name, email, age, performance,
       COALESCE(photo_url, ''), COALESCE(city, ''), COALESCE(phone, ''),
       format, progress, course_id, created_at, updated_at
FROM students
WHERE id = $1;
`
	row := r.pool.QueryRow(ctx, q, id)

	var s student.Student
	var courseID *string
	if err := row.Scan(
		&s.ID, &s.FullName, &s.Email, &s.Age, &s.Performance,
		&s.PhotoURL, &s.City, &s.Phone,
		&s.Format, &s.Progress, &courseID, &s.CreatedAt, &s.UpdatedAt,
	); err != nil {
		return nil, err
	}

	if courseID != nil {
		s.CourseID = courseID
	}

	return &s, nil
}

func (r *StudentRepository) Update(ctx context.Context, id string, in student.UpdateInput) (*student.Student, error) {
	const q = `
UPDATE students
SET full_name = $2,
    email = $3,
    age = $4,
    performance = $5,
    photo_url = $6,
    city = $7,
    phone = $8,
    format = $9,
    progress = $10,
    course_id = $11,
    updated_at = now()
WHERE id = $1
RETURNING created_at, updated_at;
`

	var createdAt, updatedAt time.Time
	if err := r.pool.QueryRow(
		ctx, q,
		id,
		in.FullName,
		in.Email,
		in.Age,
		in.Performance,
		nullIfEmpty(in.PhotoURL),
		nullIfEmpty(in.City),
		nullIfEmpty(in.Phone),
		in.Format,
		in.Progress,
		in.CourseID,
	).Scan(&createdAt, &updatedAt); err != nil {
		return nil, err
	}

	return r.GetByID(ctx, id)
}

func nullIfEmpty(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

var _ student.Repository = (*StudentRepository)(nil)
