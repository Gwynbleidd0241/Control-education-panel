package http

import (
	"net/http"

	"gateway/internal/clients/certificates"
	"gateway/internal/clients/courses"
	"gateway/internal/clients/students"
)

type Handler struct {
	students *students.Client
	courses  *courses.Client
	certs    *certificates.Client
}

func NewHandler(s *students.Client, c *courses.Client, cert *certificates.Client) *Handler {
	return &Handler{students: s, courses: c, certs: cert}
}

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", h.handleHealth)

	mux.HandleFunc("GET /api/dashboard", h.handleDashboard)
	mux.HandleFunc("GET /api/students", h.handleListStudentsAggregated)
	mux.HandleFunc("GET /api/students/", h.handleGetStudentAggregated)
	mux.HandleFunc("PUT /api/students/", h.handleProxyStudentUpdate)

	mux.HandleFunc("GET /api/courses", h.handleProxyCoursesList)
	mux.HandleFunc("GET /api/courses/", h.handleProxyCourseByID)
	mux.HandleFunc("POST /api/courses", h.handleProxyCoursesCreate)
	mux.HandleFunc("PUT /api/courses/", h.handleProxyCoursesUpdate)
	mux.HandleFunc("DELETE /api/courses/", h.handleProxyCoursesDelete)

	mux.HandleFunc("GET /api/certificates", h.handleProxyCertificatesList)
	mux.HandleFunc("GET /api/certificates/", h.handleProxyCertificateByID)
	mux.HandleFunc("POST /api/certificates/issue", h.handleProxyCertificatesIssue)

}
