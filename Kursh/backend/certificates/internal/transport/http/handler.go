package http

import (
	"net/http"

	"certificates/internal/domain/certificate"
)

type Handler struct {
	certSvc *certificate.Service
}

func NewHandler(cert *certificate.Service) *Handler {
	return &Handler{certSvc: cert}
}

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", h.handleHealth)
	mux.HandleFunc("GET /api/certificates", h.handleListCertificates)
	mux.HandleFunc("POST /api/certificates/issue", h.handleIssueCertificate)
	mux.HandleFunc("GET /api/certificates/", h.handleGetCertificateByID)
}
