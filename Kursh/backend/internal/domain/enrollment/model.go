package enrollment

import "time"

type Status string

const (
	StatusInProgress Status = "in_progress"
	StatusCompleted  Status = "completed"
	StatusDropped    Status = "dropped"
)

type Enrollment struct {
	ID          string     `json:"id"`
	StudentID   string     `json:"studentId"`
	StudentName string     `json:"studentName"`
	CourseID    string     `json:"courseId"`
	CourseTitle string     `json:"courseTitle"`
	Status      Status     `json:"status"`
	EnrolledAt  time.Time  `json:"enrolledAt"`
	CompletedAt *time.Time `json:"completedAt,omitempty"`
}

// данные для создания записи
type CreateInput struct {
	StudentID string `json:"studentId"`
	CourseID  string `json:"courseId"`
}
