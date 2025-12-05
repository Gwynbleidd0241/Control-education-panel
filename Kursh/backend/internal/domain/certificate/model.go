package certificate

import "time"

type Status string

const (
	StatusValid   Status = "valid"
	StatusExpired Status = "expired"
	StatusRevoked Status = "revoked"
)

type Certificate struct {
	ID          string     `json:"id"`
	Code        string     `json:"code"`
	StudentID   string     `json:"studentId"`
	StudentName string     `json:"studentName"`
	CourseID    string     `json:"courseId"`
	CourseTitle string     `json:"courseTitle"`
	IssuedAt    time.Time  `json:"issuedAt"`
	ExpiresAt   *time.Time `json:"expiresAt,omitempty"`
	Status      Status     `json:"status"`
	Grade       *string    `json:"grade,omitempty"`
}

type Template struct {
	ID           string    `json:"id"`
	CourseID     string    `json:"courseId"`
	CourseTitle  string    `json:"courseTitle"`
	Name         string    `json:"name"`
	ValidityDays *int      `json:"validityDays,omitempty"`
	CreatedAt    time.Time `json:"createdAt"`
}

type IssueInput struct {
	Code       string  `json:"code"`
	StudentID  string  `json:"studentId"`
	CourseID   string  `json:"courseId"`
	TemplateID *string `json:"templateId,omitempty"`
	Grade      *string `json:"grade,omitempty"`
}
