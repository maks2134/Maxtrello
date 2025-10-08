package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"user-service/internal/config"

	_ "github.com/lib/pq"
)

func NewPostgresDB(cfg *config.Config) (*sql.DB, error) {
	connStr := cfg.GetDatabaseURL()

	log.Printf("Connecting to database: %s@%s:%s/%s",
		cfg.DBUser, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	const maxAttempts = 10
	for i := 0; i < maxAttempts; i++ {
		err = db.Ping()
		if err == nil {
			break
		}

		log.Printf("Database connection attempt %d/%d failed: %v", i+1, maxAttempts, err)

		if i < maxAttempts-1 {
			time.Sleep(2 * time.Second)
		}
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database after %d attempts: %w", maxAttempts, err)
	}

	if err := createTables(db); err != nil {
		return nil, fmt.Errorf("failed to create tables: %w", err)
	}

	log.Println("Successfully connected to PostgreSQL and tables are ready!")
	return db, nil
}

func createTables(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			email VARCHAR(255) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL,
			name VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
		
		CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
	`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create tables: %w", err)
	}

	log.Println("Database tables created/verified successfully")
	return nil
}
