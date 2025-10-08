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
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * 60 * 60)

	var dbError error
	for i := 0; i < 5; i++ {
		dbError = db.Ping()
		if dbError == nil {
			break
		}
		log.Printf("Database connection attempt %d failed: %v", i+1, dbError)
		if i < 4 {
			time.Sleep(2 * time.Second)
		}
	}

	if dbError != nil {
		return nil, fmt.Errorf("failed to connect to database after 5 attempts: %v", dbError)
	}

	if err := createTables(db); err != nil {
		return nil, fmt.Errorf("failed to create tables: %v", err)
	}

	log.Println("Successfully connected !")
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
		return fmt.Errorf("failed to create tables: %v", err)
	}

	log.Println("Database tables create")
	return nil
}
