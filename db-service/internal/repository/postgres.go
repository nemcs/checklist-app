package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/nemcs/checklist-app/db-service/internal/config"
	"time"
)

func NewPostgres(cfg config.PostgresConfig) (*pgx.Conn, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)

	ctx, canel := context.WithTimeout(context.Background(), 5*time.Second)
	defer canel()
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
