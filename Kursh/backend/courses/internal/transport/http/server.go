package http

import (
	"net/http"
	"time"

	"courses/internal/config"
	"courses/internal/domain/course"
)

func NewServer(cfg config.Config, svc *course.Service) *http.Server {
	mux := http.NewServeMux()
	h := NewHandler(svc)
	h.Register(mux)

	return &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}
}
