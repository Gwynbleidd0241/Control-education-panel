package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"Gwynbleidd/internal/domain/enrollment"
)

type EnrollmentRepository struct {
	pool *pgxpool.Pool
}

func NewEnrollmentRepository(pool *pgxpool.Pool) *EnrollmentRepository {
	return &EnrollmentRepository{pool: pool}
}

func (r *EnrollmentRepository) List(ctx context.Context) ([]enrollment.Enrollment, error) {
	const q = `
SELECT e.id,
       e.student_id,
       s.full_name,
       e.course_id,
       c.title,
       e.status,
       e.enrolled_at,
       e.completed_at
FROM enrollments e
JOIN students s ON s.id = e.student_id
JOIN courses  c ON c.id = e.course_id
ORDER BY e.enrolled_at DESC;
`
	rows, err := r.pool.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []enrollment.Enrollment
	for rows.Next() {
		var e enrollment.Enrollment
		if err := rows.Scan(
			&e.ID,
			&e.StudentID,
			&e.StudentName,
			&e.CourseID,
			&e.CourseTitle,
			&e.Status,
			&e.EnrolledAt,
			&e.CompletedAt,
		); err != nil {
			return nil, err
		}
		res = append(res, e)
	}
	return res, rows.Err()
}

func (r *EnrollmentRepository) GetByID(ctx context.Context, id string) (*enrollment.Enrollment, error) {
	const q = `
SELECT e.id,
       e.student_id,
       s.full_name,
       e.course_id,
       c.title,
       e.status,
       e.enrolled_at,
       e.completed_at
FROM enrollments e
JOIN students s ON s.id = e.student_id
JOIN courses  c ON c.id = e.course_id
WHERE e.id = $1;
`
	row := r.pool.QueryRow(ctx, q, id)

	var e enrollment.Enrollment
	if err := row.Scan(
		&e.ID,
		&e.StudentID,
		&e.StudentName,
		&e.CourseID,
		&e.CourseTitle,
		&e.Status,
		&e.EnrolledAt,
		&e.CompletedAt,
	); err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *EnrollmentRepository) Create(ctx context.Context, in enrollment.CreateInput) (*enrollment.Enrollment, error) {
	const q = `
INSERT INTO enrollments (id, student_id, course_id, status)
VALUES ($1, $2, $3, $4)
RETURNING enrolled_at;
`
	id := uuid.NewString()
	status := enrollment.StatusInProgress

	var enrolledAt time.Time
	if err := r.pool.QueryRow(ctx, q, id, in.StudentID, in.CourseID, status).Scan(&enrolledAt); err != nil {
		return nil, err
	}

	const q2 = `
SELECT s.full_name, c.title
FROM students s
JOIN courses  c ON c.id = $2
WHERE s.id = $1;
`
	var studentName, courseTitle string
	if err := r.pool.QueryRow(ctx, q2, in.StudentID, in.CourseID).Scan(&studentName, &courseTitle); err != nil {
		return nil, err
	}

	return &enrollment.Enrollment{
		ID:          id,
		StudentID:   in.StudentID,
		StudentName: studentName,
		CourseID:    in.CourseID,
		CourseTitle: courseTitle,
		Status:      status,
		EnrolledAt:  enrolledAt,
		CompletedAt: nil,
	}, nil
}

func (r *EnrollmentRepository) SetStatus(ctx context.Context, id string, status enrollment.Status) error {
	const q = `
UPDATE enrollments
SET status = $2,
    completed_at = CASE
        WHEN $2 = 'completed' THEN now()
        ELSE completed_at
    END
WHERE id = $1;
`
	_, err := r.pool.Exec(ctx, q, id, status)
	return err
}
