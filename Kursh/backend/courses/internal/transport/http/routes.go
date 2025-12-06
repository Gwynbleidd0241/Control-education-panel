package http

import "net/http"

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte("ok"))
	})

	mux.HandleFunc("GET /courses", h.handleList)
	mux.HandleFunc("POST /courses", h.handleCreate)

	mux.HandleFunc("GET /courses/", h.handleGetByID)
	mux.HandleFunc("PUT /courses/", h.handleUpdate)
	mux.HandleFunc("DELETE /courses/", h.handleDelete)
}
