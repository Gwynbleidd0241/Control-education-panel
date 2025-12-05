package http

import (
	"encoding/json"
	"net/http"
	"strings"

	"Gwynbleidd/internal/domain/certificate"
)

func (h *Handler) registerCertificateRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/certificates", h.handleListCertificates)
	mux.HandleFunc("POST /api/certificates", h.handleIssueCertificate)
	mux.HandleFunc("PUT /api/certificates/", h.handleRevokeCertificate)
	mux.HandleFunc("GET /api/certificate-templates", h.handleListTemplates)
	mux.HandleFunc("GET /api/certificates/verify", h.handleVerifyCertificate)
}

func (h *Handler) handleListCertificates(w http.ResponseWriter, r *http.Request) {
	certs, err := h.certSvc.List(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list certificates")
		return
	}
	writeJSON(w, http.StatusOK, certs)
}

func (h *Handler) handleListTemplates(w http.ResponseWriter, r *http.Request) {
	templates, err := h.certSvc.ListTemplates(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list templates")
		return
	}
	writeJSON(w, http.StatusOK, templates)
}

func (h *Handler) handleIssueCertificate(w http.ResponseWriter, r *http.Request) {
	var in certificate.IssueInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
		return
	}

	created, err := h.certSvc.Issue(r.Context(), in)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to issue certificate")
		return
	}
	writeJSON(w, http.StatusCreated, created)
}

func (h *Handler) handleRevokeCertificate(w http.ResponseWriter, r *http.Request) {
	// /api/certificates/{id}/revoke
	withoutPrefix := strings.TrimPrefix(r.URL.Path, "/api/certificates/")
	if withoutPrefix == r.URL.Path || withoutPrefix == "" {
		writeError(w, http.StatusBadRequest, "invalid certificate path")
		return
	}

	parts := strings.SplitN(withoutPrefix, "/", 2)
	if len(parts) != 2 || parts[1] != "revoke" {
		writeError(w, http.StatusBadRequest, "invalid revoke path")
		return
	}
	id := parts[0]

	if err := h.certSvc.Revoke(r.Context(), id); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to revoke certificate")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) handleVerifyCertificate(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimSpace(r.URL.Query().Get("code"))
	if code == "" {
		writeError(w, http.StatusBadRequest, "code is required")
		return
	}

	cert, err := h.certSvc.VerifyByCode(r.Context(), code)
	if err != nil {
		writeError(w, http.StatusNotFound, "certificate not found")
		return
	}
	writeJSON(w, http.StatusOK, cert)
}
