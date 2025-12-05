package http

import (
	"encoding/json"
	"net/http"

	"Gwynbleidd/internal/domain/enrollment"
)

func (h *Handler) registerEnrollmentRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/enrollments", h.handleListEnrollments)
	mux.HandleFunc("POST /api/enrollments", h.handleCreateEnrollment)
	mux.HandleFunc("PUT /api/enrollments/", h.handleSetEnrollmentStatus)
}

func (h *Handler) handleListEnrollments(w http.ResponseWriter, r *http.Request) {
	enrs, err := h.enrollmentSvc.List(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list enrollments")
		return
	}
	writeJSON(w, http.StatusOK, enrs)
}

func (h *Handler) handleCreateEnrollment(w http.ResponseWriter, r *http.Request) {
	var in enrollment.CreateInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
		return
	}

	created, err := h.enrollmentSvc.Create(r.Context(), in)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to create enrollment")
		return
	}
	writeJSON(w, http.StatusCreated, created)
}

func (h *Handler) handleSetEnrollmentStatus(w http.ResponseWriter, r *http.Request) {
	id, ok := extractID(r.URL.Path, "/api/enrollments/")
	if !ok {
		writeError(w, http.StatusBadRequest, "invalid enrollment id")
		return
	}

	var in struct {
		Status enrollment.Status `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
		return
	}

	if err := h.enrollmentSvc.SetStatus(r.Context(), id, in.Status); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to update enrollment status")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
