package http

import (
	"net/http"

	"students/internal/domain/student"
)

type Handler struct {
	studentSvc *student.Service
}

func NewHandler(studentSvc *student.Service) *Handler {
	return &Handler{studentSvc: studentSvc}
}

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", h.handleHealth)
	mux.HandleFunc("GET /api/students", h.handleListStudents)
	mux.HandleFunc("GET /api/students/", h.handleGetStudentByID)
	mux.HandleFunc("PUT /api/students/", h.handleUpdateStudent)
}
