package course

import "time"

type Course struct {
	ID              string
	Title           string
	Price           float64
	Description     string
	FullDescription string
	Level           string
	DurationHours   int

	PhotoURL       *string
	MaterialsURL   *string
	Prerequisites  *string
	TargetAudience *string
	Instructor     *string
	Rating         *float64

	CreatedAt time.Time
}

type CreateInput struct {
	Title           string
	Price           float64
	Description     string
	FullDescription string
	Level           string
	Duration        int

	PhotoURL       *string
	MaterialsURL   *string
	Prerequisites  *string
	TargetAudience *string
	Instructor     *string
	Rating         *float64
}

type UpdateInput = CreateInput
