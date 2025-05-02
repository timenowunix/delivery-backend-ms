package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ClosePostgres(ctx context.Context, pool *pgxpool.Pool) {
	if pool != nil {
		pool.Close()
		log.Println("PostgreSQL connection pool closed")
	}
}
