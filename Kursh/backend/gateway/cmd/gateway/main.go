package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gateway/internal/clients/certificates"
	"gateway/internal/clients/courses"
	"gateway/internal/clients/students"
	"gateway/internal/config"
	httptransport "gateway/internal/transport/http"
	"gateway/internal/transport/http/middleware"
)

func main() {
	cfg := config.MustLoad()

	stCl := students.New(cfg.StudentsURL)
	crCl := courses.New(cfg.CoursesURL)
	ctCl := certificates.New(cfg.CertsURL)

	h := httptransport.NewHandler(stCl, crCl, ctCl)

	mux := http.NewServeMux()
	h.Register(mux)
	handler := middleware.CORS(mux)

	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		log.Printf("gateway on :%s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
			log.Fatal(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
}
