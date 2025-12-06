package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

	"courses/internal/config"
	"courses/internal/domain/course"
	pgrepo "courses/internal/storage/postgres"
	httptr "courses/internal/transport/http"
)

func main() {
	cfg := config.MustLoad()

	pool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	repo := pgrepo.NewCourseRepository(pool)
	svc := course.NewService(repo)

	srv := httptr.NewServer(cfg, svc)

	addr := ":" + cfg.Port
	log.Println("courses-service on", addr)
	log.Fatal(srv.ListenAndServe())
}
