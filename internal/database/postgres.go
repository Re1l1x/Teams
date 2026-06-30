package database

import (
	"context"
	"fmt"
	"log"

	"github.com/Re1l1x/Teams/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(cfg *config.Config) *pgxpool.Pool {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to PostgreSQL")

	return pool
}
