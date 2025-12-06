package http

import (
	"io"
	"net/http"
	"strings"
)

func (h *Handler) proxy(w http.ResponseWriter, r *http.Request, targetBase string, keepAPI bool) {
	path := r.URL.Path
	if !keepAPI {
		path = strings.TrimPrefix(path, "/api")
	}
	finalURL := targetBase + path

	req, _ := http.NewRequestWithContext(r.Context(), r.Method, finalURL, r.Body)
	req.Header = r.Header.Clone()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		writeError(w, http.StatusBadGateway, "upstream unavailable")
		return
	}
	defer resp.Body.Close()
	for k, v := range resp.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}
	w.WriteHeader(resp.StatusCode)
	_, _ = io.Copy(w, resp.Body)
}

func (h *Handler) handleProxyCoursesList(w http.ResponseWriter, r *http.Request) {
	h.proxy(w, r, h.coursesBase(), false)
}
func (h *Handler) handleProxyCourseByID(w http.ResponseWriter, r *http.Request) {
	h.proxy(w, r, h.coursesBase(), false)
}
func (h *Handler) handleProxyCoursesCreate(w http.ResponseWriter, r *http.Request) {
	h.proxy(w, r, h.coursesBase(), false)
}
func (h *Handler) handleProxyCoursesUpdate(w http.ResponseWriter, r *http.Request) {
	h.proxy(w, r, h.coursesBase(), false)
}
func (h *Handler) handleProxyCoursesDelete(w http.ResponseWriter, r *http.Request) {
	h.proxy(w, r, h.coursesBase(), false)
}

func (h *Handler) coursesBase() string {
	return h.courses.Base() // или аналогично для других сервисов!
}

func (h *Handler) handleProxyCertificatesList(w http.ResponseWriter, r *http.Request) {
	h.proxy(w, r, h.certsBase(), true)
}

func (h *Handler) handleProxyCertificateByID(w http.ResponseWriter, r *http.Request) {
	h.proxy(w, r, h.certsBase(), true)
}

func (h *Handler) handleProxyCertificatesIssue(w http.ResponseWriter, r *http.Request) {
	h.proxy(w, r, h.certsBase(), true)
}

func (h *Handler) certsBase() string {
	return h.certs.Base()
}

func (h *Handler) handleProxyStudentUpdate(w http.ResponseWriter, r *http.Request) {
	h.proxy(w, r, h.studentsBase(), true)
}

func (h *Handler) studentsBase() string {
	return h.students.Base()
}
