package certificate

import "time"

type Certificate struct {
	ID         string     `json:"id"`
	Code       string     `json:"code"`
	StudentID  string     `json:"studentId"`
	CourseID   string     `json:"courseId"`
	IssuedAt   time.Time  `json:"issuedAt"`
	Grade      string     `json:"grade,omitempty"`
	CreatedAt  time.Time  `json:"createdAt"`
}

type IssueInput struct {
	StudentID  string  `json:"studentId"`
	CourseID   string  `json:"courseId"`
	Grade      string  `json:"grade"`
}
