package student

import "time"

type Student struct {
	ID          string    `json:"id"`
	FullName    string    `json:"fullName"`
	Email       string    `json:"email"`
	Age         int       `json:"age"`
	Performance string    `json:"performance"`
	PhotoURL    string    `json:"photoUrl,omitempty"`
	City        string    `json:"city,omitempty"`
	Phone       string    `json:"phone,omitempty"`
	Format      string    `json:"format"`
	Progress    int       `json:"progress"`
	CourseID    *string   `json:"courseId,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type UpdateInput struct {
	FullName    string  `json:"fullName"`
	Email       string  `json:"email"`
	Age         int     `json:"age"`
	Performance string  `json:"performance"`
	PhotoURL    string  `json:"photoUrl"`
	City        string  `json:"city"`
	Phone       string  `json:"phone"`
	Format      string  `json:"format"`
	Progress    int     `json:"progress"`
	CourseID    *string `json:"courseId"`
}
