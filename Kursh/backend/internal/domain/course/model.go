package course

import "time"

type Level string

const (
	LevelBeginner     Level = "beginner"
	LevelIntermediate Level = "intermediate"
	LevelAdvanced     Level = "advanced"
)

type Course struct {
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	FullDescription string    `json:"full_description"`
	Level           Level     `json:"level"`
	Category        string    `json:"category"`
	DurationHours   int       `json:"durationHours"`
	Price           float64   `json:"price"`
	CreatedAt       time.Time `json:"createdAt"`
	StudentsCount   int       `json:"studentsCount"`
}

type CreateCourseInput struct {
	Title           string  `json:"title"`
	Description     string  `json:"description"`
	FullDescription string  `json:"full_description"`
	Level           Level   `json:"level"`
	Category        string  `json:"category"`
	DurationHours   int     `json:"durationHours"`
	Price           float64 `json:"price"`
}
