package http

import (
	"encoding/json"
	"net/http"

	"Gwynbleidd/internal/domain/student"
)

func (h *Handler) registerStudentRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/students", h.handleListStudents)
	mux.HandleFunc("POST /api/students", h.handleCreateStudent)
	mux.HandleFunc("GET /api/students/", h.handleGetStudentByID)
	mux.HandleFunc("PUT /api/students/", h.handleUpdateStudent)
}

func (h *Handler) handleListStudents(w http.ResponseWriter, r *http.Request) {
	students, err := h.studentSvc.List(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list students")
		return
	}
	writeJSON(w, http.StatusOK, students)
}

func (h *Handler) handleGetStudentByID(w http.ResponseWriter, r *http.Request) {
	id, ok := extractID(r.URL.Path, "/api/students/")
	if !ok {
		writeError(w, http.StatusBadRequest, "invalid student id")
		return
	}

	s, err := h.studentSvc.GetByID(r.Context(), id)
	if err != nil {
		writeError(w, http.StatusNotFound, "student not found")
		return
	}
	writeJSON(w, http.StatusOK, s)
}

func (h *Handler) handleCreateStudent(w http.ResponseWriter, r *http.Request) {
	var in student.CreateInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
		return
	}

	created, err := h.studentSvc.Create(r.Context(), in)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to create student")
		return
	}
	writeJSON(w, http.StatusCreated, created)
}

func (h *Handler) handleUpdateStudent(w http.ResponseWriter, r *http.Request) {
	id, ok := extractID(r.URL.Path, "/api/students/")
	if !ok {
		writeError(w, http.StatusBadRequest, "invalid student id")
		return
	}

	var in student.UpdateInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
		return
	}

	updated, err := h.studentSvc.Update(r.Context(), id, in)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to update student")
		return
	}
	writeJSON(w, http.StatusOK, updated)
}
