package http

import (
	"encoding/json"
	"net/http"

	"courses/internal/domain/course"
)

type Handler struct {
	svc *course.Service
}

func NewHandler(svc *course.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) handleList(w http.ResponseWriter, r *http.Request) {
	items, err := h.svc.List(r.Context())
	if err != nil {
		writeError(w, 500, "failed to list courses")
		return
	}

	out := make([]CourseDTO, 0, len(items))
	for _, c := range items {
		out = append(out, toDTO(c))
	}

	writeJSON(w, 200, out)
}

func (h *Handler) handleGetByID(w http.ResponseWriter, r *http.Request) {
	id := extractID(r.URL.Path, "/courses/")
	if id == "" {
		writeError(w, 400, "invalid course id")
		return
	}

	c, err := h.svc.GetByID(r.Context(), id)
	if err != nil {
		writeError(w, 404, "course not found")
		return
	}

	writeJSON(w, 200, toDTO(*c))
}

func (h *Handler) handleCreate(w http.ResponseWriter, r *http.Request) {
	var req CreateCourseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, 400, "invalid json")
		return
	}

	if req.Title == "" || req.Description == "" || req.FullDescription == "" || req.Level == "" {
		writeError(w, 400, "missing required fields")
		return
	}
	if req.Price <= 0 || req.Duration <= 0 {
		writeError(w, 400, "price and duration must be positive")
		return
	}

	created, err := h.svc.Create(r.Context(), toDomainCreate(req))
	if err != nil {
		writeError(w, 500, "failed to create course")
		return
	}

	writeJSON(w, 201, toDTO(*created))
}

func (h *Handler) handleUpdate(w http.ResponseWriter, r *http.Request) {
	id := extractID(r.URL.Path, "/courses/")
	if id == "" {
		writeError(w, 400, "invalid course id")
		return
	}

	var req CreateCourseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, 400, "invalid json")
		return
	}

	if req.Title == "" || req.Description == "" || req.FullDescription == "" || req.Level == "" {
		writeError(w, 400, "missing required fields")
		return
	}
	if req.Price <= 0 || req.Duration <= 0 {
		writeError(w, 400, "price and duration must be positive")
		return
	}

	updated, err := h.svc.Update(r.Context(), id, toDomainCreate(req))
	if err != nil {
		writeError(w, 500, "failed to update course")
		return
	}

	writeJSON(w, 200, toDTO(*updated))
}

func (h *Handler) handleDelete(w http.ResponseWriter, r *http.Request) {
	id := extractID(r.URL.Path, "/courses/")
	if id == "" {
		writeError(w, 400, "invalid course id")
		return
	}

	if err := h.svc.Delete(r.Context(), id); err != nil {
		writeError(w, 404, "course not found")
		return
	}

	w.WriteHeader(204)
}
