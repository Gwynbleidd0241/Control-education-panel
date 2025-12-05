package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"Gwynbleidd/internal/config"
	"Gwynbleidd/internal/domain/certificate"
	"Gwynbleidd/internal/domain/course"
	"Gwynbleidd/internal/domain/enrollment"
	"Gwynbleidd/internal/domain/student"
	"Gwynbleidd/internal/storage/postgres"
	transport "Gwynbleidd/internal/transport/http"
)

func main() {
	cfg := config.MustLoad()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pool, err := postgres.New(ctx, cfg.DSN())
	if err != nil {
		log.Fatalf("db init failed: %v", err)
	}
	defer pool.Close()

	// repos
	courseRepo := postgres.NewCourseRepository(pool)
	certRepo := postgres.NewCertificateRepository(pool)
	studentRepo := postgres.NewStudentRepository(pool)
	enrollmentRepo := postgres.NewEnrollmentRepository(pool)

	// services
	courseSvc := course.NewService(courseRepo)
	certSvc := certificate.NewService(certRepo)
	studentSvc := student.NewService(studentRepo)
	enrollmentSvc := enrollment.NewService(enrollmentRepo)

	mux := http.NewServeMux()
	handler := transport.NewHandler(courseSvc, certSvc, studentSvc, enrollmentSvc)
	handler.Register(mux)

	srv := &http.Server{
		Addr:    cfg.Addr(),
		Handler: withCORS(mux), // <--- ВАЖНО: оборачиваем mux
	}

	go func() {
		log.Printf("HTTP server listening on %s", cfg.Addr())
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen and serve error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	log.Println("Shutting down...")

	ctxShutdown, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := srv.Shutdown(ctxShutdown); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}
	log.Println("Server exited")
}

// withCORS оборачивает основной mux, чтобы браузер не ругался на CORS.
func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// откуда разрешаем фронт (CRA на 3000 порту)
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// preflight-запросы
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
