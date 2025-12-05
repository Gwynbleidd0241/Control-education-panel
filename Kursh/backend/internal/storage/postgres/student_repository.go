package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"Gwynbleidd/internal/domain/student"
)

type StudentRepository struct {
	pool *pgxpool.Pool
}

func NewStudentRepository(pool *pgxpool.Pool) *StudentRepository {
	return &StudentRepository{pool: pool}
}

// List возвращает список студентов с информацией о курсе
func (r *StudentRepository) List(ctx context.Context) ([]student.Student, error) {
	const q = `
SELECT
    s.id,
    s.full_name,
    s.email,
    s.age,
    s.performance,
    s.photo_url,
    s.city,
    s.phone,
    s.study_format,
    s.progress,
    s.created_at,
    c.id     AS course_id,
    c.title  AS course_title,
    c.price  AS course_price,
    c.level  AS course_level
FROM students s
LEFT JOIN enrollments e ON e.student_id = s.id
LEFT JOIN courses     c ON c.id = e.course_id
ORDER BY s.created_at DESC;
`
	rows, err := r.pool.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []student.Student

	for rows.Next() {
		var s student.Student

		// курс может быть NULL -> сканим в указатели
		var courseID, courseTitle, courseLevel *string
		var coursePrice *float64

		if err := rows.Scan(
			&s.ID,
			&s.FullName,
			&s.Email,
			&s.Age,
			&s.Performance,
			&s.PhotoURL,
			&s.City,
			&s.Phone,
			&s.Format,
			&s.Progress,
			&s.CreatedAt,
			&courseID,
			&courseTitle,
			&coursePrice,
			&courseLevel,
		); err != nil {
			return nil, err
		}

		if courseID != nil {
			ci := student.CourseInfo{ID: *courseID}
			if courseTitle != nil {
				ci.Title = *courseTitle
			}
			if coursePrice != nil {
				ci.Price = *coursePrice
			}
			if courseLevel != nil {
				ci.Level = *courseLevel
			}
			s.Course = &ci
		}

		res = append(res, s)
	}

	return res, rows.Err()
}

// GetByID возвращает одного студента с курсом
func (r *StudentRepository) GetByID(ctx context.Context, id string) (*student.Student, error) {
	const q = `
SELECT
    s.id,
    s.full_name,
    s.email,
    s.age,
    s.performance,
    s.photo_url,
    s.city,
    s.phone,
    s.study_format,
    s.progress,
    s.created_at,
    c.id     AS course_id,
    c.title  AS course_title,
    c.price  AS course_price,
    c.level  AS course_level
FROM students s
LEFT JOIN enrollments e ON e.student_id = s.id
LEFT JOIN courses     c ON c.id = e.course_id
WHERE s.id = $1;
`
	row := r.pool.QueryRow(ctx, q, id)

	var s student.Student
	var courseID, courseTitle, courseLevel *string
	var coursePrice *float64

	if err := row.Scan(
		&s.ID,
		&s.FullName,
		&s.Email,
		&s.Age,
		&s.Performance,
		&s.PhotoURL,
		&s.City,
		&s.Phone,
		&s.Format,
		&s.Progress,
		&s.CreatedAt,
		&courseID,
		&courseTitle,
		&coursePrice,
		&courseLevel,
	); err != nil {
		return nil, err
	}

	if courseID != nil {
		ci := student.CourseInfo{ID: *courseID}
		if courseTitle != nil {
			ci.Title = *courseTitle
		}
		if coursePrice != nil {
			ci.Price = *coursePrice
		}
		if courseLevel != nil {
			ci.Level = *courseLevel
		}
		s.Course = &ci
	}

	return &s, nil
}

// Create создаёт студента и (опционально) запись в enrollments
func (r *StudentRepository) Create(ctx context.Context, in student.CreateInput) (*student.Student, error) {
	const qStudent = `
INSERT INTO students (
    id,
    full_name,
    email,
    age,
    performance,
    photo_url,
    city,
    phone,
    study_format,
    progress
) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
RETURNING created_at;
`
	id := uuid.NewString()

	var createdAt time.Time
	if err := r.pool.QueryRow(
		ctx,
		qStudent,
		id,
		in.FullName,
		in.Email,
		in.Age,
		in.Performance,
		in.PhotoURL,
		in.City,
		in.Phone,
		in.Format,
		in.Progress,
	).Scan(&createdAt); err != nil {
		return nil, err
	}

	// Привязываем к курсу, если передан courseId
	if in.CourseID != "" {
		const qEnroll = `
INSERT INTO enrollments (id, student_id, course_id)
VALUES ($1, $2, $3);
`
		if _, err := r.pool.Exec(ctx, qEnroll, uuid.NewString(), id, in.CourseID); err != nil {
			return nil, err
		}
	}

	return &student.Student{
		ID:          id,
		FullName:    in.FullName,
		Email:       in.Email,
		Age:         in.Age,
		Performance: in.Performance,
		PhotoURL:    in.PhotoURL,
		City:        in.City,
		Phone:       in.Phone,
		Format:      in.Format,
		Progress:    in.Progress,
		CreatedAt:   createdAt,
	}, nil
}

// Update обновляет студента и (при желании) его курс
func (r *StudentRepository) Update(ctx context.Context, id string, in student.UpdateInput) (*student.Student, error) {
	const qStudent = `
UPDATE students
SET full_name    = $2,
    email        = $3,
    age          = $4,
    performance  = $5,
    photo_url    = $6,
    city         = $7,
    phone        = $8,
    study_format = $9,
    progress     = $10
WHERE id = $1
RETURNING created_at;
`

	var createdAt time.Time
	if err := r.pool.QueryRow(
		ctx,
		qStudent,
		id,
		in.FullName,
		in.Email,
		in.Age,
		in.Performance,
		in.PhotoURL,
		in.City,
		in.Phone,
		in.Format,
		in.Progress,
	).Scan(&createdAt); err != nil {
		return nil, err
	}

	// Обновляем привязку к курсу: проще всего удалить старую и вставить новую
	if in.CourseID != "" {
		const qDelete = `DELETE FROM enrollments WHERE student_id = $1;`
		if _, err := r.pool.Exec(ctx, qDelete, id); err != nil {
			return nil, err
		}

		const qEnroll = `
INSERT INTO enrollments (id, student_id, course_id)
VALUES ($1, $2, $3);
`
		if _, err := r.pool.Exec(ctx, qEnroll, uuid.NewString(), id, in.CourseID); err != nil {
			return nil, err
		}
	}

	return &student.Student{
		ID:          id,
		FullName:    in.FullName,
		Email:       in.Email,
		Age:         in.Age,
		Performance: in.Performance,
		PhotoURL:    in.PhotoURL,
		City:        in.City,
		Phone:       in.Phone,
		Format:      in.Format,
		Progress:    in.Progress,
		CreatedAt:   createdAt,
	}, nil
}
