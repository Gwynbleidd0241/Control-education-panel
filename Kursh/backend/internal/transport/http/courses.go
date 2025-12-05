package http

import (
	"encoding/json"
	"net/http"

	"Gwynbleidd/internal/domain/course"
)

func (h *Handler) registerCourseRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/courses", h.handleListCourses)
	mux.HandleFunc("POST /api/courses", h.handleCreateCourse)
	mux.HandleFunc("GET /api/courses/", h.handleGetCourseByID)
	mux.HandleFunc("PUT /api/courses/", h.handleUpdateCourse)
	mux.HandleFunc("DELETE /api/courses/", h.handleDeleteCourse)
}

func (h *Handler) handleListCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := h.courseSvc.List(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list courses")
		return
	}
	writeJSON(w, http.StatusOK, courses)
}

func (h *Handler) handleGetCourseByID(w http.ResponseWriter, r *http.Request) {
	id, ok := extractID(r.URL.Path, "/api/courses/")
	if !ok {
		writeError(w, http.StatusBadRequest, "invalid course id")
		return
	}

	c, err := h.courseSvc.GetByID(r.Context(), id)
	if err != nil {
		writeError(w, http.StatusNotFound, "course not found")
		return
	}
	writeJSON(w, http.StatusOK, c)
}

func (h *Handler) handleCreateCourse(w http.ResponseWriter, r *http.Request) {
	var in course.CreateCourseInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
		return
	}

	created, err := h.courseSvc.Create(r.Context(), in)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to create course")
		return
	}
	writeJSON(w, http.StatusCreated, created)
}

func (h *Handler) handleUpdateCourse(w http.ResponseWriter, r *http.Request) {
	id, ok := extractID(r.URL.Path, "/api/courses/")
	if !ok {
		writeError(w, http.StatusBadRequest, "invalid course id")
		return
	}

	var in course.CreateCourseInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
		return
	}

	updated, err := h.courseSvc.Update(r.Context(), id, in)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to update course")
		return
	}
	writeJSON(w, http.StatusOK, updated)
}

func (h *Handler) handleDeleteCourse(w http.ResponseWriter, r *http.Request) {
	id, ok := extractID(r.URL.Path, "/api/courses/")
	if !ok {
		writeError(w, http.StatusBadRequest, "invalid course id")
		return
	}

	if err := h.courseSvc.Delete(r.Context(), id); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to delete course")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
