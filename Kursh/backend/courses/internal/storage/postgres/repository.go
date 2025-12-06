package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"courses/internal/domain/course"
)

type CourseRepository struct {
	pool *pgxpool.Pool
}

func NewCourseRepository(pool *pgxpool.Pool) *CourseRepository {
	return &CourseRepository{pool: pool}
}

func (r *CourseRepository) List(ctx context.Context) ([]course.Course, error) {
	const q = `
SELECT id, title, price, description, full_description, level, duration_hours,
       photo_url, materials_url, prerequisites, target_audience, instructor, rating,
       created_at
FROM courses
ORDER BY created_at DESC;
`
	rows, err := r.pool.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []course.Course
	for rows.Next() {
		var c course.Course

		var photo, materials, prereq, ta, instr sql.NullString
		var rating sql.NullFloat64

		if err = rows.Scan(
			&c.ID,
			&c.Title,
			&c.Price,
			&c.Description,
			&c.FullDescription,
			&c.Level,
			&c.DurationHours,
			&photo,
			&materials,
			&prereq,
			&ta,
			&instr,
			&rating,
			&c.CreatedAt,
		); err != nil {
			return nil, err
		}

		if photo.Valid {
			c.PhotoURL = &photo.String
		}
		if materials.Valid {
			c.MaterialsURL = &materials.String
		}
		if prereq.Valid {
			c.Prerequisites = &prereq.String
		}
		if ta.Valid {
			c.TargetAudience = &ta.String
		}
		if instr.Valid {
			c.Instructor = &instr.String
		}
		if rating.Valid {
			v := rating.Float64
			c.Rating = &v
		}

		res = append(res, c)
	}

	return res, rows.Err()
}

func (r *CourseRepository) GetByID(ctx context.Context, id string) (*course.Course, error) {
	const q = `
SELECT id, title, price, description, full_description, level, duration_hours,
       photo_url, materials_url, prerequisites, target_audience, instructor, rating,
       created_at
FROM courses
WHERE id = $1;
`
	row := r.pool.QueryRow(ctx, q, id)

	var c course.Course
	var photo, materials, prereq, ta, instr sql.NullString
	var rating sql.NullFloat64

	if err := row.Scan(
		&c.ID,
		&c.Title,
		&c.Price,
		&c.Description,
		&c.FullDescription,
		&c.Level,
		&c.DurationHours,
		&photo,
		&materials,
		&prereq,
		&ta,
		&instr,
		&rating,
		&c.CreatedAt,
	); err != nil {
		return nil, err
	}

	if photo.Valid {
		c.PhotoURL = &photo.String
	}
	if materials.Valid {
		c.MaterialsURL = &materials.String
	}
	if prereq.Valid {
		c.Prerequisites = &prereq.String
	}
	if ta.Valid {
		c.TargetAudience = &ta.String
	}
	if instr.Valid {
		c.Instructor = &instr.String
	}
	if rating.Valid {
		v := rating.Float64
		c.Rating = &v
	}

	return &c, nil
}

func (r *CourseRepository) Create(ctx context.Context, in course.CreateInput) (*course.Course, error) {
	const q = `
INSERT INTO courses (
	title, price, description, full_description, level, duration_hours,
	photo_url, materials_url, prerequisites, target_audience, instructor, rating
) VALUES (
	$1,$2,$3,$4,$5,$6,
	$7,$8,$9,$10,$11,$12
)
RETURNING id, created_at;
`

	var createdAt time.Time
	var id string
	if err := r.pool.QueryRow(
		ctx,
		q,
		in.Title,
		in.Price,
		in.Description,
		in.FullDescription,
		in.Level,
		in.Duration,
		in.PhotoURL,
		in.MaterialsURL,
		in.Prerequisites,
		in.TargetAudience,
		in.Instructor,
		in.Rating,
	).Scan(&id, &createdAt); err != nil {
		return nil, err
	}

	return &course.Course{
		ID:              id,
		Title:           in.Title,
		Price:           in.Price,
		Description:     in.Description,
		FullDescription: in.FullDescription,
		Level:           in.Level,
		DurationHours:   in.Duration,
		PhotoURL:        in.PhotoURL,
		MaterialsURL:    in.MaterialsURL,
		Prerequisites:   in.Prerequisites,
		TargetAudience:  in.TargetAudience,
		Instructor:      in.Instructor,
		Rating:          in.Rating,
		CreatedAt:       createdAt,
	}, nil
}

func (r *CourseRepository) Update(ctx context.Context, id string, in course.UpdateInput) (*course.Course, error) {
	const q = `
UPDATE courses
SET title = $2,
	price = $3,
	description = $4,
	full_description = $5,
	level = $6,
	duration_hours = $7,
	photo_url = $8,
	materials_url = $9,
	prerequisites = $10,
	target_audience = $11,
	instructor = $12,
	rating = $13
WHERE id = $1
RETURNING created_at;
`
	var createdAt time.Time
	if err := r.pool.QueryRow(
		ctx,
		q,
		id,
		in.Title,
		in.Price,
		in.Description,
		in.FullDescription,
		in.Level,
		in.Duration,
		in.PhotoURL,
		in.MaterialsURL,
		in.Prerequisites,
		in.TargetAudience,
		in.Instructor,
		in.Rating,
	).Scan(&createdAt); err != nil {
		return nil, err
	}

	return &course.Course{
		ID:              id,
		Title:           in.Title,
		Price:           in.Price,
		Description:     in.Description,
		FullDescription: in.FullDescription,
		Level:           in.Level,
		DurationHours:   in.Duration,
		PhotoURL:        in.PhotoURL,
		MaterialsURL:    in.MaterialsURL,
		Prerequisites:   in.Prerequisites,
		TargetAudience:  in.TargetAudience,
		Instructor:      in.Instructor,
		Rating:          in.Rating,
		CreatedAt:       createdAt,
	}, nil
}

func (r *CourseRepository) Delete(ctx context.Context, id string) error {
	const q = `DELETE FROM courses WHERE id = $1;`
	tag, err := r.pool.Exec(ctx, q, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return errors.New("course not found")
	}
	return nil
}
