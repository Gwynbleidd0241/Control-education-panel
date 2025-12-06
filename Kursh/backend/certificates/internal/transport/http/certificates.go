package http

import (
	"encoding/json"
	"net/http"

	"certificates/internal/domain/certificate"
)

func (h *Handler) handleListCertificates(w http.ResponseWriter, r *http.Request) {
	items, err := h.certSvc.List(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list certificates")
		return
	}
	writeJSON(w, http.StatusOK, items)
}

func (h *Handler) handleGetCertificateByID(w http.ResponseWriter, r *http.Request) {
	id, ok := extractID(r.URL.Path, "/api/certificates/")
	if !ok {
		writeError(w, http.StatusBadRequest, "invalid certificate id")
		return
	}
	item, err := h.certSvc.GetByID(r.Context(), id)
	if err != nil {
		writeError(w, http.StatusNotFound, "certificate not found")
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func (h *Handler) handleIssueCertificate(w http.ResponseWriter, r *http.Request) {
	var in certificate.IssueInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
		return
	}

	created, err := h.certSvc.Issue(r.Context(), in)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, created)
}
