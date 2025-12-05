package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"Gwynbleidd/internal/domain/course"
)

type CourseRepository struct {
	pool *pgxpool.Pool
}

func NewCourseRepository(pool *pgxpool.Pool) *CourseRepository {
	return &CourseRepository{pool: pool}
}

func (r *CourseRepository) List(ctx context.Context) ([]course.Course, error) {
	const q = `
SELECT c.id,
       c.title,
       c.description,
       c.full_description,
       c.level,
       c.duration_hours,
       c.price,
       c.created_at,
       COALESCE(COUNT(e.id), 0) as students_count
FROM courses c
LEFT JOIN enrollments e ON e.course_id = c.id
GROUP BY c.id
ORDER BY c.created_at DESC;
`
	rows, err := r.pool.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []course.Course
	for rows.Next() {
		var c course.Course
		if err := rows.Scan(
			&c.ID,
			&c.Title,
			&c.Description,
			&c.FullDescription,
			&c.Level,
			&c.DurationHours,
			&c.Price,
			&c.CreatedAt,
			&c.StudentsCount,
		); err != nil {
			return nil, err
		}
		res = append(res, c)
	}
	return res, rows.Err()
}

func (r *CourseRepository) GetByID(ctx context.Context, id string) (*course.Course, error) {
	const q = `
SELECT c.id,
       c.title,
       c.description,
       c.level,
       c.duration_hours,
       c.price,
       c.created_at,
       COALESCE(COUNT(e.id), 0) as students_count
FROM courses c
LEFT JOIN enrollments e ON e.course_id = c.id
WHERE c.id = $1
GROUP BY c.id;
`
	row := r.pool.QueryRow(ctx, q, id)
	var c course.Course
	if err := row.Scan(
		&c.ID,
		&c.Title,
		&c.Description,
		&c.Level,
		&c.DurationHours,
		&c.Price,
		&c.CreatedAt,
		&c.StudentsCount,
	); err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *CourseRepository) Create(ctx context.Context, in course.CreateCourseInput) (*course.Course, error) {
	const q = `
INSERT INTO courses (id, title, description, full_description, level, duration_hours, price)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING created_at;
`
	id := uuid.NewString()

	var createdAt time.Time
	if err := r.pool.QueryRow(
		ctx,
		q,
		id,
		in.Title,
		in.Description,
		in.FullDescription,
		in.Level,
		in.DurationHours,
		in.Price,
	).Scan(&createdAt); err != nil {
		return nil, err
	}

	return &course.Course{
		ID:              id,
		Title:           in.Title,
		Description:     in.Description,
		FullDescription: in.FullDescription,
		Level:           in.Level,
		DurationHours:   in.DurationHours,
		Price:           in.Price,
		CreatedAt:       createdAt,
		StudentsCount:   0,
	}, nil
}

func (r *CourseRepository) Update(ctx context.Context, id string, in course.CreateCourseInput) (*course.Course, error) {
	const q = `
UPDATE courses
SET title = $2,
    description = $3,
    full_description = $4,
    level = $5,
    duration_hours = $6,
    price = $7
WHERE id = $1
RETURNING created_at;
`
	var createdAt time.Time
	if err := r.pool.QueryRow(
		ctx,
		q,
		id,
		in.Title,
		in.Description,
		in.FullDescription,
		in.Level,
		in.DurationHours,
		in.Price,
	).Scan(&createdAt); err != nil {
		return nil, err
	}

	return &course.Course{
		ID:              id,
		Title:           in.Title,
		Description:     in.Description,
		FullDescription: in.FullDescription,
		Level:           in.Level,
		DurationHours:   in.DurationHours,
		Price:           in.Price,
		CreatedAt:       createdAt,
	}, nil
}

func (r *CourseRepository) Delete(ctx context.Context, id string) error {
	const q = `DELETE FROM courses WHERE id = $1;`
	res, err := r.pool.Exec(ctx, q, id)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return errors.New("course not found")
	}
	return nil
}
