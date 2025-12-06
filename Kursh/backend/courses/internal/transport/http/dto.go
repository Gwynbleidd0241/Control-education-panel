package http

import (
	"time"

	"courses/internal/domain/course"
)

type CourseDTO struct {
	ID              string  `json:"id"`
	Title           string  `json:"title"`
	Price           float64 `json:"price"`
	Description     string  `json:"description"`
	FullDescription string  `json:"full_description"`
	Level           string  `json:"level"`
	Duration        int     `json:"duration"`

	PhotoURL       *string  `json:"photoUrl,omitempty"`
	MaterialsURL   *string  `json:"materialsUrl,omitempty"`
	Prerequisites  *string  `json:"prerequisites,omitempty"`
	TargetAudience *string  `json:"targetAudience,omitempty"`
	Instructor     *string  `json:"instructor,omitempty"`
	Rating         *float64 `json:"rating,omitempty"`

	CreatedAt time.Time `json:"createdAt"`
}

type CreateCourseRequest struct {
	Title           string  `json:"title"`
	Price           float64 `json:"price"`
	Description     string  `json:"description"`
	FullDescription string  `json:"full_description"`
	Level           string  `json:"level"`
	Duration        int     `json:"duration"`

	PhotoURL       *string  `json:"photoUrl"`
	MaterialsURL   *string  `json:"materialsUrl"`
	Prerequisites  *string  `json:"prerequisites"`
	TargetAudience *string  `json:"targetAudience"`
	Instructor     *string  `json:"instructor"`
	Rating         *float64 `json:"rating"`
}

func toDTO(c course.Course) CourseDTO {
	return CourseDTO{
		ID:              c.ID,
		Title:           c.Title,
		Price:           c.Price,
		Description:     c.Description,
		FullDescription: c.FullDescription,
		Level:           c.Level,
		Duration:        c.DurationHours,

		PhotoURL:       c.PhotoURL,
		MaterialsURL:   c.MaterialsURL,
		Prerequisites:  c.Prerequisites,
		TargetAudience: c.TargetAudience,
		Instructor:     c.Instructor,
		Rating:         c.Rating,

		CreatedAt: c.CreatedAt,
	}
}

func toDomainCreate(in CreateCourseRequest) course.CreateInput {
	return course.CreateInput{
		Title:           in.Title,
		Price:           in.Price,
		Description:     in.Description,
		FullDescription: in.FullDescription,
		Level:           in.Level,
		Duration:        in.Duration,

		PhotoURL:       in.PhotoURL,
		MaterialsURL:   in.MaterialsURL,
		Prerequisites:  in.Prerequisites,
		TargetAudience: in.TargetAudience,
		Instructor:     in.Instructor,
		Rating:         in.Rating,
	}
}
