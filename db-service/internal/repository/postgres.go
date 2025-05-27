package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nemcs/checklist-app/db-service/internal/config"
)

func NewPostgres(cfg config.PostgresConfig) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)

	dbpool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	return dbpool, nil
}
