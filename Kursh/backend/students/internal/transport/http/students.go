package http

import (
	"encoding/json"
	"net/http"

	"students/internal/domain/student"
)

func (h *Handler) handleListStudents(w http.ResponseWriter, r *http.Request) {
	items, err := h.studentSvc.List(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list students")
		return
	}
	writeJSON(w, http.StatusOK, items)
}

func (h *Handler) handleGetStudentByID(w http.ResponseWriter, r *http.Request) {
	id, ok := extractID(r.URL.Path, "/api/students/")
	if !ok {
		writeError(w, http.StatusBadRequest, "invalid student id")
		return
	}

	item, err := h.studentSvc.GetByID(r.Context(), id)
	if err != nil {
		writeError(w, http.StatusNotFound, "student not found")
		return
	}
	writeJSON(w, http.StatusOK, item)
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
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, updated)
}
