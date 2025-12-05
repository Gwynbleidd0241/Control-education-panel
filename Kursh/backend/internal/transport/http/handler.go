package http

import (
	"encoding/json"
	"net/http"
	"strings"

	"Gwynbleidd/internal/domain/certificate"
	"Gwynbleidd/internal/domain/course"
	"Gwynbleidd/internal/domain/enrollment"
	"Gwynbleidd/internal/domain/student"
)

type Handler struct {
	courseSvc     *course.Service
	certSvc       *certificate.Service
	studentSvc    *student.Service
	enrollmentSvc *enrollment.Service
}

func NewHandler(
	courseSvc *course.Service,
	certSvc *certificate.Service,
	studentSvc *student.Service,
	enrollmentSvc *enrollment.Service,
) *Handler {
	return &Handler{
		courseSvc:     courseSvc,
		certSvc:       certSvc,
		studentSvc:    studentSvc,
		enrollmentSvc: enrollmentSvc,
	}
}

func (h *Handler) Register(mux *http.ServeMux) {
	// health
	mux.HandleFunc("GET /api/health", h.handleHealth)

	h.registerCourseRoutes(mux)
	h.registerStudentRoutes(mux)
	h.registerEnrollmentRoutes(mux)
	h.registerCertificateRoutes(mux)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	type errResp struct {
		Error string `json:"error"`
	}
	writeJSON(w, status, errResp{Error: msg})
}

func extractID(path, prefix string) (string, bool) {
	if !strings.HasPrefix(path, prefix) {
		return "", false
	}
	id := strings.TrimPrefix(path, prefix)
	if id == "" {
		return "", false
	}
	if i := strings.IndexByte(id, '/'); i >= 0 {
		id = id[:i]
	}
	return id, true
}

// GET /api/health
func (h *Handler) handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
