package student

import "time"

// Информация о курсе, как во вложенном поле "course" в JSON
type CourseInfo struct {
	ID    string  `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
	Level string  `json:"level"`
}

// Полная модель студента под твой JSON
type Student struct {
	ID          string      `json:"id"`
	FullName    string      `json:"fullName"`
	Email       string      `json:"email"`
	Age         int         `json:"age"`
	Performance string      `json:"performance"`
	PhotoURL    string      `json:"photoUrl"`
	City        string      `json:"city"`
	Phone       string      `json:"phone"`
	Format      string      `json:"format"`   // "Очно" / "Онлайн"
	Progress    int         `json:"progress"` // 0–100
	Course      *CourseInfo `json:"course,omitempty"`
	CreatedAt   time.Time   `json:"createdAt"`
}

// Вход для создания студента
// (ожидаемый JSON для POST /students, например)
type CreateInput struct {
	FullName    string `json:"fullName"`
	Email       string `json:"email"`
	Age         int    `json:"age"`
	Performance string `json:"performance"`
	PhotoURL    string `json:"photoUrl"`
	City        string `json:"city"`
	Phone       string `json:"phone"`
	Format      string `json:"format"`
	Progress    int    `json:"progress"`
	CourseID    string `json:"courseId"` // id курса из таблицы courses
}

// Вход для обновления студента
type UpdateInput struct {
	FullName    string `json:"fullName"`
	Email       string `json:"email"`
	Age         int    `json:"age"`
	Performance string `json:"performance"`
	PhotoURL    string `json:"photoUrl"`
	City        string `json:"city"`
	Phone       string `json:"phone"`
	Format      string `json:"format"`
	Progress    int    `json:"progress"`
	CourseID    string `json:"courseId"`
}
