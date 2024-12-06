package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/okaaryanata/loan/internal/app/migration"
)

func InitPostgresDB() (*pgxpool.Pool, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Failed to parse database configuration: %v", err)
	}

	dbPool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed while create connection to database: %w", err)
	}

	if err := dbPool.Ping(ctx); err != nil {
		dbPool.Close()
		return nil, fmt.Errorf("error while ping connection to database: %w", err)
	}

	initTable, _ := strconv.ParseBool(os.Getenv("DB_INIT_TABLE"))
	if initTable {
		err = initTableDDL(ctx, dbPool)
		if err != nil {
			dbPool.Close()
			return nil, fmt.Errorf("error while init table: %w", err)
		}
	}

	log.Println("Successfully connected to the db")
	return dbPool, nil
}

func initTableDDL(ctx context.Context, db *pgxpool.Pool) error {
	queries := []string{
		migration.QueryInitTableUsers,
		migration.QueryInitTableLoans,
		migration.QueryInitTableRepayments,
	}

	for _, query := range queries {
		_, err := db.Exec(ctx, query)
		if err != nil {
			return err
		}
	}

	return nil
}
